// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ConstantsRegistry

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
)

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"AddedConstant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"RemovedConstant\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key_\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value_\",\"type\":\"bytes\"}],\"name\":\"addConstant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"constants\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_injector\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key_\",\"type\":\"string\"}],\"name\":\"removeConstant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress_\",\"type\":\"address\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_injector\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// Constants is a free data retrieval call binding the contract method 0x163f33a5.
//
// Solidity: function constants(string ) view returns(bytes)
func (_Main *MainCaller) Constants(opts *bind.CallOpts, arg0 string) ([]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "constants", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Constants is a free data retrieval call binding the contract method 0x163f33a5.
//
// Solidity: function constants(string ) view returns(bytes)
func (_Main *MainSession) Constants(arg0 string) ([]byte, error) {
	return _Main.Contract.Constants(&_Main.CallOpts, arg0)
}

// Constants is a free data retrieval call binding the contract method 0x163f33a5.
//
// Solidity: function constants(string ) view returns(bytes)
func (_Main *MainCallerSession) Constants(arg0 string) ([]byte, error) {
	return _Main.Contract.Constants(&_Main.CallOpts, arg0)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address _injector)
func (_Main *MainCaller) GetInjector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getInjector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address _injector)
func (_Main *MainSession) GetInjector() (common.Address, error) {
	return _Main.Contract.GetInjector(&_Main.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address _injector)
func (_Main *MainCallerSession) GetInjector() (common.Address, error) {
	return _Main.Contract.GetInjector(&_Main.CallOpts)
}

// AddConstant is a paid mutator transaction binding the contract method 0x4216ab3b.
//
// Solidity: function addConstant(string key_, bytes value_) returns()
func (_Main *MainTransactor) AddConstant(opts *bind.TransactOpts, key_ string, value_ []byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addConstant", key_, value_)
}

// AddConstant is a paid mutator transaction binding the contract method 0x4216ab3b.
//
// Solidity: function addConstant(string key_, bytes value_) returns()
func (_Main *MainSession) AddConstant(key_ string, value_ []byte) (*types.Transaction, error) {
	return _Main.Contract.AddConstant(&_Main.TransactOpts, key_, value_)
}

// AddConstant is a paid mutator transaction binding the contract method 0x4216ab3b.
//
// Solidity: function addConstant(string key_, bytes value_) returns()
func (_Main *MainTransactorSession) AddConstant(key_ string, value_ []byte) (*types.Transaction, error) {
	return _Main.Contract.AddConstant(&_Main.TransactOpts, key_, value_)
}

// RemoveConstant is a paid mutator transaction binding the contract method 0x965136fd.
//
// Solidity: function removeConstant(string key_) returns()
func (_Main *MainTransactor) RemoveConstant(opts *bind.TransactOpts, key_ string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeConstant", key_)
}

// RemoveConstant is a paid mutator transaction binding the contract method 0x965136fd.
//
// Solidity: function removeConstant(string key_) returns()
func (_Main *MainSession) RemoveConstant(key_ string) (*types.Transaction, error) {
	return _Main.Contract.RemoveConstant(&_Main.TransactOpts, key_)
}

// RemoveConstant is a paid mutator transaction binding the contract method 0x965136fd.
//
// Solidity: function removeConstant(string key_) returns()
func (_Main *MainTransactorSession) RemoveConstant(key_ string) (*types.Transaction, error) {
	return _Main.Contract.RemoveConstant(&_Main.TransactOpts, key_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x8389cb18.
//
// Solidity: function setDependencies(address registryAddress_) returns()
func (_Main *MainTransactor) SetDependencies(opts *bind.TransactOpts, registryAddress_ common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setDependencies", registryAddress_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x8389cb18.
//
// Solidity: function setDependencies(address registryAddress_) returns()
func (_Main *MainSession) SetDependencies(registryAddress_ common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetDependencies(&_Main.TransactOpts, registryAddress_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x8389cb18.
//
// Solidity: function setDependencies(address registryAddress_) returns()
func (_Main *MainTransactorSession) SetDependencies(registryAddress_ common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetDependencies(&_Main.TransactOpts, registryAddress_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address _injector) returns()
func (_Main *MainTransactor) SetInjector(opts *bind.TransactOpts, _injector common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setInjector", _injector)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address _injector) returns()
func (_Main *MainSession) SetInjector(_injector common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetInjector(&_Main.TransactOpts, _injector)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address _injector) returns()
func (_Main *MainTransactorSession) SetInjector(_injector common.Address) (*types.Transaction, error) {
	return _Main.Contract.SetInjector(&_Main.TransactOpts, _injector)
}

// MainAddedConstantIterator is returned from FilterAddedConstant and is used to iterate over the raw logs and unpacked data for AddedConstant events raised by the Main contract.
type MainAddedConstantIterator struct {
	Event *MainAddedConstant // Event containing the contract specifics and raw log

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
func (it *MainAddedConstantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainAddedConstant)
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
		it.Event = new(MainAddedConstant)
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
func (it *MainAddedConstantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainAddedConstantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainAddedConstant represents a AddedConstant event raised by the Main contract.
type MainAddedConstant struct {
	Name  string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddedConstant is a free log retrieval operation binding the contract event 0x265476163f06971fd1d163653ba47af9e018395e1df6b521f92de499355ee41a.
//
// Solidity: event AddedConstant(string name, bytes value)
func (_Main *MainFilterer) FilterAddedConstant(opts *bind.FilterOpts) (*MainAddedConstantIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "AddedConstant")
	if err != nil {
		return nil, err
	}
	return &MainAddedConstantIterator{contract: _Main.contract, event: "AddedConstant", logs: logs, sub: sub}, nil
}

// WatchAddedConstant is a free log subscription operation binding the contract event 0x265476163f06971fd1d163653ba47af9e018395e1df6b521f92de499355ee41a.
//
// Solidity: event AddedConstant(string name, bytes value)
func (_Main *MainFilterer) WatchAddedConstant(opts *bind.WatchOpts, sink chan<- *MainAddedConstant) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "AddedConstant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainAddedConstant)
				if err := _Main.contract.UnpackLog(event, "AddedConstant", log); err != nil {
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

// ParseAddedConstant is a log parse operation binding the contract event 0x265476163f06971fd1d163653ba47af9e018395e1df6b521f92de499355ee41a.
//
// Solidity: event AddedConstant(string name, bytes value)
func (_Main *MainFilterer) ParseAddedConstant(log types.Log) (*MainAddedConstant, error) {
	event := new(MainAddedConstant)
	if err := _Main.contract.UnpackLog(event, "AddedConstant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRemovedConstantIterator is returned from FilterRemovedConstant and is used to iterate over the raw logs and unpacked data for RemovedConstant events raised by the Main contract.
type MainRemovedConstantIterator struct {
	Event *MainRemovedConstant // Event containing the contract specifics and raw log

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
func (it *MainRemovedConstantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRemovedConstant)
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
		it.Event = new(MainRemovedConstant)
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
func (it *MainRemovedConstantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRemovedConstantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRemovedConstant represents a RemovedConstant event raised by the Main contract.
type MainRemovedConstant struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedConstant is a free log retrieval operation binding the contract event 0x756b8e626fd3514f249eda10fac2a5342e23cbabdb29803c2c8743bd3489afb1.
//
// Solidity: event RemovedConstant(string name)
func (_Main *MainFilterer) FilterRemovedConstant(opts *bind.FilterOpts) (*MainRemovedConstantIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RemovedConstant")
	if err != nil {
		return nil, err
	}
	return &MainRemovedConstantIterator{contract: _Main.contract, event: "RemovedConstant", logs: logs, sub: sub}, nil
}

// WatchRemovedConstant is a free log subscription operation binding the contract event 0x756b8e626fd3514f249eda10fac2a5342e23cbabdb29803c2c8743bd3489afb1.
//
// Solidity: event RemovedConstant(string name)
func (_Main *MainFilterer) WatchRemovedConstant(opts *bind.WatchOpts, sink chan<- *MainRemovedConstant) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RemovedConstant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRemovedConstant)
				if err := _Main.contract.UnpackLog(event, "RemovedConstant", log); err != nil {
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

// ParseRemovedConstant is a log parse operation binding the contract event 0x756b8e626fd3514f249eda10fac2a5342e23cbabdb29803c2c8743bd3489afb1.
//
// Solidity: event RemovedConstant(string name)
func (_Main *MainFilterer) ParseRemovedConstant(log types.Log) (*MainRemovedConstant, error) {
	event := new(MainRemovedConstant)
	if err := _Main.contract.UnpackLog(event, "RemovedConstant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
