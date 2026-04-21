// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package liquiditypool

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

// LiquidityPoolMetaData contains all meta data concerning the LiquidityPool contract.
var LiquidityPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ChannelFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalReturned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGain\",\"type\":\"uint256\"}],\"name\":\"ChannelSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sharesMinted\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"OperatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sharesBurned\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mintedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"lockedByChannel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"}],\"name\":\"previewDepositShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesAmount\",\"type\":\"uint256\"}],\"name\":\"previewWithdrawETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"}],\"name\":\"settleChannel\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLockedETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesToBurn\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethAmountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawableETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161105b38038061105b83398101604081905261002f91610101565b600160005561003d336100af565b6001600160a01b03811661008a5760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b21037b832b930ba37b960811b604482015260640160405180910390fd5b600480546001600160a01b0319166001600160a01b0392909216919091179055610131565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006020828403121561011357600080fd5b81516001600160a01b038116811461012a57600080fd5b9392505050565b610f1b806101406000396000f3fe6080604052600436106101025760003560e01c806378f3bfe211610095578063b6cc49da11610064578063b6cc49da146102cd578063d0e30db0146102fa578063f2fde38b14610302578063f5eb42dc14610322578063fd3d31991461035857600080fd5b806378f3bfe2146102595780638afe47cf146102795780638da5cb5b1461028f578063b3ab15fb146102ad57600080fd5b8063570ca735116100d1578063570ca735146101d95780636574bcd4146102115780636913e7cc14610231578063715018a61461024457600080fd5b806301e1d11414610159578063260b6086146101815780632e1a7d4d146101a35780633a98ef39146101c357600080fd5b366101545760405162461bcd60e51b815260206004820152601c60248201527f557365206465706f736974206f7220736574746c654368616e6e656c0000000060448201526064015b60405180910390fd5b600080fd5b34801561016557600080fd5b5061016e61036b565b6040519081526020015b60405180910390f35b34801561018d57600080fd5b506101a161019c366004610d99565b610380565b005b3480156101af57600080fd5b5061016e6101be366004610dbb565b61056e565b3480156101cf57600080fd5b5061016e60025481565b3480156101e557600080fd5b506004546101f9906001600160a01b031681565b6040516001600160a01b039091168152602001610178565b34801561021d57600080fd5b5061016e61022c366004610dbb565b610713565b34801561023d57600080fd5b504761016e565b34801561025057600080fd5b506101a161075a565b34801561026557600080fd5b5061016e610274366004610dbb565b610790565b34801561028557600080fd5b5061016e60065481565b34801561029b57600080fd5b506001546001600160a01b03166101f9565b3480156102b957600080fd5b506101a16102c8366004610dd4565b6107bd565b3480156102d957600080fd5b5061016e6102e8366004610dbb565b60056020526000908152604090205481565b61016e610882565b34801561030e57600080fd5b506101a161031d366004610dd4565b6109fb565b34801561032e57600080fd5b5061016e61033d366004610dd4565b6001600160a01b031660009081526003602052604090205490565b6101a1610366366004610dbb565b610a96565b60006006544761037b9190610e1a565b905090565b6004546001600160a01b031633146103d35760405162461bcd60e51b815260206004820152601660248201527513db9b1e481bdc195c985d1bdc8818d85b8818d85b1b60521b604482015260640161014b565b6002600054036103f55760405162461bcd60e51b815260040161014b90610e2d565b60026000558061043c5760405162461bcd60e51b81526020600482015260126024820152710416d6f756e74206d757374206265203e20360741b604482015260640161014b565b6000828152600560205260409020546001116104935760405162461bcd60e51b815260206004820152601660248201527510da185b9b995b08185b1c9958591e48199d5b99195960521b604482015260640161014b565b478111156104e35760405162461bcd60e51b815260206004820152601b60248201527f496e73756666696369656e742066726565206c69717569646974790000000000604482015260640161014b565b600082815260056020526040812082905560068054839290610506908490610e1a565b9091555050600454610521906001600160a01b031682610c29565b6004546040518281526001600160a01b039091169083907f174dd9cbed5b79ced30198029c21a27b4721fd0f4dd46ac813de82798e9098189060200160405180910390a350506001600055565b60006002600054036105925760405162461bcd60e51b815260040161014b90610e2d565b6002600055816105d95760405162461bcd60e51b81526020600482015260126024820152710536861726573206d757374206265203e20360741b604482015260640161014b565b336000908152600360205260409020548281101561062f5760405162461bcd60e51b8152602060048201526013602482015272496e73756666696369656e742073686172657360681b604482015260640161014b565b60025461063c4785610e64565b6106469190610e7b565b91506000821161068e5760405162461bcd60e51b81526020600482015260136024820152720576974686472617720616d6f756e74203d203606c1b604482015260640161014b565b6106988382610e9d565b33600090815260036020526040812091909155600280548592906106bd908490610e9d565b909155506106cd90503383610c29565b604080518381526020810185905233917f92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6910160405180910390a2506001600055919050565b6000600182101561072657506000919050565b60016002541015610735575090565b61073d61036b565b60025461074a9084610e64565b6107549190610e7b565b92915050565b6001546001600160a01b031633146107845760405162461bcd60e51b815260040161014b90610eb0565b61078e6000610d47565b565b600060018210806107a357506001600254105b156107b057506000919050565b60025461074a4784610e64565b6001546001600160a01b031633146107e75760405162461bcd60e51b815260040161014b90610eb0565b6001600160a01b0381166108305760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b21037b832b930ba37b960811b604482015260640161014b565b600480546001600160a01b038381166001600160a01b0319831681179093556040519116919082907ffbe5b6cbafb274f445d7fed869dc77a838d8243a22c460de156560e8857cad0390600090a35050565b60006002600054036108a65760405162461bcd60e51b815260040161014b90610e2d565b600260005534806108ef5760405162461bcd60e51b815260206004820152601360248201527204465706f736974206d757374206265203e203606c1b604482015260640161014b565b6001600254101561090257809150610936565b60008161090d61036b565b6109179190610e9d565b905080600254836109289190610e64565b6109329190610e7b565b9250505b6000821161097a5760405162461bcd60e51b815260206004820152601160248201527004d696e74656420736861726573203d203607c1b604482015260640161014b565b3360009081526003602052604081208054849290610999908490610e1a565b9250508190555081600260008282546109b29190610e1a565b9091555050604080518281526020810184905233917f73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca910160405180910390a250600160005590565b6001546001600160a01b03163314610a255760405162461bcd60e51b815260040161014b90610eb0565b6001600160a01b038116610a8a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161014b565b610a9381610d47565b50565b6004546001600160a01b03163314610ae95760405162461bcd60e51b815260206004820152601660248201527513db9b1e481bdc195c985d1bdc8818d85b8818d85b1b60521b604482015260640161014b565b600260005403610b0b5760405162461bcd60e51b815260040161014b90610e2d565b600260009081558181526005602052604090205480610b5e5760405162461bcd60e51b815260206004820152600f60248201526e155b9adb9bdddb8818da185b9b995b608a1b604482015260640161014b565b80341015610bae5760405162461bcd60e51b815260206004820152601c60248201527f52657475726e6564204554482062656c6f77207072696e636970616c00000000604482015260640161014b565b8060066000828254610bc09190610e9d565b9091555050600082815260056020526040812055817fbb7a5a625951cb72f1e443bc7a3cd146dcdc284df4a5ed440993e5df803d2bc78234610c028282610e9d565b6040805193845260208401929092529082015260600160405180910390a250506001600055565b80471015610c795760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604482015260640161014b565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114610cc6576040519150601f19603f3d011682016040523d82523d6000602084013e610ccb565b606091505b5050905080610d425760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d61792068617665207265766572746564000000000000606482015260840161014b565b505050565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60008060408385031215610dac57600080fd5b50508035926020909101359150565b600060208284031215610dcd57600080fd5b5035919050565b600060208284031215610de657600080fd5b81356001600160a01b0381168114610dfd57600080fd5b9392505050565b634e487b7160e01b600052601160045260246000fd5b8082018082111561075457610754610e04565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b808202811582820484141761075457610754610e04565b600082610e9857634e487b7160e01b600052601260045260246000fd5b500490565b8181038181111561075457610754610e04565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260408201526060019056fea264697066735822122064bad81d53577447c729f6956d6b8e7b8e508c815b1a99c87344558b0071226b64736f6c63430008220033",
}

// LiquidityPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidityPoolMetaData.ABI instead.
var LiquidityPoolABI = LiquidityPoolMetaData.ABI

// LiquidityPoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LiquidityPoolMetaData.Bin instead.
var LiquidityPoolBin = LiquidityPoolMetaData.Bin

// DeployLiquidityPool deploys a new Ethereum contract, binding an instance of LiquidityPool to it.
func DeployLiquidityPool(auth *bind.TransactOpts, backend bind.ContractBackend, _operator common.Address) (common.Address, *types.Transaction, *LiquidityPool, error) {
	parsed, err := LiquidityPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LiquidityPoolBin), backend, _operator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LiquidityPool{LiquidityPoolCaller: LiquidityPoolCaller{contract: contract}, LiquidityPoolTransactor: LiquidityPoolTransactor{contract: contract}, LiquidityPoolFilterer: LiquidityPoolFilterer{contract: contract}}, nil
}

// LiquidityPool is an auto generated Go binding around an Ethereum contract.
type LiquidityPool struct {
	LiquidityPoolCaller     // Read-only binding to the contract
	LiquidityPoolTransactor // Write-only binding to the contract
	LiquidityPoolFilterer   // Log filterer for contract events
}

// LiquidityPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidityPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidityPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidityPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidityPoolSession struct {
	Contract     *LiquidityPool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidityPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidityPoolCallerSession struct {
	Contract *LiquidityPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LiquidityPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidityPoolTransactorSession struct {
	Contract     *LiquidityPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LiquidityPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidityPoolRaw struct {
	Contract *LiquidityPool // Generic contract binding to access the raw methods on
}

// LiquidityPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidityPoolCallerRaw struct {
	Contract *LiquidityPoolCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidityPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidityPoolTransactorRaw struct {
	Contract *LiquidityPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidityPool creates a new instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPool(address common.Address, backend bind.ContractBackend) (*LiquidityPool, error) {
	contract, err := bindLiquidityPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidityPool{LiquidityPoolCaller: LiquidityPoolCaller{contract: contract}, LiquidityPoolTransactor: LiquidityPoolTransactor{contract: contract}, LiquidityPoolFilterer: LiquidityPoolFilterer{contract: contract}}, nil
}

// NewLiquidityPoolCaller creates a new read-only instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPoolCaller(address common.Address, caller bind.ContractCaller) (*LiquidityPoolCaller, error) {
	contract, err := bindLiquidityPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolCaller{contract: contract}, nil
}

// NewLiquidityPoolTransactor creates a new write-only instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidityPoolTransactor, error) {
	contract, err := bindLiquidityPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolTransactor{contract: contract}, nil
}

// NewLiquidityPoolFilterer creates a new log filterer instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidityPoolFilterer, error) {
	contract, err := bindLiquidityPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolFilterer{contract: contract}, nil
}

