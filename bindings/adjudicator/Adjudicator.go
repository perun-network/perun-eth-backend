// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adjudicator

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AdjudicatorSignedState is an auto generated low-level Go binding around an user-defined struct.
type AdjudicatorSignedState struct {
	Params ChannelParams
	State  ChannelState
	Sigs   [][]byte
}

// ChannelAllocation is an auto generated low-level Go binding around an user-defined struct.
type ChannelAllocation struct {
	Assets   []ChannelAsset
	Backends []*big.Int
	Balances [][]*big.Int
	Locked   []ChannelSubAlloc
}

// ChannelAsset is an auto generated low-level Go binding around an user-defined struct.
type ChannelAsset struct {
	ChainID   *big.Int
	EthHolder common.Address
	CcHolder  []byte
}

// ChannelParams is an auto generated low-level Go binding around an user-defined struct.
type ChannelParams struct {
	ChallengeDuration *big.Int
	Nonce             *big.Int
	Participants      []ChannelParticipant
	App               common.Address
	LedgerChannel     bool
	VirtualChannel    bool
	Coordinator       common.Address
}

// ChannelParticipant is an auto generated low-level Go binding around an user-defined struct.
type ChannelParticipant struct {
	EthAddress common.Address
	CcAddress  []byte
}

// ChannelState is an auto generated low-level Go binding around an user-defined struct.
type ChannelState struct {
	ChannelID [32]byte
	Version   uint64
	Outcome   ChannelAllocation
	AppData   []byte
	IsFinal   bool
}

// ChannelSubAlloc is an auto generated low-level Go binding around an user-defined struct.
type ChannelSubAlloc struct {
	ID       [32]byte
	Balances []*big.Int
	IndexMap []uint16
}

