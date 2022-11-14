// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package constants_registry

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

// ConstantsRegistryMetaData contains all meta data concerning the ConstantsRegistry contract.
var ConstantsRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"AddedConstant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"RemovedConstant\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key_\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value_\",\"type\":\"bytes\"}],\"name\":\"addConstant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"constants\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_injector\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key_\",\"type\":\"string\"}],\"name\":\"removeConstant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress_\",\"type\":\"address\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_injector\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ConstantsRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ConstantsRegistryMetaData.ABI instead.
var ConstantsRegistryABI = ConstantsRegistryMetaData.ABI

// ConstantsRegistry is an auto generated Go binding around an Ethereum contract.
type ConstantsRegistry struct {
	ConstantsRegistryCaller     // Read-only binding to the contract
	ConstantsRegistryTransactor // Write-only binding to the contract
	ConstantsRegistryFilterer   // Log filterer for contract events
}

// ConstantsRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConstantsRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConstantsRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConstantsRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConstantsRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConstantsRegistrySession struct {
	Contract     *ConstantsRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConstantsRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConstantsRegistryCallerSession struct {
	Contract *ConstantsRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ConstantsRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConstantsRegistryTransactorSession struct {
	Contract     *ConstantsRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ConstantsRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConstantsRegistryRaw struct {
	Contract *ConstantsRegistry // Generic contract binding to access the raw methods on
}

// ConstantsRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConstantsRegistryCallerRaw struct {
	Contract *ConstantsRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ConstantsRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConstantsRegistryTransactorRaw struct {
	Contract *ConstantsRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConstantsRegistry creates a new instance of ConstantsRegistry, bound to a specific deployed contract.
func NewConstantsRegistry(address common.Address, backend bind.ContractBackend) (*ConstantsRegistry, error) {
	contract, err := bindConstantsRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConstantsRegistry{ConstantsRegistryCaller: ConstantsRegistryCaller{contract: contract}, ConstantsRegistryTransactor: ConstantsRegistryTransactor{contract: contract}, ConstantsRegistryFilterer: ConstantsRegistryFilterer{contract: contract}}, nil
}

// NewConstantsRegistryCaller creates a new read-only instance of ConstantsRegistry, bound to a specific deployed contract.
func NewConstantsRegistryCaller(address common.Address, caller bind.ContractCaller) (*ConstantsRegistryCaller, error) {
	contract, err := bindConstantsRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantsRegistryCaller{contract: contract}, nil
}

// NewConstantsRegistryTransactor creates a new write-only instance of ConstantsRegistry, bound to a specific deployed contract.
func NewConstantsRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ConstantsRegistryTransactor, error) {
	contract, err := bindConstantsRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConstantsRegistryTransactor{contract: contract}, nil
}

// NewConstantsRegistryFilterer creates a new log filterer instance of ConstantsRegistry, bound to a specific deployed contract.
func NewConstantsRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ConstantsRegistryFilterer, error) {
	contract, err := bindConstantsRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConstantsRegistryFilterer{contract: contract}, nil
}

// bindConstantsRegistry binds a generic wrapper to an already deployed contract.
func bindConstantsRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConstantsRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConstantsRegistry *ConstantsRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConstantsRegistry.Contract.ConstantsRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConstantsRegistry *ConstantsRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConstantsRegistry.Contract.ConstantsRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConstantsRegistry *ConstantsRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConstantsRegistry.Contract.ConstantsRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConstantsRegistry *ConstantsRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConstantsRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConstantsRegistry *ConstantsRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConstantsRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConstantsRegistry *ConstantsRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConstantsRegistry.Contract.contract.Transact(opts, method, params...)
}

