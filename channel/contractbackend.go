// Copyright 2025 - See NOTICE file for copyright holders.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package channel

import (
	"context"
	"math/big"
	"sync"

	"perun.network/go-perun/channel/multi"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"

	"perun.network/go-perun/log"
	pcontext "polycry.pt/poly-go/context"
)

const (
	// How many blocks we query into the past for events.
	startBlockOffset            = 100
	contractBackendHeadBuffSize = 10
)

// errTxTimedOut is an internal named error that with an empty message.
// Because calling function is expected to check for this error and
// create a TxTimedoutError with additional context.
var errTxTimedOut = errors.New("")

var (
	// SharedExpectedNonces is a map of each expected next nonce of all clients.
	SharedExpectedNonces map[multi.LedgerIDMapKey]map[common.Address]uint64
	// SharedExpectedNoncesMutex is a mutex to protect the shared expected nonces map.
	SharedExpectedNoncesMutex = &sync.Mutex{}
)

// ContractInterface provides all functions needed by an ethereum backend.
// Both test.SimulatedBackend and ethclient.Client implement this interface.
type ContractInterface interface {
	bind.ContractBackend
	ethereum.ChainReader
	ethereum.TransactionReader
}

// Transactor can be used to make transactOpts for a given account.
type Transactor interface {
	NewTransactor(account accounts.Account) (*bind.TransactOpts, error)
}

// ContractBackend adds a keystore and an on-chain account to the ContractInterface.
// This is needed to send on-chain transaction to interact with the smart contracts.
type ContractBackend struct {
	ContractInterface
	tr                Transactor
	expectedNextNonce map[common.Address]uint64
	txFinalityDepth   uint64
	chainID           ChainID
}

// NewContractBackend creates a new ContractBackend with the given parameters.
// txFinalityDepth defines in how many consecutive blocks a TX has to be
// included to be considered final. Must be at least 1.
func NewContractBackend(cf ContractInterface, chainID ChainID, tr Transactor, txFinalityDepth uint64) ContractBackend {
	SharedExpectedNoncesMutex.Lock()
	defer SharedExpectedNoncesMutex.Unlock()
	// Check if the shared maps are initialized, if not, initialize them.
	if SharedExpectedNonces == nil {
		SharedExpectedNonces = make(map[multi.LedgerIDMapKey]map[common.Address]uint64)
	}

	// Check if the specific chainID entry exists in the shared maps, if not, create it.
	if _, exists := SharedExpectedNonces[chainID.MapKey()]; !exists {
		SharedExpectedNonces[chainID.MapKey()] = make(map[common.Address]uint64)
	}
	return ContractBackend{
		ContractInterface: cf,
		tr:                tr,
		expectedNextNonce: SharedExpectedNonces[chainID.MapKey()],
		txFinalityDepth:   txFinalityDepth,
		chainID:           chainID,
	}
}

// ChainID returns the chain identifier of the contract backend.
func (c *ContractBackend) ChainID() ChainID {
	return c.chainID
}

// NewWatchOpts returns bind.WatchOpts with the field Start set to the current
// block number and the ctx field set to the passed context.
func (c *ContractBackend) NewWatchOpts(ctx context.Context) (*bind.WatchOpts, error) {
	blockNum, err := c.pastOffsetBlockNum(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "new watch opts")
	}

	return &bind.WatchOpts{
		Start:   &blockNum,
		Context: ctx,
	}, nil
}

// NewFilterOpts returns bind.FilterOpts with the field Start set to the block
// number 100 blocks ago (or 1) and the field End set to nil and the ctx field
// set to the passed context.
func (c *ContractBackend) NewFilterOpts(ctx context.Context) (*bind.FilterOpts, error) {
	blockNum, err := c.pastOffsetBlockNum(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "new filter opts")
	}
	return &bind.FilterOpts{
		Start:   blockNum,
		End:     nil,
		Context: ctx,
	}, nil
}

func (c *ContractBackend) pastOffsetBlockNum(ctx context.Context) (uint64, error) {
	h, err := c.HeaderByNumber(ctx, nil)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return uint64(0), errors.WithMessage(err, "retrieving latest block")
	}

	// max(1, latestBlock - offset)
	if h.Number.Uint64() <= startBlockOffset {
		return 1, nil
	}
	return h.Number.Uint64() - startBlockOffset, nil
}

// NewTransactor returns bind.TransactOpts with the context, gas limit and
// account set as specified, using the ContractBackend's Transactor.
//
// The gas price and nonce are not set and will be set by go-ethereum
// automatically when not manually specified by the caller. The caller must also
// set the value manually afterwards if it should be different from 0.
func (c *ContractBackend) NewTransactor(ctx context.Context, gasLimit uint64, acc accounts.Account) (*bind.TransactOpts, error) {
	auth, err := c.tr.NewTransactor(acc)
	if err != nil {
		return nil, errors.WithMessage(err, "creating transactor")
	}

	// Set context and gas limit.
	auth.GasLimit = gasLimit
	auth.Context = ctx

	// Set and store nonce.
	nonce, err := c.nonce(ctx, auth.From)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	return auth, nil
}