// AdjudicatorMetaData contains all meta data concerning the Adjudicator contract.
var AdjudicatorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timeout\",\"type\":\"uint64\"}],\"name\":\"ChannelUpdate\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccAddress\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Participant[]\",\"name\":\"participants\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"channelID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccAddress\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Participant[]\",\"name\":\"participants\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethHolder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccHolder\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"backends\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethHolder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccHolder\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"backends\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State[]\",\"name\":\"subStates\",\"type\":\"tuple[]\"}],\"name\":\"conclude\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccAddress\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Participant[]\",\"name\":\"participants\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethHolder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccHolder\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"backends\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"name\":\"concludeFinal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccAddress\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Participant[]\",\"name\":\"participants\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethHolder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccHolder\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"backends\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"internalType\":\"structAdjudicator.SignedState\",\"name\":\"channel\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccAddress\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Participant[]\",\"name\":\"participants\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethHolder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccHolder\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"backends\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"internalType\":\"structAdjudicator.SignedState[]\",\"name\":\"subChannels\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"coordSigs\",\"type\":\"bytes[]\"}],\"name\":\"coordinate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputes\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timeout\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"challengeDuration\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"hasApp\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethHolder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccHolder\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"backends\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"hashState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccAddress\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Participant[]\",\"name\":\"participants\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethHolder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccHolder\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"backends\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"stateOld\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethHolder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccHolder\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"backends\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"actorIdx\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"progress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccAddress\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Participant[]\",\"name\":\"participants\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethHolder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccHolder\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"backends\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"internalType\":\"structAdjudicator.SignedState\",\"name\":\"channel\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"challengeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccAddress\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Participant[]\",\"name\":\"participants\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"app\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ledgerChannel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualChannel\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"internalType\":\"structChannel.Params\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethHolder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccHolder\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Asset[]\",\"name\":\"assets\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"backends\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"balances\",\"type\":\"uint256[][]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"indexMap\",\"type\":\"uint16[]\"}],\"internalType\":\"structChannel.SubAlloc[]\",\"name\":\"locked\",\"type\":\"tuple[]\"}],\"internalType\":\"structChannel.Allocation\",\"name\":\"outcome\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"appData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isFinal\",\"type\":\"bool\"}],\"internalType\":\"structChannel.State\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"}],\"internalType\":\"structAdjudicator.SignedState[]\",\"name\":\"subChannels\",\"type\":\"tuple[]\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50613ba98061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610085575f3560e01c806394f23fbe1161005857806394f23fbe146101725780639620eee314610185578063cd91b8ac14610198578063ec53f6c8146101ab575f5ffd5b806311be1997146100895780637a114d3d146101295780637a8380b21461014a5780637c39037d1461015f575b5f5ffd5b6100e36100973660046129f2565b5f602081905290815260409020805460018201546002909201546001600160401b0380831693600160401b8404821693600160801b81049092169260ff600160c01b9093048316921686565b604080516001600160401b0397881681529587166020870152939095169284019290925260ff166060830152608082015290151560a082015260c0015b60405180910390f35b61013c61013736600461302c565b6101be565b604051908152602001610120565b61015d61015836600461335b565b6101d6565b005b61013c61016d3660046133be565b610212565b61015d6101803660046133ef565b61021c565b61015d6101933660046134ad565b6104a6565b61015d6101a63660046135b0565b610502565b61015d6101b936600461363b565b61066c565b5f6101c882610707565b805190602001209050919050565b8151608001516102015760405162461bcd60e51b81526004016101f890613695565b60405180910390fd5b61020c82825f610730565b50505050565b5f6101c8826108a7565b5f610229845f01516108ba565b8054909150600160c01b900460ff166102905780546001600160401b031642101561028b5760405162461bcd60e51b81526020600482015260126024820152711d1a5b595bdd5d081b9bdd081c185cdcd95960721b60448201526064016101f8565b610326565b80545f19600160c01b90910460ff16016102ee5780546001600160401b0316421061028b5760405162461bcd60e51b815260206004820152600e60248201526d1d1a5b595bdd5d081c185cdcd95960921b60448201526064016101f8565b60405162461bcd60e51b815260206004820152600d60248201526c696e76616c696420706861736560981b60448201526064016101f8565b60608601516001600160a01b03166103705760405162461bcd60e51b815260206004820152600d60248201526c06d75737420686176652061707609c1b60448201526064016101f8565b85604001515183106103bc5760405162461bcd60e51b81526020600482015260156024820152746163746f72496478206f7574206f662072616e676560581b60448201526064016101f8565b6103c68685610912565b6103cf856101be565b8160010154146104135760405162461bcd60e51b815260206004820152600f60248201526e77726f6e67206f6c6420737461746560881b60448201526064016101f8565b61044661041f85610707565b8388604001518681518110610436576104366136cd565b60200260200101515f015161095f565b6104865760405162461bcd60e51b8152602060048201526011602482015270696e76616c6964207369676e617475726560781b60448201526064016101f8565b610492868686866109e5565b61049e86856001610b2d565b505050505050565b82608001516104c75760405162461bcd60e51b81526004016101f890613695565b6104d18383610912565b6104db8383610ce4565b5f6104e783835f610f29565b508351604080860151519087015192935061020c928461127e565b82608001516105235760405162461bcd60e51b81526004016101f890613695565b6080820151151560011461056b5760405162461bcd60e51b815260206004820152600f60248201526e1cdd185d19481b9bdd08199a5b985b608a1b60448201526064016101f8565b60408201516060015151156105b55760405162461bcd60e51b815260206004820152601060248201526f686173207375622d6368616e6e656c7360801b60448201526064016101f8565b6105bf8383610912565b6105ca838383611347565b81515f908152602081905260409020600181015415801590610637578154600219600160c01b90910460ff16016106375760405162461bcd60e51b8152602060048201526011602482015270636f6e636c7564656420616c726561647960781b60448201526064016101f8565b61064385856003610b2d565b610665845f015185604001515f0151876040015187604001516040015161127e565b5050505050565b82516080015161068e5760405162461bcd60e51b81526004016101f890613695565b61069f835f01518460200151610912565b81516106ac9060016136f5565b8151146106fb5760405162461bcd60e51b815260206004820152601960248201527f636f6f726453696773206c656e677468206d69736d617463680000000000000060448201526064016101f8565b6106658383835f61142f565b60608160405160200161071a91906139be565b6040516020818303038152906040529050919050565b60208301516040015180516060918391610749876115c8565b6107568260400151611822565b93505f826060015190508651815111156107aa5760405162461bcd60e51b81526020600482015260156024820152741cdd5890da185b9b995b1cc81d1bdbc81cda1bdc9d605a1b60448201526064016101f8565b5f5b815181101561089b575f88866107c1816139d0565b9750815181106107d3576107d36136cd565b602002602001015190505f5f8484815181106107f1576107f16136cd565b6020026020010151836020015191509150805f0151825f0151146108505760405162461bcd60e51b81526020600482015260166024820152751a5b9d985b1a59081cdd588b58da185b9b995b081a5960521b60448201526064016101f8565b606061085d848d8b610730565b604084015151909a5090915061087490889061190c565b6108828360200151826119ac565b61088c8a82611a91565b505050508060010190506107ac565b50505050935093915050565b60608160405160200161071a9190613ac4565b5f818152602081905260409020600181015415158061090c5760405162461bcd60e51b815260206004820152600e60248201526d1b9bdd081c9959da5cdd195c995960921b60448201526064016101f8565b50919050565b61091b82610212565b81511461095b5760405162461bcd60e51b815260206004820152600e60248201526d696e76616c696420706172616d7360901b60448201526064016101f8565b5050565b5f5f6109be85805190602001206040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018290525f90605c01604051602081830303815290604052805190602001209050919050565b90505f6109cb8286611aff565b6001600160a01b0390811690851614925050509392505050565b60208301516109f5906001613ad6565b6001600160401b031682602001516001600160401b031614610a595760405162461bcd60e51b815260206004820152601d60248201527f76657273696f6e206d75737420696e6372656d656e74206279206f6e6500000060448201526064016101f8565b608083015115610aab5760405162461bcd60e51b815260206004820181905260248201527f63616e6e6f742070726f67726573732066726f6d2066696e616c20737461746560448201526064016101f8565b610ac383604001518360400151866040015151611b23565b60608401516040516319e5ab7960e11b81526001600160a01b038216906333cb56f290610afa908890889088908890600401613af5565b5f6040518083038186803b158015610b10575f5ffd5b505afa158015610b22573d5f5f3e3d5ffd5b505050505050505050565b81515f908152602081815260409091206001810154855182549386015177ffffffffffffffffffffffffffffffff000000000000000019909416600160401b6001600160401b039283160267ffffffffffffffff60801b191617600160801b9190941602929092178155606085015190911515906001600160a01b0316151560028301805460ff1916911515919091179055826003811115610bd157610bd16136b9565b825460ff91909116600160c01b0260ff60c01b19909116178255610bf4846101be565b6001830155608084015115610c2257815467ffffffffffffffff1916426001600160401b0316178255610c79565b801580610c3b5750815460ff600160c01b909104166001145b15610c79578154610c5c90600160401b90046001600160401b031642613ad6565b825467ffffffffffffffff19166001600160401b03919091161782555b83516040805160c08101825284546001600160401b038082168352600160401b820481166020840152600160801b82041692820192909252600160c01b90910460ff908116606083015260018501546080830152600285015416151560a08201526106659190611df0565b5f610cf1825f01516108ba565b9050610cfc826101be565b816001015414610d3e5760405162461bcd60e51b815260206004820152600d60248201526c696e76616c696420737461746560981b60448201526064016101f8565b8054600219600160c01b90910460ff1601610d8f5760405162461bcd60e51b8152602060048201526011602482015270185b1c9958591e4818dbdb98db1d591959607a1b60448201526064016101f8565b805460ff600160c01b90910416600214610eaa57610dad8383611ee1565b15610dfa5760405162461bcd60e51b815260206004820152601f60248201527f636f6f7264696e6174656420736574746c656d656e742072657175697265640060448201526064016101f8565b8054600160c01b900460ff16158015610e175750600281015460ff165b15610e57578054610e3a906001600160401b03600160401b820481169116613ad6565b815467ffffffffffffffff19166001600160401b03919091161781555b80546001600160401b0316421015610eaa5760405162461bcd60e51b81526020600482015260166024820152751d1a5b595bdd5d081b9bdd081c185cdcd959081e595d60521b60448201526064016101f8565b8054600360c01b60ff60c01b1982161780835583516040805160c0810182526001600160401b039485168152600160401b840485166020820152600160801b840490941690840152600160c01b90910460ff908116606084015260018401546080840152600284015416151560a0830152610f2491611df0565b505050565b60605f610f3585611f0f565b60408501515180516001600160401b03811115610f5457610f54612a09565b604051908082528060200260200182016040528015610f8757816020015b6060815260200190600190039081610f725790505b5092505f5b81518110156110a8575f8760400151604001518281518110610fb057610fb06136cd565b6020026020010151905080516001600160401b03811115610fd357610fd3612a09565b604051908082528060200260200182016040528015610ffc578160200160208202803683370190505b5085838151811061100f5761100f6136cd565b60209081029190910101525f5b815181101561109e57886040015160400151838151811061103f5761103f6136cd565b60200260200101518181518110611058576110586136cd565b6020026020010151868481518110611072576110726136cd565b6020026020010151828151811061108b5761108b6136cd565b602090810291909101015260010161101c565b5050600101610f8c565b506040860151606001518492505f5b8151811015611273575f8282815181106110d3576110d36136cd565b602002602001015190505f8886806110ea906139d0565b9750815181106110fc576110fc6136cd565b60200260200101519050805f0151825f0151146111535760405162461bcd60e51b81526020600482015260156024820152741a5b9d985b1a59081cdd5898da185b9b995b081a59605a1b60448201526064016101f8565b6060611160828b89610f29565b60408501519098509091505f5b8751811015611263575f5b825181101561125a575f848381518110611194576111946136cd565b602002602001015182815181106111ad576111ad6136cd565b602002602001015190505f8483815181106111ca576111ca6136cd565b60200260200101519050818d85815181106111e7576111e76136cd565b60200260200101518261ffff1681518110611204576112046136cd565b602002602001015161121691906136f5565b8d8581518110611228576112286136cd565b60200260200101518261ffff1681518110611245576112456136cd565b60209081029190910101525050600101611178565b5060010161116d565b50505050508060010190506110b7565b505050935093915050565b5f5b8351811015610665575f84828151811061129c5761129c6136cd565b6020026020010151905046815f01510361133e5760208101516001600160a01b03161561133e5780602001516001600160a01b031663295482ce87868686815181106112ea576112ea6136cd565b60200260200101516040518463ffffffff1660e01b815260040161131093929190613b3f565b5f604051808303815f87803b158015611327575f5ffd5b505af1158015611339573d5f5f3e3d5ffd5b505050505b50600101611280565b5f61135183610707565b90508151846040015151146113a85760405162461bcd60e51b815260206004820152601a60248201527f7369676e617475726573206c656e677468206d69736d6174636800000000000060448201526064016101f8565b5f5b8251811015610665576113e7828483815181106113c9576113c96136cd565b602002602001015187604001518481518110610436576104366136cd565b6114275760405162461bcd60e51b8152602060048201526011602482015270696e76616c6964207369676e617475726560781b60448201526064016101f8565b6001016113aa565b602084015160400151805183516060928492909161146890899088908690811061145b5761145b6136cd565b6020026020010151611ff7565b6114758260400151611822565b93505f826060015190508751815111156114c95760405162461bcd60e51b81526020600482015260156024820152741cdd5890da185b9b995b1cc81d1bdbc81cda1bdc9d605a1b60448201526064016101f8565b5f5b81518110156115bb575f89866114e0816139d0565b9750815181106114f2576114f26136cd565b602002602001015190505f5f848481518110611510576115106136cd565b6020026020010151836020015191509150805f0151825f01511461156f5760405162461bcd60e51b81526020600482015260166024820152751a5b9d985b1a59081cdd588b58da185b9b995b081a5960521b60448201526064016101f8565b606061157d848e8e8c61142f565b604084015151909a5090915061159490889061190c565b6115a28360200151826119ac565b6115ac8a82611a91565b505050508060010190506114cb565b5050505094509492505050565b805160208201516115d98282610912565b6115e882828560400151611347565b8160a00151156116855760608201516001600160a01b03161561163f5760405162461bcd60e51b815260206004820152600f60248201526e063616e6e6f7420686176652061707608c1b60448201526064016101f8565b60408101516060015151156116855760405162461bcd60e51b815260206004820152600c60248201526b199d5b991cc81b1bd8dad95960a21b60448201526064016101f8565b80515f90815260208181526040918290206001810154835160c08101855282546001600160401b038082168352600160401b8204811695830195909552600160801b810490941694810194909452600160c01b90920460ff90811660608501526080840183905260029091015416151560a0830152158015906118175761170b836101be565b82608001510361171c575050505050565b82602001516001600160401b031682604001516001600160401b0316106117775760405162461bcd60e51b815260206004820152600f60248201526e34b73b30b634b2103b32b939b4b7b760891b60448201526064016101f8565b606082015160ff16156117be5760405162461bcd60e51b815260206004820152600f60248201526e696e636f727265637420706861736560881b60448201526064016101f8565b81516001600160401b031642106118175760405162461bcd60e51b815260206004820152601960248201527f72656675746174696f6e2074696d656f7574207061737365640000000000000060448201526064016101f8565b61066584845f610b2d565b606081516001600160401b0381111561183d5761183d612a09565b604051908082528060200260200182016040528015611866578160200160208202803683370190505b5090505f5b825181101561090c575f838281518110611887576118876136cd565b602002602001015190505f5f90505b8151811015611902578181815181106118b1576118b16136cd565b60200260200101518484815181106118cb576118cb6136cd565b60200260200101516118dd91906136f5565b8484815181106118ef576118ef6136cd565b6020908102919091010152600101611896565b505060010161186b565b805182511461195d5760405162461bcd60e51b815260206004820152601760248201527f41737365745b5d3a20756e657175616c206c656e67746800000000000000000060448201526064016101f8565b5f5b8251811015610f24576119a483828151811061197d5761197d6136cd565b6020026020010151838381518110611997576119976136cd565b6020026020010151612255565b60010161195f565b80518251146119fd5760405162461bcd60e51b815260206004820152601960248201527f75696e743235365b5d3a20756e657175616c206c656e6774680000000000000060448201526064016101f8565b5f5b8251811015610f2457818181518110611a1a57611a1a6136cd565b6020026020010151838281518110611a3457611a346136cd565b602002602001015114611a895760405162461bcd60e51b815260206004820152601760248201527f75696e743235365b5d3a20756e657175616c206974656d00000000000000000060448201526064016101f8565b6001016119ff565b5f5b8251811015610f2457818181518110611aae57611aae6136cd565b6020026020010151838281518110611ac857611ac86136cd565b6020026020010151611ada91906136f5565b838281518110611aec57611aec6136cd565b6020908102919091010152600101611a93565b5f5f5f611b0c858561234d565b91509150611b19816123b8565b5090505b92915050565b81604001515183604001515114611b7c5760405162461bcd60e51b815260206004820152601860248201527f62616c616e636573206c656e677468206d69736d61746368000000000000000060448201526064016101f8565b81515183515114611bc85760405162461bcd60e51b81526020600482015260166024820152750c2e6e6cae8e640d8cadccee8d040dad2e6dac2e8c6d60531b60448201526064016101f8565b611bda83606001518360600151612570565b5f5b82515181101561020c57611c1b845f01518281518110611bfe57611bfe6136cd565b6020026020010151845f01518381518110611997576119976136cd565b604084015180515f91829185919085908110611c3957611c396136cd565b60200260200101515114611c8f5760405162461bcd60e51b815260206004820152601c60248201527f6f6c642062616c616e636573206c656e677468206d69736d617463680000000060448201526064016101f8565b8385604001518481518110611ca657611ca66136cd565b60200260200101515114611cfc5760405162461bcd60e51b815260206004820152601c60248201527f6e65772062616c616e636573206c656e677468206d69736d617463680000000060448201526064016101f8565b5f5b84811015611d965786604001518481518110611d1c57611d1c6136cd565b60200260200101518181518110611d3557611d356136cd565b602002602001015183611d4891906136f5565b925085604001518481518110611d6057611d606136cd565b60200260200101518181518110611d7957611d796136cd565b602002602001015182611d8c91906136f5565b9150600101611cfe565b50808214611de65760405162461bcd60e51b815260206004820152601860248201527f73756d206f662062616c616e636573206d69736d61746368000000000000000060448201526064016101f8565b5050600101611bdc565b5f82815260208181526040918290208351815485840151868601516060808901516001600160401b039586166fffffffffffffffffffffffffffffffff199095168517600160401b948716949094029390931768ffffffffffffffffff60801b1916600160801b9590921694850260ff60c01b191691909117600160c01b60ff9093169283021785556080880151600186015560a08801516002909501805460ff1916951515959095179094558551928352938201939093529283019190915283917f895ef5a5fc3efd313a300b006d6ce97ff0670dfe04f6eea90417edf924fa786b910160405180910390a25050565b5f611ef88360c001516001600160a01b0316151590565b8015611f085750611f0882612610565b9392505050565b5f611f1c825f01516108ba565b9050611f27826101be565b816001015414611f695760405162461bcd60e51b815260206004820152600d60248201526c696e76616c696420737461746560981b60448201526064016101f8565b805460ff600160c01b9091041660031461095b578054600360c01b60ff60c01b1982161780835583516040805160c0810182526001600160401b039485168152600160401b840485166020820152600160801b840490941690840152600160c01b90910460ff908116606084015260018401546080840152600284015416151560a083015261095b91611df0565b815160208301516120088282610912565b80515f908152602081905260409020600181015415158061205c5760405162461bcd60e51b815260206004820152600e60248201526d1b9bdd081c9959da5cdd195c995960921b60448201526064016101f8565b815461207390600160c01b900460ff168585612702565b6120b15760405162461bcd60e51b815260206004820152600f60248201526e696e636f727265637420706861736560881b60448201526064016101f8565b81546001600160401b031642101561210b5760405162461bcd60e51b815260206004820152601d60248201527f72656675746174696f6e2074696d656f7574206e6f742070617373656400000060448201526064016101f8565b815460208401516001600160401b03600160801b9092048216911610156121665760405162461bcd60e51b815260206004820152600f60248201526e34b73b30b634b2103b32b939b4b7b760891b60448201526064016101f8565b61217584848860400151611347565b612180848487612721565b815460208401516001600160401b03600160801b9092048216911611156121db57602083015182546001600160401b03909116600160801b0267ffffffffffffffff60801b199091161782556121d5836101be565b60018301555b8154600160c11b60ff60c01b1982161780845584516040805160c0810182526001600160401b039485168152600160401b840485166020820152600160801b840490941690840152600160c01b90910460ff908116606084015260018501546080840152600285015416151560a083015261049e91611df0565b80518251146122985760405162461bcd60e51b815260206004820152600f60248201526e1d5b995c5d585b0818da185a5b9251608a1b60448201526064016101f8565b80602001516001600160a01b031682602001516001600160a01b0316146122f55760405162461bcd60e51b81526020600482015260116024820152703ab732b8bab0b61032ba342437b63232b960791b60448201526064016101f8565b8060400151805190602001208260400151805190602001201461095b5760405162461bcd60e51b815260206004820152601060248201526f3ab732b8bab0b61031b1a437b63232b960811b60448201526064016101f8565b5f5f8251604103612381576020830151604084015160608501515f1a61237587828585612788565b945094505050506123b1565b82516040036123aa576020830151604084015161239f86838361286d565b9350935050506123b1565b505f905060025b9250929050565b5f8160048111156123cb576123cb6136b9565b036123d35750565b60018160048111156123e7576123e76136b9565b036124345760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016101f8565b6002816004811115612448576124486136b9565b036124955760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016101f8565b60038160048111156124a9576124a96136b9565b036125015760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016101f8565b6004816004811115612515576125156136b9565b0361256d5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016101f8565b50565b80518251146125c15760405162461bcd60e51b815260206004820152601a60248201527f537562416c6c6f635b5d3a20756e657175616c206c656e67746800000000000060448201526064016101f8565b5f5b8251811015610f24576126088382815181106125e1576125e16136cd565b60200260200101518383815181106125fb576125fb6136cd565b60200260200101516128a5565b6001016125c3565b60408101518051515f91906001811161262c57505f9392505050565b808260200151511461264157505f9392505050565b5f82602001515f81518110612658576126586136cd565b602002602001015190505f835f01515f81518110612678576126786136cd565b602090810291909101015151905060015b838110156126f65782856020015182815181106126a8576126a86136cd565b60200260200101511415806126dc575081855f015182815181106126ce576126ce6136cd565b60200260200101515f015114155b156126ee575060019695505050505050565b600101612689565b505f9695505050505050565b5f60ff841615801561271957506127198383611ee1565b949350505050565b5f61272b83610707565b905061273c81838660c0015161095f565b61020c5760405162461bcd60e51b815260206004820152601d60248201527f696e76616c696420636f6f7264696e61746f72207369676e617475726500000060448201526064016101f8565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156127bd57505f90506003612864565b8460ff16601b141580156127d557508460ff16601c14155b156127e557505f90506004612864565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612836573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661285e575f60019250925050612864565b91505f90505b94509492505050565b5f806001600160ff1b0383168161288960ff86901c601b6136f5565b905061289787828885612788565b935093505050935093915050565b80518251146128ed5760405162461bcd60e51b815260206004820152601460248201527314dd58905b1b1bd8ce881d5b995c5d585b08125160621b60448201526064016101f8565b6128ff826020015182602001516119ac565b61095b82604001518260400151805182511461295d5760405162461bcd60e51b815260206004820152601860248201527f75696e7431365b5d3a20756e657175616c206c656e677468000000000000000060448201526064016101f8565b5f5b8251811015610f245781818151811061297a5761297a6136cd565b602002602001015161ffff16838281518110612998576129986136cd565b602002602001015161ffff16146129ea5760405162461bcd60e51b815260206004820152601660248201527575696e7431365b5d3a20756e657175616c206974656d60501b60448201526064016101f8565b60010161295f565b5f60208284031215612a02575f5ffd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715612a3f57612a3f612a09565b60405290565b604051608081016001600160401b0381118282101715612a3f57612a3f612a09565b60405160a081016001600160401b0381118282101715612a3f57612a3f612a09565b604080519081016001600160401b0381118282101715612a3f57612a3f612a09565b60405160e081016001600160401b0381118282101715612a3f57612a3f612a09565b604051601f8201601f191681016001600160401b0381118282101715612af557612af5612a09565b604052919050565b5f6001600160401b03821115612b1557612b15612a09565b5060051b60200190565b80356001600160a01b0381168114612b35575f5ffd5b919050565b5f82601f830112612b49575f5ffd5b81356001600160401b03811115612b6257612b62612a09565b612b75601f8201601f1916602001612acd565b818152846020838601011115612b89575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f82601f830112612bb4575f5ffd5b8135612bc7612bc282612afd565b612acd565b8082825260208201915060208360051b860101925085831115612be8575f5ffd5b602085015b83811015612c7b5780356001600160401b03811115612c0a575f5ffd5b86016060818903601f19011215612c1f575f5ffd5b612c27612a1d565b60208201358152612c3a60408301612b1f565b602082015260608201356001600160401b03811115612c57575f5ffd5b612c668a602083860101612b3a565b60408301525084525060209283019201612bed565b5095945050505050565b5f82601f830112612c94575f5ffd5b8135612ca2612bc282612afd565b8082825260208201915060208360051b860101925085831115612cc3575f5ffd5b602085015b83811015612c7b578035835260209283019201612cc8565b5f82601f830112612cef575f5ffd5b8135612cfd612bc282612afd565b8082825260208201915060208360051b860101925085831115612d1e575f5ffd5b602085015b83811015612c7b5780356001600160401b03811115612d40575f5ffd5b612d4f886020838a0101612c85565b84525060209283019201612d23565b5f82601f830112612d6d575f5ffd5b8135612d7b612bc282612afd565b8082825260208201915060208360051b860101925085831115612d9c575f5ffd5b602085015b83811015612c7b5780356001600160401b03811115612dbe575f5ffd5b86016060818903601f19011215612dd3575f5ffd5b612ddb612a1d565b6020820135815260408201356001600160401b03811115612dfa575f5ffd5b612e098a602083860101612c85565b60208301525060608201356001600160401b03811115612e27575f5ffd5b60208184010192505088601f830112612e3e575f5ffd5b8135612e4c612bc282612afd565b8082825260208201915060208360051b86010192508b831115612e6d575f5ffd5b6020850194505b82851015612e9e57843561ffff81168114612e8d575f5ffd5b825260209485019490910190612e74565b6040840152505084525060209283019201612da1565b5f60808284031215612ec4575f5ffd5b612ecc612a45565b905081356001600160401b03811115612ee3575f5ffd5b612eef84828501612ba5565b82525060208201356001600160401b03811115612f0a575f5ffd5b612f1684828501612c85565b60208301525060408201356001600160401b03811115612f34575f5ffd5b612f4084828501612ce0565b60408301525060608201356001600160401b03811115612f5e575f5ffd5b612f6a84828501612d5e565b60608301525092915050565b80358015158114612b35575f5ffd5b5f60a08284031215612f95575f5ffd5b612f9d612a67565b82358152905060208201356001600160401b0381168114612fbc575f5ffd5b602082015260408201356001600160401b03811115612fd9575f5ffd5b612fe584828501612eb4565b60408301525060608201356001600160401b03811115613003575f5ffd5b61300f84828501612b3a565b60608301525061302160808301612f76565b608082015292915050565b5f6020828403121561303c575f5ffd5b81356001600160401b03811115613051575f5ffd5b61271984828501612f85565b5f82601f83011261306c575f5ffd5b813561307a612bc282612afd565b8082825260208201915060208360051b86010192508583111561309b575f5ffd5b602085015b83811015612c7b5780356001600160401b038111156130bd575f5ffd5b86016040818903601f190112156130d2575f5ffd5b6130da612a89565b6130e660208301612b1f565b815260408201356001600160401b03811115613100575f5ffd5b61310f8a602083860101612b3a565b60208301525080855250506020830192506020810190506130a0565b5f60e0828403121561313b575f5ffd5b613143612aab565b8235815260208084013590820152905060408201356001600160401b0381111561316b575f5ffd5b6131778482850161305d565b60408301525061318960608301612b1f565b606082015261319a60808301612f76565b60808201526131ab60a08301612f76565b60a08201526131bc60c08301612b1f565b60c082015292915050565b5f82601f8301126131d6575f5ffd5b81356131e4612bc282612afd565b8082825260208201915060208360051b860101925085831115613205575f5ffd5b602085015b83811015612c7b5780356001600160401b03811115613227575f5ffd5b613236886020838a0101612b3a565b8452506020928301920161320a565b5f60608284031215613255575f5ffd5b61325d612a1d565b905081356001600160401b03811115613274575f5ffd5b6132808482850161312b565b82525060208201356001600160401b0381111561329b575f5ffd5b6132a784828501612f85565b60208301525060408201356001600160401b038111156132c5575f5ffd5b6132d1848285016131c7565b60408301525092915050565b5f82601f8301126132ec575f5ffd5b81356132fa612bc282612afd565b8082825260208201915060208360051b86010192508583111561331b575f5ffd5b602085015b83811015612c7b5780356001600160401b0381111561333d575f5ffd5b61334c886020838a0101613245565b84525060209283019201613320565b5f5f6040838503121561336c575f5ffd5b82356001600160401b03811115613381575f5ffd5b61338d85828601613245565b92505060208301356001600160401b038111156133a8575f5ffd5b6133b4858286016132dd565b9150509250929050565b5f602082840312156133ce575f5ffd5b81356001600160401b038111156133e3575f5ffd5b6127198482850161312b565b5f5f5f5f5f60a08688031215613403575f5ffd5b85356001600160401b03811115613418575f5ffd5b6134248882890161312b565b95505060208601356001600160401b0381111561343f575f5ffd5b61344b88828901612f85565b94505060408601356001600160401b03811115613466575f5ffd5b61347288828901612f85565b9350506060860135915060808601356001600160401b03811115613494575f5ffd5b6134a088828901612b3a565b9150509295509295909350565b5f5f5f606084860312156134bf575f5ffd5b83356001600160401b038111156134d4575f5ffd5b6134e08682870161312b565b93505060208401356001600160401b038111156134fb575f5ffd5b61350786828701612f85565b92505060408401356001600160401b03811115613522575f5ffd5b8401601f81018613613532575f5ffd5b8035613540612bc282612afd565b8082825260208201915060208360051b850101925088831115613561575f5ffd5b602084015b838110156135a15780356001600160401b03811115613583575f5ffd5b6135928b602083890101612f85565b84525060209283019201613566565b50809450505050509250925092565b5f5f5f606084860312156135c2575f5ffd5b83356001600160401b038111156135d7575f5ffd5b6135e38682870161312b565b93505060208401356001600160401b038111156135fe575f5ffd5b61360a86828701612f85565b92505060408401356001600160401b03811115613625575f5ffd5b613631868287016131c7565b9150509250925092565b5f5f5f6060848603121561364d575f5ffd5b83356001600160401b03811115613662575f5ffd5b61366e86828701613245565b93505060208401356001600160401b03811115613689575f5ffd5b61360a868287016132dd565b6020808252600a90820152693737ba103632b233b2b960b11b604082015260600190565b634e487b7160e01b5f52602160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b80820180821115611b1d57611b1d6136e1565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f8151808452602084019350602083015f5b82811015613766578151865260209586019590910190600101613748565b5093949350505050565b5f82825180855260208501945060208160051b830101602085015f5b838110156137be57601f198584030188526137a8838351613736565b602098890198909350919091019060010161378c565b50909695505050505050565b5f82825180855260208501945060208160051b830101602085015f5b838110156137be57601f198584030188528151805184526020810151606060208601526138166060860182613736565b6040928301518682039387019390935282518082526020938401935f9350909101905b808310156138605761ffff8451168252602082019150602084019350600183019250613839565b5060209a8b019a909550939093019250506001016137e6565b805182526001600160401b0360208201511660208301525f604082015160a0604085015261012084018151608060a0870152818151808452610140880191506101408160051b89010193506020830192505f5b818110156139285788850361013f190183528351805186526020808201516001600160a01b03169087015260409081015160609187018290529061391290870182613708565b95505060209384019392909201916001016138cc565b505050506020820151858203609f190160c08701526139478282613736565b9150506040820151609f198683030160e08701526139658282613770565b91505060608201519150609f198582030161010086015261398681836137ca565b915050606083015184820360608601526139a08282613708565b91505060808301516139b6608086018215159052565b509392505050565b602081525f611f086020830184613879565b5f600182016139e1576139e16136e1565b5060010190565b5f82825180855260208501945060208160051b830101602085015f5b838110156137be57848303601f19018852815180516001600160a01b03168452602090810151604091850182905290613a3f90850182613708565b6020998a0199909450929092019150600101613a04565b80518252602081015160208301525f604082015160e06040850152613a7e60e08501826139e8565b905060018060a01b03606084015116606085015260808301511515608085015260a0830151151560a085015260c08301516139b660c08601826001600160a01b03169052565b602081525f611f086020830184613a56565b6001600160401b038181168382160190811115611b1d57611b1d6136e1565b608081525f613b076080830187613a56565b8281036020840152613b198187613879565b90508281036040840152613b2d8186613879565b91505082606083015295945050505050565b838152606060208201525f613b5760608301856139e8565b8281036040840152613b698185613736565b969550505050505056fea264697066735822122020b9f36b12fba765630294538a4c2f16eec7b8ce8ea453a00a137960dc0e435764736f6c63430008220033",
}

