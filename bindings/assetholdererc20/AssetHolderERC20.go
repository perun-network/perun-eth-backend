// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package assetholdererc20

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

// AssetHolderWithdrawalAuth is an auto generated low-level Go binding around an user-defined struct.
type AssetHolderWithdrawalAuth struct {
	ChannelID   [32]byte
	Participant ChannelParticipant
	Receiver    common.Address
	Amount      *big.Int
}

// ChannelParticipant is an auto generated low-level Go binding around an user-defined struct.
type ChannelParticipant struct {
	EthAddress common.Address
	CcAddress  []byte
}

// Assetholdererc20MetaData contains all meta data concerning the Assetholdererc20 contract.
var Assetholdererc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adjudicator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"}],\"name\":\"OutcomeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adjudicator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"holdings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccAddress\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Participant[]\",\"name\":\"parts\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newBals\",\"type\":\"uint256[]\"}],\"name\":\"setOutcome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"settled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"ccAddress\",\"type\":\"bytes\"}],\"internalType\":\"structChannel.Participant\",\"name\":\"participant\",\"type\":\"tuple\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAssetHolder.WithdrawalAuth\",\"name\":\"authorization\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161129d38038061129d83398101604081905261002f91610070565b600280546001600160a01b0319166001600160a01b03938416179055166080526100a3565b80516001600160a01b038116811461006b57600080fd5b919050565b6000806040838503121561008357600080fd5b61008c83610054565b915061009a60208401610054565b90509250929050565b6080516111d16100cc600039600081816101740152818161080301526109a301526111d16000f3fe6080604052600436106100705760003560e01c8063ae9ee18c1161004e578063ae9ee18c146100e7578063d945af1d14610122578063fc0c546a14610162578063fca0f7781461019657600080fd5b80631de26e1614610075578063295482ce1461008a57806353c2ed8e146100aa575b600080fd5b610088610083366004610dfb565b6101b6565b005b34801561009657600080fd5b506100886100a5366004610e62565b610230565b3480156100b657600080fd5b506002546100ca906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156100f357600080fd5b50610114610102366004610edc565b60006020819052908152604090205481565b6040519081526020016100de565b34801561012e57600080fd5b5061015261013d366004610edc565b60016020526000908152604090205460ff1681565b60405190151581526020016100de565b34801561016e57600080fd5b506100ca7f000000000000000000000000000000000000000000000000000000000000000081565b3480156101a257600080fd5b506100886101b1366004610ef5565b610554565b6101c0828261076a565b6000828152602081905260409020546101d990826107ce565b6000838152602081905260409020556101f282826107e1565b817fcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a98260405161022491815260200190565b60405180910390a25050565b6002546001600160a01b0316331461029d5760405162461bcd60e51b815260206004820152602560248201527f63616e206f6e6c792062652063616c6c6564206279207468652061646a75646960448201526431b0ba37b960d91b60648201526084015b60405180910390fd5b8281146102fe5760405162461bcd60e51b815260206004820152602960248201527f7061727469636970616e7473206c656e6774682073686f756c6420657175616c6044820152682062616c616e63657360b81b6064820152608401610294565b60008581526001602052604090205460ff161561036b5760405162461bcd60e51b815260206004820152602560248201527f747279696e6720746f2073657420616c726561647920736574746c6564206368604482015264185b9b995b60da1b6064820152608401610294565b600085815260208190526040812080549082905590808567ffffffffffffffff81111561039a5761039a610f94565b6040519080825280602002602001820160405280156103c3578160200160208202803683370190505b50905060005b8681101561049757600061040e8a8a8a858181106103e9576103e9610faa565b90506020028101906103fb9190610fc0565b610409906020810190610ff5565b6108ba565b90508083838151811061042357610423610faa565b60200260200101818152505061045460008083815260200190815260200160002054866107ce90919063ffffffff16565b945061048187878481811061046b5761046b610faa565b90506020020135856107ce90919063ffffffff16565b935050808061048f90611028565b9150506103c9565b508183106105075760005b86811015610505578585828181106104bc576104bc610faa565b905060200201356000808484815181106104d8576104d8610faa565b602002602001015181526020019081526020016000208190555080806104fd90611028565b9150506104a2565b505b6000888152600160208190526040808320805460ff19169092179091555189917fef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b891a25050505050505050565b823560009081526001602052604090205460ff166105aa5760405162461bcd60e51b815260206004820152601360248201527218da185b9b995b081b9bdd081cd95d1d1b1959606a1b6044820152606401610294565b61061d836040516020016105be9190611051565b60408051601f198184030181526020601f86018190048102840181019092528483529190859085908190840183828082843760009201919091525061060a925050506020870187610fc0565b610618906020810190610ff5565b6108ff565b6106695760405162461bcd60e51b815260206004820152601d60248201527f7369676e617475726520766572696669636174696f6e206661696c65640000006044820152606401610294565b600061067d84356103fb6020870187610fc0565b600081815260208190526040902054909150606085013511156106d75760405162461bcd60e51b8152602060048201526012602482015271696e73756666696369656e742066756e647360701b6044820152606401610294565b6000818152602081905260409020546106f490606086013561098d565b60008281526020819052604090205561070e848484610999565b807fd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81606086018035906107449060408901610ff5565b604080519283526001600160a01b0390911660208301520160405180910390a250505050565b34156107ca5760405162461bcd60e51b815260206004820152602960248201527f6d6573736167652076616c7565206d757374206265203020666f7220746f6b656044820152681b8819195c1bdcda5d60ba1b6064820152608401610294565b5050565b60006107da8284611134565b9392505050565b6040516323b872dd60e01b8152336004820152306024820152604481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906323b872dd906064016020604051808303816000875af1158015610854573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610878919061114c565b6107ca5760405162461bcd60e51b81526020600482015260136024820152721d1c985b9cd9995c919c9bdb4819985a5b1959606a1b6044820152606401610294565b600082826040516020016108e19291909182526001600160a01b0316602082015260400190565b60405160208183030381529060405280519060200120905092915050565b60008061096085805190602001206040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050600061096e8286610a8a565b6001600160a01b0390811690851614925050509392505050565b505050565b60006107da828461116e565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663a9059cbb6109d86060860160408701610ff5565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152606086013560248201526044016020604051808303816000875af1158015610a28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4c919061114c565b6109885760405162461bcd60e51b815260206004820152600f60248201526e1d1c985b9cd9995c8819985a5b1959608a1b6044820152606401610294565b6000806000610a998585610aae565b91509150610aa681610b1c565b509392505050565b6000808251604103610ae45760208301516040840151606085015160001a610ad887828585610cd5565b94509450505050610b15565b8251604003610b0d5760208301516040840151610b02868383610dc2565b935093505050610b15565b506000905060025b9250929050565b6000816004811115610b3057610b30611185565b03610b385750565b6001816004811115610b4c57610b4c611185565b03610b995760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610294565b6002816004811115610bad57610bad611185565b03610bfa5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610294565b6003816004811115610c0e57610c0e611185565b03610c665760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610294565b6004816004811115610c7a57610c7a611185565b03610cd25760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610294565b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610d0c5750600090506003610db9565b8460ff16601b14158015610d2457508460ff16601c14155b15610d355750600090506004610db9565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610d89573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610db257600060019250925050610db9565b9150600090505b94509492505050565b6000806001600160ff1b03831681610ddf60ff86901c601b611134565b9050610ded87828885610cd5565b935093505050935093915050565b60008060408385031215610e0e57600080fd5b50508035926020909101359150565b60008083601f840112610e2f57600080fd5b50813567ffffffffffffffff811115610e4757600080fd5b6020830191508360208260051b8501011115610b1557600080fd5b600080600080600060608688031215610e7a57600080fd5b85359450602086013567ffffffffffffffff80821115610e9957600080fd5b610ea589838a01610e1d565b90965094506040880135915080821115610ebe57600080fd5b50610ecb88828901610e1d565b969995985093965092949392505050565b600060208284031215610eee57600080fd5b5035919050565b600080600060408486031215610f0a57600080fd5b833567ffffffffffffffff80821115610f2257600080fd5b9085019060808288031215610f3657600080fd5b90935060208501359080821115610f4c57600080fd5b818601915086601f830112610f6057600080fd5b813581811115610f6f57600080fd5b876020828501011115610f8157600080fd5b6020830194508093505050509250925092565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60008235603e19833603018112610fd657600080fd5b9190910192915050565b6001600160a01b0381168114610cd257600080fd5b60006020828403121561100757600080fd5b81356107da81610fe0565b634e487b7160e01b600052601160045260246000fd5b60006001820161103a5761103a611012565b5060010190565b803561104c81610fe0565b919050565b602081528135602082015260006020830135603e1984360301811261107557600080fd5b608060408401528301803561108981610fe0565b6001600160a01b031660a0840152602081013536829003601e190181126110af57600080fd5b0160208101903567ffffffffffffffff8111156110cb57600080fd5b8036038213156110da57600080fd5b604060c08501528060e0850152610100818382870137600081838701015261110460408701611041565b6001600160a01b03811660608701529250606095909501356080850152601f01601f191690920190920192915050565b6000821982111561114757611147611012565b500190565b60006020828403121561115e57600080fd5b815180151581146107da57600080fd5b60008282101561118057611180611012565b500390565b634e487b7160e01b600052602160045260246000fdfea2646970667358221220ad0874027a67427bfc782242eb9ed3252b2e96dbb6f2b27da3f355e3da35f48564736f6c634300080f0033",
}