// bindLiquidityPool binds a generic wrapper to an already deployed contract.
func bindLiquidityPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquidityPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityPool *LiquidityPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidityPool.Contract.LiquidityPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityPool *LiquidityPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.Contract.LiquidityPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityPool *LiquidityPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityPool.Contract.LiquidityPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityPool *LiquidityPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidityPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityPool *LiquidityPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityPool *LiquidityPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityPool.Contract.contract.Transact(opts, method, params...)
}

// LockedByChannel is a free data retrieval call binding the contract method 0xb6cc49da.
//
// Solidity: function lockedByChannel(bytes32 ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) LockedByChannel(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "lockedByChannel", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedByChannel is a free data retrieval call binding the contract method 0xb6cc49da.
//
// Solidity: function lockedByChannel(bytes32 ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) LockedByChannel(arg0 [32]byte) (*big.Int, error) {
	return _LiquidityPool.Contract.LockedByChannel(&_LiquidityPool.CallOpts, arg0)
}

// LockedByChannel is a free data retrieval call binding the contract method 0xb6cc49da.
//
// Solidity: function lockedByChannel(bytes32 ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) LockedByChannel(arg0 [32]byte) (*big.Int, error) {
	return _LiquidityPool.Contract.LockedByChannel(&_LiquidityPool.CallOpts, arg0)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_LiquidityPool *LiquidityPoolCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_LiquidityPool *LiquidityPoolSession) Operator() (common.Address, error) {
	return _LiquidityPool.Contract.Operator(&_LiquidityPool.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_LiquidityPool *LiquidityPoolCallerSession) Operator() (common.Address, error) {
	return _LiquidityPool.Contract.Operator(&_LiquidityPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityPool *LiquidityPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityPool *LiquidityPoolSession) Owner() (common.Address, error) {
	return _LiquidityPool.Contract.Owner(&_LiquidityPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityPool *LiquidityPoolCallerSession) Owner() (common.Address, error) {
	return _LiquidityPool.Contract.Owner(&_LiquidityPool.CallOpts)
}

// PreviewDepositShares is a free data retrieval call binding the contract method 0x6574bcd4.
//
// Solidity: function previewDepositShares(uint256 ethAmount) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) PreviewDepositShares(opts *bind.CallOpts, ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "previewDepositShares", ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDepositShares is a free data retrieval call binding the contract method 0x6574bcd4.
//
// Solidity: function previewDepositShares(uint256 ethAmount) view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) PreviewDepositShares(ethAmount *big.Int) (*big.Int, error) {
	return _LiquidityPool.Contract.PreviewDepositShares(&_LiquidityPool.CallOpts, ethAmount)
}

// PreviewDepositShares is a free data retrieval call binding the contract method 0x6574bcd4.
//
// Solidity: function previewDepositShares(uint256 ethAmount) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) PreviewDepositShares(ethAmount *big.Int) (*big.Int, error) {
	return _LiquidityPool.Contract.PreviewDepositShares(&_LiquidityPool.CallOpts, ethAmount)
}

// PreviewWithdrawETH is a free data retrieval call binding the contract method 0x78f3bfe2.
//
// Solidity: function previewWithdrawETH(uint256 sharesAmount) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) PreviewWithdrawETH(opts *bind.CallOpts, sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "previewWithdrawETH", sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdrawETH is a free data retrieval call binding the contract method 0x78f3bfe2.
//
// Solidity: function previewWithdrawETH(uint256 sharesAmount) view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) PreviewWithdrawETH(sharesAmount *big.Int) (*big.Int, error) {
	return _LiquidityPool.Contract.PreviewWithdrawETH(&_LiquidityPool.CallOpts, sharesAmount)
}

