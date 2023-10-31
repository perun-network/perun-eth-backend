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

// keep track ob the increse allowance and deposit processes
var mu sync.Mutex
var locks = make(map[string]*sync.Mutex)

// create key from account address and asset to only lock the process when the hub is deposits the same asset at the same time
func lockKey(account common.Address, asset common.Address) string {
	return fmt.Sprintf("%s-%s", account.Hex(), asset.Hex())
}

// NewERC20Depositor creates a new ERC20Depositor.
func NewERC20Depositor(token common.Address) *ERC20Depositor {
	return &ERC20Depositor{Token: token}
}

// Deposit deposits ERC20 tokens into the ERC20 AssetHolder specified at the
// request's asset address.
func (d *ERC20Depositor) Deposit(ctx context.Context, req DepositReq) (types.Transactions, error) {
	lockKey := lockKey(req.Account.Address, req.Asset.EthAddress())
	// Lock the mutex associated with this key.
	mu.Lock()
	if _, exists := locks[lockKey]; !exists {
		locks[lockKey] = &sync.Mutex{}
	}
	mu.Unlock()

	// Bind a `AssetHolderERC20` instance.
	assetholder, err := assetholdererc20.NewAssetholdererc20(req.Asset.EthAddress(), req.CB)
	if err != nil {
		return nil, errors.Wrapf(err, "binding AssetHolderERC20 contract at: %x", req.Asset)
	}
	// Bind an `ERC20` instance.
	token, err := peruntoken.NewPeruntoken(d.Token, req.CB)
	if err != nil {
		return nil, errors.Wrapf(err, "binding ERC20 contract at: %x", d.Token)
	}
	//get Allowance
	callOpts := bind.CallOpts{
		Pending: false,
		Context: ctx,
	}
	// Lock the specific mutex for this combination.
	locks[lockKey].Lock()

	allowance, err := token.Allowance(&callOpts, req.Account.Address, req.Asset.EthAddress())

	if err != nil {
		return nil, errors.WithMessagef(err, "could not get Allowance for asset: %x", req.Asset)
	}
	result := new(big.Int).Add(req.Balance, allowance)

	// Increase the allowance.
	opts, err := req.CB.NewTransactor(ctx, ERC20DepositorTXGasLimit, req.Account)
	if err != nil {
		return nil, errors.WithMessagef(err, "creating transactor for asset: %x", req.Asset)
	}
	// Create a channel for receiving PeruntokenApproval events
	eventSink := make(chan *peruntoken.PeruntokenApproval)

	// Create a channel for receiving the Approval event
	eventReceived := make(chan bool)

	// Watch for Approval events and send them to the eventSink
	subscription, err := token.WatchApproval(&bind.WatchOpts{Start: nil, Context: ctx}, eventSink, []common.Address{req.Account.Address}, []common.Address{req.Asset.EthAddress()})
	if err != nil {
		return nil, errors.WithMessagef(err, "Cannot listen for event")
	}
	// tx0, err := token.IncreaseAllowance(opts, req.Asset.EthAddress(), req.Balance)
	fmt.Println("Approve: ", result)
	tx1, err := token.Approve(opts, req.Asset.EthAddress(), result)

	if err != nil {
		err = cherrors.CheckIsChainNotReachableError(err)
		return nil, errors.WithMessagef(err, "increasing allowance for asset: %x", req.Asset)
	}

	go func() {
		select {
		case event := <-eventSink:
			// Handle the event
			log.Printf("Received Approval event: Owner: %s, Spender: %s, Value: %s\n", event.Owner.Hex(), event.Spender.Hex(), event.Value.String())
			// Set the flag to true to indicate event received
			eventReceived <- true
		case err := <-subscription.Err():
			// Handle subscription error
			log.Println("Subscription error:", err)
		}
	}()

	// Wait for the Approval event to be received
	approvalReceived := <-eventReceived
	defer locks[lockKey].Unlock()
	if approvalReceived {
		_, err1 := token.Allowance(&callOpts, req.Account.Address, req.Asset.EthAddress())

		if err1 != nil {
			return nil, errors.WithMessagef(err, "could not get Allowance for asset: %x", req.Asset)
		}
		// Deposit.
		opts, err := req.CB.NewTransactor(ctx, ERC20DepositorTXGasLimit, req.Account)
		if err != nil {
			return nil, errors.WithMessagef(err, "creating transactor for asset: %x", req.Asset)
		}
		tx2, err := assetholder.Deposit(opts, req.FundingID, req.Balance)
		err = cherrors.CheckIsChainNotReachableError(err)
		return []*types.Transaction{tx1, tx2}, errors.WithMessage(err, "AssetHolderERC20 depositing")
	} else {
		return nil, errors.WithMessagef(err, "could not approve")
	}
}

// NumTX returns 2 since it does IncreaseAllowance and Deposit.
func (*ERC20Depositor) NumTX() uint32 {
	return erc20DepositorNumTx
}
