package channel

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/perun-network/perun-eth-backend/bindings"
	"github.com/perun-network/perun-eth-backend/bindings/adjudicator"
	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"
	"github.com/pkg/errors"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/client"
	"perun.network/go-perun/log"
	"perun.network/go-perun/wallet"
	psync "polycry.pt/poly-go/sync"
)

// Ensure Coordinator implements CoordinatorBackend.
var _ channel.CoordinatorSubscriber = (*Coordinator)(nil)

// The Coordinator struct implements the channel.Coordinator interface.
// It provides all	functionality to coordinate a channel on-chain.
type Coordinator struct {
	ContractBackend
	// chainID specifies the chain the funder is living on.
	chainID  ChainID
	contract *adjudicator.Adjudicator
	bound    *bind.BoundContract
	// The address to which we send all funds.
	Receiver common.Address
	// Structured logger
	log log.Logger
	// Transaction mutex
	mu psync.Mutex
	// txSender is sending the TX.
	txSender accounts.Account
	// gasLimit is the gas limit for all transactions to the Adjudicator.
	gasLimit uint64
}

// NewCoordinator creates a new ethereum coordinator. The receiver is the
// on-chain address that receives withdrawals.
func NewCoordinator(backend ContractBackend, contract common.Address, receiver common.Address, txSender accounts.Account, gasLimit uint64) *Coordinator {
	contr, err := adjudicator.NewAdjudicator(contract, backend)
	if err != nil {
		panic("Could not create a new instance of adjudicator")
	}
	bound := bind.NewBoundContract(contract, bindings.ABI.Adjudicator, backend, backend, backend)
	return &Coordinator{
		ContractBackend: backend,
		chainID:         backend.chainID,
		contract:        contr,
		bound:           bound,
		Receiver:        receiver,
		txSender:        txSender,
		log:             log.WithField("txSender", txSender.Address),
		gasLimit:        gasLimit,
	}
}

// Coordinate coordinates the state of the channel directly on the blockchain.
func (c *Coordinator) Coordinate(
	ctx context.Context,
	req channel.AdjudicatorReq,
	subChannels []channel.SignedState,
	coordSigs []wallet.Sig) error {
	// In the case of final states, we already call concludeFinal on the
	// adjudicator. Method ensureCoordinated calls concludeFinal for final states.
	if err := c.ensureCoordinated(ctx, req, subChannels, coordSigs); err != nil {
		return errors.WithMessage(err, "ensuring Coordinated")
	}

	return nil
}

func needCoordinate(ctx context.Context, req channel.AdjudicatorReq) bool {
	if !channel.IsNoApp(req.Params.App) && false /* !channel.HasCoordinator(req.Params)*/ {
		return true
	}
	return false
}

// call calls the given contract function `fn` with the data from `req`.
// `fn` should be a method of `a.contract`, like `a.contract.Register`.
// `txType` should be one of the valid transaction types defined in the client package.
func (c *Coordinator) call(ctx context.Context, req channel.AdjudicatorReq, fn adjFunc, txType OnChainTxType) error {
	ethParams := ToEthParams(req.Params)
	ethState := ToEthState(req.Tx.State)
	tx, err := func() (*types.Transaction, error) {
		if !c.mu.TryLockCtx(ctx) {
			return nil, errors.Wrap(ctx.Err(), "context canceled while acquiring tx lock")
		}
		defer c.mu.Unlock()

		trans, err := c.NewTransactor(ctx, c.gasLimit, c.txSender)
		if err != nil {
			return nil, errors.WithMessage(err, "creating transactor")
		}
		tx, err := fn(trans, ethParams, ethState, req.Tx.Sigs)
		if err != nil {
			err = cherrors.CheckIsChainNotReachableError(err)
			return nil, errors.WithMessage(err, "calling adjudicator function")
		}
		log.Debugf("Sent transaction %v", tx.Hash().Hex())
		return tx, nil
	}()
	if err != nil {
		return err
	}

	_, err = c.ConfirmTransaction(ctx, tx, c.txSender)
	if errors.Is(err, errTxTimedOut) {
		err = client.NewTxTimedoutError(txType.String(), tx.Hash().Hex(), err.Error())
	}
	return errors.WithMessage(err, "mining transaction")
}