// Assetholdererc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Assetholdererc20MetaData.ABI instead.
var Assetholdererc20ABI = Assetholdererc20MetaData.ABI

// Assetholdererc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Assetholdererc20MetaData.Bin instead.
var Assetholdererc20Bin = Assetholdererc20MetaData.Bin

// DeployAssetholdererc20 deploys a new Ethereum contract, binding an instance of Assetholdererc20 to it.
func DeployAssetholdererc20(auth *bind.TransactOpts, backend bind.ContractBackend, _adjudicator common.Address, _token common.Address) (common.Address, *types.Transaction, *Assetholdererc20, error) {
	parsed, err := Assetholdererc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Assetholdererc20Bin), backend, _adjudicator, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Assetholdererc20{Assetholdererc20Caller: Assetholdererc20Caller{contract: contract}, Assetholdererc20Transactor: Assetholdererc20Transactor{contract: contract}, Assetholdererc20Filterer: Assetholdererc20Filterer{contract: contract}}, nil
}

// Assetholdererc20 is an auto generated Go binding around an Ethereum contract.
type Assetholdererc20 struct {
	Assetholdererc20Caller     // Read-only binding to the contract
	Assetholdererc20Transactor // Write-only binding to the contract
	Assetholdererc20Filterer   // Log filterer for contract events
}

