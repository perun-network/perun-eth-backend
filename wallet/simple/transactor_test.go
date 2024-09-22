// Copyright 2021 - See NOTICE file for copyright holders.
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

package simple_test

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/perun-network/perun-eth-backend/channel/test"
	"github.com/perun-network/perun-eth-backend/wallet"
	"github.com/perun-network/perun-eth-backend/wallet/simple"
	pkgtest "polycry.pt/poly-go/test"
)

// Missing address for which key will not be contained in the wallet.
const missingAddr = "0x1"

func TestTransactor(t *testing.T) {
	rng := pkgtest.Prng(t)
	chainID := rng.Int63()

	tests := []struct {
		title   string
		signer  types.Signer
		chainID int64
		txType  test.TxType
	}{
		{
			title:  "FrontierSigner",
			signer: &types.FrontierSigner{},
			txType: test.LegacyTx,
		},
		{
			title:  "HomesteadSigner",
			signer: &types.HomesteadSigner{},
			txType: test.LegacyTx,
		},
		{
			title:   "EIP155Signer",
			signer:  types.NewEIP155Signer(big.NewInt(chainID)),
			txType:  test.EIP155Tx,
			chainID: chainID,
		},
		{
			title:   "LatestSigner",
			signer:  types.LatestSignerForChainID(big.NewInt(chainID)),
			txType:  test.EIP1559Tx,
			chainID: chainID,
		},
	}

	for _, _t := range tests {
		_t := _t
		t.Run(_t.title, func(t *testing.T) {
			s := newTransactorSetup(t, rng, _t.signer, _t.chainID, _t.txType)
			test.GenericSignerTest(t, rng, s)
		})
	}
}

func newTransactorSetup(t require.TestingT, prng *rand.Rand, signer types.Signer, chainID int64, txType test.TxType) test.TransactorSetup {
	simpleWallet := simple.NewWallet()
	require.NotNil(t, simpleWallet)
	validAcc := simpleWallet.NewRandomAccount(prng)
	require.NotNil(t, validAcc)

	return test.TransactorSetup{
		Signer:     signer,
		ChainID:    chainID,
		TxType:     txType,
		Tr:         simple.NewTransactor(simpleWallet, signer),
		ValidAcc:   accounts.Account{Address: wallet.AsEthAddr(validAcc.Address())},
		MissingAcc: accounts.Account{Address: common.HexToAddress(missingAddr)},
	}
}