// AdjudicatorABI is the input ABI used to generate the binding from.
// Deprecated: Use AdjudicatorMetaData.ABI instead.
var AdjudicatorABI = AdjudicatorMetaData.ABI

// AdjudicatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AdjudicatorMetaData.Bin instead.
var AdjudicatorBin = AdjudicatorMetaData.Bin

// DeployAdjudicator deploys a new Ethereum contract, binding an instance of Adjudicator to it.
func DeployAdjudicator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Adjudicator, error) {
	parsed, err := AdjudicatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AdjudicatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Adjudicator{AdjudicatorCaller: AdjudicatorCaller{contract: contract}, AdjudicatorTransactor: AdjudicatorTransactor{contract: contract}, AdjudicatorFilterer: AdjudicatorFilterer{contract: contract}}, nil
}

// Adjudicator is an auto generated Go binding around an Ethereum contract.
type Adjudicator struct {
	AdjudicatorCaller     // Read-only binding to the contract
	AdjudicatorTransactor // Write-only binding to the contract
	AdjudicatorFilterer   // Log filterer for contract events
}

// AdjudicatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdjudicatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdjudicatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdjudicatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdjudicatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdjudicatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdjudicatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdjudicatorSession struct {
	Contract     *Adjudicator      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdjudicatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdjudicatorCallerSession struct {
	Contract *AdjudicatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AdjudicatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdjudicatorTransactorSession struct {
	Contract     *AdjudicatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AdjudicatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdjudicatorRaw struct {
	Contract *Adjudicator // Generic contract binding to access the raw methods on
}

// AdjudicatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdjudicatorCallerRaw struct {
	Contract *AdjudicatorCaller // Generic read-only contract binding to access the raw methods on
}

// AdjudicatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdjudicatorTransactorRaw struct {
	Contract *AdjudicatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdjudicator creates a new instance of Adjudicator, bound to a specific deployed contract.
func NewAdjudicator(address common.Address, backend bind.ContractBackend) (*Adjudicator, error) {
	contract, err := bindAdjudicator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Adjudicator{AdjudicatorCaller: AdjudicatorCaller{contract: contract}, AdjudicatorTransactor: AdjudicatorTransactor{contract: contract}, AdjudicatorFilterer: AdjudicatorFilterer{contract: contract}}, nil
}

// NewAdjudicatorCaller creates a new read-only instance of Adjudicator, bound to a specific deployed contract.
func NewAdjudicatorCaller(address common.Address, caller bind.ContractCaller) (*AdjudicatorCaller, error) {
	contract, err := bindAdjudicator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdjudicatorCaller{contract: contract}, nil
}

// NewAdjudicatorTransactor creates a new write-only instance of Adjudicator, bound to a specific deployed contract.
func NewAdjudicatorTransactor(address common.Address, transactor bind.ContractTransactor) (*AdjudicatorTransactor, error) {
	contract, err := bindAdjudicator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdjudicatorTransactor{contract: contract}, nil
}

// NewAdjudicatorFilterer creates a new log filterer instance of Adjudicator, bound to a specific deployed contract.
func NewAdjudicatorFilterer(address common.Address, filterer bind.ContractFilterer) (*AdjudicatorFilterer, error) {
	contract, err := bindAdjudicator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdjudicatorFilterer{contract: contract}, nil
}

// bindAdjudicator binds a generic wrapper to an already deployed contract.
func bindAdjudicator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AdjudicatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Adjudicator *AdjudicatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Adjudicator.Contract.AdjudicatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Adjudicator *AdjudicatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adjudicator.Contract.AdjudicatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Adjudicator *AdjudicatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Adjudicator.Contract.AdjudicatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Adjudicator *AdjudicatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Adjudicator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Adjudicator *AdjudicatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Adjudicator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Adjudicator *AdjudicatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Adjudicator.Contract.contract.Transact(opts, method, params...)
}

