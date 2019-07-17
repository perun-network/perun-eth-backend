// Copyright (c) 2019 The Perun Authors. All rights reserved.
// This file is part of go-perun. Use of this source code is governed by a
// MIT-style license that can be found in the LICENSE file.

package wallet

import (
	"strconv"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	perun "perun.network/go-perun/wallet"
)

// Backend implements the utility interface defined in the wallet package.
type Backend struct{}

// NewAddressFromString creates a new address from a string.
func (h *Backend) NewAddressFromString(s string) (perun.Address, error) {
	addr, err := common.NewMixedcaseAddressFromString(s)
	if err != nil {
		zeroAddr := common.BytesToAddress(make([]byte, 20, 20))
		return &Address{zeroAddr}, err
	}
	return &Address{addr.Address()}, nil
}

// NewAddressFromBytes creates a new address from a byte array.
func (h *Backend) NewAddressFromBytes(data []byte) (perun.Address, error) {
	if len(data) != 20 {
		errString := "could not create address from bytes of length: " + strconv.Itoa(len(data))
		return &Address{ZeroAddr}, errors.New(errString)
	}
	return &Address{common.BytesToAddress(data)}, nil
}

// VerifySignature verifies if a signature was made by this account.
func (h *Backend) VerifySignature(msg, sig []byte, a perun.Address) (bool, error) {
	hash := crypto.Keccak256(msg)
	pk, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return false, err
	}
	addr := crypto.PubkeyToAddress(*pk)
	return a.Equals(&Address{addr}), nil
}
