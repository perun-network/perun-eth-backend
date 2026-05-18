package channel

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/perun-network/perun-eth-backend/bindings"
	"github.com/perun-network/perun-eth-backend/bindings/adjudicator"
	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"
	"github.com/perun-network/perun-eth-backend/subscription"
	"github.com/perun-network/perun-eth-backend/wallet"
	"github.com/pkg/errors"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/log"
)

func (c *Coordinator) Subscribe(ctx context.Context, chID channel.ID) (channel.AdjudicatorSubscription, error) {
	subErr := make(chan error, 1)
	events := make(chan *subscription.Event, adjEventBuffSize)
	eFact := func() *subscription.Event {
		return &subscription.Event{
			Name:   bindings.Events.AdjChannelUpdate,
			Data:   new(adjudicator.AdjudicatorChannelUpdate),
			Filter: [][]interface{}{{chID}},
		}
	}
	pastBlocks := uint64(0)
	if dispute, err := c.dispute(ctx, chID); err != nil {
		return nil, errors.WithMessage(err, "querying dispute")
	} else if hasRecordedDispute(dispute) {
		pastBlocks = startBlockOffset
	}
	sub, err := subscription.Subscribe(ctx, c.ContractBackend, c.bound, eFact, pastBlocks, c.txFinalityDepth)
	if err != nil {
		return nil, errors.WithMessage(err, "creating filter-watch event subscription")
	}
	// The resistant event subscription already replays past events before it
	// streams future ones, and RegisteredSub keeps only the newest matching event
	// for the caller.
	go func() {
		subErr <- sub.Read(ctx, events)
	}()
	rsub := &RegisteredSub{
		cr:     c.ContractInterface,
		sub:    sub,
		subErr: subErr,
		next:   make(chan channel.AdjudicatorEvent, 1),
		err:    make(chan error, 1),
	}
	go rsub.updateNextCoord(ctx, events, c)

	return rsub, nil
}

func (r *RegisteredSub) updateNextCoord(ctx context.Context, events chan *subscription.Event, c *Coordinator) {
evloop:
	for {
		select {
		case _next := <-events:
			err := r.processNextCoord(ctx, c, _next)
			if err != nil {
				r.err <- err
				break evloop
			}
		case err := <-r.subErr:
			if err != nil {
				r.err <- errors.WithMessage(err, "EventSub closed")
			} else {
				// Normal closing should produce no error
				close(r.err)
			}
			break evloop
		}
	}

	// subscription got closed, close next channel and return
	select {
	case <-r.next:
	default:
	}
	close(r.next)
}

func (r *RegisteredSub) processNextCoord(ctx context.Context, c *Coordinator, _next *subscription.Event) (err error) {
	next, ok := _next.Data.(*adjudicator.AdjudicatorChannelUpdate)
	next.Raw = _next.Log
	if !ok {
		log.Panicf("unexpected event type: %T", _next.Data)
	}

	select {
	// drain next-channel on new event
	case current := <-r.next:
		currentTimeout, ok := current.Timeout().(*BlockTimeout)
		if !ok {
			log.Panic("wrong timeout type")
		}
		// if newer version or same version and newer timeout, replace
		if current.Version() < next.Version || current.Version() == next.Version && currentTimeout.Time < next.Timeout {
			var e channel.AdjudicatorEvent
			e, err = c.convertEvent(ctx, next)
			if err != nil {
				return
			}

			r.next <- e
		} else { // otherwise, reuse old
			r.next <- current
		}
	default: // next-channel is empty
		var e channel.AdjudicatorEvent
		e, err = c.convertEvent(ctx, next)
		if err != nil {
			return
		}

		r.next <- e
	}
	return err
}

