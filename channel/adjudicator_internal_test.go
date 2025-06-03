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
	"math/rand"
	"testing"

	"github.com/perun-network/perun-eth-backend/wallet"

	"github.com/perun-network/perun-eth-backend/bindings/adjudicator"
	"github.com/stretchr/testify/assert"
	"perun.network/go-perun/channel"
	channeltest "perun.network/go-perun/channel/test"
	pkgtest "polycry.pt/poly-go/test"
)

func Test_toEthSubStates(t *testing.T) {
	var (
		rng    = pkgtest.Prng(t)
		assert = assert.New(t)
	)

	tests := []struct {
		title string
		setup func() (state *channel.State, subStates channel.StateMap, expected []adjudicator.ChannelState)
	}{
		{
			title: "nil map gives nil slice",
			setup: func() (state *channel.State, subStates channel.StateMap, expected []adjudicator.ChannelState) {
				return channeltest.NewRandomState(rng, channeltest.WithBackend(wallet.BackendID)), nil, nil
			},
		},
		{
			title: "fresh map gives nil slice",
			setup: func() (state *channel.State, subStates channel.StateMap, expected []adjudicator.ChannelState) {
				return channeltest.NewRandomState(rng, channeltest.WithBackend(wallet.BackendID)), nil, nil
			},
		},
		{
			title: "1 layer of sub-channels",
			setup: func() (state *channel.State, subStates channel.StateMap, expected []adjudicator.ChannelState) {
				// ch[0]( ch[1], ch[2], ch[3] )
				ch := genStates(rng, 4)
				ch[0].State.AddSubAlloc(*ch[1].State.ToSubAlloc())
				ch[0].State.AddSubAlloc(*ch[2].State.ToSubAlloc())
				ch[0].State.AddSubAlloc(*ch[3].State.ToSubAlloc())
				var ethStates []adjudicator.ChannelState
				for i := 1; i < len(ch); i++ {
					ethStates = append(ethStates, ToEthState(ch[i].State))
				}
				return ch[0].State, toStateMap(ch[1:]...), ethStates
			},
		},
		{
			title: "2 layers of sub-channels",
			setup: func() (state *channel.State, subStates channel.StateMap, expected []adjudicator.ChannelState) {
				// ch[0]( ch[1]( ch[2], ch[3] ), ch[4], ch[5] (ch[6] ) )
				ch := genStates(rng, 7)
				ch[0].State.AddSubAlloc(*ch[1].State.ToSubAlloc())
				ch[0].State.AddSubAlloc(*ch[4].State.ToSubAlloc())
				ch[0].State.AddSubAlloc(*ch[5].State.ToSubAlloc())
				ch[1].State.AddSubAlloc(*ch[2].State.ToSubAlloc())
				ch[1].State.AddSubAlloc(*ch[3].State.ToSubAlloc())
				ch[5].State.AddSubAlloc(*ch[6].State.ToSubAlloc())
				var ethStates []adjudicator.ChannelState
				for i := 1; i < len(ch); i++ {
					ethStates = append(ethStates, ToEthState(ch[i].State))
				}
				return ch[0].State, toStateMap(ch[1:]...), ethStates
			},
		},
	}

	for _, tc := range tests {
		state, subStates, expected := tc.setup()
		got := toEthSubStates(state, subStates)
		assert.Equal(expected, got, tc.title)
	}
}

func genStates(rng *rand.Rand, n int) (states []*channel.SignedState) {
	states = make([]*channel.SignedState, n)
	for i := range states {
		params, state := channeltest.NewRandomParamsAndState(rng, channeltest.WithBackend(wallet.BackendID))
		states[i] = &channel.SignedState{Params: params, State: state}
	}
	return
}

func toStateMap(states ...*channel.SignedState) (_states channel.StateMap) {
	_states = channel.MakeStateMap()
	_states.Add(states...)
	return
}

func toEthStates(states ...*channel.State) (_states []adjudicator.ChannelState) {
	_states = make([]adjudicator.ChannelState, len(states))
	for i, s := range states {
		_states[i] = ToEthState(s)
	}
	return
}