// Assetholdererc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Assetholdererc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Assetholdererc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Assetholdererc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Assetholdererc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Assetholdererc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Assetholdererc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Assetholdererc20Session struct {
	Contract     *Assetholdererc20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Assetholdererc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Assetholdererc20CallerSession struct {
	Contract *Assetholdererc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Assetholdererc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Assetholdererc20TransactorSession struct {
	Contract     *Assetholdererc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Assetholdererc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Assetholdererc20Raw struct {
	Contract *Assetholdererc20 // Generic contract binding to access the raw methods on
}

// Assetholdererc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Assetholdererc20CallerRaw struct {
	Contract *Assetholdererc20Caller // Generic read-only contract binding to access the raw methods on
}

// Assetholdererc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Assetholdererc20TransactorRaw struct {
	Contract *Assetholdererc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetholdererc20 creates a new instance of Assetholdererc20, bound to a specific deployed contract.
func NewAssetholdererc20(address common.Address, backend bind.ContractBackend) (*Assetholdererc20, error) {
	contract, err := bindAssetholdererc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Assetholdererc20{Assetholdererc20Caller: Assetholdererc20Caller{contract: contract}, Assetholdererc20Transactor: Assetholdererc20Transactor{contract: contract}, Assetholdererc20Filterer: Assetholdererc20Filterer{contract: contract}}, nil
}

// NewAssetholdererc20Caller creates a new read-only instance of Assetholdererc20, bound to a specific deployed contract.
func NewAssetholdererc20Caller(address common.Address, caller bind.ContractCaller) (*Assetholdererc20Caller, error) {
	contract, err := bindAssetholdererc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Assetholdererc20Caller{contract: contract}, nil
}

// NewAssetholdererc20Transactor creates a new write-only instance of Assetholdererc20, bound to a specific deployed contract.
func NewAssetholdererc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Assetholdererc20Transactor, error) {
	contract, err := bindAssetholdererc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Assetholdererc20Transactor{contract: contract}, nil
}

