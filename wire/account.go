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

// Package wire contains the wire account implementation for the Ethereum Backend.
package wire

import (
	"math/rand"

	"perun.network/go-perun/wire"
)

// Account is a wire account.
type Account struct {
	addr *Address
}

// Address returns the account's address.
func (acc *Account) Address() wire.Address {
	return acc.addr
}

// NewRandomAccount generates a new random account.
func NewRandomAccount(rng *rand.Rand) *Account {
	return &Account{
		addr: NewRandomAddress(rng),
	}
}
