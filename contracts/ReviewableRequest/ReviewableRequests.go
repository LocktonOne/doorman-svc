// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ReviewableRequest

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"acceptData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rejectData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestDropped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"acceptData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rejectData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"RequestUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"}],\"name\":\"acceptRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"acceptData_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rejectData_\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"description_\",\"type\":\"string\"}],\"name\":\"createRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"}],\"name\":\"dropRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_injector\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"}],\"name\":\"rejectRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"enumIReviewableRequests.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"acceptData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rejectData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress_\",\"type\":\"address\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_injector\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"executor_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"acceptData_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rejectData_\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"description_\",\"type\":\"string\"}],\"name\":\"updateRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_Main *MainCaller) NextRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "nextRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_Main *MainSession) NextRequestId() (*big.Int, error) {
	return _Main.Contract.NextRequestId(&_Main.CallOpts)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_Main *MainCallerSession) NextRequestId() (*big.Int, error) {
	return _Main.Contract.NextRequestId(&_Main.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint8 status, address creator, address executor, bytes acceptData, bytes rejectData)
func (_Main *MainCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status     uint8
	Creator    common.Address
	Executor   common.Address
	AcceptData []byte
	RejectData []byte
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		Status     uint8
		Creator    common.Address
		Executor   common.Address
		AcceptData []byte
		RejectData []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Creator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Executor = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.AcceptData = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.RejectData = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint8 status, address creator, address executor, bytes acceptData, bytes rejectData)
func (_Main *MainSession) Requests(arg0 *big.Int) (struct {
	Status     uint8
	Creator    common.Address
	Executor   common.Address
	AcceptData []byte
	RejectData []byte
}, error) {
	return _Main.Contract.Requests(&_Main.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint8 status, address creator, address executor, bytes acceptData, bytes rejectData)
func (_Main *MainCallerSession) Requests(arg0 *big.Int) (struct {
	Status     uint8
	Creator    common.Address
	Executor   common.Address
	AcceptData []byte
	RejectData []byte
}, error) {
	return _Main.Contract.Requests(&_Main.CallOpts, arg0)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0x4ba1f098.
//
// Solidity: function acceptRequest(uint256 requestId_) returns()
func (_Main *MainTransactor) AcceptRequest(opts *bind.TransactOpts, requestId_ *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "acceptRequest", requestId_)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0x4ba1f098.
//
// Solidity: function acceptRequest(uint256 requestId_) returns()
func (_Main *MainSession) AcceptRequest(requestId_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AcceptRequest(&_Main.TransactOpts, requestId_)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0x4ba1f098.
//
// Solidity: function acceptRequest(uint256 requestId_) returns()
func (_Main *MainTransactorSession) AcceptRequest(requestId_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.AcceptRequest(&_Main.TransactOpts, requestId_)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x3b7a7ff0.
//
// Solidity: function createRequest(address executor_, bytes acceptData_, bytes rejectData_, string description_) returns()
func (_Main *MainTransactor) CreateRequest(opts *bind.TransactOpts, executor_ common.Address, acceptData_ []byte, rejectData_ []byte, description_ string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "createRequest", executor_, acceptData_, rejectData_, description_)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x3b7a7ff0.
//
// Solidity: function createRequest(address executor_, bytes acceptData_, bytes rejectData_, string description_) returns()
func (_Main *MainSession) CreateRequest(executor_ common.Address, acceptData_ []byte, rejectData_ []byte, description_ string) (*types.Transaction, error) {
	return _Main.Contract.CreateRequest(&_Main.TransactOpts, executor_, acceptData_, rejectData_, description_)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x3b7a7ff0.
//
// Solidity: function createRequest(address executor_, bytes acceptData_, bytes rejectData_, string description_) returns()
func (_Main *MainTransactorSession) CreateRequest(executor_ common.Address, acceptData_ []byte, rejectData_ []byte, description_ string) (*types.Transaction, error) {
	return _Main.Contract.CreateRequest(&_Main.TransactOpts, executor_, acceptData_, rejectData_, description_)
}

// DropRequest is a paid mutator transaction binding the contract method 0x2938ad27.
//
// Solidity: function dropRequest(uint256 requestId_) returns()
func (_Main *MainTransactor) DropRequest(opts *bind.TransactOpts, requestId_ *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "dropRequest", requestId_)
}

// DropRequest is a paid mutator transaction binding the contract method 0x2938ad27.
//
// Solidity: function dropRequest(uint256 requestId_) returns()
func (_Main *MainSession) DropRequest(requestId_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.DropRequest(&_Main.TransactOpts, requestId_)
}

// DropRequest is a paid mutator transaction binding the contract method 0x2938ad27.
//
// Solidity: function dropRequest(uint256 requestId_) returns()
func (_Main *MainTransactorSession) DropRequest(requestId_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.DropRequest(&_Main.TransactOpts, requestId_)
}

// RejectRequest is a paid mutator transaction binding the contract method 0x2d7788db.
//
// Solidity: function rejectRequest(uint256 requestId_) returns()
func (_Main *MainTransactor) RejectRequest(opts *bind.TransactOpts, requestId_ *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "rejectRequest", requestId_)
}

// RejectRequest is a paid mutator transaction binding the contract method 0x2d7788db.
//
// Solidity: function rejectRequest(uint256 requestId_) returns()
func (_Main *MainSession) RejectRequest(requestId_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RejectRequest(&_Main.TransactOpts, requestId_)
}

// RejectRequest is a paid mutator transaction binding the contract method 0x2d7788db.
//
// Solidity: function rejectRequest(uint256 requestId_) returns()
func (_Main *MainTransactorSession) RejectRequest(requestId_ *big.Int) (*types.Transaction, error) {
	return _Main.Contract.RejectRequest(&_Main.TransactOpts, requestId_)
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

// UpdateRequest is a paid mutator transaction binding the contract method 0xab29dc68.
//
// Solidity: function updateRequest(uint256 requestId_, address executor_, bytes acceptData_, bytes rejectData_, string description_) returns()
func (_Main *MainTransactor) UpdateRequest(opts *bind.TransactOpts, requestId_ *big.Int, executor_ common.Address, acceptData_ []byte, rejectData_ []byte, description_ string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "updateRequest", requestId_, executor_, acceptData_, rejectData_, description_)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xab29dc68.
//
// Solidity: function updateRequest(uint256 requestId_, address executor_, bytes acceptData_, bytes rejectData_, string description_) returns()
func (_Main *MainSession) UpdateRequest(requestId_ *big.Int, executor_ common.Address, acceptData_ []byte, rejectData_ []byte, description_ string) (*types.Transaction, error) {
	return _Main.Contract.UpdateRequest(&_Main.TransactOpts, requestId_, executor_, acceptData_, rejectData_, description_)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xab29dc68.
//
// Solidity: function updateRequest(uint256 requestId_, address executor_, bytes acceptData_, bytes rejectData_, string description_) returns()
func (_Main *MainTransactorSession) UpdateRequest(requestId_ *big.Int, executor_ common.Address, acceptData_ []byte, rejectData_ []byte, description_ string) (*types.Transaction, error) {
	return _Main.Contract.UpdateRequest(&_Main.TransactOpts, requestId_, executor_, acceptData_, rejectData_, description_)
}

// MainRequestAcceptedIterator is returned from FilterRequestAccepted and is used to iterate over the raw logs and unpacked data for RequestAccepted events raised by the Main contract.
type MainRequestAcceptedIterator struct {
	Event *MainRequestAccepted // Event containing the contract specifics and raw log

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
func (it *MainRequestAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRequestAccepted)
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
		it.Event = new(MainRequestAccepted)
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
func (it *MainRequestAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRequestAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRequestAccepted represents a RequestAccepted event raised by the Main contract.
type MainRequestAccepted struct {
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestAccepted is a free log retrieval operation binding the contract event 0xfb12b5b7f0394ce9e6aef75030196cd5bc5a07eacd25ab060968f67586901267.
//
// Solidity: event RequestAccepted(uint256 requestId)
func (_Main *MainFilterer) FilterRequestAccepted(opts *bind.FilterOpts) (*MainRequestAcceptedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RequestAccepted")
	if err != nil {
		return nil, err
	}
	return &MainRequestAcceptedIterator{contract: _Main.contract, event: "RequestAccepted", logs: logs, sub: sub}, nil
}

// WatchRequestAccepted is a free log subscription operation binding the contract event 0xfb12b5b7f0394ce9e6aef75030196cd5bc5a07eacd25ab060968f67586901267.
//
// Solidity: event RequestAccepted(uint256 requestId)
func (_Main *MainFilterer) WatchRequestAccepted(opts *bind.WatchOpts, sink chan<- *MainRequestAccepted) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RequestAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRequestAccepted)
				if err := _Main.contract.UnpackLog(event, "RequestAccepted", log); err != nil {
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

// ParseRequestAccepted is a log parse operation binding the contract event 0xfb12b5b7f0394ce9e6aef75030196cd5bc5a07eacd25ab060968f67586901267.
//
// Solidity: event RequestAccepted(uint256 requestId)
func (_Main *MainFilterer) ParseRequestAccepted(log types.Log) (*MainRequestAccepted, error) {
	event := new(MainRequestAccepted)
	if err := _Main.contract.UnpackLog(event, "RequestAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the Main contract.
type MainRequestCreatedIterator struct {
	Event *MainRequestCreated // Event containing the contract specifics and raw log

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
func (it *MainRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRequestCreated)
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
		it.Event = new(MainRequestCreated)
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
func (it *MainRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRequestCreated represents a RequestCreated event raised by the Main contract.
type MainRequestCreated struct {
	RequestId   *big.Int
	Creator     common.Address
	Executor    common.Address
	AcceptData  []byte
	RejectData  []byte
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0xa1e746593d17a65b84771d8c234ca178a9e5500aa5ce58c2960b36c16975a40b.
//
// Solidity: event RequestCreated(uint256 requestId, address creator, address executor, bytes acceptData, bytes rejectData, string description)
func (_Main *MainFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*MainRequestCreatedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &MainRequestCreatedIterator{contract: _Main.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0xa1e746593d17a65b84771d8c234ca178a9e5500aa5ce58c2960b36c16975a40b.
//
// Solidity: event RequestCreated(uint256 requestId, address creator, address executor, bytes acceptData, bytes rejectData, string description)
func (_Main *MainFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *MainRequestCreated) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRequestCreated)
				if err := _Main.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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

// ParseRequestCreated is a log parse operation binding the contract event 0xa1e746593d17a65b84771d8c234ca178a9e5500aa5ce58c2960b36c16975a40b.
//
// Solidity: event RequestCreated(uint256 requestId, address creator, address executor, bytes acceptData, bytes rejectData, string description)
func (_Main *MainFilterer) ParseRequestCreated(log types.Log) (*MainRequestCreated, error) {
	event := new(MainRequestCreated)
	if err := _Main.contract.UnpackLog(event, "RequestCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRequestDroppedIterator is returned from FilterRequestDropped and is used to iterate over the raw logs and unpacked data for RequestDropped events raised by the Main contract.
type MainRequestDroppedIterator struct {
	Event *MainRequestDropped // Event containing the contract specifics and raw log

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
func (it *MainRequestDroppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRequestDropped)
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
		it.Event = new(MainRequestDropped)
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
func (it *MainRequestDroppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRequestDroppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRequestDropped represents a RequestDropped event raised by the Main contract.
type MainRequestDropped struct {
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestDropped is a free log retrieval operation binding the contract event 0x51470b97eff85c6e931350460943bd60fb87a4cd33a5eae3b4c3b686ac28df8b.
//
// Solidity: event RequestDropped(uint256 requestId)
func (_Main *MainFilterer) FilterRequestDropped(opts *bind.FilterOpts) (*MainRequestDroppedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RequestDropped")
	if err != nil {
		return nil, err
	}
	return &MainRequestDroppedIterator{contract: _Main.contract, event: "RequestDropped", logs: logs, sub: sub}, nil
}

// WatchRequestDropped is a free log subscription operation binding the contract event 0x51470b97eff85c6e931350460943bd60fb87a4cd33a5eae3b4c3b686ac28df8b.
//
// Solidity: event RequestDropped(uint256 requestId)
func (_Main *MainFilterer) WatchRequestDropped(opts *bind.WatchOpts, sink chan<- *MainRequestDropped) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RequestDropped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRequestDropped)
				if err := _Main.contract.UnpackLog(event, "RequestDropped", log); err != nil {
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

// ParseRequestDropped is a log parse operation binding the contract event 0x51470b97eff85c6e931350460943bd60fb87a4cd33a5eae3b4c3b686ac28df8b.
//
// Solidity: event RequestDropped(uint256 requestId)
func (_Main *MainFilterer) ParseRequestDropped(log types.Log) (*MainRequestDropped, error) {
	event := new(MainRequestDropped)
	if err := _Main.contract.UnpackLog(event, "RequestDropped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRequestRejectedIterator is returned from FilterRequestRejected and is used to iterate over the raw logs and unpacked data for RequestRejected events raised by the Main contract.
type MainRequestRejectedIterator struct {
	Event *MainRequestRejected // Event containing the contract specifics and raw log

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
func (it *MainRequestRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRequestRejected)
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
		it.Event = new(MainRequestRejected)
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
func (it *MainRequestRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRequestRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRequestRejected represents a RequestRejected event raised by the Main contract.
type MainRequestRejected struct {
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestRejected is a free log retrieval operation binding the contract event 0x1fdac08e93f72ff55b24ebdf4152b7f56b98982fe9249639c1010cc52aebccd6.
//
// Solidity: event RequestRejected(uint256 requestId)
func (_Main *MainFilterer) FilterRequestRejected(opts *bind.FilterOpts) (*MainRequestRejectedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RequestRejected")
	if err != nil {
		return nil, err
	}
	return &MainRequestRejectedIterator{contract: _Main.contract, event: "RequestRejected", logs: logs, sub: sub}, nil
}

// WatchRequestRejected is a free log subscription operation binding the contract event 0x1fdac08e93f72ff55b24ebdf4152b7f56b98982fe9249639c1010cc52aebccd6.
//
// Solidity: event RequestRejected(uint256 requestId)
func (_Main *MainFilterer) WatchRequestRejected(opts *bind.WatchOpts, sink chan<- *MainRequestRejected) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RequestRejected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRequestRejected)
				if err := _Main.contract.UnpackLog(event, "RequestRejected", log); err != nil {
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

// ParseRequestRejected is a log parse operation binding the contract event 0x1fdac08e93f72ff55b24ebdf4152b7f56b98982fe9249639c1010cc52aebccd6.
//
// Solidity: event RequestRejected(uint256 requestId)
func (_Main *MainFilterer) ParseRequestRejected(log types.Log) (*MainRequestRejected, error) {
	event := new(MainRequestRejected)
	if err := _Main.contract.UnpackLog(event, "RequestRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRequestUpdatedIterator is returned from FilterRequestUpdated and is used to iterate over the raw logs and unpacked data for RequestUpdated events raised by the Main contract.
type MainRequestUpdatedIterator struct {
	Event *MainRequestUpdated // Event containing the contract specifics and raw log

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
func (it *MainRequestUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRequestUpdated)
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
		it.Event = new(MainRequestUpdated)
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
func (it *MainRequestUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRequestUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRequestUpdated represents a RequestUpdated event raised by the Main contract.
type MainRequestUpdated struct {
	RequestId    *big.Int
	NewRequestId *big.Int
	Executor     common.Address
	AcceptData   []byte
	RejectData   []byte
	Description  string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRequestUpdated is a free log retrieval operation binding the contract event 0x4cc21576703dec37e543715963bb2a4469a0e340e29e56e22821fbc8e0663f18.
//
// Solidity: event RequestUpdated(uint256 requestId, uint256 newRequestId, address executor, bytes acceptData, bytes rejectData, string description)
func (_Main *MainFilterer) FilterRequestUpdated(opts *bind.FilterOpts) (*MainRequestUpdatedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RequestUpdated")
	if err != nil {
		return nil, err
	}
	return &MainRequestUpdatedIterator{contract: _Main.contract, event: "RequestUpdated", logs: logs, sub: sub}, nil
}

// WatchRequestUpdated is a free log subscription operation binding the contract event 0x4cc21576703dec37e543715963bb2a4469a0e340e29e56e22821fbc8e0663f18.
//
// Solidity: event RequestUpdated(uint256 requestId, uint256 newRequestId, address executor, bytes acceptData, bytes rejectData, string description)
func (_Main *MainFilterer) WatchRequestUpdated(opts *bind.WatchOpts, sink chan<- *MainRequestUpdated) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RequestUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRequestUpdated)
				if err := _Main.contract.UnpackLog(event, "RequestUpdated", log); err != nil {
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

// ParseRequestUpdated is a log parse operation binding the contract event 0x4cc21576703dec37e543715963bb2a4469a0e340e29e56e22821fbc8e0663f18.
//
// Solidity: event RequestUpdated(uint256 requestId, uint256 newRequestId, address executor, bytes acceptData, bytes rejectData, string description)
func (_Main *MainFilterer) ParseRequestUpdated(log types.Log) (*MainRequestUpdated, error) {
	event := new(MainRequestUpdated)
	if err := _Main.contract.UnpackLog(event, "RequestUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