// PreviewWithdrawETH is a free data retrieval call binding the contract method 0x78f3bfe2.
//
// Solidity: function previewWithdrawETH(uint256 sharesAmount) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) PreviewWithdrawETH(sharesAmount *big.Int) (*big.Int, error) {
	return _LiquidityPool.Contract.PreviewWithdrawETH(&_LiquidityPool.CallOpts, sharesAmount)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address provider) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) SharesOf(opts *bind.CallOpts, provider common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "sharesOf", provider)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address provider) view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) SharesOf(provider common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.SharesOf(&_LiquidityPool.CallOpts, provider)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address provider) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) SharesOf(provider common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.SharesOf(&_LiquidityPool.CallOpts, provider)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) TotalAssets() (*big.Int, error) {
	return _LiquidityPool.Contract.TotalAssets(&_LiquidityPool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) TotalAssets() (*big.Int, error) {
	return _LiquidityPool.Contract.TotalAssets(&_LiquidityPool.CallOpts)
}

// TotalLockedETH is a free data retrieval call binding the contract method 0x8afe47cf.
//
// Solidity: function totalLockedETH() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) TotalLockedETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "totalLockedETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLockedETH is a free data retrieval call binding the contract method 0x8afe47cf.
//
// Solidity: function totalLockedETH() view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) TotalLockedETH() (*big.Int, error) {
	return _LiquidityPool.Contract.TotalLockedETH(&_LiquidityPool.CallOpts)
}