// nonce tries to determine the correct nonce by comparing local and chain nonce
// expectations.
func (c *ContractBackend) nonce(ctx context.Context, sender common.Address) (uint64, error) {
	// Look up pending nonce from backend.
	nonce, err := c.PendingNonceAt(ctx, sender)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return 0, errors.WithMessage(err, "fetching nonce")
	}
	SharedExpectedNoncesMutex.Lock()
	defer SharedExpectedNoncesMutex.Unlock()
	expectedNextNonce, found := c.expectedNextNonce[sender]
	if !found {
		c.expectedNextNonce[sender] = 0
	}

	// Compare nonces and use larger.
	if nonce < expectedNextNonce {
		nonce = expectedNextNonce
	}

	// Update local expectation.
	c.expectedNextNonce[sender] = nonce + 1
	return nonce, nil
}

// ConfirmTransaction returns the receipt of the transaction if it was
// included in at least `TxFinalityDepth` many blocks at one point in time.
// Returns `txTimedOutError` on context timeout or cancel.
func (c *ContractBackend) ConfirmTransaction(ctx context.Context, tx *types.Transaction, acc accounts.Account) (*types.Receipt, error) {
	receipt, err := c.confirmNTimes(ctx, tx, c.txFinalityDepth)
	if err != nil {
		if pcontext.IsContextError(err) {
			err = errTxTimedOut
		}
		return nil, errors.WithMessage(err, "sending transaction")
	}

	if receipt.Status == types.ReceiptStatusFailed {
		reason, err := errorReason(ctx, c, tx, receipt.BlockNumber, acc)
		if err != nil {
			log.Error("TX failed; error determining reason: ", err)
			// There is no way in ethereum to really decide this, but since we
			// do it in the error case only, it should be fine.
			// The limit of 1000 was determined by trial-and-error.
			if receipt.GasUsed+1000 > tx.Gas() {
				log.WithFields(log.Fields{"Used": receipt.GasUsed, "Limit": tx.Gas()}).Warn("TX could be out of gas")
			}
		} else {
			log.Warn("TX failed with reason: ", reason)
		}
		return receipt, errors.WithStack(ErrTxFailed)
	}
	return receipt, nil
}

// confirmNTimes waits for a transaction to be included in `finalityDepth`
// many consecutive blocks. `finalityDepth` must be at least one.
func (c *ContractBackend) confirmNTimes(ctx context.Context, tx *types.Transaction, finalityDepth uint64) (*types.Receipt, error) {
	if finalityDepth < 1 {
		return nil, errors.New("finalityDepth was less than 1")
	}
	// Wait to be included at least once.
	head, err := c.waitMined(ctx, tx)
	if err != nil {
		return nil, errors.WithMessage(err, "waiting for TX to be mined")
	}

	// Set up header sub for future blocks.
	heads := make(chan *types.Header, contractBackendHeadBuffSize)
	heads <- head // Include the current head.
	hsub, err := c.SubscribeNewHead(ctx, heads)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return nil, errors.WithMessage(err, "subscribing to heads")
	}
	defer hsub.Unsubscribe()

	for {
		select {
		case head := <-heads:
			// Poll the receipt of the TX.
			receipt, err := c.ContractInterface.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				err = cherrors.CheckIsChainNotReachableError(err)
				log.Warnf("Failed to get tx receipt: %v", err)
				break
			}
			if receipt != nil && isFinal(receipt, head, finalityDepth) {
				return receipt, nil
			}
			// TX is either not included in the canonical chain anymore
			// or not yet final; wait for next head.
		case err := <-hsub.Err():
			err = cherrors.CheckIsChainNotReachableError(err)
			return nil, errors.WithMessage(err, "header subscription")
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

// waitMined waits for a TX to be mined and returns the latest head.
func (c *ContractBackend) waitMined(ctx context.Context, tx *types.Transaction) (*types.Header, error) {
	_, err := bind.WaitMined(ctx, c, tx)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return nil, errors.WithMessage(err, "waiting for mined")
	}
	head, err := c.HeaderByNumber(ctx, nil)
	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return nil, errors.WithMessage(err, "subscribing to heads")
	}
	return head, nil
}

// TxFinalityDepth returns the transaction finality depth of the contract backend.
func (c *ContractBackend) TxFinalityDepth() uint64 {
	return c.txFinalityDepth
}

// Returns ((head.number - receipt.number) + 1) >= finalityDepth.
func isFinal(receipt *types.Receipt, head *types.Header, _finalityDepth uint64) bool {
	finalityDepth := new(big.Int)
	finalityDepth.SetUint64(_finalityDepth)

	diff := new(big.Int).Sub(head.Number, receipt.BlockNumber)
	included := new(big.Int).Add(diff, big.NewInt(1))
	return included.Cmp(finalityDepth) >= 0
}
