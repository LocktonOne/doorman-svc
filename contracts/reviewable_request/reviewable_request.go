// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reviewable_request

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

// ReviewableRequestMetaData contains all meta data concerning the ReviewableRequest contract.
var ReviewableRequestMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"acceptData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rejectData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestDropped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"acceptData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rejectData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"RequestUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"}],\"name\":\"acceptRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executor_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"acceptData_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rejectData_\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"description_\",\"type\":\"string\"}],\"name\":\"createRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"}],\"name\":\"dropRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_injector\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"}],\"name\":\"rejectRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"enumIReviewableRequests.RequestStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"acceptData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rejectData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress_\",\"type\":\"address\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_injector\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"executor_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"acceptData_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rejectData_\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"description_\",\"type\":\"string\"}],\"name\":\"updateRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ReviewableRequestABI is the input ABI used to generate the binding from.
// Deprecated: Use ReviewableRequestMetaData.ABI instead.
var ReviewableRequestABI = ReviewableRequestMetaData.ABI

// ReviewableRequest is an auto generated Go binding around an Ethereum contract.
type ReviewableRequest struct {
	ReviewableRequestCaller     // Read-only binding to the contract
	ReviewableRequestTransactor // Write-only binding to the contract
	ReviewableRequestFilterer   // Log filterer for contract events
}

// ReviewableRequestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReviewableRequestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReviewableRequestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReviewableRequestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReviewableRequestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReviewableRequestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReviewableRequestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReviewableRequestSession struct {
	Contract     *ReviewableRequest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReviewableRequestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReviewableRequestCallerSession struct {
	Contract *ReviewableRequestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ReviewableRequestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReviewableRequestTransactorSession struct {
	Contract     *ReviewableRequestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ReviewableRequestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReviewableRequestRaw struct {
	Contract *ReviewableRequest // Generic contract binding to access the raw methods on
}

// ReviewableRequestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReviewableRequestCallerRaw struct {
	Contract *ReviewableRequestCaller // Generic read-only contract binding to access the raw methods on
}

// ReviewableRequestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReviewableRequestTransactorRaw struct {
	Contract *ReviewableRequestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReviewableRequest creates a new instance of ReviewableRequest, bound to a specific deployed contract.
func NewReviewableRequest(address common.Address, backend bind.ContractBackend) (*ReviewableRequest, error) {
	contract, err := bindReviewableRequest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReviewableRequest{ReviewableRequestCaller: ReviewableRequestCaller{contract: contract}, ReviewableRequestTransactor: ReviewableRequestTransactor{contract: contract}, ReviewableRequestFilterer: ReviewableRequestFilterer{contract: contract}}, nil
}

// NewReviewableRequestCaller creates a new read-only instance of ReviewableRequest, bound to a specific deployed contract.
func NewReviewableRequestCaller(address common.Address, caller bind.ContractCaller) (*ReviewableRequestCaller, error) {
	contract, err := bindReviewableRequest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReviewableRequestCaller{contract: contract}, nil
}

// NewReviewableRequestTransactor creates a new write-only instance of ReviewableRequest, bound to a specific deployed contract.
func NewReviewableRequestTransactor(address common.Address, transactor bind.ContractTransactor) (*ReviewableRequestTransactor, error) {
	contract, err := bindReviewableRequest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReviewableRequestTransactor{contract: contract}, nil
}

// NewReviewableRequestFilterer creates a new log filterer instance of ReviewableRequest, bound to a specific deployed contract.
func NewReviewableRequestFilterer(address common.Address, filterer bind.ContractFilterer) (*ReviewableRequestFilterer, error) {
	contract, err := bindReviewableRequest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReviewableRequestFilterer{contract: contract}, nil
}

// bindReviewableRequest binds a generic wrapper to an already deployed contract.
func bindReviewableRequest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReviewableRequestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReviewableRequest *ReviewableRequestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReviewableRequest.Contract.ReviewableRequestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReviewableRequest *ReviewableRequestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.ReviewableRequestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReviewableRequest *ReviewableRequestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.ReviewableRequestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReviewableRequest *ReviewableRequestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReviewableRequest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReviewableRequest *ReviewableRequestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReviewableRequest *ReviewableRequestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.contract.Transact(opts, method, params...)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address _injector)
func (_ReviewableRequest *ReviewableRequestCaller) GetInjector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReviewableRequest.contract.Call(opts, &out, "getInjector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address _injector)
func (_ReviewableRequest *ReviewableRequestSession) GetInjector() (common.Address, error) {
	return _ReviewableRequest.Contract.GetInjector(&_ReviewableRequest.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address _injector)
func (_ReviewableRequest *ReviewableRequestCallerSession) GetInjector() (common.Address, error) {
	return _ReviewableRequest.Contract.GetInjector(&_ReviewableRequest.CallOpts)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_ReviewableRequest *ReviewableRequestCaller) NextRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReviewableRequest.contract.Call(opts, &out, "nextRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_ReviewableRequest *ReviewableRequestSession) NextRequestId() (*big.Int, error) {
	return _ReviewableRequest.Contract.NextRequestId(&_ReviewableRequest.CallOpts)
}

// NextRequestId is a free data retrieval call binding the contract method 0x6a84a985.
//
// Solidity: function nextRequestId() view returns(uint256)
func (_ReviewableRequest *ReviewableRequestCallerSession) NextRequestId() (*big.Int, error) {
	return _ReviewableRequest.Contract.NextRequestId(&_ReviewableRequest.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint8 status, address creator, address executor, bytes acceptData, bytes rejectData)
func (_ReviewableRequest *ReviewableRequestCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status     uint8
	Creator    common.Address
	Executor   common.Address
	AcceptData []byte
	RejectData []byte
}, error) {
	var out []interface{}
	err := _ReviewableRequest.contract.Call(opts, &out, "requests", arg0)

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
func (_ReviewableRequest *ReviewableRequestSession) Requests(arg0 *big.Int) (struct {
	Status     uint8
	Creator    common.Address
	Executor   common.Address
	AcceptData []byte
	RejectData []byte
}, error) {
	return _ReviewableRequest.Contract.Requests(&_ReviewableRequest.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint8 status, address creator, address executor, bytes acceptData, bytes rejectData)
func (_ReviewableRequest *ReviewableRequestCallerSession) Requests(arg0 *big.Int) (struct {
	Status     uint8
	Creator    common.Address
	Executor   common.Address
	AcceptData []byte
	RejectData []byte
}, error) {
	return _ReviewableRequest.Contract.Requests(&_ReviewableRequest.CallOpts, arg0)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0x4ba1f098.
//
// Solidity: function acceptRequest(uint256 requestId_) returns()
func (_ReviewableRequest *ReviewableRequestTransactor) AcceptRequest(opts *bind.TransactOpts, requestId_ *big.Int) (*types.Transaction, error) {
	return _ReviewableRequest.contract.Transact(opts, "acceptRequest", requestId_)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0x4ba1f098.
//
// Solidity: function acceptRequest(uint256 requestId_) returns()
func (_ReviewableRequest *ReviewableRequestSession) AcceptRequest(requestId_ *big.Int) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.AcceptRequest(&_ReviewableRequest.TransactOpts, requestId_)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0x4ba1f098.
//
// Solidity: function acceptRequest(uint256 requestId_) returns()
func (_ReviewableRequest *ReviewableRequestTransactorSession) AcceptRequest(requestId_ *big.Int) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.AcceptRequest(&_ReviewableRequest.TransactOpts, requestId_)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x3b7a7ff0.
//
// Solidity: function createRequest(address executor_, bytes acceptData_, bytes rejectData_, string description_) returns()
func (_ReviewableRequest *ReviewableRequestTransactor) CreateRequest(opts *bind.TransactOpts, executor_ common.Address, acceptData_ []byte, rejectData_ []byte, description_ string) (*types.Transaction, error) {
	return _ReviewableRequest.contract.Transact(opts, "createRequest", executor_, acceptData_, rejectData_, description_)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x3b7a7ff0.
//
// Solidity: function createRequest(address executor_, bytes acceptData_, bytes rejectData_, string description_) returns()
func (_ReviewableRequest *ReviewableRequestSession) CreateRequest(executor_ common.Address, acceptData_ []byte, rejectData_ []byte, description_ string) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.CreateRequest(&_ReviewableRequest.TransactOpts, executor_, acceptData_, rejectData_, description_)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x3b7a7ff0.
//
// Solidity: function createRequest(address executor_, bytes acceptData_, bytes rejectData_, string description_) returns()
func (_ReviewableRequest *ReviewableRequestTransactorSession) CreateRequest(executor_ common.Address, acceptData_ []byte, rejectData_ []byte, description_ string) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.CreateRequest(&_ReviewableRequest.TransactOpts, executor_, acceptData_, rejectData_, description_)
}

// DropRequest is a paid mutator transaction binding the contract method 0x2938ad27.
//
// Solidity: function dropRequest(uint256 requestId_) returns()
func (_ReviewableRequest *ReviewableRequestTransactor) DropRequest(opts *bind.TransactOpts, requestId_ *big.Int) (*types.Transaction, error) {
	return _ReviewableRequest.contract.Transact(opts, "dropRequest", requestId_)
}

// DropRequest is a paid mutator transaction binding the contract method 0x2938ad27.
//
// Solidity: function dropRequest(uint256 requestId_) returns()
func (_ReviewableRequest *ReviewableRequestSession) DropRequest(requestId_ *big.Int) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.DropRequest(&_ReviewableRequest.TransactOpts, requestId_)
}

// DropRequest is a paid mutator transaction binding the contract method 0x2938ad27.
//
// Solidity: function dropRequest(uint256 requestId_) returns()
func (_ReviewableRequest *ReviewableRequestTransactorSession) DropRequest(requestId_ *big.Int) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.DropRequest(&_ReviewableRequest.TransactOpts, requestId_)
}

// RejectRequest is a paid mutator transaction binding the contract method 0x2d7788db.
//
// Solidity: function rejectRequest(uint256 requestId_) returns()
func (_ReviewableRequest *ReviewableRequestTransactor) RejectRequest(opts *bind.TransactOpts, requestId_ *big.Int) (*types.Transaction, error) {
	return _ReviewableRequest.contract.Transact(opts, "rejectRequest", requestId_)
}

// RejectRequest is a paid mutator transaction binding the contract method 0x2d7788db.
//
// Solidity: function rejectRequest(uint256 requestId_) returns()
func (_ReviewableRequest *ReviewableRequestSession) RejectRequest(requestId_ *big.Int) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.RejectRequest(&_ReviewableRequest.TransactOpts, requestId_)
}

// RejectRequest is a paid mutator transaction binding the contract method 0x2d7788db.
//
// Solidity: function rejectRequest(uint256 requestId_) returns()
func (_ReviewableRequest *ReviewableRequestTransactorSession) RejectRequest(requestId_ *big.Int) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.RejectRequest(&_ReviewableRequest.TransactOpts, requestId_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x8389cb18.
//
// Solidity: function setDependencies(address registryAddress_) returns()
func (_ReviewableRequest *ReviewableRequestTransactor) SetDependencies(opts *bind.TransactOpts, registryAddress_ common.Address) (*types.Transaction, error) {
	return _ReviewableRequest.contract.Transact(opts, "setDependencies", registryAddress_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x8389cb18.
//
// Solidity: function setDependencies(address registryAddress_) returns()
func (_ReviewableRequest *ReviewableRequestSession) SetDependencies(registryAddress_ common.Address) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.SetDependencies(&_ReviewableRequest.TransactOpts, registryAddress_)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x8389cb18.
//
// Solidity: function setDependencies(address registryAddress_) returns()
func (_ReviewableRequest *ReviewableRequestTransactorSession) SetDependencies(registryAddress_ common.Address) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.SetDependencies(&_ReviewableRequest.TransactOpts, registryAddress_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address _injector) returns()
func (_ReviewableRequest *ReviewableRequestTransactor) SetInjector(opts *bind.TransactOpts, _injector common.Address) (*types.Transaction, error) {
	return _ReviewableRequest.contract.Transact(opts, "setInjector", _injector)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address _injector) returns()
func (_ReviewableRequest *ReviewableRequestSession) SetInjector(_injector common.Address) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.SetInjector(&_ReviewableRequest.TransactOpts, _injector)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address _injector) returns()
func (_ReviewableRequest *ReviewableRequestTransactorSession) SetInjector(_injector common.Address) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.SetInjector(&_ReviewableRequest.TransactOpts, _injector)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xab29dc68.
//
// Solidity: function updateRequest(uint256 requestId_, address executor_, bytes acceptData_, bytes rejectData_, string description_) returns()
func (_ReviewableRequest *ReviewableRequestTransactor) UpdateRequest(opts *bind.TransactOpts, requestId_ *big.Int, executor_ common.Address, acceptData_ []byte, rejectData_ []byte, description_ string) (*types.Transaction, error) {
	return _ReviewableRequest.contract.Transact(opts, "updateRequest", requestId_, executor_, acceptData_, rejectData_, description_)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xab29dc68.
//
// Solidity: function updateRequest(uint256 requestId_, address executor_, bytes acceptData_, bytes rejectData_, string description_) returns()
func (_ReviewableRequest *ReviewableRequestSession) UpdateRequest(requestId_ *big.Int, executor_ common.Address, acceptData_ []byte, rejectData_ []byte, description_ string) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.UpdateRequest(&_ReviewableRequest.TransactOpts, requestId_, executor_, acceptData_, rejectData_, description_)
}

// UpdateRequest is a paid mutator transaction binding the contract method 0xab29dc68.
//
// Solidity: function updateRequest(uint256 requestId_, address executor_, bytes acceptData_, bytes rejectData_, string description_) returns()
func (_ReviewableRequest *ReviewableRequestTransactorSession) UpdateRequest(requestId_ *big.Int, executor_ common.Address, acceptData_ []byte, rejectData_ []byte, description_ string) (*types.Transaction, error) {
	return _ReviewableRequest.Contract.UpdateRequest(&_ReviewableRequest.TransactOpts, requestId_, executor_, acceptData_, rejectData_, description_)
}

// ReviewableRequestRequestAcceptedIterator is returned from FilterRequestAccepted and is used to iterate over the raw logs and unpacked data for RequestAccepted events raised by the ReviewableRequest contract.
type ReviewableRequestRequestAcceptedIterator struct {
	Event *ReviewableRequestRequestAccepted // Event containing the contract specifics and raw log

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
func (it *ReviewableRequestRequestAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReviewableRequestRequestAccepted)
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
		it.Event = new(ReviewableRequestRequestAccepted)
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
func (it *ReviewableRequestRequestAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReviewableRequestRequestAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReviewableRequestRequestAccepted represents a RequestAccepted event raised by the ReviewableRequest contract.
type ReviewableRequestRequestAccepted struct {
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestAccepted is a free log retrieval operation binding the contract event 0xfb12b5b7f0394ce9e6aef75030196cd5bc5a07eacd25ab060968f67586901267.
//
// Solidity: event RequestAccepted(uint256 requestId)
func (_ReviewableRequest *ReviewableRequestFilterer) FilterRequestAccepted(opts *bind.FilterOpts) (*ReviewableRequestRequestAcceptedIterator, error) {

	logs, sub, err := _ReviewableRequest.contract.FilterLogs(opts, "RequestAccepted")
	if err != nil {
		return nil, err
	}
	return &ReviewableRequestRequestAcceptedIterator{contract: _ReviewableRequest.contract, event: "RequestAccepted", logs: logs, sub: sub}, nil
}

// WatchRequestAccepted is a free log subscription operation binding the contract event 0xfb12b5b7f0394ce9e6aef75030196cd5bc5a07eacd25ab060968f67586901267.
//
// Solidity: event RequestAccepted(uint256 requestId)
func (_ReviewableRequest *ReviewableRequestFilterer) WatchRequestAccepted(opts *bind.WatchOpts, sink chan<- *ReviewableRequestRequestAccepted) (event.Subscription, error) {

	logs, sub, err := _ReviewableRequest.contract.WatchLogs(opts, "RequestAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReviewableRequestRequestAccepted)
				if err := _ReviewableRequest.contract.UnpackLog(event, "RequestAccepted", log); err != nil {
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
func (_ReviewableRequest *ReviewableRequestFilterer) ParseRequestAccepted(log types.Log) (*ReviewableRequestRequestAccepted, error) {
	event := new(ReviewableRequestRequestAccepted)
	if err := _ReviewableRequest.contract.UnpackLog(event, "RequestAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReviewableRequestRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the ReviewableRequest contract.
type ReviewableRequestRequestCreatedIterator struct {
	Event *ReviewableRequestRequestCreated // Event containing the contract specifics and raw log

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
func (it *ReviewableRequestRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReviewableRequestRequestCreated)
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
		it.Event = new(ReviewableRequestRequestCreated)
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
func (it *ReviewableRequestRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReviewableRequestRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReviewableRequestRequestCreated represents a RequestCreated event raised by the ReviewableRequest contract.
type ReviewableRequestRequestCreated struct {
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
func (_ReviewableRequest *ReviewableRequestFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*ReviewableRequestRequestCreatedIterator, error) {

	logs, sub, err := _ReviewableRequest.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &ReviewableRequestRequestCreatedIterator{contract: _ReviewableRequest.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0xa1e746593d17a65b84771d8c234ca178a9e5500aa5ce58c2960b36c16975a40b.
//
// Solidity: event RequestCreated(uint256 requestId, address creator, address executor, bytes acceptData, bytes rejectData, string description)
func (_ReviewableRequest *ReviewableRequestFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *ReviewableRequestRequestCreated) (event.Subscription, error) {

	logs, sub, err := _ReviewableRequest.contract.WatchLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReviewableRequestRequestCreated)
				if err := _ReviewableRequest.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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
func (_ReviewableRequest *ReviewableRequestFilterer) ParseRequestCreated(log types.Log) (*ReviewableRequestRequestCreated, error) {
	event := new(ReviewableRequestRequestCreated)
	if err := _ReviewableRequest.contract.UnpackLog(event, "RequestCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReviewableRequestRequestDroppedIterator is returned from FilterRequestDropped and is used to iterate over the raw logs and unpacked data for RequestDropped events raised by the ReviewableRequest contract.
type ReviewableRequestRequestDroppedIterator struct {
	Event *ReviewableRequestRequestDropped // Event containing the contract specifics and raw log

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
func (it *ReviewableRequestRequestDroppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReviewableRequestRequestDropped)
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
		it.Event = new(ReviewableRequestRequestDropped)
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
func (it *ReviewableRequestRequestDroppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReviewableRequestRequestDroppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReviewableRequestRequestDropped represents a RequestDropped event raised by the ReviewableRequest contract.
type ReviewableRequestRequestDropped struct {
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestDropped is a free log retrieval operation binding the contract event 0x51470b97eff85c6e931350460943bd60fb87a4cd33a5eae3b4c3b686ac28df8b.
//
// Solidity: event RequestDropped(uint256 requestId)
func (_ReviewableRequest *ReviewableRequestFilterer) FilterRequestDropped(opts *bind.FilterOpts) (*ReviewableRequestRequestDroppedIterator, error) {

	logs, sub, err := _ReviewableRequest.contract.FilterLogs(opts, "RequestDropped")
	if err != nil {
		return nil, err
	}
	return &ReviewableRequestRequestDroppedIterator{contract: _ReviewableRequest.contract, event: "RequestDropped", logs: logs, sub: sub}, nil
}

// WatchRequestDropped is a free log subscription operation binding the contract event 0x51470b97eff85c6e931350460943bd60fb87a4cd33a5eae3b4c3b686ac28df8b.
//
// Solidity: event RequestDropped(uint256 requestId)
func (_ReviewableRequest *ReviewableRequestFilterer) WatchRequestDropped(opts *bind.WatchOpts, sink chan<- *ReviewableRequestRequestDropped) (event.Subscription, error) {

	logs, sub, err := _ReviewableRequest.contract.WatchLogs(opts, "RequestDropped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReviewableRequestRequestDropped)
				if err := _ReviewableRequest.contract.UnpackLog(event, "RequestDropped", log); err != nil {
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
func (_ReviewableRequest *ReviewableRequestFilterer) ParseRequestDropped(log types.Log) (*ReviewableRequestRequestDropped, error) {
	event := new(ReviewableRequestRequestDropped)
	if err := _ReviewableRequest.contract.UnpackLog(event, "RequestDropped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReviewableRequestRequestRejectedIterator is returned from FilterRequestRejected and is used to iterate over the raw logs and unpacked data for RequestRejected events raised by the ReviewableRequest contract.
type ReviewableRequestRequestRejectedIterator struct {
	Event *ReviewableRequestRequestRejected // Event containing the contract specifics and raw log

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
func (it *ReviewableRequestRequestRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReviewableRequestRequestRejected)
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
		it.Event = new(ReviewableRequestRequestRejected)
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
func (it *ReviewableRequestRequestRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReviewableRequestRequestRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReviewableRequestRequestRejected represents a RequestRejected event raised by the ReviewableRequest contract.
type ReviewableRequestRequestRejected struct {
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestRejected is a free log retrieval operation binding the contract event 0x1fdac08e93f72ff55b24ebdf4152b7f56b98982fe9249639c1010cc52aebccd6.
//
// Solidity: event RequestRejected(uint256 requestId)
func (_ReviewableRequest *ReviewableRequestFilterer) FilterRequestRejected(opts *bind.FilterOpts) (*ReviewableRequestRequestRejectedIterator, error) {

	logs, sub, err := _ReviewableRequest.contract.FilterLogs(opts, "RequestRejected")
	if err != nil {
		return nil, err
	}
	return &ReviewableRequestRequestRejectedIterator{contract: _ReviewableRequest.contract, event: "RequestRejected", logs: logs, sub: sub}, nil
}

// WatchRequestRejected is a free log subscription operation binding the contract event 0x1fdac08e93f72ff55b24ebdf4152b7f56b98982fe9249639c1010cc52aebccd6.
//
// Solidity: event RequestRejected(uint256 requestId)
func (_ReviewableRequest *ReviewableRequestFilterer) WatchRequestRejected(opts *bind.WatchOpts, sink chan<- *ReviewableRequestRequestRejected) (event.Subscription, error) {

	logs, sub, err := _ReviewableRequest.contract.WatchLogs(opts, "RequestRejected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReviewableRequestRequestRejected)
				if err := _ReviewableRequest.contract.UnpackLog(event, "RequestRejected", log); err != nil {
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
func (_ReviewableRequest *ReviewableRequestFilterer) ParseRequestRejected(log types.Log) (*ReviewableRequestRequestRejected, error) {
	event := new(ReviewableRequestRequestRejected)
	if err := _ReviewableRequest.contract.UnpackLog(event, "RequestRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReviewableRequestRequestUpdatedIterator is returned from FilterRequestUpdated and is used to iterate over the raw logs and unpacked data for RequestUpdated events raised by the ReviewableRequest contract.
type ReviewableRequestRequestUpdatedIterator struct {
	Event *ReviewableRequestRequestUpdated // Event containing the contract specifics and raw log

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
func (it *ReviewableRequestRequestUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReviewableRequestRequestUpdated)
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
		it.Event = new(ReviewableRequestRequestUpdated)
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
func (it *ReviewableRequestRequestUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReviewableRequestRequestUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReviewableRequestRequestUpdated represents a RequestUpdated event raised by the ReviewableRequest contract.
type ReviewableRequestRequestUpdated struct {
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
func (_ReviewableRequest *ReviewableRequestFilterer) FilterRequestUpdated(opts *bind.FilterOpts) (*ReviewableRequestRequestUpdatedIterator, error) {

	logs, sub, err := _ReviewableRequest.contract.FilterLogs(opts, "RequestUpdated")
	if err != nil {
		return nil, err
	}
	return &ReviewableRequestRequestUpdatedIterator{contract: _ReviewableRequest.contract, event: "RequestUpdated", logs: logs, sub: sub}, nil
}

// WatchRequestUpdated is a free log subscription operation binding the contract event 0x4cc21576703dec37e543715963bb2a4469a0e340e29e56e22821fbc8e0663f18.
//
// Solidity: event RequestUpdated(uint256 requestId, uint256 newRequestId, address executor, bytes acceptData, bytes rejectData, string description)
func (_ReviewableRequest *ReviewableRequestFilterer) WatchRequestUpdated(opts *bind.WatchOpts, sink chan<- *ReviewableRequestRequestUpdated) (event.Subscription, error) {

	logs, sub, err := _ReviewableRequest.contract.WatchLogs(opts, "RequestUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReviewableRequestRequestUpdated)
				if err := _ReviewableRequest.contract.UnpackLog(event, "RequestUpdated", log); err != nil {
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
func (_ReviewableRequest *ReviewableRequestFilterer) ParseRequestUpdated(log types.Log) (*ReviewableRequestRequestUpdated, error) {
	event := new(ReviewableRequestRequestUpdated)
	if err := _ReviewableRequest.contract.UnpackLog(event, "RequestUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