// Constants is a free data retrieval call binding the contract method 0x163f33a5.
//
// Solidity: function constants(string ) view returns(bytes)
func (_ConstantsRegistry *ConstantsRegistryCaller) Constants(opts *bind.CallOpts, arg0 string) ([]byte, error) {
	var out []interface{}
	err := _ConstantsRegistry.contract.Call(opts, &out, "constants", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Constants is a free data retrieval call binding the contract method 0x163f33a5.
//
// Solidity: function constants(string ) view returns(bytes)
func (_ConstantsRegistry *ConstantsRegistrySession) Constants(arg0 string) ([]byte, error) {
	return _ConstantsRegistry.Contract.Constants(&_ConstantsRegistry.CallOpts, arg0)
}

// Constants is a free data retrieval call binding the contract method 0x163f33a5.
//
// Solidity: function constants(string ) view returns(bytes)
func (_ConstantsRegistry *ConstantsRegistryCallerSession) Constants(arg0 string) ([]byte, error) {
	return _ConstantsRegistry.Contract.Constants(&_ConstantsRegistry.CallOpts, arg0)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address _injector)
func (_ConstantsRegistry *ConstantsRegistryCaller) GetInjector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConstantsRegistry.contract.Call(opts, &out, "getInjector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address _injector)
func (_ConstantsRegistry *ConstantsRegistrySession) GetInjector() (common.Address, error) {
	return _ConstantsRegistry.Contract.GetInjector(&_ConstantsRegistry.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address _injector)
func (_ConstantsRegistry *ConstantsRegistryCallerSession) GetInjector() (common.Address, error) {
	return _ConstantsRegistry.Contract.GetInjector(&_ConstantsRegistry.CallOpts)
}

// AddConstant is a paid mutator transaction binding the contract method 0x4216ab3b.
//
// Solidity: function addConstant(string key_, bytes value_) returns()
func (_ConstantsRegistry *ConstantsRegistryTransactor) AddConstant(opts *bind.TransactOpts, key_ string, value_ []byte) (*types.Transaction, error) {
	return _ConstantsRegistry.contract.Transact(opts, "addConstant", key_, value_)
}

// AddConstant is a paid mutator transaction binding the contract method 0x4216ab3b.
//
// Solidity: function addConstant(string key_, bytes value_) returns()
func (_ConstantsRegistry *ConstantsRegistrySession) AddConstant(key_ string, value_ []byte) (*types.Transaction, error) {
	return _ConstantsRegistry.Contract.AddConstant(&_ConstantsRegistry.TransactOpts, key_, value_)
}

// AddConstant is a paid mutator transaction binding the contract method 0x4216ab3b.
//
// Solidity: function addConstant(string key_, bytes value_) returns()
func (_ConstantsRegistry *ConstantsRegistryTransactorSession) AddConstant(key_ string, value_ []byte) (*types.Transaction, error) {
	return _ConstantsRegistry.Contract.AddConstant(&_ConstantsRegistry.TransactOpts, key_, value_)
}

// RemoveConstant is a paid mutator transaction binding the contract method 0x965136fd.
//
// Solidity: function removeConstant(string key_) returns()
func (_ConstantsRegistry *ConstantsRegistryTransactor) RemoveConstant(opts *bind.TransactOpts, key_ string) (*types.Transaction, error) {
	return _ConstantsRegistry.contract.Transact(opts, "removeConstant", key_)
}

// RemoveConstant is a paid mutator transaction binding the contract method 0x965136fd.
//
// Solidity: function removeConstant(string key_) returns()
func (_ConstantsRegistry *ConstantsRegistrySession) RemoveConstant(key_ string) (*types.Transaction, error) {
	return _ConstantsRegistry.Contract.RemoveConstant(&_ConstantsRegistry.TransactOpts, key_)
}

// RemoveConstant is a paid mutator transaction binding the contract method 0x965136fd.
//
// Solidity: function removeConstant(string key_) returns()
func (_ConstantsRegistry *ConstantsRegistryTransactorSession) RemoveConstant(key_ string) (*types.Transaction, error) {
	return _ConstantsRegistry.Contract.RemoveConstant(&_ConstantsRegistry.TransactOpts, key_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x8389cb18.
//
// Solidity: function setDependencies(address registryAddress_) returns()
func (_ConstantsRegistry *ConstantsRegistryTransactor) SetDependencies(opts *bind.TransactOpts, registryAddress_ common.Address) (*types.Transaction, error) {
	return _ConstantsRegistry.contract.Transact(opts, "setDependencies", registryAddress_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x8389cb18.
//
// Solidity: function setDependencies(address registryAddress_) returns()
func (_ConstantsRegistry *ConstantsRegistrySession) SetDependencies(registryAddress_ common.Address) (*types.Transaction, error) {
	return _ConstantsRegistry.Contract.SetDependencies(&_ConstantsRegistry.TransactOpts, registryAddress_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x8389cb18.
//
// Solidity: function setDependencies(address registryAddress_) returns()
func (_ConstantsRegistry *ConstantsRegistryTransactorSession) SetDependencies(registryAddress_ common.Address) (*types.Transaction, error) {
	return _ConstantsRegistry.Contract.SetDependencies(&_ConstantsRegistry.TransactOpts, registryAddress_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address _injector) returns()
func (_ConstantsRegistry *ConstantsRegistryTransactor) SetInjector(opts *bind.TransactOpts, _injector common.Address) (*types.Transaction, error) {
	return _ConstantsRegistry.contract.Transact(opts, "setInjector", _injector)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address _injector) returns()
func (_ConstantsRegistry *ConstantsRegistrySession) SetInjector(_injector common.Address) (*types.Transaction, error) {
	return _ConstantsRegistry.Contract.SetInjector(&_ConstantsRegistry.TransactOpts, _injector)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address _injector) returns()
func (_ConstantsRegistry *ConstantsRegistryTransactorSession) SetInjector(_injector common.Address) (*types.Transaction, error) {
	return _ConstantsRegistry.Contract.SetInjector(&_ConstantsRegistry.TransactOpts, _injector)
}

// ConstantsRegistryAddedConstantIterator is returned from FilterAddedConstant and is used to iterate over the raw logs and unpacked data for AddedConstant events raised by the ConstantsRegistry contract.
type ConstantsRegistryAddedConstantIterator struct {
	Event *ConstantsRegistryAddedConstant // Event containing the contract specifics and raw log

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
func (it *ConstantsRegistryAddedConstantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstantsRegistryAddedConstant)
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
		it.Event = new(ConstantsRegistryAddedConstant)
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
func (it *ConstantsRegistryAddedConstantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstantsRegistryAddedConstantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstantsRegistryAddedConstant represents a AddedConstant event raised by the ConstantsRegistry contract.
type ConstantsRegistryAddedConstant struct {
	Name  string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddedConstant is a free log retrieval operation binding the contract event 0x265476163f06971fd1d163653ba47af9e018395e1df6b521f92de499355ee41a.
//
// Solidity: event AddedConstant(string name, bytes value)
func (_ConstantsRegistry *ConstantsRegistryFilterer) FilterAddedConstant(opts *bind.FilterOpts) (*ConstantsRegistryAddedConstantIterator, error) {

	logs, sub, err := _ConstantsRegistry.contract.FilterLogs(opts, "AddedConstant")
	if err != nil {
		return nil, err
	}
	return &ConstantsRegistryAddedConstantIterator{contract: _ConstantsRegistry.contract, event: "AddedConstant", logs: logs, sub: sub}, nil
}

// WatchAddedConstant is a free log subscription operation binding the contract event 0x265476163f06971fd1d163653ba47af9e018395e1df6b521f92de499355ee41a.
//
// Solidity: event AddedConstant(string name, bytes value)
func (_ConstantsRegistry *ConstantsRegistryFilterer) WatchAddedConstant(opts *bind.WatchOpts, sink chan<- *ConstantsRegistryAddedConstant) (event.Subscription, error) {

	logs, sub, err := _ConstantsRegistry.contract.WatchLogs(opts, "AddedConstant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstantsRegistryAddedConstant)
				if err := _ConstantsRegistry.contract.UnpackLog(event, "AddedConstant", log); err != nil {
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
func (_ConstantsRegistry *ConstantsRegistryFilterer) ParseAddedConstant(log types.Log) (*ConstantsRegistryAddedConstant, error) {
	event := new(ConstantsRegistryAddedConstant)
	if err := _ConstantsRegistry.contract.UnpackLog(event, "AddedConstant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConstantsRegistryRemovedConstantIterator is returned from FilterRemovedConstant and is used to iterate over the raw logs and unpacked data for RemovedConstant events raised by the ConstantsRegistry contract.
type ConstantsRegistryRemovedConstantIterator struct {
	Event *ConstantsRegistryRemovedConstant // Event containing the contract specifics and raw log

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
func (it *ConstantsRegistryRemovedConstantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConstantsRegistryRemovedConstant)
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
		it.Event = new(ConstantsRegistryRemovedConstant)
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
func (it *ConstantsRegistryRemovedConstantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConstantsRegistryRemovedConstantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConstantsRegistryRemovedConstant represents a RemovedConstant event raised by the ConstantsRegistry contract.
type ConstantsRegistryRemovedConstant struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedConstant is a free log retrieval operation binding the contract event 0x756b8e626fd3514f249eda10fac2a5342e23cbabdb29803c2c8743bd3489afb1.
//
// Solidity: event RemovedConstant(string name)
func (_ConstantsRegistry *ConstantsRegistryFilterer) FilterRemovedConstant(opts *bind.FilterOpts) (*ConstantsRegistryRemovedConstantIterator, error) {

	logs, sub, err := _ConstantsRegistry.contract.FilterLogs(opts, "RemovedConstant")
	if err != nil {
		return nil, err
	}
	return &ConstantsRegistryRemovedConstantIterator{contract: _ConstantsRegistry.contract, event: "RemovedConstant", logs: logs, sub: sub}, nil
}

// WatchRemovedConstant is a free log subscription operation binding the contract event 0x756b8e626fd3514f249eda10fac2a5342e23cbabdb29803c2c8743bd3489afb1.
//
// Solidity: event RemovedConstant(string name)
func (_ConstantsRegistry *ConstantsRegistryFilterer) WatchRemovedConstant(opts *bind.WatchOpts, sink chan<- *ConstantsRegistryRemovedConstant) (event.Subscription, error) {

	logs, sub, err := _ConstantsRegistry.contract.WatchLogs(opts, "RemovedConstant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConstantsRegistryRemovedConstant)
				if err := _ConstantsRegistry.contract.UnpackLog(event, "RemovedConstant", log); err != nil {
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
func (_ConstantsRegistry *ConstantsRegistryFilterer) ParseRemovedConstant(log types.Log) (*ConstantsRegistryRemovedConstant, error) {
	event := new(ConstantsRegistryRemovedConstant)
	if err := _ConstantsRegistry.contract.UnpackLog(event, "RemovedConstant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
