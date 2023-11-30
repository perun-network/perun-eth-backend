// Copyright 2020 - See NOTICE file for copyright holders.
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
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"perun.network/go-perun/log"

	"github.com/perun-network/perun-eth-backend/bindings/assetholdererc20"
	"github.com/perun-network/perun-eth-backend/bindings/peruntoken"
	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"
)

// ERC20Depositor deposits tokens into the `AssetHolderERC20` contract.
// It is bound to a token but can be reused to deposit multiple times.
type ERC20Depositor struct {
	Token common.Address
}

// ERC20DepositorTXGasLimit is the limit of Gas that an `ERC20Depositor` will
// spend per transaction when depositing funds.
// An `IncreaseAllowance` uses ~45kGas and a `Deposit` call ~84kGas on average.
const ERC20DepositorTXGasLimit = 100000

// Return value of ERC20Depositor.NumTx.
const erc20DepositorNumTx = 2

// Keep track of the increase allowance and deposit processes.
var mu sync.Mutex
var locks = make(map[string]*sync.Mutex)

// DepositResult is created to keep track of the returned values.
type DepositResult struct {
	Transactions types.Transactions
	Error        error
}

// Create key from account address and asset to only lock the process when hub deposits the same asset at the same time.
func lockKey(account common.Address, asset common.Address) string {
	return fmt.Sprintf("%s-%s", account.Hex(), asset.Hex())
}

// Retrieves Lock for specific key.
func handleLock(lockKey string) *sync.Mutex {
	mu.Lock()
	defer mu.Unlock()

	if lock, exists := locks[lockKey]; exists {
		return lock
	}

	lock := &sync.Mutex{}
	locks[lockKey] = lock
	return lock
}

// Locks the lock argument, runs the given function and then unlocks the lock argument.
func lockAndUnlock(lock *sync.Mutex, fn func()) {
	mu.Lock()
	defer mu.Unlock()
	lock.Lock()
	defer lock.Unlock()
	fn()
}

// NewERC20Depositor creates a new ERC20Depositor.
func NewERC20Depositor(token common.Address) *ERC20Depositor {
	return &ERC20Depositor{Token: token}
}

// Deposit approves the value to be swapped and calls DepositOnly.
// nolint:funlen
func (d *ERC20Depositor) Deposit(ctx context.Context, req DepositReq) (types.Transactions, error) {
	lockKey := lockKey(req.Account.Address, req.Asset.EthAddress())
	lock := handleLock(lockKey)

	// Bind an `ERC20` instance.
	token, err := peruntoken.NewPeruntoken(d.Token, req.CB)
	if err != nil {
		return nil, errors.Wrapf(err, "binding ERC20 contract at: %x", d.Token)
	}
	callOpts := bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	// variables for the return value.
	var depResult DepositResult
	var approvalReceived bool
	var tx1 *types.Transaction
	var err1 error
	lockAndUnlock(lock, func() {
		allowance, err := token.Allowance(&callOpts, req.Account.Address, req.Asset.EthAddress())
		if err != nil {
			depResult.Transactions = nil
			depResult.Error = errors.WithMessagef(err, "could not get Allowance for asset: %x", req.Asset)
		}
		result := new(big.Int).Add(req.Balance, allowance)

		// Increase the allowance.
		opts, err := req.CB.NewTransactor(ctx, ERC20DepositorTXGasLimit, req.Account)
		if err != nil {
			depResult.Transactions = nil
			depResult.Error = errors.WithMessagef(err, "creating transactor for asset: %x", req.Asset)
		}
		// Create a channel for receiving PeruntokenApproval events
		eventSink := make(chan *peruntoken.PeruntokenApproval)

		// Create a channel for receiving the Approval event
		eventReceived := make(chan bool)

		// Watch for Approval events and send them to the eventSink
		subscription, err := token.WatchApproval(&bind.WatchOpts{Start: nil, Context: ctx}, eventSink, []common.Address{req.Account.Address}, []common.Address{req.Asset.EthAddress()})
		if err != nil {
			depResult.Transactions = nil
			depResult.Error = errors.WithMessagef(err, "Cannot listen for event")
		}
		tx1, err1 = token.Approve(opts, req.Asset.EthAddress(), result)
		if err1 != nil {
			err = cherrors.CheckIsChainNotReachableError(err)
			depResult.Transactions = nil
			depResult.Error = errors.WithMessagef(err, "increasing allowance for asset: %x", req.Asset)
		}

		go func() {
			select {
			case event := <-eventSink:
				log.Printf("Received Approval event: Owner: %s, Spender: %s, Value: %s\n", event.Owner.Hex(), event.Spender.Hex(), event.Value.String())
				eventReceived <- true
			case err := <-subscription.Err():
				log.Println("Subscription error:", err)
			}
		}()
		approvalReceived = <-eventReceived
	})
	if approvalReceived {
		tx2, err := d.DepositOnly(ctx, req)
		depResult.Transactions = []*types.Transaction{tx1, tx2}
		depResult.Error = errors.WithMessage(err, "AssetHolderERC20 depositing")
	}
	return depResult.Transactions, depResult.Error
}

// DepositOnly deposits ERC20 tokens into the ERC20 AssetHolder specified at the
// requests asset address.
func (d *ERC20Depositor) DepositOnly(ctx context.Context, req DepositReq) (*types.Transaction, error) {
	lockKey := lockKey(req.Account.Address, req.Asset.EthAddress())
	lock := handleLock(lockKey)

	var result *types.Transaction
	var resError error

	lockAndUnlock(lock, func() {
		// Bind a `AssetHolderERC20` instance.
		assetholder, err := assetholdererc20.NewAssetholdererc20(req.Asset.EthAddress(), req.CB)
		if err != nil {
			resError = errors.Wrapf(err, "binding AssetHolderERC20 contract at: %x", req.Asset)
			result = nil
		}
		// Deposit.
		opts, err := req.CB.NewTransactor(ctx, ERC20DepositorTXGasLimit, req.Account)
		if err != nil {
			resError = errors.WithMessagef(err, "creating transactor for asset: %x", req.Asset)
			result = nil
		}

		tx2, err := assetholder.Deposit(opts, req.FundingID, req.Balance)
		err = cherrors.CheckIsChainNotReachableError(err)
		result = tx2
		resError = errors.WithMessage(err, "AssetHolderERC20 depositing")
	})
	return result, resError
}

// NumTX returns 2 since it does IncreaseAllowance and Deposit.
func (*ERC20Depositor) NumTX() uint32 {
	return erc20DepositorNumTx
}
