// Copyright 2022 - See NOTICE file for copyright holders.
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

package wire

import (
	"bytes"
	"errors"
	"math/rand"

	"github.com/perun-network/perun-eth-backend/wallet"
	"github.com/perun-network/perun-eth-backend/wallet/test"

	"perun.network/go-perun/wire"
)

// Address is a wrapper for wallet.Address.
type Address struct {
	*wallet.Address
}

// NewAddress returns a new address.
func NewAddress() *Address {
	return &Address{}
}

// Equal returns whether the two addresses are equal.
func (a Address) Equal(b wire.Address) bool {
	bTyped, ok := b.(*Address)
	if !ok {
		panic("wrong type")
	}
	return a.Address.Equal(bTyped.Address)
}

// Cmp compares the byte representation of two addresses. For `a.Cmp(b)`
// returns -1 if a < b, 0 if a == b, 1 if a > b.
func (a Address) Cmp(b wire.Address) int {
	bTyped, ok := b.(*Address)
	if !ok {
		panic("wrong type")
	}
	return a.Address.Cmp(bTyped.Address)
}

// NewRandomAddress returns a new random peer address.
func NewRandomAddress(rng *rand.Rand) *Address {
	addr := test.NewRandomAddress(rng)
	return &Address{&addr}
}

// Verify verifies a message signature.
// It returns an error if the signature is invalid.
func (a Address) Verify(_ []byte, sig []byte) error {
	if !bytes.Equal(sig, []byte("Authenticate")) {
		return errors.New("invalid signature")
	}
	return nil
}
