// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reverseregistrar

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ReverseregistrarABI is the input ABI used to generate the binding from.
const ReverseregistrarABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"claimForAddr\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"claimWithResolver\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"node\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setDefaultResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setNameForAddr\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Reverseregistrar is an auto generated Go binding around an Ethereum contract.
type Reverseregistrar struct {
	ReverseregistrarCaller     // Read-only binding to the contract
	ReverseregistrarTransactor // Write-only binding to the contract
	ReverseregistrarFilterer   // Log filterer for contract events
}

// ReverseregistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReverseregistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverseregistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReverseregistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverseregistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReverseregistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReverseregistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReverseregistrarSession struct {
	Contract     *Reverseregistrar // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReverseregistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReverseregistrarCallerSession struct {
	Contract *ReverseregistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ReverseregistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReverseregistrarTransactorSession struct {
	Contract     *ReverseregistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ReverseregistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReverseregistrarRaw struct {
	Contract *Reverseregistrar // Generic contract binding to access the raw methods on
}

// ReverseregistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReverseregistrarCallerRaw struct {
	Contract *ReverseregistrarCaller // Generic read-only contract binding to access the raw methods on
}

// ReverseregistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReverseregistrarTransactorRaw struct {
	Contract *ReverseregistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReverseregistrar creates a new instance of Reverseregistrar, bound to a specific deployed contract.
func NewReverseregistrar(address common.Address, backend bind.ContractBackend) (*Reverseregistrar, error) {
	contract, err := bindReverseregistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reverseregistrar{ReverseregistrarCaller: ReverseregistrarCaller{contract: contract}, ReverseregistrarTransactor: ReverseregistrarTransactor{contract: contract}, ReverseregistrarFilterer: ReverseregistrarFilterer{contract: contract}}, nil
}

// NewReverseregistrarCaller creates a new read-only instance of Reverseregistrar, bound to a specific deployed contract.
func NewReverseregistrarCaller(address common.Address, caller bind.ContractCaller) (*ReverseregistrarCaller, error) {
	contract, err := bindReverseregistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReverseregistrarCaller{contract: contract}, nil
}

// NewReverseregistrarTransactor creates a new write-only instance of Reverseregistrar, bound to a specific deployed contract.
func NewReverseregistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*ReverseregistrarTransactor, error) {
	contract, err := bindReverseregistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReverseregistrarTransactor{contract: contract}, nil
}

// NewReverseregistrarFilterer creates a new log filterer instance of Reverseregistrar, bound to a specific deployed contract.
func NewReverseregistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*ReverseregistrarFilterer, error) {
	contract, err := bindReverseregistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReverseregistrarFilterer{contract: contract}, nil
}

// bindReverseregistrar binds a generic wrapper to an already deployed contract.
func bindReverseregistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReverseregistrarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reverseregistrar *ReverseregistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reverseregistrar.Contract.ReverseregistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reverseregistrar *ReverseregistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.ReverseregistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reverseregistrar *ReverseregistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.ReverseregistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reverseregistrar *ReverseregistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reverseregistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reverseregistrar *ReverseregistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reverseregistrar *ReverseregistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.contract.Transact(opts, method, params...)
}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(address addr) pure returns(bytes32)
func (_Reverseregistrar *ReverseregistrarCaller) Node(opts *bind.CallOpts, addr common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Reverseregistrar.contract.Call(opts, &out, "node", addr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(address addr) pure returns(bytes32)
func (_Reverseregistrar *ReverseregistrarSession) Node(addr common.Address) ([32]byte, error) {
	return _Reverseregistrar.Contract.Node(&_Reverseregistrar.CallOpts, addr)
}

// Node is a free data retrieval call binding the contract method 0xbffbe61c.
//
// Solidity: function node(address addr) pure returns(bytes32)
func (_Reverseregistrar *ReverseregistrarCallerSession) Node(addr common.Address) ([32]byte, error) {
	return _Reverseregistrar.Contract.Node(&_Reverseregistrar.CallOpts, addr)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address owner) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarTransactor) Claim(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Reverseregistrar.contract.Transact(opts, "claim", owner)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address owner) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarSession) Claim(owner common.Address) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.Claim(&_Reverseregistrar.TransactOpts, owner)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address owner) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarTransactorSession) Claim(owner common.Address) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.Claim(&_Reverseregistrar.TransactOpts, owner)
}