// ChannelID is a free data retrieval call binding the contract method 0x7c39037d.
//
// Solidity: function channelID((uint256,uint256,(address,bytes)[],address,bool,bool,address) params) pure returns(bytes32)
func (_Adjudicator *AdjudicatorCaller) ChannelID(opts *bind.CallOpts, params ChannelParams) ([32]byte, error) {
	var out []interface{}
	err := _Adjudicator.contract.Call(opts, &out, "channelID", params)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChannelID is a free data retrieval call binding the contract method 0x7c39037d.
//
// Solidity: function channelID((uint256,uint256,(address,bytes)[],address,bool,bool,address) params) pure returns(bytes32)
func (_Adjudicator *AdjudicatorSession) ChannelID(params ChannelParams) ([32]byte, error) {
	return _Adjudicator.Contract.ChannelID(&_Adjudicator.CallOpts, params)
}

// ChannelID is a free data retrieval call binding the contract method 0x7c39037d.
//
// Solidity: function channelID((uint256,uint256,(address,bytes)[],address,bool,bool,address) params) pure returns(bytes32)
func (_Adjudicator *AdjudicatorCallerSession) ChannelID(params ChannelParams) ([32]byte, error) {
	return _Adjudicator.Contract.ChannelID(&_Adjudicator.CallOpts, params)
}

// Disputes is a free data retrieval call binding the contract method 0x11be1997.
//
// Solidity: function disputes(bytes32 ) view returns(uint64 timeout, uint64 challengeDuration, uint64 version, uint8 phase, bytes32 stateHash, bool hasApp)
func (_Adjudicator *AdjudicatorCaller) Disputes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Timeout           uint64
	ChallengeDuration uint64
	Version           uint64
	Phase             uint8
	StateHash         [32]byte
	HasApp            bool
}, error) {
	var out []interface{}
	err := _Adjudicator.contract.Call(opts, &out, "disputes", arg0)

	outstruct := new(struct {
		Timeout           uint64
		ChallengeDuration uint64
		Version           uint64
		Phase             uint8
		StateHash         [32]byte
		HasApp            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timeout = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.ChallengeDuration = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Version = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Phase = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.StateHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.HasApp = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Disputes is a free data retrieval call binding the contract method 0x11be1997.
//
// Solidity: function disputes(bytes32 ) view returns(uint64 timeout, uint64 challengeDuration, uint64 version, uint8 phase, bytes32 stateHash, bool hasApp)
func (_Adjudicator *AdjudicatorSession) Disputes(arg0 [32]byte) (struct {
	Timeout           uint64
	ChallengeDuration uint64
	Version           uint64
	Phase             uint8
	StateHash         [32]byte
	HasApp            bool
}, error) {
	return _Adjudicator.Contract.Disputes(&_Adjudicator.CallOpts, arg0)
}

// Disputes is a free data retrieval call binding the contract method 0x11be1997.
//
// Solidity: function disputes(bytes32 ) view returns(uint64 timeout, uint64 challengeDuration, uint64 version, uint8 phase, bytes32 stateHash, bool hasApp)
func (_Adjudicator *AdjudicatorCallerSession) Disputes(arg0 [32]byte) (struct {
	Timeout           uint64
	ChallengeDuration uint64
	Version           uint64
	Phase             uint8
	StateHash         [32]byte
	HasApp            bool
}, error) {
	return _Adjudicator.Contract.Disputes(&_Adjudicator.CallOpts, arg0)
}

// HashState is a free data retrieval call binding the contract method 0x7a114d3d.
//
// Solidity: function hashState((bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) state) pure returns(bytes32)
func (_Adjudicator *AdjudicatorCaller) HashState(opts *bind.CallOpts, state ChannelState) ([32]byte, error) {
	var out []interface{}
	err := _Adjudicator.contract.Call(opts, &out, "hashState", state)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashState is a free data retrieval call binding the contract method 0x7a114d3d.
//
// Solidity: function hashState((bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) state) pure returns(bytes32)
func (_Adjudicator *AdjudicatorSession) HashState(state ChannelState) ([32]byte, error) {
	return _Adjudicator.Contract.HashState(&_Adjudicator.CallOpts, state)
}

// HashState is a free data retrieval call binding the contract method 0x7a114d3d.
//
// Solidity: function hashState((bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) state) pure returns(bytes32)
func (_Adjudicator *AdjudicatorCallerSession) HashState(state ChannelState) ([32]byte, error) {
	return _Adjudicator.Contract.HashState(&_Adjudicator.CallOpts, state)
}

// Conclude is a paid mutator transaction binding the contract method 0x9620eee3.
//
// Solidity: function conclude((uint256,uint256,(address,bytes)[],address,bool,bool,address) params, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) state, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool)[] subStates) returns()
func (_Adjudicator *AdjudicatorTransactor) Conclude(opts *bind.TransactOpts, params ChannelParams, state ChannelState, subStates []ChannelState) (*types.Transaction, error) {
	return _Adjudicator.contract.Transact(opts, "conclude", params, state, subStates)
}

// Conclude is a paid mutator transaction binding the contract method 0x9620eee3.
//
// Solidity: function conclude((uint256,uint256,(address,bytes)[],address,bool,bool,address) params, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) state, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool)[] subStates) returns()
func (_Adjudicator *AdjudicatorSession) Conclude(params ChannelParams, state ChannelState, subStates []ChannelState) (*types.Transaction, error) {
	return _Adjudicator.Contract.Conclude(&_Adjudicator.TransactOpts, params, state, subStates)
}

// Conclude is a paid mutator transaction binding the contract method 0x9620eee3.
//
// Solidity: function conclude((uint256,uint256,(address,bytes)[],address,bool,bool,address) params, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) state, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool)[] subStates) returns()
func (_Adjudicator *AdjudicatorTransactorSession) Conclude(params ChannelParams, state ChannelState, subStates []ChannelState) (*types.Transaction, error) {
	return _Adjudicator.Contract.Conclude(&_Adjudicator.TransactOpts, params, state, subStates)
}

// ConcludeFinal is a paid mutator transaction binding the contract method 0xcd91b8ac.
//
// Solidity: function concludeFinal((uint256,uint256,(address,bytes)[],address,bool,bool,address) params, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) state, bytes[] sigs) returns()
func (_Adjudicator *AdjudicatorTransactor) ConcludeFinal(opts *bind.TransactOpts, params ChannelParams, state ChannelState, sigs [][]byte) (*types.Transaction, error) {
	return _Adjudicator.contract.Transact(opts, "concludeFinal", params, state, sigs)
}

// ConcludeFinal is a paid mutator transaction binding the contract method 0xcd91b8ac.
//
// Solidity: function concludeFinal((uint256,uint256,(address,bytes)[],address,bool,bool,address) params, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) state, bytes[] sigs) returns()
func (_Adjudicator *AdjudicatorSession) ConcludeFinal(params ChannelParams, state ChannelState, sigs [][]byte) (*types.Transaction, error) {
	return _Adjudicator.Contract.ConcludeFinal(&_Adjudicator.TransactOpts, params, state, sigs)
}

// ConcludeFinal is a paid mutator transaction binding the contract method 0xcd91b8ac.
//
// Solidity: function concludeFinal((uint256,uint256,(address,bytes)[],address,bool,bool,address) params, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) state, bytes[] sigs) returns()
func (_Adjudicator *AdjudicatorTransactorSession) ConcludeFinal(params ChannelParams, state ChannelState, sigs [][]byte) (*types.Transaction, error) {
	return _Adjudicator.Contract.ConcludeFinal(&_Adjudicator.TransactOpts, params, state, sigs)
}

// Coordinate is a paid mutator transaction binding the contract method 0xec53f6c8.
//
// Solidity: function coordinate(((uint256,uint256,(address,bytes)[],address,bool,bool,address),(bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),bytes[]) channel, ((uint256,uint256,(address,bytes)[],address,bool,bool,address),(bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),bytes[])[] subChannels, bytes[] coordSigs) returns()
func (_Adjudicator *AdjudicatorTransactor) Coordinate(opts *bind.TransactOpts, channel AdjudicatorSignedState, subChannels []AdjudicatorSignedState, coordSigs [][]byte) (*types.Transaction, error) {
	return _Adjudicator.contract.Transact(opts, "coordinate", channel, subChannels, coordSigs)
}

// Coordinate is a paid mutator transaction binding the contract method 0xec53f6c8.
//
// Solidity: function coordinate(((uint256,uint256,(address,bytes)[],address,bool,bool,address),(bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),bytes[]) channel, ((uint256,uint256,(address,bytes)[],address,bool,bool,address),(bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),bytes[])[] subChannels, bytes[] coordSigs) returns()
func (_Adjudicator *AdjudicatorSession) Coordinate(channel AdjudicatorSignedState, subChannels []AdjudicatorSignedState, coordSigs [][]byte) (*types.Transaction, error) {
	return _Adjudicator.Contract.Coordinate(&_Adjudicator.TransactOpts, channel, subChannels, coordSigs)
}

// Coordinate is a paid mutator transaction binding the contract method 0xec53f6c8.
//
// Solidity: function coordinate(((uint256,uint256,(address,bytes)[],address,bool,bool,address),(bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),bytes[]) channel, ((uint256,uint256,(address,bytes)[],address,bool,bool,address),(bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),bytes[])[] subChannels, bytes[] coordSigs) returns()
func (_Adjudicator *AdjudicatorTransactorSession) Coordinate(channel AdjudicatorSignedState, subChannels []AdjudicatorSignedState, coordSigs [][]byte) (*types.Transaction, error) {
	return _Adjudicator.Contract.Coordinate(&_Adjudicator.TransactOpts, channel, subChannels, coordSigs)
}

// Progress is a paid mutator transaction binding the contract method 0x94f23fbe.
//
// Solidity: function progress((uint256,uint256,(address,bytes)[],address,bool,bool,address) params, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) stateOld, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) state, uint256 actorIdx, bytes sig) returns()
func (_Adjudicator *AdjudicatorTransactor) Progress(opts *bind.TransactOpts, params ChannelParams, stateOld ChannelState, state ChannelState, actorIdx *big.Int, sig []byte) (*types.Transaction, error) {
	return _Adjudicator.contract.Transact(opts, "progress", params, stateOld, state, actorIdx, sig)
}

// Progress is a paid mutator transaction binding the contract method 0x94f23fbe.
//
// Solidity: function progress((uint256,uint256,(address,bytes)[],address,bool,bool,address) params, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) stateOld, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) state, uint256 actorIdx, bytes sig) returns()
func (_Adjudicator *AdjudicatorSession) Progress(params ChannelParams, stateOld ChannelState, state ChannelState, actorIdx *big.Int, sig []byte) (*types.Transaction, error) {
	return _Adjudicator.Contract.Progress(&_Adjudicator.TransactOpts, params, stateOld, state, actorIdx, sig)
}

// Progress is a paid mutator transaction binding the contract method 0x94f23fbe.
//
// Solidity: function progress((uint256,uint256,(address,bytes)[],address,bool,bool,address) params, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) stateOld, (bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool) state, uint256 actorIdx, bytes sig) returns()
func (_Adjudicator *AdjudicatorTransactorSession) Progress(params ChannelParams, stateOld ChannelState, state ChannelState, actorIdx *big.Int, sig []byte) (*types.Transaction, error) {
	return _Adjudicator.Contract.Progress(&_Adjudicator.TransactOpts, params, stateOld, state, actorIdx, sig)
}

// Register is a paid mutator transaction binding the contract method 0x7a8380b2.
//
// Solidity: function register(((uint256,uint256,(address,bytes)[],address,bool,bool,address),(bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),bytes[]) channel, ((uint256,uint256,(address,bytes)[],address,bool,bool,address),(bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),bytes[])[] subChannels) returns()
func (_Adjudicator *AdjudicatorTransactor) Register(opts *bind.TransactOpts, channel AdjudicatorSignedState, subChannels []AdjudicatorSignedState) (*types.Transaction, error) {
	return _Adjudicator.contract.Transact(opts, "register", channel, subChannels)
}

// Register is a paid mutator transaction binding the contract method 0x7a8380b2.
//
// Solidity: function register(((uint256,uint256,(address,bytes)[],address,bool,bool,address),(bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),bytes[]) channel, ((uint256,uint256,(address,bytes)[],address,bool,bool,address),(bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),bytes[])[] subChannels) returns()
func (_Adjudicator *AdjudicatorSession) Register(channel AdjudicatorSignedState, subChannels []AdjudicatorSignedState) (*types.Transaction, error) {
	return _Adjudicator.Contract.Register(&_Adjudicator.TransactOpts, channel, subChannels)
}

// Register is a paid mutator transaction binding the contract method 0x7a8380b2.
//
// Solidity: function register(((uint256,uint256,(address,bytes)[],address,bool,bool,address),(bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),bytes[]) channel, ((uint256,uint256,(address,bytes)[],address,bool,bool,address),(bytes32,uint64,((uint256,address,bytes)[],uint256[],uint256[][],(bytes32,uint256[],uint16[])[]),bytes,bool),bytes[])[] subChannels) returns()
func (_Adjudicator *AdjudicatorTransactorSession) Register(channel AdjudicatorSignedState, subChannels []AdjudicatorSignedState) (*types.Transaction, error) {
	return _Adjudicator.Contract.Register(&_Adjudicator.TransactOpts, channel, subChannels)
}

// AdjudicatorChannelUpdateIterator is returned from FilterChannelUpdate and is used to iterate over the raw logs and unpacked data for ChannelUpdate events raised by the Adjudicator contract.
type AdjudicatorChannelUpdateIterator struct {
	Event *AdjudicatorChannelUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AdjudicatorChannelUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdjudicatorChannelUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AdjudicatorChannelUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AdjudicatorChannelUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdjudicatorChannelUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdjudicatorChannelUpdate represents a ChannelUpdate event raised by the Adjudicator contract.
type AdjudicatorChannelUpdate struct {
	ChannelID [32]byte
	Version   uint64
	Phase     uint8
	Timeout   uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelUpdate is a free log retrieval operation binding the contract event 0x895ef5a5fc3efd313a300b006d6ce97ff0670dfe04f6eea90417edf924fa786b.
//
// Solidity: event ChannelUpdate(bytes32 indexed channelID, uint64 version, uint8 phase, uint64 timeout)
func (_Adjudicator *AdjudicatorFilterer) FilterChannelUpdate(opts *bind.FilterOpts, channelID [][32]byte) (*AdjudicatorChannelUpdateIterator, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	logs, sub, err := _Adjudicator.contract.FilterLogs(opts, "ChannelUpdate", channelIDRule)
	if err != nil {
		return nil, err
	}
	return &AdjudicatorChannelUpdateIterator{contract: _Adjudicator.contract, event: "ChannelUpdate", logs: logs, sub: sub}, nil
}

// WatchChannelUpdate is a free log subscription operation binding the contract event 0x895ef5a5fc3efd313a300b006d6ce97ff0670dfe04f6eea90417edf924fa786b.
//
// Solidity: event ChannelUpdate(bytes32 indexed channelID, uint64 version, uint8 phase, uint64 timeout)
func (_Adjudicator *AdjudicatorFilterer) WatchChannelUpdate(opts *bind.WatchOpts, sink chan<- *AdjudicatorChannelUpdate, channelID [][32]byte) (event.Subscription, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	logs, sub, err := _Adjudicator.contract.WatchLogs(opts, "ChannelUpdate", channelIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdjudicatorChannelUpdate)
				if err := _Adjudicator.contract.UnpackLog(event, "ChannelUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChannelUpdate is a log parse operation binding the contract event 0x895ef5a5fc3efd313a300b006d6ce97ff0670dfe04f6eea90417edf924fa786b.
//
// Solidity: event ChannelUpdate(bytes32 indexed channelID, uint64 version, uint8 phase, uint64 timeout)
func (_Adjudicator *AdjudicatorFilterer) ParseChannelUpdate(log types.Log) (*AdjudicatorChannelUpdate, error) {
	event := new(AdjudicatorChannelUpdate)
	if err := _Adjudicator.contract.UnpackLog(event, "ChannelUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
