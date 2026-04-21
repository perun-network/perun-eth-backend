// Copyright 2026 - See NOTICE file for copyright holders.
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

package adapter

import (
	"context"
	"fmt"
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/perun-network/perun-eth-backend/bindings/liquiditypool"
	"github.com/perun-network/perun-eth-backend/channel"
)

type liquidityPoolContract interface {
	WithdrawableETH(opts *bind.CallOpts) (*big.Int, error)
	LockedByChannel(opts *bind.CallOpts, channelID [32]byte) (*big.Int, error)
	FundChannel(opts *bind.TransactOpts, channelID [32]byte, amount *big.Int) (*types.Transaction, error)
	SettleChannel(opts *bind.TransactOpts, channelID [32]byte) (*types.Transaction, error)
	Operator(opts *bind.CallOpts) (common.Address, error)
	TotalAssets(opts *bind.CallOpts) (*big.Int, error)
	TotalLockedETH(opts *bind.CallOpts) (*big.Int, error)
}

type adapterBackend interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	NewTransactor(ctx context.Context, gasLimit uint64, acc accounts.Account) (*bind.TransactOpts, error)
	TxFinalityDepth() uint64
}

type txOptsFactory func(ctx context.Context) (*bind.TransactOpts, error)
type txConfirmer func(ctx context.Context, tx *types.Transaction) (*types.Receipt, error)

// AdapterOption configures runtime behavior of the adapter.
type AdapterOption func(*adapterConfig)

type adapterConfig struct {
	retryInitial time.Duration
	retryMax     time.Duration
	finality     uint64
}

const (
	defaultRetryInitial = 250 * time.Millisecond
	defaultRetryMax     = 5 * time.Second
)

// WithRetryBackoff configures reconnect backoff bounds for event subscriptions.
func WithRetryBackoff(initial, max time.Duration) AdapterOption {
	return func(c *adapterConfig) {
		if initial > 0 {
			c.retryInitial = initial
		}
		if max > 0 {
			c.retryMax = max
		}
	}
}

// WithFinalityDepth overrides transaction finality depth for tx confirmation.
// A zero value keeps backend default finality depth.
func WithFinalityDepth(depth uint64) AdapterOption {
	return func(c *adapterConfig) {
		if depth > 0 {
			c.finality = depth
		}
	}
}

// LiquidityPoolAdapter is the ETH execution adapter between Hub/Websocket and LiquidityPool.
type LiquidityPoolAdapter struct {
	backend      adapterBackend
	contract     liquidityPoolContract
	poolAddress  common.Address
	newTxOpts    txOptsFactory
	confirmTx    txConfirmer
	retryInitial time.Duration
	retryMax     time.Duration
	finality     uint64
}

