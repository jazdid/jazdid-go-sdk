// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package controller

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

// IBulkRegistrarControllerRentPrice is an auto generated low-level Go binding around an user-defined struct.
type IBulkRegistrarControllerRentPrice struct {
	Currency common.Address
	Cost     *big.Int
}

// ControllerABI is the input ABI used to generate the binding from.
const ControllerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"labelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameRegisterFailed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"registrars\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"names\",\"type\":\"string[]\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"registrars\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"names\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"durations\",\"type\":\"uint256[]\"}],\"name\":\"bulkRenew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"}],\"name\":\"nameRecords\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"registrars\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"names\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"durations\",\"type\":\"uint256[]\"}],\"name\":\"rentPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"internalType\":\"structIBulkRegistrarController.RentPrice[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ControllerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// Available is a free data retrieval call binding the contract method 0xc75b3b61.
//
// Solidity: function available(address[] registrars, string[] names) view returns(bool[], uint256[])
func (_Controller *ControllerCaller) Available(opts *bind.CallOpts, registrars []common.Address, names []string) ([]bool, []*big.Int, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "available", registrars, names)

	if err != nil {
		return *new([]bool), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// Available is a free data retrieval call binding the contract method 0xc75b3b61.
//
// Solidity: function available(address[] registrars, string[] names) view returns(bool[], uint256[])
func (_Controller *ControllerSession) Available(registrars []common.Address, names []string) ([]bool, []*big.Int, error) {
	return _Controller.Contract.Available(&_Controller.CallOpts, registrars, names)
}

// Available is a free data retrieval call binding the contract method 0xc75b3b61.
//
// Solidity: function available(address[] registrars, string[] names) view returns(bool[], uint256[])
func (_Controller *ControllerCallerSession) Available(registrars []common.Address, names []string) ([]bool, []*big.Int, error) {
	return _Controller.Contract.Available(&_Controller.CallOpts, registrars, names)
}

// NameRecords is a free data retrieval call binding the contract method 0x0eb52e20.
//
// Solidity: function nameRecords(address registrar, bytes32 node, string[] keys) view returns(address, string[])
func (_Controller *ControllerCaller) NameRecords(opts *bind.CallOpts, registrar common.Address, node [32]byte, keys []string) (common.Address, []string, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "nameRecords", registrar, node, keys)

	if err != nil {
		return *new(common.Address), *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([]string)).(*[]string)

	return out0, out1, err

}

// NameRecords is a free data retrieval call binding the contract method 0x0eb52e20.
//
// Solidity: function nameRecords(address registrar, bytes32 node, string[] keys) view returns(address, string[])
func (_Controller *ControllerSession) NameRecords(registrar common.Address, node [32]byte, keys []string) (common.Address, []string, error) {
	return _Controller.Contract.NameRecords(&_Controller.CallOpts, registrar, node, keys)
}

// NameRecords is a free data retrieval call binding the contract method 0x0eb52e20.
//
// Solidity: function nameRecords(address registrar, bytes32 node, string[] keys) view returns(address, string[])
func (_Controller *ControllerCallerSession) NameRecords(registrar common.Address, node [32]byte, keys []string) (common.Address, []string, error) {
	return _Controller.Contract.NameRecords(&_Controller.CallOpts, registrar, node, keys)
}

// RentPrice is a free data retrieval call binding the contract method 0x1aeb5460.
//
// Solidity: function rentPrice(address[] registrars, string[] names, uint256[] durations) view returns((address,uint256)[])
func (_Controller *ControllerCaller) RentPrice(opts *bind.CallOpts, registrars []common.Address, names []string, durations []*big.Int) ([]IBulkRegistrarControllerRentPrice, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "rentPrice", registrars, names, durations)

	if err != nil {
		return *new([]IBulkRegistrarControllerRentPrice), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBulkRegistrarControllerRentPrice)).(*[]IBulkRegistrarControllerRentPrice)

	return out0, err

}

// RentPrice is a free data retrieval call binding the contract method 0x1aeb5460.
//
// Solidity: function rentPrice(address[] registrars, string[] names, uint256[] durations) view returns((address,uint256)[])
func (_Controller *ControllerSession) RentPrice(registrars []common.Address, names []string, durations []*big.Int) ([]IBulkRegistrarControllerRentPrice, error) {
	return _Controller.Contract.RentPrice(&_Controller.CallOpts, registrars, names, durations)
}

