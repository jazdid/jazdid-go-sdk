// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registrar

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

// RegistrarABI is the input ABI used to generate the binding from.
const RegistrarABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"labelId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"labelId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"labelId\",\"type\":\"uint256\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseNode\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gracePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"labelId\",\"type\":\"uint256\"}],\"name\":\"nameExpires\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"nameOf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"labelId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"reclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"labelId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tld\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"labelId\",\"type\":\"uint256\"}],\"name\":\"tokenOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Registrar is an auto generated Go binding around an Ethereum contract.
type Registrar struct {
	RegistrarCaller     // Read-only binding to the contract
	RegistrarTransactor // Write-only binding to the contract
	RegistrarFilterer   // Log filterer for contract events
}

// RegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrarSession struct {
	Contract     *Registrar        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrarCallerSession struct {
	Contract *RegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrarTransactorSession struct {
	Contract     *RegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrarRaw struct {
	Contract *Registrar // Generic contract binding to access the raw methods on
}

// RegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrarCallerRaw struct {
	Contract *RegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrarTransactorRaw struct {
	Contract *RegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistrar creates a new instance of Registrar, bound to a specific deployed contract.
func NewRegistrar(address common.Address, backend bind.ContractBackend) (*Registrar, error) {
	contract, err := bindRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registrar{RegistrarCaller: RegistrarCaller{contract: contract}, RegistrarTransactor: RegistrarTransactor{contract: contract}, RegistrarFilterer: RegistrarFilterer{contract: contract}}, nil
}

// NewRegistrarCaller creates a new read-only instance of Registrar, bound to a specific deployed contract.
func NewRegistrarCaller(address common.Address, caller bind.ContractCaller) (*RegistrarCaller, error) {
	contract, err := bindRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarCaller{contract: contract}, nil
}

// NewRegistrarTransactor creates a new write-only instance of Registrar, bound to a specific deployed contract.
func NewRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrarTransactor, error) {
	contract, err := bindRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarTransactor{contract: contract}, nil
}

// NewRegistrarFilterer creates a new log filterer instance of Registrar, bound to a specific deployed contract.
func NewRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrarFilterer, error) {
	contract, err := bindRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrarFilterer{contract: contract}, nil
}

// bindRegistrar binds a generic wrapper to an already deployed contract.
func bindRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistrarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrar *RegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registrar.Contract.RegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrar *RegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrar.Contract.RegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrar *RegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrar.Contract.RegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrar *RegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrar *RegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrar *RegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrar.Contract.contract.Transact(opts, method, params...)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 labelId) view returns(bool)
func (_Registrar *RegistrarCaller) Available(opts *bind.CallOpts, labelId *big.Int) (bool, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "available", labelId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 labelId) view returns(bool)
func (_Registrar *RegistrarSession) Available(labelId *big.Int) (bool, error) {
	return _Registrar.Contract.Available(&_Registrar.CallOpts, labelId)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 labelId) view returns(bool)
func (_Registrar *RegistrarCallerSession) Available(labelId *big.Int) (bool, error) {
	return _Registrar.Contract.Available(&_Registrar.CallOpts, labelId)
}

// BaseNode is a free data retrieval call binding the contract method 0xddf7fcb0.
//
// Solidity: function baseNode() view returns(bytes32)
func (_Registrar *RegistrarCaller) BaseNode(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "baseNode")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BaseNode is a free data retrieval call binding the contract method 0xddf7fcb0.
//
// Solidity: function baseNode() view returns(bytes32)
func (_Registrar *RegistrarSession) BaseNode() ([32]byte, error) {
	return _Registrar.Contract.BaseNode(&_Registrar.CallOpts)
}

