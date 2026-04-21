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
	"errors"
	"fmt"
	"strings"

	ch "github.com/perun-network/perun-eth-backend/channel"
	cherrors "github.com/perun-network/perun-eth-backend/channel/errors"
)

var (
	// ErrRetriable is a transient transport/RPC failure.
	ErrRetriable = errors.New("eth adapter: retriable error")
	// ErrDeterministic is a permanent input/state mismatch.
	ErrDeterministic = errors.New("eth adapter: deterministic error")
	// ErrContractRevert is an on-chain revert condition.
	ErrContractRevert = errors.New("eth adapter: contract revert")
)

func classifyEthError(err error) error {
	if err == nil {
		return nil
	}
	if isRPCTimeout(err) || isConnectionRefused(err) || cherrors.IsChainNotReachableError(err) {
		return fmt.Errorf("%w: %v", ErrRetriable, err)
	}
	if isRevertError(err) || ch.IsErrTxFailed(err) {
		return fmt.Errorf("%w: %v", ErrContractRevert, err)
	}
	return fmt.Errorf("%w: %v", ErrDeterministic, err)
}

func isRPCTimeout(err error) bool {
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		return true
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "timeout") || strings.Contains(msg, "deadline exceeded")
}

func isConnectionRefused(err error) bool {
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "connection refused") || strings.Contains(msg, "connection reset") || strings.Contains(msg, "eof")
}

func isRevertError(err error) bool {
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "revert") || strings.Contains(msg, "execution reverted")
}