// RentPrice is a free data retrieval call binding the contract method 0x1aeb5460.
//
// Solidity: function rentPrice(address[] registrars, string[] names, uint256[] durations) view returns((address,uint256)[])
func (_Controller *ControllerCallerSession) RentPrice(registrars []common.Address, names []string, durations []*big.Int) ([]IBulkRegistrarControllerRentPrice, error) {
	return _Controller.Contract.RentPrice(&_Controller.CallOpts, registrars, names, durations)
}

// BulkRenew is a paid mutator transaction binding the contract method 0x5d28558b.
//
// Solidity: function bulkRenew(address[] registrars, string[] names, uint256[] durations) payable returns()
func (_Controller *ControllerTransactor) BulkRenew(opts *bind.TransactOpts, registrars []common.Address, names []string, durations []*big.Int) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "bulkRenew", registrars, names, durations)
}

// BulkRenew is a paid mutator transaction binding the contract method 0x5d28558b.
//
// Solidity: function bulkRenew(address[] registrars, string[] names, uint256[] durations) payable returns()
func (_Controller *ControllerSession) BulkRenew(registrars []common.Address, names []string, durations []*big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BulkRenew(&_Controller.TransactOpts, registrars, names, durations)
}

// BulkRenew is a paid mutator transaction binding the contract method 0x5d28558b.
//
// Solidity: function bulkRenew(address[] registrars, string[] names, uint256[] durations) payable returns()
func (_Controller *ControllerTransactorSession) BulkRenew(registrars []common.Address, names []string, durations []*big.Int) (*types.Transaction, error) {
	return _Controller.Contract.BulkRenew(&_Controller.TransactOpts, registrars, names, durations)
}

// ControllerNameRegisterFailedIterator is returned from FilterNameRegisterFailed and is used to iterate over the raw logs and unpacked data for NameRegisterFailed events raised by the Controller contract.
type ControllerNameRegisterFailedIterator struct {
	Event *ControllerNameRegisterFailed // Event containing the contract specifics and raw log

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
func (it *ControllerNameRegisterFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerNameRegisterFailed)
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
		it.Event = new(ControllerNameRegisterFailed)
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
func (it *ControllerNameRegisterFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerNameRegisterFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerNameRegisterFailed represents a NameRegisterFailed event raised by the Controller contract.
type ControllerNameRegisterFailed struct {
	Registrar common.Address
	LabelId   [32]byte
	Name      string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNameRegisterFailed is a free log retrieval operation binding the contract event 0xc741833b4fde08945e725ee970f949d8bb04c43c5c3346ee67cc4e960cbfb33d.
//
// Solidity: event NameRegisterFailed(address indexed registrar, bytes32 indexed labelId, string name)
func (_Controller *ControllerFilterer) FilterNameRegisterFailed(opts *bind.FilterOpts, registrar []common.Address, labelId [][32]byte) (*ControllerNameRegisterFailedIterator, error) {

	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}
	var labelIdRule []interface{}
	for _, labelIdItem := range labelId {
		labelIdRule = append(labelIdRule, labelIdItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "NameRegisterFailed", registrarRule, labelIdRule)
	if err != nil {
		return nil, err
	}
	return &ControllerNameRegisterFailedIterator{contract: _Controller.contract, event: "NameRegisterFailed", logs: logs, sub: sub}, nil
}

// WatchNameRegisterFailed is a free log subscription operation binding the contract event 0xc741833b4fde08945e725ee970f949d8bb04c43c5c3346ee67cc4e960cbfb33d.
//
// Solidity: event NameRegisterFailed(address indexed registrar, bytes32 indexed labelId, string name)
func (_Controller *ControllerFilterer) WatchNameRegisterFailed(opts *bind.WatchOpts, sink chan<- *ControllerNameRegisterFailed, registrar []common.Address, labelId [][32]byte) (event.Subscription, error) {

	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}
	var labelIdRule []interface{}
	for _, labelIdItem := range labelId {
		labelIdRule = append(labelIdRule, labelIdItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "NameRegisterFailed", registrarRule, labelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerNameRegisterFailed)
				if err := _Controller.contract.UnpackLog(event, "NameRegisterFailed", log); err != nil {
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

// ParseNameRegisterFailed is a log parse operation binding the contract event 0xc741833b4fde08945e725ee970f949d8bb04c43c5c3346ee67cc4e960cbfb33d.
//
// Solidity: event NameRegisterFailed(address indexed registrar, bytes32 indexed labelId, string name)
func (_Controller *ControllerFilterer) ParseNameRegisterFailed(log types.Log) (*ControllerNameRegisterFailed, error) {
	event := new(ControllerNameRegisterFailed)
	if err := _Controller.contract.UnpackLog(event, "NameRegisterFailed", log); err != nil {
		return nil, err
	}
	return event, nil
}