// ClaimForAddr is a paid mutator transaction binding the contract method 0x65669631.
//
// Solidity: function claimForAddr(address addr, address owner, address resolver) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarTransactor) ClaimForAddr(opts *bind.TransactOpts, addr common.Address, owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _Reverseregistrar.contract.Transact(opts, "claimForAddr", addr, owner, resolver)
}

// ClaimForAddr is a paid mutator transaction binding the contract method 0x65669631.
//
// Solidity: function claimForAddr(address addr, address owner, address resolver) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarSession) ClaimForAddr(addr common.Address, owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.ClaimForAddr(&_Reverseregistrar.TransactOpts, addr, owner, resolver)
}

// ClaimForAddr is a paid mutator transaction binding the contract method 0x65669631.
//
// Solidity: function claimForAddr(address addr, address owner, address resolver) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarTransactorSession) ClaimForAddr(addr common.Address, owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.ClaimForAddr(&_Reverseregistrar.TransactOpts, addr, owner, resolver)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(address owner, address resolver) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarTransactor) ClaimWithResolver(opts *bind.TransactOpts, owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _Reverseregistrar.contract.Transact(opts, "claimWithResolver", owner, resolver)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(address owner, address resolver) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarSession) ClaimWithResolver(owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.ClaimWithResolver(&_Reverseregistrar.TransactOpts, owner, resolver)
}

// ClaimWithResolver is a paid mutator transaction binding the contract method 0x0f5a5466.
//
// Solidity: function claimWithResolver(address owner, address resolver) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarTransactorSession) ClaimWithResolver(owner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.ClaimWithResolver(&_Reverseregistrar.TransactOpts, owner, resolver)
}

// SetDefaultResolver is a paid mutator transaction binding the contract method 0xc66485b2.
//
// Solidity: function setDefaultResolver(address resolver) returns()
func (_Reverseregistrar *ReverseregistrarTransactor) SetDefaultResolver(opts *bind.TransactOpts, resolver common.Address) (*types.Transaction, error) {
	return _Reverseregistrar.contract.Transact(opts, "setDefaultResolver", resolver)
}

// SetDefaultResolver is a paid mutator transaction binding the contract method 0xc66485b2.
//
// Solidity: function setDefaultResolver(address resolver) returns()
func (_Reverseregistrar *ReverseregistrarSession) SetDefaultResolver(resolver common.Address) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.SetDefaultResolver(&_Reverseregistrar.TransactOpts, resolver)
}

// SetDefaultResolver is a paid mutator transaction binding the contract method 0xc66485b2.
//
// Solidity: function setDefaultResolver(address resolver) returns()
func (_Reverseregistrar *ReverseregistrarTransactorSession) SetDefaultResolver(resolver common.Address) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.SetDefaultResolver(&_Reverseregistrar.TransactOpts, resolver)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarTransactor) SetName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Reverseregistrar.contract.Transact(opts, "setName", name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarSession) SetName(name string) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.SetName(&_Reverseregistrar.TransactOpts, name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarTransactorSession) SetName(name string) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.SetName(&_Reverseregistrar.TransactOpts, name)
}

// SetNameForAddr is a paid mutator transaction binding the contract method 0x7a806d6b.
//
// Solidity: function setNameForAddr(address addr, address owner, address resolver, string name) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarTransactor) SetNameForAddr(opts *bind.TransactOpts, addr common.Address, owner common.Address, resolver common.Address, name string) (*types.Transaction, error) {
	return _Reverseregistrar.contract.Transact(opts, "setNameForAddr", addr, owner, resolver, name)
}

// SetNameForAddr is a paid mutator transaction binding the contract method 0x7a806d6b.
//
// Solidity: function setNameForAddr(address addr, address owner, address resolver, string name) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarSession) SetNameForAddr(addr common.Address, owner common.Address, resolver common.Address, name string) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.SetNameForAddr(&_Reverseregistrar.TransactOpts, addr, owner, resolver, name)
}

// SetNameForAddr is a paid mutator transaction binding the contract method 0x7a806d6b.
//
// Solidity: function setNameForAddr(address addr, address owner, address resolver, string name) returns(bytes32)
func (_Reverseregistrar *ReverseregistrarTransactorSession) SetNameForAddr(addr common.Address, owner common.Address, resolver common.Address, name string) (*types.Transaction, error) {
	return _Reverseregistrar.Contract.SetNameForAddr(&_Reverseregistrar.TransactOpts, addr, owner, resolver, name)
}