// TotalLockedETH is a free data retrieval call binding the contract method 0x8afe47cf.
//
// Solidity: function totalLockedETH() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) TotalLockedETH() (*big.Int, error) {
	return _LiquidityPool.Contract.TotalLockedETH(&_LiquidityPool.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) TotalShares() (*big.Int, error) {
	return _LiquidityPool.Contract.TotalShares(&_LiquidityPool.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) TotalShares() (*big.Int, error) {
	return _LiquidityPool.Contract.TotalShares(&_LiquidityPool.CallOpts)
}

// WithdrawableETH is a free data retrieval call binding the contract method 0x6913e7cc.
//
// Solidity: function withdrawableETH() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) WithdrawableETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "withdrawableETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableETH is a free data retrieval call binding the contract method 0x6913e7cc.
//
// Solidity: function withdrawableETH() view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) WithdrawableETH() (*big.Int, error) {
	return _LiquidityPool.Contract.WithdrawableETH(&_LiquidityPool.CallOpts)
}

// WithdrawableETH is a free data retrieval call binding the contract method 0x6913e7cc.
//
// Solidity: function withdrawableETH() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) WithdrawableETH() (*big.Int, error) {
	return _LiquidityPool.Contract.WithdrawableETH(&_LiquidityPool.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256 mintedShares)
func (_LiquidityPool *LiquidityPoolTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256 mintedShares)
func (_LiquidityPool *LiquidityPoolSession) Deposit() (*types.Transaction, error) {
	return _LiquidityPool.Contract.Deposit(&_LiquidityPool.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256 mintedShares)
func (_LiquidityPool *LiquidityPoolTransactorSession) Deposit() (*types.Transaction, error) {
	return _LiquidityPool.Contract.Deposit(&_LiquidityPool.TransactOpts)
}

// FundChannel is a paid mutator transaction binding the contract method 0x260b6086.
//
// Solidity: function fundChannel(bytes32 channelId, uint256 amount) returns()
func (_LiquidityPool *LiquidityPoolTransactor) FundChannel(opts *bind.TransactOpts, channelId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "fundChannel", channelId, amount)
}

// FundChannel is a paid mutator transaction binding the contract method 0x260b6086.
//
// Solidity: function fundChannel(bytes32 channelId, uint256 amount) returns()
func (_LiquidityPool *LiquidityPoolSession) FundChannel(channelId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.FundChannel(&_LiquidityPool.TransactOpts, channelId, amount)
}

// FundChannel is a paid mutator transaction binding the contract method 0x260b6086.
//
// Solidity: function fundChannel(bytes32 channelId, uint256 amount) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) FundChannel(channelId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.FundChannel(&_LiquidityPool.TransactOpts, channelId, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityPool *LiquidityPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityPool *LiquidityPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidityPool.Contract.RenounceOwnership(&_LiquidityPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidityPool.Contract.RenounceOwnership(&_LiquidityPool.TransactOpts)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address newOperator) returns()
func (_LiquidityPool *LiquidityPoolTransactor) SetOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "setOperator", newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address newOperator) returns()
func (_LiquidityPool *LiquidityPoolSession) SetOperator(newOperator common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SetOperator(&_LiquidityPool.TransactOpts, newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address newOperator) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) SetOperator(newOperator common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SetOperator(&_LiquidityPool.TransactOpts, newOperator)
}

// SettleChannel is a paid mutator transaction binding the contract method 0xfd3d3199.
//
// Solidity: function settleChannel(bytes32 channelId) payable returns()
func (_LiquidityPool *LiquidityPoolTransactor) SettleChannel(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "settleChannel", channelId)
}

// SettleChannel is a paid mutator transaction binding the contract method 0xfd3d3199.
//
// Solidity: function settleChannel(bytes32 channelId) payable returns()
func (_LiquidityPool *LiquidityPoolSession) SettleChannel(channelId [32]byte) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SettleChannel(&_LiquidityPool.TransactOpts, channelId)
}