//nolint:funlen
func (c *Coordinator) convertEvent(ctx context.Context, e *adjudicator.AdjudicatorChannelUpdate) (channel.AdjudicatorEvent, error) {
	base := channel.NewAdjudicatorEventBase(e.ChannelID, NewBlockTimeout(c.ContractInterface, e.Timeout), e.Version)
	switch e.Phase {
	case phaseDispute:
		args, err := c.fetchRegisterCallData(ctx, e.Raw.TxHash)
		if err != nil {
			return nil, errors.WithMessage(err, "fetching call data")
		}

		ch, ok := args.signedState(e.ChannelID)
		if !ok {
			return nil, errors.Errorf("channel not found in calldata: %v", e.ChannelID)
		}

		var app channel.App
		var zeroAddress common.Address
		if ch.Params.App == zeroAddress {
			app = channel.NoApp()
		} else {
			appAddr := wallet.AsWalletAddr(ch.Params.App)
			appID := &AppID{
				Address: appAddr,
			}

			app, err = channel.Resolve(appID)
			if err != nil {
				return nil, err
			}
		}
		state := FromEthState(app, &ch.State)

		return &channel.RegisteredEvent{
			AdjudicatorEventBase: *base,
			State:                &state,
			Sigs:                 ch.Sigs,
		}, nil

	case phaseForceExec:
		args, err := c.fetchProgressCallData(ctx, e.Raw.TxHash)
		if err != nil {
			return nil, errors.WithMessage(err, "fetching call data")
		}
		appAddr := wallet.AsWalletAddr(args.Params.App)
		appID := &AppID{
			Address: appAddr,
		}
		app, err := channel.Resolve(appID)
		if err != nil {
			return nil, errors.WithMessage(err, "resolving app")
		}
		newState := FromEthState(app, &args.State)
		return &channel.ProgressedEvent{
			AdjudicatorEventBase: *base,
			State:                &newState,
			Idx:                  channel.Index(args.ActorIdx.Uint64()),
		}, nil

	case phaseCoordinated:
		args, err := c.fetchCoordinateCallData(ctx, e.Raw.TxHash)
		if err != nil {
			return nil, errors.WithMessage(err, "fetching call data")
		}

		ch, ok := args.signedState(e.ChannelID)
		if !ok {
			return nil, errors.Errorf("channel not found in calldata: %v", e.ChannelID)
		}

		var app channel.App
		var zeroAddress common.Address
		if ch.Params.App == zeroAddress {
			app = channel.NoApp()
		} else {
			appAddr := wallet.AsWalletAddr(ch.Params.App)
			appID := &AppID{
				Address: appAddr,
			}

			app, err = channel.Resolve(appID)
			if err != nil {
				return nil, err
			}
		}
		state := FromEthState(app, &ch.State)

		return &channel.CoordinatedEvent{
			AdjudicatorEventBase: *base,
			State:                &state,
			Sigs:                 ch.Sigs,
		}, nil
	case phaseConcluded:
		return &channel.ConcludedEvent{AdjudicatorEventBase: *base}, nil

	default:
		panic("unknown phase")
	}
}

func (c *Coordinator) fetchCallData(ctx context.Context, txHash common.Hash, method abi.Method, args interface{}) error {
	tx, _, err := c.TransactionByHash(ctx, txHash)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return errors.WithMessage(err, "getting transaction")
	}

	argsData := tx.Data()[len(method.ID):]

	argsI, err := method.Inputs.UnpackValues(argsData)
	if err != nil {
		return errors.WithMessage(err, "unpacking")
	}

	err = method.Inputs.Copy(args, argsI)
	if err != nil {
		return errors.WithMessage(err, "copying into struct")
	}

	return nil
}

func (c *Coordinator) fetchProgressCallData(ctx context.Context, txHash common.Hash) (*progressCallData, error) {
	var args progressCallData
	err := c.fetchCallData(ctx, txHash, abiProgress, &args)
	return &args, errors.WithMessage(err, "fetching call data")
}

func (c *Coordinator) fetchRegisterCallData(ctx context.Context, txHash common.Hash) (*registerCallData, error) {
	var args registerCallData
	err := c.fetchCallData(ctx, txHash, abiRegister, &args)
	return &args, errors.WithMessage(err, "fetching call data")
}

func (c *Coordinator) fetchCoordinateCallData(ctx context.Context, txHash common.Hash) (*coordinateCallData, error) {
	var args coordinateCallData
	err := c.fetchCallData(ctx, txHash, abiCoordinate, &args)
	return &args, errors.WithMessage(err, "fetching call data")
}
