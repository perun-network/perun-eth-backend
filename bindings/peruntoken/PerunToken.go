// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package peruntoken

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

// PeruntokenMetaData contains all meta data concerning the Peruntoken contract.
var PeruntokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"initBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051610cd3380380610cd383398101604081905261002e916101e5565b6040518060400160405280600a8152602001692832b93ab72a37b5b2b760b11b8152506040518060400160405280600381526020016228292760e91b815250816003908161007c9190610347565b5060046100898282610347565b505f9150505b82518110156100c9576100c18382815181106100ad576100ad610405565b6020026020010151836100d160201b60201c565b60010161008f565b50505061043e565b6001600160a01b03821661012b5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b8060025f82825461013c9190610419565b90915550506001600160a01b0382165f9081526020819052604081208054839290610168908490610419565b90915550506040518181526001600160a01b038316905f907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b505050565b634e487b7160e01b5f52604160045260245ffd5b80516001600160a01b03811681146101e0575f5ffd5b919050565b5f5f604083850312156101f6575f5ffd5b82516001600160401b0381111561020b575f5ffd5b8301601f8101851361021b575f5ffd5b80516001600160401b03811115610234576102346101b6565b604051600582901b90603f8201601f191681016001600160401b0381118282101715610262576102626101b6565b60405291825260208184018101929081018884111561027f575f5ffd5b6020850194505b838510156102a557610297856101ca565b815260209485019401610286565b506020969096015195979596505050505050565b600181811c908216806102cd57607f821691505b6020821081036102eb57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156101b157828211156101b157805f5260205f20601f840160051c602085101561031c57505f5b90810190601f840160051c035f5b8181101561033f575f8382015560010161032a565b505050505050565b81516001600160401b03811115610360576103606101b6565b6103748161036e84546102b9565b846102f1565b6020601f8211600181146103a6575f831561038f5750848201515b5f19600385901b1c1916600184901b1784556103fe565b5f84815260208120601f198516915b828110156103d557878501518255602094850194600190920191016103b5565b50848210156103f257868401515f19600387901b60f8161c191681555b505060018360011b0184555b5050505050565b634e487b7160e01b5f52603260045260245ffd5b8082018082111561043857634e487b7160e01b5f52601160045260245ffd5b92915050565b6108888061044b5f395ff3fe608060405234801561000f575f5ffd5b50600436106100a6575f3560e01c8063395093511161006e578063395093511461011f57806370a082311461013257806395d89b411461015a578063a457c2d714610162578063a9059cbb14610175578063dd62ed3e14610188575f5ffd5b806306fdde03146100aa578063095ea7b3146100c857806318160ddd146100eb57806323b872dd146100fd578063313ce56714610110575b5f5ffd5b6100b26101c0565b6040516100bf91906106f8565b60405180910390f35b6100db6100d6366004610748565b610250565b60405190151581526020016100bf565b6002545b6040519081526020016100bf565b6100db61010b366004610770565b610266565b604051601281526020016100bf565b6100db61012d366004610748565b61031a565b6100ef6101403660046107aa565b6001600160a01b03165f9081526020819052604090205490565b6100b2610355565b6100db610170366004610748565b610364565b6100db610183366004610748565b6103fc565b6100ef6101963660046107ca565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6060600380546101cf906107fb565b80601f01602080910402602001604051908101604052809291908181526020018280546101fb906107fb565b80156102465780601f1061021d57610100808354040283529160200191610246565b820191905f5260205f20905b81548152906001019060200180831161022957829003601f168201915b5050505050905090565b5f61025c338484610408565b5060015b92915050565b6001600160a01b0383165f9081526001602090815260408083203384529091528120545f19811461030457828110156102f75760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b60648201526084015b60405180910390fd5b6103048533858403610408565b61030f85858561052b565b506001949350505050565b335f8181526001602090815260408083206001600160a01b0387168452909152812054909161025c918590610350908690610833565b610408565b6060600480546101cf906107fb565b335f9081526001602090815260408083206001600160a01b0386168452909152812054828110156103e55760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016102ee565b6103f23385858403610408565b5060019392505050565b5f61025c33848461052b565b6001600160a01b03831661046a5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016102ee565b6001600160a01b0382166104cb5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016102ee565b6001600160a01b038381165f8181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b03831661058f5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016102ee565b6001600160a01b0382166105f15760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016102ee565b6001600160a01b0383165f90815260208190526040902054818110156106685760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016102ee565b6001600160a01b038085165f9081526020819052604080822085850390559185168152908120805484929061069e908490610833565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516106ea91815260200190565b60405180910390a350505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80356001600160a01b0381168114610743575f5ffd5b919050565b5f5f60408385031215610759575f5ffd5b6107628361072d565b946020939093013593505050565b5f5f5f60608486031215610782575f5ffd5b61078b8461072d565b92506107996020850161072d565b929592945050506040919091013590565b5f602082840312156107ba575f5ffd5b6107c38261072d565b9392505050565b5f5f604083850312156107db575f5ffd5b6107e48361072d565b91506107f26020840161072d565b90509250929050565b600181811c9082168061080f57607f821691505b60208210810361082d57634e487b7160e01b5f52602260045260245ffd5b50919050565b8082018082111561026057634e487b7160e01b5f52601160045260245ffdfea264697066735822122019ea585083333ab51ae65dd18d3ae17fb0dedaa40dac45a11ee42b402290f75964736f6c63430008220033",
}

