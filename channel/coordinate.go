package channel

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/perun-network/perun-eth-backend/bindings/adjudicator"
	"github.com/perun-network/perun-eth-backend/subscription"
	"github.com/pkg/errors"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/log"
	"perun.network/go-perun/wallet"
)

func (c *Coordinator) ensureCoordinated(ctx context.Context, req channel.AdjudicatorReq, subChannels []channel.SignedState, coordSigs []wallet.Sig) error {
	// Check whether it is already coordinated or concluded. If it is already concluded, we can return early. If it is not coordinated, the call to Coordinate will fail with an error.
	if coordinated, err := c.isCoordinated(ctx, req); err != nil {
		return errors.WithMessage(err, "isCoordinated")
	} else if coordinated {
		return nil
	}
	if concluded, err := c.isConcluded(ctx, req.Tx.ID); err != nil {
		return errors.WithMessage(err, "isConcluded")
	} else if concluded {
		return nil
	}

	// Wait until we can conclude.
	err := c.waitCoordinable(ctx, req)
	if err != nil {
		return fmt.Errorf("waiting for concludability: %w", err)
	}

	// No coordinate event found in the past, send transaction.
	err = c.coordinate(ctx, req, subChannels, coordSigs)
	if err != nil {
		return errors.WithMessage(err, "coordinating")
	}

	// A concurrent coordinate can make our own transaction revert. In that case,
	// re-check the current on-chain state before waiting for a fresh event.
	if coordinated, err := c.isCoordinated(ctx, req); err != nil {
		return errors.WithMessage(err, "re-checking coordinated state")
	} else if coordinated {
		return nil
	}
	if coordinated, err := c.hasCoordinatedEvent(ctx, req.Tx.ID); err != nil {
		return errors.WithMessage(err, "checking coordinated event")
	} else if coordinated {
		return nil
	}

	// Wait for coordinated event.
	sub, events, subErr, err := c.createEventSub(ctx, req.Tx.ID, false, startBlockOffset)
	if err != nil {
		return errors.WithMessage(err, "subscribing")
	}
	defer sub.Close()
	for {
		select {
		case _e := <-events:
			e, ok := _e.Data.(*adjudicator.AdjudicatorChannelUpdate)
			if !ok {
				log.Panic("wrong event type")
			}
			if e.Phase == phaseCoordinated || e.Phase == phaseConcluded {
				return nil
			}
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "context cancelled")
		case err = <-subErr:
			if err != nil {
				return errors.WithMessage(err, "subscription error")
			}
			return errors.New("subscription closed")
		}
	}
}

func (c *Coordinator) createEventSub(
	ctx context.Context,
	ch channel.ID,
	past bool,
	pastBlocks uint64,
) (
	*subscription.ResistantEventSub,
	<-chan *subscription.Event,
	<-chan error,
	error,
) {
	sub, err := subscription.Subscribe(
		ctx,
		c.ContractBackend,
		c.bound,
		updateEventType(ch),
		pastBlocks,
		c.txFinalityDepth,
	)
	if err != nil {
		return nil, nil, nil, errors.WithMessage(err, "subscribing")
	}

	events := make(chan *subscription.Event, adjEventBuffSize)
	subErr := make(chan error, 1)
	go func() {
		defer close(events)
		if past {
			subErr <- sub.ReadPast(ctx, events)
		} else {
			subErr <- sub.Read(ctx, events)
		}
	}()
	return sub, events, subErr, nil
}

// isCoordinated returns whether a channel is already coordinated.
func (c *Coordinator) isCoordinated(ctx context.Context, req channel.AdjudicatorReq) (bool, error) {
	dispute, err := c.dispute(ctx, req.Tx.ID)
	if err != nil {
		return false, errors.WithMessage(err, "querying dispute")
	}
	return hasRecordedDispute(dispute) && dispute.Phase == phaseCoordinated, nil
}