// BaseNode is a free data retrieval call binding the contract method 0xddf7fcb0.
//
// Solidity: function baseNode() view returns(bytes32)
func (_Registrar *RegistrarCallerSession) BaseNode() ([32]byte, error) {
	return _Registrar.Contract.BaseNode(&_Registrar.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Registrar *RegistrarCaller) Exists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "exists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Registrar *RegistrarSession) Exists(tokenId *big.Int) (bool, error) {
	return _Registrar.Contract.Exists(&_Registrar.CallOpts, tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Registrar *RegistrarCallerSession) Exists(tokenId *big.Int) (bool, error) {
	return _Registrar.Contract.Exists(&_Registrar.CallOpts, tokenId)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Registrar *RegistrarCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Registrar *RegistrarSession) FeeRecipient() (common.Address, error) {
	return _Registrar.Contract.FeeRecipient(&_Registrar.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Registrar *RegistrarCallerSession) FeeRecipient() (common.Address, error) {
	return _Registrar.Contract.FeeRecipient(&_Registrar.CallOpts)
}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() pure returns(uint256)
func (_Registrar *RegistrarCaller) GracePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "gracePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() pure returns(uint256)
func (_Registrar *RegistrarSession) GracePeriod() (*big.Int, error) {
	return _Registrar.Contract.GracePeriod(&_Registrar.CallOpts)
}

// GracePeriod is a free data retrieval call binding the contract method 0xa06db7dc.
//
// Solidity: function gracePeriod() pure returns(uint256)
func (_Registrar *RegistrarCallerSession) GracePeriod() (*big.Int, error) {
	return _Registrar.Contract.GracePeriod(&_Registrar.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_Registrar *RegistrarCaller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "issuer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_Registrar *RegistrarSession) Issuer() (common.Address, error) {
	return _Registrar.Contract.Issuer(&_Registrar.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_Registrar *RegistrarCallerSession) Issuer() (common.Address, error) {
	return _Registrar.Contract.Issuer(&_Registrar.CallOpts)
}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 labelId) view returns(uint256)
func (_Registrar *RegistrarCaller) NameExpires(opts *bind.CallOpts, labelId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "nameExpires", labelId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 labelId) view returns(uint256)
func (_Registrar *RegistrarSession) NameExpires(labelId *big.Int) (*big.Int, error) {
	return _Registrar.Contract.NameExpires(&_Registrar.CallOpts, labelId)
}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 labelId) view returns(uint256)
func (_Registrar *RegistrarCallerSession) NameExpires(labelId *big.Int) (*big.Int, error) {
	return _Registrar.Contract.NameExpires(&_Registrar.CallOpts, labelId)
}

// NameOf is a free data retrieval call binding the contract method 0x051a2664.
//
// Solidity: function nameOf(uint256 tokenId) view returns(string)
func (_Registrar *RegistrarCaller) NameOf(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "nameOf", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NameOf is a free data retrieval call binding the contract method 0x051a2664.
//
// Solidity: function nameOf(uint256 tokenId) view returns(string)
func (_Registrar *RegistrarSession) NameOf(tokenId *big.Int) (string, error) {
	return _Registrar.Contract.NameOf(&_Registrar.CallOpts, tokenId)
}

// NameOf is a free data retrieval call binding the contract method 0x051a2664.
//
// Solidity: function nameOf(uint256 tokenId) view returns(string)
func (_Registrar *RegistrarCallerSession) NameOf(tokenId *big.Int) (string, error) {
	return _Registrar.Contract.NameOf(&_Registrar.CallOpts, tokenId)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Registrar *RegistrarCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Registrar *RegistrarSession) NextTokenId() (*big.Int, error) {
	return _Registrar.Contract.NextTokenId(&_Registrar.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Registrar *RegistrarCallerSession) NextTokenId() (*big.Int, error) {
	return _Registrar.Contract.NextTokenId(&_Registrar.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_Registrar *RegistrarCaller) PriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "priceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_Registrar *RegistrarSession) PriceOracle() (common.Address, error) {
	return _Registrar.Contract.PriceOracle(&_Registrar.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x2630c12f.
//
// Solidity: function priceOracle() view returns(address)
func (_Registrar *RegistrarCallerSession) PriceOracle() (common.Address, error) {
	return _Registrar.Contract.PriceOracle(&_Registrar.CallOpts)
}

// Tld is a free data retrieval call binding the contract method 0x2d551432.
//
// Solidity: function tld() view returns(string)
func (_Registrar *RegistrarCaller) Tld(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "tld")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Tld is a free data retrieval call binding the contract method 0x2d551432.
//
// Solidity: function tld() view returns(string)
func (_Registrar *RegistrarSession) Tld() (string, error) {
	return _Registrar.Contract.Tld(&_Registrar.CallOpts)
}

// Tld is a free data retrieval call binding the contract method 0x2d551432.
//
// Solidity: function tld() view returns(string)
func (_Registrar *RegistrarCallerSession) Tld() (string, error) {
	return _Registrar.Contract.Tld(&_Registrar.CallOpts)
}

// TokenOf is a free data retrieval call binding the contract method 0xea78803f.
//
// Solidity: function tokenOf(uint256 labelId) view returns(uint256)
func (_Registrar *RegistrarCaller) TokenOf(opts *bind.CallOpts, labelId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "tokenOf", labelId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOf is a free data retrieval call binding the contract method 0xea78803f.
//
// Solidity: function tokenOf(uint256 labelId) view returns(uint256)
func (_Registrar *RegistrarSession) TokenOf(labelId *big.Int) (*big.Int, error) {
	return _Registrar.Contract.TokenOf(&_Registrar.CallOpts, labelId)
}

// TokenOf is a free data retrieval call binding the contract method 0xea78803f.
//
// Solidity: function tokenOf(uint256 labelId) view returns(uint256)
func (_Registrar *RegistrarCallerSession) TokenOf(labelId *big.Int) (*big.Int, error) {
	return _Registrar.Contract.TokenOf(&_Registrar.CallOpts, labelId)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 labelId, address owner) returns()
func (_Registrar *RegistrarTransactor) Reclaim(opts *bind.TransactOpts, labelId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "reclaim", labelId, owner)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 labelId, address owner) returns()
func (_Registrar *RegistrarSession) Reclaim(labelId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.Reclaim(&_Registrar.TransactOpts, labelId, owner)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 labelId, address owner) returns()
func (_Registrar *RegistrarTransactorSession) Reclaim(labelId *big.Int, owner common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.Reclaim(&_Registrar.TransactOpts, labelId, owner)
}

// Register is a paid mutator transaction binding the contract method 0xe0aade54.
//
// Solidity: function register(string name, address owner, uint256 duration, address resolver) returns(uint256 tokenId, uint256 expires)
func (_Registrar *RegistrarTransactor) Register(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, resolver common.Address) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "register", name, owner, duration, resolver)
}

// Register is a paid mutator transaction binding the contract method 0xe0aade54.
//
// Solidity: function register(string name, address owner, uint256 duration, address resolver) returns(uint256 tokenId, uint256 expires)
func (_Registrar *RegistrarSession) Register(name string, owner common.Address, duration *big.Int, resolver common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.Register(&_Registrar.TransactOpts, name, owner, duration, resolver)
}

// Register is a paid mutator transaction binding the contract method 0xe0aade54.
//
// Solidity: function register(string name, address owner, uint256 duration, address resolver) returns(uint256 tokenId, uint256 expires)
func (_Registrar *RegistrarTransactorSession) Register(name string, owner common.Address, duration *big.Int, resolver common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.Register(&_Registrar.TransactOpts, name, owner, duration, resolver)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 labelId, uint256 duration) returns(uint256 tokenId, uint256 expires)
func (_Registrar *RegistrarTransactor) Renew(opts *bind.TransactOpts, labelId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "renew", labelId, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 labelId, uint256 duration) returns(uint256 tokenId, uint256 expires)
func (_Registrar *RegistrarSession) Renew(labelId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Registrar.Contract.Renew(&_Registrar.TransactOpts, labelId, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 labelId, uint256 duration) returns(uint256 tokenId, uint256 expires)
func (_Registrar *RegistrarTransactorSession) Renew(labelId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Registrar.Contract.Renew(&_Registrar.TransactOpts, labelId, duration)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_Registrar *RegistrarTransactor) SetResolver(opts *bind.TransactOpts, resolver common.Address) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "setResolver", resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_Registrar *RegistrarSession) SetResolver(resolver common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.SetResolver(&_Registrar.TransactOpts, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_Registrar *RegistrarTransactorSession) SetResolver(resolver common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.SetResolver(&_Registrar.TransactOpts, resolver)
}

// RegistrarNameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the Registrar contract.
type RegistrarNameRegisteredIterator struct {
	Event *RegistrarNameRegistered // Event containing the contract specifics and raw log

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
func (it *RegistrarNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarNameRegistered)
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
		it.Event = new(RegistrarNameRegistered)
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
func (it *RegistrarNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarNameRegistered represents a NameRegistered event raised by the Registrar contract.
type RegistrarNameRegistered struct {
	TokenId *big.Int
	LabelId *big.Int
	Owner   common.Address
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0x9683eb0f3e3193de922967d3019feef591895f5a1bcca02a2bd0e60b51b33dc6.
//
// Solidity: event NameRegistered(uint256 indexed tokenId, uint256 indexed labelId, address indexed owner, uint256 expires)
func (_Registrar *RegistrarFilterer) FilterNameRegistered(opts *bind.FilterOpts, tokenId []*big.Int, labelId []*big.Int, owner []common.Address) (*RegistrarNameRegisteredIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var labelIdRule []interface{}
	for _, labelIdItem := range labelId {
		labelIdRule = append(labelIdRule, labelIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Registrar.contract.FilterLogs(opts, "NameRegistered", tokenIdRule, labelIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarNameRegisteredIterator{contract: _Registrar.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0x9683eb0f3e3193de922967d3019feef591895f5a1bcca02a2bd0e60b51b33dc6.
//
// Solidity: event NameRegistered(uint256 indexed tokenId, uint256 indexed labelId, address indexed owner, uint256 expires)
func (_Registrar *RegistrarFilterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *RegistrarNameRegistered, tokenId []*big.Int, labelId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var labelIdRule []interface{}
	for _, labelIdItem := range labelId {
		labelIdRule = append(labelIdRule, labelIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Registrar.contract.WatchLogs(opts, "NameRegistered", tokenIdRule, labelIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarNameRegistered)
				if err := _Registrar.contract.UnpackLog(event, "NameRegistered", log); err != nil {
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

// ParseNameRegistered is a log parse operation binding the contract event 0x9683eb0f3e3193de922967d3019feef591895f5a1bcca02a2bd0e60b51b33dc6.
//
// Solidity: event NameRegistered(uint256 indexed tokenId, uint256 indexed labelId, address indexed owner, uint256 expires)
func (_Registrar *RegistrarFilterer) ParseNameRegistered(log types.Log) (*RegistrarNameRegistered, error) {
	event := new(RegistrarNameRegistered)
	if err := _Registrar.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistrarNameRenewedIterator is returned from FilterNameRenewed and is used to iterate over the raw logs and unpacked data for NameRenewed events raised by the Registrar contract.
type RegistrarNameRenewedIterator struct {
	Event *RegistrarNameRenewed // Event containing the contract specifics and raw log

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
func (it *RegistrarNameRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarNameRenewed)
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
		it.Event = new(RegistrarNameRenewed)
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
func (it *RegistrarNameRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarNameRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarNameRenewed represents a NameRenewed event raised by the Registrar contract.
type RegistrarNameRenewed struct {
	TokenId *big.Int
	LabelId *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRenewed is a free log retrieval operation binding the contract event 0xd4b61d6768ffa07c4e38f85dca93cda5b258136f967cbe4efcfe2a4ed9047c20.
//
// Solidity: event NameRenewed(uint256 indexed tokenId, uint256 indexed labelId, uint256 expires)
func (_Registrar *RegistrarFilterer) FilterNameRenewed(opts *bind.FilterOpts, tokenId []*big.Int, labelId []*big.Int) (*RegistrarNameRenewedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var labelIdRule []interface{}
	for _, labelIdItem := range labelId {
		labelIdRule = append(labelIdRule, labelIdItem)
	}

	logs, sub, err := _Registrar.contract.FilterLogs(opts, "NameRenewed", tokenIdRule, labelIdRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarNameRenewedIterator{contract: _Registrar.contract, event: "NameRenewed", logs: logs, sub: sub}, nil
}

// WatchNameRenewed is a free log subscription operation binding the contract event 0xd4b61d6768ffa07c4e38f85dca93cda5b258136f967cbe4efcfe2a4ed9047c20.
//
// Solidity: event NameRenewed(uint256 indexed tokenId, uint256 indexed labelId, uint256 expires)
func (_Registrar *RegistrarFilterer) WatchNameRenewed(opts *bind.WatchOpts, sink chan<- *RegistrarNameRenewed, tokenId []*big.Int, labelId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var labelIdRule []interface{}
	for _, labelIdItem := range labelId {
		labelIdRule = append(labelIdRule, labelIdItem)
	}

	logs, sub, err := _Registrar.contract.WatchLogs(opts, "NameRenewed", tokenIdRule, labelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarNameRenewed)
				if err := _Registrar.contract.UnpackLog(event, "NameRenewed", log); err != nil {
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

// ParseNameRenewed is a log parse operation binding the contract event 0xd4b61d6768ffa07c4e38f85dca93cda5b258136f967cbe4efcfe2a4ed9047c20.
//
// Solidity: event NameRenewed(uint256 indexed tokenId, uint256 indexed labelId, uint256 expires)
func (_Registrar *RegistrarFilterer) ParseNameRenewed(log types.Log) (*RegistrarNameRenewed, error) {
	event := new(RegistrarNameRenewed)
	if err := _Registrar.contract.UnpackLog(event, "NameRenewed", log); err != nil {
		return nil, err
	}
	return event, nil
}