// PeruntokenABI is the input ABI used to generate the binding from.
// Deprecated: Use PeruntokenMetaData.ABI instead.
var PeruntokenABI = PeruntokenMetaData.ABI

// PeruntokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PeruntokenMetaData.Bin instead.
var PeruntokenBin = PeruntokenMetaData.Bin

// DeployPeruntoken deploys a new Ethereum contract, binding an instance of Peruntoken to it.
func DeployPeruntoken(auth *bind.TransactOpts, backend bind.ContractBackend, accounts []common.Address, initBalance *big.Int) (common.Address, *types.Transaction, *Peruntoken, error) {
	parsed, err := PeruntokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PeruntokenBin), backend, accounts, initBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Peruntoken{PeruntokenCaller: PeruntokenCaller{contract: contract}, PeruntokenTransactor: PeruntokenTransactor{contract: contract}, PeruntokenFilterer: PeruntokenFilterer{contract: contract}}, nil
}

// Peruntoken is an auto generated Go binding around an Ethereum contract.
type Peruntoken struct {
	PeruntokenCaller     // Read-only binding to the contract
	PeruntokenTransactor // Write-only binding to the contract
	PeruntokenFilterer   // Log filterer for contract events
}

// PeruntokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PeruntokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeruntokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PeruntokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeruntokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PeruntokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeruntokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PeruntokenSession struct {
	Contract     *Peruntoken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PeruntokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PeruntokenCallerSession struct {
	Contract *PeruntokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PeruntokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PeruntokenTransactorSession struct {
	Contract     *PeruntokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PeruntokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PeruntokenRaw struct {
	Contract *Peruntoken // Generic contract binding to access the raw methods on
}

// PeruntokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PeruntokenCallerRaw struct {
	Contract *PeruntokenCaller // Generic read-only contract binding to access the raw methods on
}

// PeruntokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PeruntokenTransactorRaw struct {
	Contract *PeruntokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeruntoken creates a new instance of Peruntoken, bound to a specific deployed contract.
func NewPeruntoken(address common.Address, backend bind.ContractBackend) (*Peruntoken, error) {
	contract, err := bindPeruntoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Peruntoken{PeruntokenCaller: PeruntokenCaller{contract: contract}, PeruntokenTransactor: PeruntokenTransactor{contract: contract}, PeruntokenFilterer: PeruntokenFilterer{contract: contract}}, nil
}

// NewPeruntokenCaller creates a new read-only instance of Peruntoken, bound to a specific deployed contract.
func NewPeruntokenCaller(address common.Address, caller bind.ContractCaller) (*PeruntokenCaller, error) {
	contract, err := bindPeruntoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PeruntokenCaller{contract: contract}, nil
}

// NewPeruntokenTransactor creates a new write-only instance of Peruntoken, bound to a specific deployed contract.
func NewPeruntokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PeruntokenTransactor, error) {
	contract, err := bindPeruntoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PeruntokenTransactor{contract: contract}, nil
}

// NewPeruntokenFilterer creates a new log filterer instance of Peruntoken, bound to a specific deployed contract.
func NewPeruntokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PeruntokenFilterer, error) {
	contract, err := bindPeruntoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PeruntokenFilterer{contract: contract}, nil
}