// NewLiquidityPoolAdapter creates a new adapter over a deployed LiquidityPool contract.
func NewLiquidityPoolAdapter(
	backend *channel.ContractBackend,
	poolAddress common.Address,
	txSender accounts.Account,
	gasLimit uint64,
	opts ...AdapterOption,
) (*LiquidityPoolAdapter, error) {
	contract, err := liquiditypool.NewLiquidityPool(poolAddress, backend)
	if err != nil {
		return nil, fmt.Errorf("binding LiquidityPool: %w", err)
	}

	cfg := adapterConfig{
		retryInitial: defaultRetryInitial,
		retryMax:     defaultRetryMax,
		finality:     backend.TxFinalityDepth(),
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.retryInitial <= 0 {
		cfg.retryInitial = defaultRetryInitial
	}
	if cfg.retryMax < cfg.retryInitial {
		cfg.retryMax = cfg.retryInitial
	}
	if cfg.finality < 1 {
		cfg.finality = 1
	}

	return &LiquidityPoolAdapter{
		backend:     backend,
		contract:    contract,
		poolAddress: poolAddress,
		newTxOpts: func(ctx context.Context) (*bind.TransactOpts, error) {
			return backend.NewTransactor(ctx, gasLimit, txSender)
		},
		confirmTx: func(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
			return confirmWithFinality(ctx, backend, tx, cfg.finality)
		},
		retryInitial: cfg.retryInitial,
		retryMax:     cfg.retryMax,
		finality:     cfg.finality,
	}, nil
}

func newTestAdapter(
	backend adapterBackend,
	contract liquidityPoolContract,
	newTxOpts txOptsFactory,
	confirm txConfirmer,
) *LiquidityPoolAdapter {
	return &LiquidityPoolAdapter{
		backend:      backend,
		contract:     contract,
		newTxOpts:    newTxOpts,
		confirmTx:    confirm,
		retryInitial: 5 * time.Millisecond,
		retryMax:     50 * time.Millisecond,
		finality:     1,
	}
}

func confirmWithFinality(ctx context.Context, backend adapterBackend, tx *types.Transaction, finalityDepth uint64) (*types.Receipt, error) {
	if finalityDepth < 1 {
		finalityDepth = 1
	}

	head, err := waitMinedHead(ctx, backend, tx)
	if err != nil {
		return nil, err
	}

	heads := make(chan *types.Header, 10)
	heads <- head
	hsub, err := backend.SubscribeNewHead(ctx, heads)
	if err != nil {
		return nil, err
	}
	defer hsub.Unsubscribe()

	for {
		select {
		case head := <-heads:
			receipt, err := backend.TransactionReceipt(ctx, tx.Hash())
			if err != nil || receipt == nil {
				continue
			}
			if isReceiptFinal(receipt, head, finalityDepth) {
				return receipt, nil
			}
		case err := <-hsub.Err():
			if err != nil {
				return nil, err
			}
			return nil, context.Canceled
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func waitMinedHead(ctx context.Context, backend adapterBackend, tx *types.Transaction) (*types.Header, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		receipt, err := backend.TransactionReceipt(ctx, tx.Hash())
		if err == nil && receipt != nil {
			return backend.HeaderByNumber(ctx, nil)
		}

		timer := time.NewTimer(100 * time.Millisecond)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, ctx.Err()
		case <-timer.C:
		}
	}
}

func isReceiptFinal(receipt *types.Receipt, head *types.Header, finalityDepth uint64) bool {
	depth := new(big.Int).SetUint64(finalityDepth)
	diff := new(big.Int).Sub(head.Number, receipt.BlockNumber)
	included := new(big.Int).Add(diff, big.NewInt(1))
	return included.Cmp(depth) >= 0
}

// FundChannel calls fundChannel(channelId, amount) on LiquidityPool.
func (a *LiquidityPoolAdapter) FundChannel(ctx context.Context, channelID [32]byte, amount *big.Int) error {
	free, err := a.contract.WithdrawableETH(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("%w: cannot read free liquidity: %v", ErrRetriable, err)
	}
	if amount.Cmp(free) > 0 {
		return fmt.Errorf("%w: amount %s > free liquidity %s", ErrDeterministic, amount, free)
	}

	existing, err := a.contract.LockedByChannel(&bind.CallOpts{Context: ctx}, channelID)
	if err != nil {
		return fmt.Errorf("%w: cannot read channel lock: %v", ErrRetriable, err)
	}
	if existing.Sign() > 0 {
		return fmt.Errorf("%w: channel already funded with %s", ErrDeterministic, existing)
	}

	opts, err := a.newTxOpts(ctx)
	if err != nil {
		return classifyEthError(err)
	}
	tx, err := a.contract.FundChannel(opts, channelID, amount)
	if err != nil {
		return classifyEthError(err)
	}
	receipt, err := a.confirmTx(ctx, tx)
	if err != nil {
		return fmt.Errorf("%w: tx wait failed: %v", ErrRetriable, err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("%w: fundChannel reverted tx=%s", ErrContractRevert, tx.Hash())
	}
	return nil
}

// SettleChannel calls settleChannel(channelId) payable on LiquidityPool.
func (a *LiquidityPoolAdapter) SettleChannel(ctx context.Context, channelID [32]byte, returnAmount *big.Int) error {
	locked, err := a.contract.LockedByChannel(&bind.CallOpts{Context: ctx}, channelID)
	if err != nil {
		return fmt.Errorf("%w: cannot read channel lock: %v", ErrRetriable, err)
	}
	if locked.Sign() == 0 {
		return fmt.Errorf("%w: channel %x not found in pool", ErrDeterministic, channelID)
	}
	if returnAmount.Cmp(locked) < 0 {
		return fmt.Errorf("%w: returnAmount %s < locked principal %s - contract would revert", ErrDeterministic, returnAmount, locked)
	}

	opts, err := a.newTxOpts(ctx)
	if err != nil {
		return classifyEthError(err)
	}
	opts.Value = returnAmount
	tx, err := a.contract.SettleChannel(opts, channelID)
	if err != nil {
		return classifyEthError(err)
	}
	receipt, err := a.confirmTx(ctx, tx)
	if err != nil {
		return fmt.Errorf("%w: tx wait failed: %v", ErrRetriable, err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("%w: settleChannel reverted tx=%s", ErrContractRevert, tx.Hash())
	}
	return nil
}

// GetOperator returns the delegated operator address from LiquidityPool.
func (a *LiquidityPoolAdapter) GetOperator(ctx context.Context) (common.Address, error) {
	op, err := a.contract.Operator(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, fmt.Errorf("%w: operator: %v", ErrRetriable, err)
	}
	return op, nil
}

// GetPoolState returns total assets and locked ETH values for Hub accounting.
func (a *LiquidityPoolAdapter) GetPoolState(ctx context.Context) (reserve uint64, locked uint64, err error) {
	total, err := a.contract.TotalAssets(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, 0, fmt.Errorf("%w: totalAssets: %v", ErrRetriable, err)
	}
	lockedBig, err := a.contract.TotalLockedETH(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, 0, fmt.Errorf("%w: totalLockedETH: %v", ErrRetriable, err)
	}
	return total.Uint64(), lockedBig.Uint64(), nil
}
