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
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"perun.network/go-perun/wallet"

	"github.com/perun-network/perun-eth-backend/bindings/adjudicator"
	ethwallet "github.com/perun-network/perun-eth-backend/wallet"
)

// A cross-chain channel carries assets of other backends (e.g. CKB) whose
// holder is an opaque byte string, not an Ethereum asset encoding. Decoding
// them as *Asset panicked when such a state came back from a contract event.
func Test_fromEthAssets_crossChain(t *testing.T) {
	const otherBackend wallet.BackendID = 2
	ethHolder := common.HexToAddress("0x3f5ea58ff28b18f1547c2b624e88ef9a761c0442")
	ccHolder := []byte{0x49, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0xde, 0xad, 0xbe, 0xef}

	assets := []adjudicator.ChannelAsset{
		{ChainID: big.NewInt(11155111), EthHolder: ethHolder},
		{ChainID: big.NewInt(2), CcHolder: ccHolder},
	}
	bIDs := []wallet.BackendID{ethwallet.BackendID, otherBackend}

	var decoded []interface{ Address() []byte }
	require.NotPanics(t, func() {
		for _, a := range fromEthAssets(assets, bIDs) {
			decoded = append(decoded, a.(interface{ Address() []byte }))
		}
	})
	require.Len(t, decoded, 2)

	ethAsset, ok := decoded[0].(*Asset)
	require.True(t, ok, "Ethereum asset must decode as *Asset")
	require.Equal(t, ethHolder, ethAsset.EthAddress())

	ccAsset, ok := decoded[1].(*CCAsset)
	require.True(t, ok, "foreign asset must decode as *CCAsset")
	require.Equal(t, ccHolder, ccAsset.Address())
	require.True(t, ccAsset.Equal(NewCCAsset(big.NewInt(2), ccHolder)))
	require.False(t, ccAsset.Equal(NewCCAsset(big.NewInt(3), ccHolder)), "different ledger")
}