// bindPeruntoken binds a generic wrapper to an already deployed contract.
func bindPeruntoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PeruntokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Peruntoken *PeruntokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Peruntoken.Contract.PeruntokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Peruntoken *PeruntokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peruntoken.Contract.PeruntokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Peruntoken *PeruntokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Peruntoken.Contract.PeruntokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Peruntoken *PeruntokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Peruntoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Peruntoken *PeruntokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Peruntoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Peruntoken *PeruntokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Peruntoken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Peruntoken *PeruntokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Peruntoken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Peruntoken *PeruntokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Peruntoken.Contract.Allowance(&_Peruntoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Peruntoken *PeruntokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Peruntoken.Contract.Allowance(&_Peruntoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Peruntoken *PeruntokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Peruntoken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Peruntoken *PeruntokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Peruntoken.Contract.BalanceOf(&_Peruntoken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Peruntoken *PeruntokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Peruntoken.Contract.BalanceOf(&_Peruntoken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Peruntoken *PeruntokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Peruntoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Peruntoken *PeruntokenSession) Decimals() (uint8, error) {
	return _Peruntoken.Contract.Decimals(&_Peruntoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Peruntoken *PeruntokenCallerSession) Decimals() (uint8, error) {
	return _Peruntoken.Contract.Decimals(&_Peruntoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Peruntoken *PeruntokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Peruntoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Peruntoken *PeruntokenSession) Name() (string, error) {
	return _Peruntoken.Contract.Name(&_Peruntoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Peruntoken *PeruntokenCallerSession) Name() (string, error) {
	return _Peruntoken.Contract.Name(&_Peruntoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Peruntoken *PeruntokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Peruntoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Peruntoken *PeruntokenSession) Symbol() (string, error) {
	return _Peruntoken.Contract.Symbol(&_Peruntoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Peruntoken *PeruntokenCallerSession) Symbol() (string, error) {
	return _Peruntoken.Contract.Symbol(&_Peruntoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Peruntoken *PeruntokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Peruntoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Peruntoken *PeruntokenSession) TotalSupply() (*big.Int, error) {
	return _Peruntoken.Contract.TotalSupply(&_Peruntoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Peruntoken *PeruntokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Peruntoken.Contract.TotalSupply(&_Peruntoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.Approve(&_Peruntoken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.Approve(&_Peruntoken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Peruntoken *PeruntokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Peruntoken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Peruntoken *PeruntokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.DecreaseAllowance(&_Peruntoken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Peruntoken *PeruntokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.DecreaseAllowance(&_Peruntoken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Peruntoken *PeruntokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Peruntoken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Peruntoken *PeruntokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.IncreaseAllowance(&_Peruntoken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Peruntoken *PeruntokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.IncreaseAllowance(&_Peruntoken.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.Transfer(&_Peruntoken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.Transfer(&_Peruntoken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.TransferFrom(&_Peruntoken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Peruntoken *PeruntokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Peruntoken.Contract.TransferFrom(&_Peruntoken.TransactOpts, sender, recipient, amount)
}

// PeruntokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Peruntoken contract.
type PeruntokenApprovalIterator struct {
	Event *PeruntokenApproval // Event containing the contract specifics and raw log

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
func (it *PeruntokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeruntokenApproval)
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
		it.Event = new(PeruntokenApproval)
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
func (it *PeruntokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeruntokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeruntokenApproval represents a Approval event raised by the Peruntoken contract.
type PeruntokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Peruntoken *PeruntokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PeruntokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Peruntoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PeruntokenApprovalIterator{contract: _Peruntoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Peruntoken *PeruntokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PeruntokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Peruntoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeruntokenApproval)
				if err := _Peruntoken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Peruntoken *PeruntokenFilterer) ParseApproval(log types.Log) (*PeruntokenApproval, error) {
	event := new(PeruntokenApproval)
	if err := _Peruntoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PeruntokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Peruntoken contract.
type PeruntokenTransferIterator struct {
	Event *PeruntokenTransfer // Event containing the contract specifics and raw log

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
func (it *PeruntokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeruntokenTransfer)
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
		it.Event = new(PeruntokenTransfer)
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
func (it *PeruntokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeruntokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeruntokenTransfer represents a Transfer event raised by the Peruntoken contract.
type PeruntokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Peruntoken *PeruntokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PeruntokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Peruntoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PeruntokenTransferIterator{contract: _Peruntoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Peruntoken *PeruntokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PeruntokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Peruntoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeruntokenTransfer)
				if err := _Peruntoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Peruntoken *PeruntokenFilterer) ParseTransfer(log types.Log) (*PeruntokenTransfer, error) {
	event := new(PeruntokenTransfer)
	if err := _Peruntoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