// NewAssetholdererc20Filterer creates a new log filterer instance of Assetholdererc20, bound to a specific deployed contract.
func NewAssetholdererc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Assetholdererc20Filterer, error) {
	contract, err := bindAssetholdererc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Assetholdererc20Filterer{contract: contract}, nil
}

// bindAssetholdererc20 binds a generic wrapper to an already deployed contract.
func bindAssetholdererc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Assetholdererc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assetholdererc20 *Assetholdererc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Assetholdererc20.Contract.Assetholdererc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assetholdererc20 *Assetholdererc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.Assetholdererc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assetholdererc20 *Assetholdererc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.Assetholdererc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assetholdererc20 *Assetholdererc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Assetholdererc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assetholdererc20 *Assetholdererc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assetholdererc20 *Assetholdererc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.contract.Transact(opts, method, params...)
}

// Adjudicator is a free data retrieval call binding the contract method 0x53c2ed8e.
//
// Solidity: function adjudicator() view returns(address)
func (_Assetholdererc20 *Assetholdererc20Caller) Adjudicator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Assetholdererc20.contract.Call(opts, &out, "adjudicator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Adjudicator is a free data retrieval call binding the contract method 0x53c2ed8e.
//
// Solidity: function adjudicator() view returns(address)
func (_Assetholdererc20 *Assetholdererc20Session) Adjudicator() (common.Address, error) {
	return _Assetholdererc20.Contract.Adjudicator(&_Assetholdererc20.CallOpts)
}

// Adjudicator is a free data retrieval call binding the contract method 0x53c2ed8e.
//
// Solidity: function adjudicator() view returns(address)
func (_Assetholdererc20 *Assetholdererc20CallerSession) Adjudicator() (common.Address, error) {
	return _Assetholdererc20.Contract.Adjudicator(&_Assetholdererc20.CallOpts)
}

// Holdings is a free data retrieval call binding the contract method 0xae9ee18c.
//
// Solidity: function holdings(bytes32 ) view returns(uint256)
func (_Assetholdererc20 *Assetholdererc20Caller) Holdings(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Assetholdererc20.contract.Call(opts, &out, "holdings", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Holdings is a free data retrieval call binding the contract method 0xae9ee18c.
//
// Solidity: function holdings(bytes32 ) view returns(uint256)
func (_Assetholdererc20 *Assetholdererc20Session) Holdings(arg0 [32]byte) (*big.Int, error) {
	return _Assetholdererc20.Contract.Holdings(&_Assetholdererc20.CallOpts, arg0)
}

// Holdings is a free data retrieval call binding the contract method 0xae9ee18c.
//
// Solidity: function holdings(bytes32 ) view returns(uint256)
func (_Assetholdererc20 *Assetholdererc20CallerSession) Holdings(arg0 [32]byte) (*big.Int, error) {
	return _Assetholdererc20.Contract.Holdings(&_Assetholdererc20.CallOpts, arg0)
}

// Settled is a free data retrieval call binding the contract method 0xd945af1d.
//
// Solidity: function settled(bytes32 ) view returns(bool)
func (_Assetholdererc20 *Assetholdererc20Caller) Settled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Assetholdererc20.contract.Call(opts, &out, "settled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Settled is a free data retrieval call binding the contract method 0xd945af1d.
//
// Solidity: function settled(bytes32 ) view returns(bool)
func (_Assetholdererc20 *Assetholdererc20Session) Settled(arg0 [32]byte) (bool, error) {
	return _Assetholdererc20.Contract.Settled(&_Assetholdererc20.CallOpts, arg0)
}

// Settled is a free data retrieval call binding the contract method 0xd945af1d.
//
// Solidity: function settled(bytes32 ) view returns(bool)
func (_Assetholdererc20 *Assetholdererc20CallerSession) Settled(arg0 [32]byte) (bool, error) {
	return _Assetholdererc20.Contract.Settled(&_Assetholdererc20.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Assetholdererc20 *Assetholdererc20Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Assetholdererc20.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Assetholdererc20 *Assetholdererc20Session) Token() (common.Address, error) {
	return _Assetholdererc20.Contract.Token(&_Assetholdererc20.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Assetholdererc20 *Assetholdererc20CallerSession) Token() (common.Address, error) {
	return _Assetholdererc20.Contract.Token(&_Assetholdererc20.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x1de26e16.
//
// Solidity: function deposit(bytes32 fundingID, uint256 amount) payable returns()
func (_Assetholdererc20 *Assetholdererc20Transactor) Deposit(opts *bind.TransactOpts, fundingID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Assetholdererc20.contract.Transact(opts, "deposit", fundingID, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1de26e16.
//
// Solidity: function deposit(bytes32 fundingID, uint256 amount) payable returns()
func (_Assetholdererc20 *Assetholdererc20Session) Deposit(fundingID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.Deposit(&_Assetholdererc20.TransactOpts, fundingID, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1de26e16.
//
// Solidity: function deposit(bytes32 fundingID, uint256 amount) payable returns()
func (_Assetholdererc20 *Assetholdererc20TransactorSession) Deposit(fundingID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.Deposit(&_Assetholdererc20.TransactOpts, fundingID, amount)
}

// SetOutcome is a paid mutator transaction binding the contract method 0x295482ce.
//
// Solidity: function setOutcome(bytes32 channelID, (address,bytes)[] parts, uint256[] newBals) returns()
func (_Assetholdererc20 *Assetholdererc20Transactor) SetOutcome(opts *bind.TransactOpts, channelID [32]byte, parts []ChannelParticipant, newBals []*big.Int) (*types.Transaction, error) {
	return _Assetholdererc20.contract.Transact(opts, "setOutcome", channelID, parts, newBals)
}

// SetOutcome is a paid mutator transaction binding the contract method 0x295482ce.
//
// Solidity: function setOutcome(bytes32 channelID, (address,bytes)[] parts, uint256[] newBals) returns()
func (_Assetholdererc20 *Assetholdererc20Session) SetOutcome(channelID [32]byte, parts []ChannelParticipant, newBals []*big.Int) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.SetOutcome(&_Assetholdererc20.TransactOpts, channelID, parts, newBals)
}

// SetOutcome is a paid mutator transaction binding the contract method 0x295482ce.
//
// Solidity: function setOutcome(bytes32 channelID, (address,bytes)[] parts, uint256[] newBals) returns()
func (_Assetholdererc20 *Assetholdererc20TransactorSession) SetOutcome(channelID [32]byte, parts []ChannelParticipant, newBals []*big.Int) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.SetOutcome(&_Assetholdererc20.TransactOpts, channelID, parts, newBals)
}

// Withdraw is a paid mutator transaction binding the contract method 0xfca0f778.
//
// Solidity: function withdraw((bytes32,(address,bytes),address,uint256) authorization, bytes signature) returns()
func (_Assetholdererc20 *Assetholdererc20Transactor) Withdraw(opts *bind.TransactOpts, authorization AssetHolderWithdrawalAuth, signature []byte) (*types.Transaction, error) {
	return _Assetholdererc20.contract.Transact(opts, "withdraw", authorization, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0xfca0f778.
//
// Solidity: function withdraw((bytes32,(address,bytes),address,uint256) authorization, bytes signature) returns()
func (_Assetholdererc20 *Assetholdererc20Session) Withdraw(authorization AssetHolderWithdrawalAuth, signature []byte) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.Withdraw(&_Assetholdererc20.TransactOpts, authorization, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0xfca0f778.
//
// Solidity: function withdraw((bytes32,(address,bytes),address,uint256) authorization, bytes signature) returns()
func (_Assetholdererc20 *Assetholdererc20TransactorSession) Withdraw(authorization AssetHolderWithdrawalAuth, signature []byte) (*types.Transaction, error) {
	return _Assetholdererc20.Contract.Withdraw(&_Assetholdererc20.TransactOpts, authorization, signature)
}

// Assetholdererc20DepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Assetholdererc20 contract.
type Assetholdererc20DepositedIterator struct {
	Event *Assetholdererc20Deposited // Event containing the contract specifics and raw log

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
func (it *Assetholdererc20DepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Assetholdererc20Deposited)
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
		it.Event = new(Assetholdererc20Deposited)
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
func (it *Assetholdererc20DepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Assetholdererc20DepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Assetholdererc20Deposited represents a Deposited event raised by the Assetholdererc20 contract.
type Assetholdererc20Deposited struct {
	FundingID [32]byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a9.
//
// Solidity: event Deposited(bytes32 indexed fundingID, uint256 amount)
func (_Assetholdererc20 *Assetholdererc20Filterer) FilterDeposited(opts *bind.FilterOpts, fundingID [][32]byte) (*Assetholdererc20DepositedIterator, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assetholdererc20.contract.FilterLogs(opts, "Deposited", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return &Assetholdererc20DepositedIterator{contract: _Assetholdererc20.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a9.
//
// Solidity: event Deposited(bytes32 indexed fundingID, uint256 amount)
func (_Assetholdererc20 *Assetholdererc20Filterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *Assetholdererc20Deposited, fundingID [][32]byte) (event.Subscription, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assetholdererc20.contract.WatchLogs(opts, "Deposited", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Assetholdererc20Deposited)
				if err := _Assetholdererc20.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a9.
//
// Solidity: event Deposited(bytes32 indexed fundingID, uint256 amount)
func (_Assetholdererc20 *Assetholdererc20Filterer) ParseDeposited(log types.Log) (*Assetholdererc20Deposited, error) {
	event := new(Assetholdererc20Deposited)
	if err := _Assetholdererc20.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Assetholdererc20OutcomeSetIterator is returned from FilterOutcomeSet and is used to iterate over the raw logs and unpacked data for OutcomeSet events raised by the Assetholdererc20 contract.
type Assetholdererc20OutcomeSetIterator struct {
	Event *Assetholdererc20OutcomeSet // Event containing the contract specifics and raw log

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
func (it *Assetholdererc20OutcomeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Assetholdererc20OutcomeSet)
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
		it.Event = new(Assetholdererc20OutcomeSet)
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
func (it *Assetholdererc20OutcomeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Assetholdererc20OutcomeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Assetholdererc20OutcomeSet represents a OutcomeSet event raised by the Assetholdererc20 contract.
type Assetholdererc20OutcomeSet struct {
	ChannelID [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOutcomeSet is a free log retrieval operation binding the contract event 0xef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b8.
//
// Solidity: event OutcomeSet(bytes32 indexed channelID)
func (_Assetholdererc20 *Assetholdererc20Filterer) FilterOutcomeSet(opts *bind.FilterOpts, channelID [][32]byte) (*Assetholdererc20OutcomeSetIterator, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	logs, sub, err := _Assetholdererc20.contract.FilterLogs(opts, "OutcomeSet", channelIDRule)
	if err != nil {
		return nil, err
	}
	return &Assetholdererc20OutcomeSetIterator{contract: _Assetholdererc20.contract, event: "OutcomeSet", logs: logs, sub: sub}, nil
}

// WatchOutcomeSet is a free log subscription operation binding the contract event 0xef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b8.
//
// Solidity: event OutcomeSet(bytes32 indexed channelID)
func (_Assetholdererc20 *Assetholdererc20Filterer) WatchOutcomeSet(opts *bind.WatchOpts, sink chan<- *Assetholdererc20OutcomeSet, channelID [][32]byte) (event.Subscription, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	logs, sub, err := _Assetholdererc20.contract.WatchLogs(opts, "OutcomeSet", channelIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Assetholdererc20OutcomeSet)
				if err := _Assetholdererc20.contract.UnpackLog(event, "OutcomeSet", log); err != nil {
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

// ParseOutcomeSet is a log parse operation binding the contract event 0xef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b8.
//
// Solidity: event OutcomeSet(bytes32 indexed channelID)
func (_Assetholdererc20 *Assetholdererc20Filterer) ParseOutcomeSet(log types.Log) (*Assetholdererc20OutcomeSet, error) {
	event := new(Assetholdererc20OutcomeSet)
	if err := _Assetholdererc20.contract.UnpackLog(event, "OutcomeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Assetholdererc20WithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Assetholdererc20 contract.
type Assetholdererc20WithdrawnIterator struct {
	Event *Assetholdererc20Withdrawn // Event containing the contract specifics and raw log

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
func (it *Assetholdererc20WithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Assetholdererc20Withdrawn)
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
		it.Event = new(Assetholdererc20Withdrawn)
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
func (it *Assetholdererc20WithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Assetholdererc20WithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Assetholdererc20Withdrawn represents a Withdrawn event raised by the Assetholdererc20 contract.
type Assetholdererc20Withdrawn struct {
	FundingID [32]byte
	Amount    *big.Int
	Receiver  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81.
//
// Solidity: event Withdrawn(bytes32 indexed fundingID, uint256 amount, address receiver)
func (_Assetholdererc20 *Assetholdererc20Filterer) FilterWithdrawn(opts *bind.FilterOpts, fundingID [][32]byte) (*Assetholdererc20WithdrawnIterator, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assetholdererc20.contract.FilterLogs(opts, "Withdrawn", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return &Assetholdererc20WithdrawnIterator{contract: _Assetholdererc20.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81.
//
// Solidity: event Withdrawn(bytes32 indexed fundingID, uint256 amount, address receiver)
func (_Assetholdererc20 *Assetholdererc20Filterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *Assetholdererc20Withdrawn, fundingID [][32]byte) (event.Subscription, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assetholdererc20.contract.WatchLogs(opts, "Withdrawn", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Assetholdererc20Withdrawn)
				if err := _Assetholdererc20.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81.
//
// Solidity: event Withdrawn(bytes32 indexed fundingID, uint256 amount, address receiver)
func (_Assetholdererc20 *Assetholdererc20Filterer) ParseWithdrawn(log types.Log) (*Assetholdererc20Withdrawn, error) {
	event := new(Assetholdererc20Withdrawn)
	if err := _Assetholdererc20.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
