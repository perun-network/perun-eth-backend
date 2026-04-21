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
	"math/big"
	"sync"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/perun-network/perun-eth-backend/bindings/liquiditypool"
)

const eventStartBlockOffset uint64 = 100

type liquidityPoolEventParser interface {
	ParseChannelFunded(log types.Log) (*liquiditypool.LiquidityPoolChannelFunded, error)
	ParseChannelSettled(log types.Log) (*liquiditypool.LiquidityPoolChannelSettled, error)
	ParseOperatorUpdated(log types.Log) (*liquiditypool.LiquidityPoolOperatorUpdated, error)
}

type managedSubscription struct {
	errCh  chan error
	cancel context.CancelFunc
	once   sync.Once
}

func (s *managedSubscription) Err() <-chan error {
	return s.errCh
}

func (s *managedSubscription) Unsubscribe() {
	s.once.Do(s.cancel)
}

func (a *LiquidityPoolAdapter) parser() (liquidityPoolEventParser, bool) {
	p, ok := a.contract.(liquidityPoolEventParser)
	return p, ok
}

// SubscribeChannelFunded subscribes typed ChannelFunded events with reconnect/backoff.
func (a *LiquidityPoolAdapter) SubscribeChannelFunded(ctx context.Context, ch chan<- ChannelFundedEvent) (Subscription, error) {
	p, ok := a.parser()
	if !ok {
		return nil, classifyEthError(bind.ErrNoCode)
	}
	return a.subscribeLoop(ctx, func(log types.Log) error {
		e, err := p.ParseChannelFunded(log)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ch <- ChannelFundedEvent{
			ChannelID: e.ChannelId,
			Principal: e.Principal,
			Operator:  e.Operator,
			BlockNum:  e.Raw.BlockNumber,
		}:
			return nil
		}
	}, liquiditypool.LiquidityPoolChannelFundedTopic, func() { close(ch) })
}

// SubscribeChannelSettled subscribes typed ChannelSettled events with reconnect/backoff.
func (a *LiquidityPoolAdapter) SubscribeChannelSettled(ctx context.Context, ch chan<- ChannelSettledEvent) (Subscription, error) {
	p, ok := a.parser()
	if !ok {
		return nil, classifyEthError(bind.ErrNoCode)
	}
	return a.subscribeLoop(ctx, func(log types.Log) error {
		e, err := p.ParseChannelSettled(log)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ch <- ChannelSettledEvent{
			ChannelID:     e.ChannelId,
			Principal:     e.Principal,
			TotalReturned: e.TotalReturned,
			FeeGain:       e.FeeGain,
			BlockNum:      e.Raw.BlockNumber,
		}:
			return nil
		}
	}, liquiditypool.LiquidityPoolChannelSettledTopic, func() { close(ch) })
}

// SubscribeOperatorUpdated subscribes typed OperatorUpdated events with reconnect/backoff.
func (a *LiquidityPoolAdapter) SubscribeOperatorUpdated(ctx context.Context, ch chan<- OperatorUpdatedEvent) (Subscription, error) {
	p, ok := a.parser()
	if !ok {
		return nil, classifyEthError(bind.ErrNoCode)
	}
	return a.subscribeLoop(ctx, func(log types.Log) error {
		e, err := p.ParseOperatorUpdated(log)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ch <- OperatorUpdatedEvent{
			Previous: e.PreviousOperator,
			New:      e.NewOperator,
		}:
			return nil
		}
	}, liquiditypool.LiquidityPoolOperatorUpdatedTopic, func() { close(ch) })
}

func (a *LiquidityPoolAdapter) subscribeLoop(
	parent context.Context,
	onLog func(log types.Log) error,
	topic common.Hash,
	onClose func(),
) (Subscription, error) {
	if a.backend == nil {
		return nil, classifyEthError(bind.ErrNoCode)
	}

	ctx, cancel := context.WithCancel(parent)
	sub := &managedSubscription{errCh: make(chan error, 1), cancel: cancel}

	start, err := a.startBlock(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	go func() {
		defer close(sub.errCh)
		defer onClose()

		nextBlock := start
		delay := a.retryInitial
		maxDelay := a.retryMax
		if delay <= 0 {
			delay = 250 * time.Millisecond
		}
		if maxDelay < delay {
			maxDelay = 5 * time.Second
		}

		for {
			if ctx.Err() != nil {
				return
			}

			query := ethereum.FilterQuery{
				Addresses: []common.Address{a.poolAddress},
				Topics:    [][]common.Hash{{topic}},
				FromBlock: new(big.Int).SetUint64(nextBlock),
			}

			logs, err := a.backend.FilterLogs(ctx, query)
			if err != nil {
				if !sendErr(sub.errCh, classifyEthError(err)) {
					return
				}
				time.Sleep(delay)
				delay = minDuration(delay*2, maxDelay)
				continue
			}
			for _, lg := range logs {
				if lg.Removed {
					continue
				}
				if err := onLog(lg); err != nil {
					if ctx.Err() != nil {
						return
					}
					if !sendErr(sub.errCh, classifyEthError(err)) {
						return
					}
				}
				if lg.BlockNumber >= nextBlock {
					nextBlock = lg.BlockNumber + 1
				}
			}

			live := make(chan types.Log, 32)
			liveSub, err := a.backend.SubscribeFilterLogs(ctx, query, live)
			if err != nil {
				if !sendErr(sub.errCh, classifyEthError(err)) {
					return
				}
				time.Sleep(delay)
				delay = minDuration(delay*2, maxDelay)
				continue
			}
			delay = a.retryInitial
			if delay <= 0 {
				delay = 250 * time.Millisecond
			}

			reconnect := false
			for !reconnect {
				select {
				case <-ctx.Done():
					liveSub.Unsubscribe()
					return
				case err := <-liveSub.Err():
					liveSub.Unsubscribe()
					if ctx.Err() != nil {
						return
					}
					if err != nil {
						if !sendErr(sub.errCh, classifyEthError(err)) {
							return
						}
					}
					reconnect = true
				case lg := <-live:
					if lg.Removed {
						continue
					}
					if err := onLog(lg); err != nil {
						if ctx.Err() != nil {
							liveSub.Unsubscribe()
							return
						}
						if !sendErr(sub.errCh, classifyEthError(err)) {
							liveSub.Unsubscribe()
							return
						}
					}
					if lg.BlockNumber >= nextBlock {
						nextBlock = lg.BlockNumber + 1
					}
				}
			}
			time.Sleep(delay)
			delay = minDuration(delay*2, maxDelay)
		}
	}()

	return sub, nil
}

func (a *LiquidityPoolAdapter) startBlock(ctx context.Context) (uint64, error) {
	head, err := a.backend.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, classifyEthError(err)
	}
	h := head.Number.Uint64()
	if h <= eventStartBlockOffset {
		return 1, nil
	}
	return h - eventStartBlockOffset, nil
}

func minDuration(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}

func sendErr(ch chan<- error, err error) bool {
	select {
	case ch <- err:
		return true
	default:
		return true
	}
}