// SettleChannel is a paid mutator transaction binding the contract method 0xfd3d3199.
//
// Solidity: function settleChannel(bytes32 channelId) payable returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) SettleChannel(channelId [32]byte) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SettleChannel(&_LiquidityPool.TransactOpts, channelId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityPool *LiquidityPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityPool *LiquidityPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.TransferOwnership(&_LiquidityPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.TransferOwnership(&_LiquidityPool.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 sharesToBurn) returns(uint256 ethAmountOut)
func (_LiquidityPool *LiquidityPoolTransactor) Withdraw(opts *bind.TransactOpts, sharesToBurn *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "withdraw", sharesToBurn)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 sharesToBurn) returns(uint256 ethAmountOut)
func (_LiquidityPool *LiquidityPoolSession) Withdraw(sharesToBurn *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Withdraw(&_LiquidityPool.TransactOpts, sharesToBurn)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 sharesToBurn) returns(uint256 ethAmountOut)
func (_LiquidityPool *LiquidityPoolTransactorSession) Withdraw(sharesToBurn *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Withdraw(&_LiquidityPool.TransactOpts, sharesToBurn)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityPool *LiquidityPoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityPool *LiquidityPoolSession) Receive() (*types.Transaction, error) {
	return _LiquidityPool.Contract.Receive(&_LiquidityPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) Receive() (*types.Transaction, error) {
	return _LiquidityPool.Contract.Receive(&_LiquidityPool.TransactOpts)
}

// LiquidityPoolChannelFundedIterator is returned from FilterChannelFunded and is used to iterate over the raw logs and unpacked data for ChannelFunded events raised by the LiquidityPool contract.
type LiquidityPoolChannelFundedIterator struct {
	Event *LiquidityPoolChannelFunded // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolChannelFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolChannelFunded)
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
		it.Event = new(LiquidityPoolChannelFunded)
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
func (it *LiquidityPoolChannelFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolChannelFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolChannelFunded represents a ChannelFunded event raised by the LiquidityPool contract.
type LiquidityPoolChannelFunded struct {
	ChannelId [32]byte
	Principal *big.Int
	Operator  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelFunded is a free log retrieval operation binding the contract event 0x174dd9cbed5b79ced30198029c21a27b4721fd0f4dd46ac813de82798e909818.
//
// Solidity: event ChannelFunded(bytes32 indexed channelId, uint256 principal, address indexed operator)
func (_LiquidityPool *LiquidityPoolFilterer) FilterChannelFunded(opts *bind.FilterOpts, channelId [][32]byte, operator []common.Address) (*LiquidityPoolChannelFundedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "ChannelFunded", channelIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolChannelFundedIterator{contract: _LiquidityPool.contract, event: "ChannelFunded", logs: logs, sub: sub}, nil
}

// WatchChannelFunded is a free log subscription operation binding the contract event 0x174dd9cbed5b79ced30198029c21a27b4721fd0f4dd46ac813de82798e909818.
//
// Solidity: event ChannelFunded(bytes32 indexed channelId, uint256 principal, address indexed operator)
func (_LiquidityPool *LiquidityPoolFilterer) WatchChannelFunded(opts *bind.WatchOpts, sink chan<- *LiquidityPoolChannelFunded, channelId [][32]byte, operator []common.Address) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "ChannelFunded", channelIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolChannelFunded)
				if err := _LiquidityPool.contract.UnpackLog(event, "ChannelFunded", log); err != nil {
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

// ParseChannelFunded is a log parse operation binding the contract event 0x174dd9cbed5b79ced30198029c21a27b4721fd0f4dd46ac813de82798e909818.
//
// Solidity: event ChannelFunded(bytes32 indexed channelId, uint256 principal, address indexed operator)
func (_LiquidityPool *LiquidityPoolFilterer) ParseChannelFunded(log types.Log) (*LiquidityPoolChannelFunded, error) {
	event := new(LiquidityPoolChannelFunded)
	if err := _LiquidityPool.contract.UnpackLog(event, "ChannelFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolChannelSettledIterator is returned from FilterChannelSettled and is used to iterate over the raw logs and unpacked data for ChannelSettled events raised by the LiquidityPool contract.
type LiquidityPoolChannelSettledIterator struct {
	Event *LiquidityPoolChannelSettled // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolChannelSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolChannelSettled)
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
		it.Event = new(LiquidityPoolChannelSettled)
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
func (it *LiquidityPoolChannelSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolChannelSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolChannelSettled represents a ChannelSettled event raised by the LiquidityPool contract.
type LiquidityPoolChannelSettled struct {
	ChannelId     [32]byte
	Principal     *big.Int
	TotalReturned *big.Int
	FeeGain       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChannelSettled is a free log retrieval operation binding the contract event 0xbb7a5a625951cb72f1e443bc7a3cd146dcdc284df4a5ed440993e5df803d2bc7.
//
// Solidity: event ChannelSettled(bytes32 indexed channelId, uint256 principal, uint256 totalReturned, uint256 feeGain)
func (_LiquidityPool *LiquidityPoolFilterer) FilterChannelSettled(opts *bind.FilterOpts, channelId [][32]byte) (*LiquidityPoolChannelSettledIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "ChannelSettled", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolChannelSettledIterator{contract: _LiquidityPool.contract, event: "ChannelSettled", logs: logs, sub: sub}, nil
}

// WatchChannelSettled is a free log subscription operation binding the contract event 0xbb7a5a625951cb72f1e443bc7a3cd146dcdc284df4a5ed440993e5df803d2bc7.
//
// Solidity: event ChannelSettled(bytes32 indexed channelId, uint256 principal, uint256 totalReturned, uint256 feeGain)
func (_LiquidityPool *LiquidityPoolFilterer) WatchChannelSettled(opts *bind.WatchOpts, sink chan<- *LiquidityPoolChannelSettled, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "ChannelSettled", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolChannelSettled)
				if err := _LiquidityPool.contract.UnpackLog(event, "ChannelSettled", log); err != nil {
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

// ParseChannelSettled is a log parse operation binding the contract event 0xbb7a5a625951cb72f1e443bc7a3cd146dcdc284df4a5ed440993e5df803d2bc7.
//
// Solidity: event ChannelSettled(bytes32 indexed channelId, uint256 principal, uint256 totalReturned, uint256 feeGain)
func (_LiquidityPool *LiquidityPoolFilterer) ParseChannelSettled(log types.Log) (*LiquidityPoolChannelSettled, error) {
	event := new(LiquidityPoolChannelSettled)
	if err := _LiquidityPool.contract.UnpackLog(event, "ChannelSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the LiquidityPool contract.
type LiquidityPoolDepositedIterator struct {
	Event *LiquidityPoolDeposited // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolDeposited)
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
		it.Event = new(LiquidityPoolDeposited)
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
func (it *LiquidityPoolDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolDeposited represents a Deposited event raised by the LiquidityPool contract.
type LiquidityPoolDeposited struct {
	Provider     common.Address
	EthAmount    *big.Int
	SharesMinted *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed provider, uint256 ethAmount, uint256 sharesMinted)
func (_LiquidityPool *LiquidityPoolFilterer) FilterDeposited(opts *bind.FilterOpts, provider []common.Address) (*LiquidityPoolDepositedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Deposited", providerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolDepositedIterator{contract: _LiquidityPool.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed provider, uint256 ethAmount, uint256 sharesMinted)
func (_LiquidityPool *LiquidityPoolFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *LiquidityPoolDeposited, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Deposited", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolDeposited)
				if err := _LiquidityPool.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed provider, uint256 ethAmount, uint256 sharesMinted)
func (_LiquidityPool *LiquidityPoolFilterer) ParseDeposited(log types.Log) (*LiquidityPoolDeposited, error) {
	event := new(LiquidityPoolDeposited)
	if err := _LiquidityPool.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolOperatorUpdatedIterator is returned from FilterOperatorUpdated and is used to iterate over the raw logs and unpacked data for OperatorUpdated events raised by the LiquidityPool contract.
type LiquidityPoolOperatorUpdatedIterator struct {
	Event *LiquidityPoolOperatorUpdated // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolOperatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolOperatorUpdated)
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
		it.Event = new(LiquidityPoolOperatorUpdated)
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
func (it *LiquidityPoolOperatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolOperatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolOperatorUpdated represents a OperatorUpdated event raised by the LiquidityPool contract.
type LiquidityPoolOperatorUpdated struct {
	PreviousOperator common.Address
	NewOperator      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOperatorUpdated is a free log retrieval operation binding the contract event 0xfbe5b6cbafb274f445d7fed869dc77a838d8243a22c460de156560e8857cad03.
//
// Solidity: event OperatorUpdated(address indexed previousOperator, address indexed newOperator)
func (_LiquidityPool *LiquidityPoolFilterer) FilterOperatorUpdated(opts *bind.FilterOpts, previousOperator []common.Address, newOperator []common.Address) (*LiquidityPoolOperatorUpdatedIterator, error) {

	var previousOperatorRule []interface{}
	for _, previousOperatorItem := range previousOperator {
		previousOperatorRule = append(previousOperatorRule, previousOperatorItem)
	}
	var newOperatorRule []interface{}
	for _, newOperatorItem := range newOperator {
		newOperatorRule = append(newOperatorRule, newOperatorItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "OperatorUpdated", previousOperatorRule, newOperatorRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolOperatorUpdatedIterator{contract: _LiquidityPool.contract, event: "OperatorUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorUpdated is a free log subscription operation binding the contract event 0xfbe5b6cbafb274f445d7fed869dc77a838d8243a22c460de156560e8857cad03.
//
// Solidity: event OperatorUpdated(address indexed previousOperator, address indexed newOperator)
func (_LiquidityPool *LiquidityPoolFilterer) WatchOperatorUpdated(opts *bind.WatchOpts, sink chan<- *LiquidityPoolOperatorUpdated, previousOperator []common.Address, newOperator []common.Address) (event.Subscription, error) {

	var previousOperatorRule []interface{}
	for _, previousOperatorItem := range previousOperator {
		previousOperatorRule = append(previousOperatorRule, previousOperatorItem)
	}
	var newOperatorRule []interface{}
	for _, newOperatorItem := range newOperator {
		newOperatorRule = append(newOperatorRule, newOperatorItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "OperatorUpdated", previousOperatorRule, newOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolOperatorUpdated)
				if err := _LiquidityPool.contract.UnpackLog(event, "OperatorUpdated", log); err != nil {
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

// ParseOperatorUpdated is a log parse operation binding the contract event 0xfbe5b6cbafb274f445d7fed869dc77a838d8243a22c460de156560e8857cad03.
//
// Solidity: event OperatorUpdated(address indexed previousOperator, address indexed newOperator)
func (_LiquidityPool *LiquidityPoolFilterer) ParseOperatorUpdated(log types.Log) (*LiquidityPoolOperatorUpdated, error) {
	event := new(LiquidityPoolOperatorUpdated)
	if err := _LiquidityPool.contract.UnpackLog(event, "OperatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LiquidityPool contract.
type LiquidityPoolOwnershipTransferredIterator struct {
	Event *LiquidityPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolOwnershipTransferred)
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
		it.Event = new(LiquidityPoolOwnershipTransferred)
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
func (it *LiquidityPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolOwnershipTransferred represents a OwnershipTransferred event raised by the LiquidityPool contract.
type LiquidityPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidityPool *LiquidityPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LiquidityPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolOwnershipTransferredIterator{contract: _LiquidityPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidityPool *LiquidityPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LiquidityPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolOwnershipTransferred)
				if err := _LiquidityPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidityPool *LiquidityPoolFilterer) ParseOwnershipTransferred(log types.Log) (*LiquidityPoolOwnershipTransferred, error) {
	event := new(LiquidityPoolOwnershipTransferred)
	if err := _LiquidityPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the LiquidityPool contract.
type LiquidityPoolWithdrawnIterator struct {
	Event *LiquidityPoolWithdrawn // Event containing the contract specifics and raw log

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
func (it *LiquidityPoolWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolWithdrawn)
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
		it.Event = new(LiquidityPoolWithdrawn)
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
func (it *LiquidityPoolWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolWithdrawn represents a Withdrawn event raised by the LiquidityPool contract.
type LiquidityPoolWithdrawn struct {
	Provider     common.Address
	EthAmount    *big.Int
	SharesBurned *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed provider, uint256 ethAmount, uint256 sharesBurned)
func (_LiquidityPool *LiquidityPoolFilterer) FilterWithdrawn(opts *bind.FilterOpts, provider []common.Address) (*LiquidityPoolWithdrawnIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Withdrawn", providerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolWithdrawnIterator{contract: _LiquidityPool.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed provider, uint256 ethAmount, uint256 sharesBurned)
func (_LiquidityPool *LiquidityPoolFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *LiquidityPoolWithdrawn, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Withdrawn", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolWithdrawn)
				if err := _LiquidityPool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed provider, uint256 ethAmount, uint256 sharesBurned)
func (_LiquidityPool *LiquidityPoolFilterer) ParseWithdrawn(log types.Log) (*LiquidityPoolWithdrawn, error) {
	event := new(LiquidityPoolWithdrawn)
	if err := _LiquidityPool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