// waitCoordinable waits until the specified channel is coordinable.
func (c *Coordinator) waitCoordinable(ctx context.Context, req channel.AdjudicatorReq) error {
	dispute, err := c.dispute(ctx, req.Tx.ID)
	if err != nil {
		return errors.WithMessage(err, "querying dispute")
	}
	if !hasRecordedDispute(dispute) {
		return errors.New("channel not registered")
	}

	switch dispute.Phase {
	case phaseDispute:
		t := dispute.Timeout
		if dispute.HasApp && !channel.IsNoApp(req.Params.App) {
			return errors.New("cannot coordinate app channel in dispute phase") //TODO: we could allow coordinate in dispute phase for ledger channels, but this is not needed for our current use cases and can be added later if needed.
		}
		t += concludeWaitSlack
		return NewBlockTimeout(c.ContractBackend, t).Wait(ctx)
	case phaseForceExec:
		return errors.New("cannot coordinate app channel") //TODO: we could allow coordinate in force exec phase for ledger channels, but this is not needed for our current use cases and can be added later if needed.
	case phaseCoordinated:
		return nil
	case phaseConcluded:
		return nil
	default:
		return nil
	}
}

func (c *Coordinator) coordinate(ctx context.Context, req channel.AdjudicatorReq, subChannels []channel.SignedState,
	coordSigs []wallet.Sig) error {
	// coordinate will only succeed if the channel is already registered and the state is not yet concluded.
	// If the channel is already concluded, we can return early. If the channel is not yet registered, the call to coordinate will fail with an error. If the channel is registered but not yet coordinated, we can call coordinate. If the channel is already coordinated, the call to coordinate will fail with an error,
	// but we can ignore this error and wait for the event instead.
	// coordinate is not working with forced execution, so if the channel is force executed, we can directly call concludeFinal without checking whether the channel is already coordinated.
	// If the channel is already coordinated, we can return early. If the channel is not yet coordinated, we can call coordinate. If the channel is already coordinated, the call to coordinate will fail with an error, but we can ignore this error and wait for the event instead.
	err := errors.WithMessage(c.callCoordinate(ctx, req, subChannels, coordSigs), "calling coordinate")
	if isBenignCoordinateTxFailure(err) {
		c.log.WithError(err).Warn("Calling coordinate failed, waiting for event anyways...")
	} else if err != nil {
		return err
	}
	return nil
}

func (c *Coordinator) callCoordinate(ctx context.Context, req channel.AdjudicatorReq, subChannels []channel.SignedState, coordSigs []wallet.Sig) error {
	return c.call(ctx, req,
		func(opts *bind.TransactOpts, params adjudicator.ChannelParams, state adjudicator.ChannelState, sigs [][]byte) (*types.Transaction, error) {
			ch := adjudicator.AdjudicatorSignedState{
				Params: params,
				State:  state,
				Sigs:   sigs,
			}
			sub := toEthSignedStates(subChannels)
			return c.contract.Coordinate(opts, ch, sub, coordSigs)
		}, Coordinate)
}

// isConcluded returns whether a channel is already concluded.
func (c *Coordinator) isConcluded(ctx context.Context, ch channel.ID) (bool, error) {
	dispute, err := c.dispute(ctx, ch)
	if err != nil {
		return false, errors.WithMessage(err, "querying dispute")
	}
	return hasRecordedDispute(dispute) && dispute.Phase == phaseConcluded, nil
}

func (c *Coordinator) dispute(ctx context.Context, ch channel.ID) (disputeInfo, error) {
	return c.contract.Disputes(&bind.CallOpts{Context: ctx}, ch)
}

func isBenignCoordinateTxFailure(err error) bool {
	if !IsErrTxFailed(err) {
		return false
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "coordinated already") || strings.Contains(msg, "concluded already")
}

func (c *Coordinator) hasCoordinatedEvent(ctx context.Context, ch channel.ID) (bool, error) {

	sub, err := subscription.NewEventSub(ctx, c.ContractBackend, c.bound, updateEventType(ch), startBlockOffset)
	if err != nil {
		return false, errors.WithMessage(err, "subscribing")
	}
	defer sub.Close()
	events := make(chan *subscription.Event, adjEventBuffSize)
	subErr := make(chan error, 1)
	go func() {
		defer close(events)
		subErr <- sub.ReadPast(ctx, events)
	}()

	// Read all events and check for coordinated.
	for _e := range events {
		e, ok := _e.Data.(*adjudicator.AdjudicatorChannelUpdate)
		if !ok {
			log.Panic("wrong event type")
		}
		if e.Phase == phaseCoordinated {
			return true, nil
		}
	}
	return false, errors.WithMessage(<-subErr, "reading past events")
}
