// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package role_managed_registry

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

// RoleManagedRegistryMetaData contains all meta data concerning the RoleManagedRegistry contract.
var RoleManagedRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isProxy\",\"type\":\"bool\"}],\"name\":\"AddedContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"RemovedContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MASTER_ACCESS_MANAGEMENT_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"addContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"addProxyContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyUpgrader\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"hasContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"injectDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"justAddProxyContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"removeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"}],\"name\":\"upgradeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"upgradeContractAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RoleManagedRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RoleManagedRegistryMetaData.ABI instead.
var RoleManagedRegistryABI = RoleManagedRegistryMetaData.ABI

// RoleManagedRegistry is an auto generated Go binding around an Ethereum contract.
type RoleManagedRegistry struct {
	RoleManagedRegistryCaller     // Read-only binding to the contract
	RoleManagedRegistryTransactor // Write-only binding to the contract
	RoleManagedRegistryFilterer   // Log filterer for contract events
}

// RoleManagedRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoleManagedRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleManagedRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoleManagedRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleManagedRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoleManagedRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleManagedRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoleManagedRegistrySession struct {
	Contract     *RoleManagedRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RoleManagedRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoleManagedRegistryCallerSession struct {
	Contract *RoleManagedRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RoleManagedRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoleManagedRegistryTransactorSession struct {
	Contract     *RoleManagedRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RoleManagedRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoleManagedRegistryRaw struct {
	Contract *RoleManagedRegistry // Generic contract binding to access the raw methods on
}

// RoleManagedRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoleManagedRegistryCallerRaw struct {
	Contract *RoleManagedRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RoleManagedRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoleManagedRegistryTransactorRaw struct {
	Contract *RoleManagedRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoleManagedRegistry creates a new instance of RoleManagedRegistry, bound to a specific deployed contract.
func NewRoleManagedRegistry(address common.Address, backend bind.ContractBackend) (*RoleManagedRegistry, error) {
	contract, err := bindRoleManagedRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoleManagedRegistry{RoleManagedRegistryCaller: RoleManagedRegistryCaller{contract: contract}, RoleManagedRegistryTransactor: RoleManagedRegistryTransactor{contract: contract}, RoleManagedRegistryFilterer: RoleManagedRegistryFilterer{contract: contract}}, nil
}

// NewRoleManagedRegistryCaller creates a new read-only instance of RoleManagedRegistry, bound to a specific deployed contract.
func NewRoleManagedRegistryCaller(address common.Address, caller bind.ContractCaller) (*RoleManagedRegistryCaller, error) {
	contract, err := bindRoleManagedRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoleManagedRegistryCaller{contract: contract}, nil
}

// NewRoleManagedRegistryTransactor creates a new write-only instance of RoleManagedRegistry, bound to a specific deployed contract.
func NewRoleManagedRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RoleManagedRegistryTransactor, error) {
	contract, err := bindRoleManagedRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoleManagedRegistryTransactor{contract: contract}, nil
}

// NewRoleManagedRegistryFilterer creates a new log filterer instance of RoleManagedRegistry, bound to a specific deployed contract.
func NewRoleManagedRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RoleManagedRegistryFilterer, error) {
	contract, err := bindRoleManagedRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoleManagedRegistryFilterer{contract: contract}, nil
}

// bindRoleManagedRegistry binds a generic wrapper to an already deployed contract.
func bindRoleManagedRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoleManagedRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleManagedRegistry *RoleManagedRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoleManagedRegistry.Contract.RoleManagedRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleManagedRegistry *RoleManagedRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.RoleManagedRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleManagedRegistry *RoleManagedRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.RoleManagedRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleManagedRegistry *RoleManagedRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoleManagedRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleManagedRegistry *RoleManagedRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleManagedRegistry *RoleManagedRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.contract.Transact(opts, method, params...)
}

// MASTERACCESSMANAGEMENTNAME is a free data retrieval call binding the contract method 0xa4ce3439.
//
// Solidity: function MASTER_ACCESS_MANAGEMENT_NAME() view returns(string)
func (_RoleManagedRegistry *RoleManagedRegistryCaller) MASTERACCESSMANAGEMENTNAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RoleManagedRegistry.contract.Call(opts, &out, "MASTER_ACCESS_MANAGEMENT_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MASTERACCESSMANAGEMENTNAME is a free data retrieval call binding the contract method 0xa4ce3439.
//
// Solidity: function MASTER_ACCESS_MANAGEMENT_NAME() view returns(string)
func (_RoleManagedRegistry *RoleManagedRegistrySession) MASTERACCESSMANAGEMENTNAME() (string, error) {
	return _RoleManagedRegistry.Contract.MASTERACCESSMANAGEMENTNAME(&_RoleManagedRegistry.CallOpts)
}

// MASTERACCESSMANAGEMENTNAME is a free data retrieval call binding the contract method 0xa4ce3439.
//
// Solidity: function MASTER_ACCESS_MANAGEMENT_NAME() view returns(string)
func (_RoleManagedRegistry *RoleManagedRegistryCallerSession) MASTERACCESSMANAGEMENTNAME() (string, error) {
	return _RoleManagedRegistry.Contract.MASTERACCESSMANAGEMENTNAME(&_RoleManagedRegistry.CallOpts)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string name) view returns(address)
func (_RoleManagedRegistry *RoleManagedRegistryCaller) GetContract(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _RoleManagedRegistry.contract.Call(opts, &out, "getContract", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string name) view returns(address)
func (_RoleManagedRegistry *RoleManagedRegistrySession) GetContract(name string) (common.Address, error) {
	return _RoleManagedRegistry.Contract.GetContract(&_RoleManagedRegistry.CallOpts, name)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string name) view returns(address)
func (_RoleManagedRegistry *RoleManagedRegistryCallerSession) GetContract(name string) (common.Address, error) {
	return _RoleManagedRegistry.Contract.GetContract(&_RoleManagedRegistry.CallOpts, name)
}

// GetImplementation is a free data retrieval call binding the contract method 0x6b683896.
//
// Solidity: function getImplementation(string name) view returns(address)
func (_RoleManagedRegistry *RoleManagedRegistryCaller) GetImplementation(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _RoleManagedRegistry.contract.Call(opts, &out, "getImplementation", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0x6b683896.
//
// Solidity: function getImplementation(string name) view returns(address)
func (_RoleManagedRegistry *RoleManagedRegistrySession) GetImplementation(name string) (common.Address, error) {
	return _RoleManagedRegistry.Contract.GetImplementation(&_RoleManagedRegistry.CallOpts, name)
}

// GetImplementation is a free data retrieval call binding the contract method 0x6b683896.
//
// Solidity: function getImplementation(string name) view returns(address)
func (_RoleManagedRegistry *RoleManagedRegistryCallerSession) GetImplementation(name string) (common.Address, error) {
	return _RoleManagedRegistry.Contract.GetImplementation(&_RoleManagedRegistry.CallOpts, name)
}

// GetProxyUpgrader is a free data retrieval call binding the contract method 0xd10611fc.
//
// Solidity: function getProxyUpgrader() view returns(address)
func (_RoleManagedRegistry *RoleManagedRegistryCaller) GetProxyUpgrader(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoleManagedRegistry.contract.Call(opts, &out, "getProxyUpgrader")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyUpgrader is a free data retrieval call binding the contract method 0xd10611fc.
//
// Solidity: function getProxyUpgrader() view returns(address)
func (_RoleManagedRegistry *RoleManagedRegistrySession) GetProxyUpgrader() (common.Address, error) {
	return _RoleManagedRegistry.Contract.GetProxyUpgrader(&_RoleManagedRegistry.CallOpts)
}

// GetProxyUpgrader is a free data retrieval call binding the contract method 0xd10611fc.
//
// Solidity: function getProxyUpgrader() view returns(address)
func (_RoleManagedRegistry *RoleManagedRegistryCallerSession) GetProxyUpgrader() (common.Address, error) {
	return _RoleManagedRegistry.Contract.GetProxyUpgrader(&_RoleManagedRegistry.CallOpts)
}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(string name) view returns(bool)
func (_RoleManagedRegistry *RoleManagedRegistryCaller) HasContract(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _RoleManagedRegistry.contract.Call(opts, &out, "hasContract", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(string name) view returns(bool)
func (_RoleManagedRegistry *RoleManagedRegistrySession) HasContract(name string) (bool, error) {
	return _RoleManagedRegistry.Contract.HasContract(&_RoleManagedRegistry.CallOpts, name)
}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(string name) view returns(bool)
func (_RoleManagedRegistry *RoleManagedRegistryCallerSession) HasContract(name string) (bool, error) {
	return _RoleManagedRegistry.Contract.HasContract(&_RoleManagedRegistry.CallOpts, name)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RoleManagedRegistry *RoleManagedRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RoleManagedRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RoleManagedRegistry *RoleManagedRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _RoleManagedRegistry.Contract.ProxiableUUID(&_RoleManagedRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RoleManagedRegistry *RoleManagedRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _RoleManagedRegistry.Contract.ProxiableUUID(&_RoleManagedRegistry.CallOpts)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string name_, address contractAddress_) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactor) AddContract(opts *bind.TransactOpts, name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.contract.Transact(opts, "addContract", name_, contractAddress_)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string name_, address contractAddress_) returns()
func (_RoleManagedRegistry *RoleManagedRegistrySession) AddContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.AddContract(&_RoleManagedRegistry.TransactOpts, name_, contractAddress_)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string name_, address contractAddress_) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactorSession) AddContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.AddContract(&_RoleManagedRegistry.TransactOpts, name_, contractAddress_)
}

// AddProxyContract is a paid mutator transaction binding the contract method 0xe0e084f8.
//
// Solidity: function addProxyContract(string name_, address contractAddress_) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactor) AddProxyContract(opts *bind.TransactOpts, name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.contract.Transact(opts, "addProxyContract", name_, contractAddress_)
}

// AddProxyContract is a paid mutator transaction binding the contract method 0xe0e084f8.
//
// Solidity: function addProxyContract(string name_, address contractAddress_) returns()
func (_RoleManagedRegistry *RoleManagedRegistrySession) AddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.AddProxyContract(&_RoleManagedRegistry.TransactOpts, name_, contractAddress_)
}

// AddProxyContract is a paid mutator transaction binding the contract method 0xe0e084f8.
//
// Solidity: function addProxyContract(string name_, address contractAddress_) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactorSession) AddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.AddProxyContract(&_RoleManagedRegistry.TransactOpts, name_, contractAddress_)
}

// InjectDependencies is a paid mutator transaction binding the contract method 0x1adad8cf.
//
// Solidity: function injectDependencies(string name_) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactor) InjectDependencies(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _RoleManagedRegistry.contract.Transact(opts, "injectDependencies", name_)
}

// InjectDependencies is a paid mutator transaction binding the contract method 0x1adad8cf.
//
// Solidity: function injectDependencies(string name_) returns()
func (_RoleManagedRegistry *RoleManagedRegistrySession) InjectDependencies(name_ string) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.InjectDependencies(&_RoleManagedRegistry.TransactOpts, name_)
}

// InjectDependencies is a paid mutator transaction binding the contract method 0x1adad8cf.
//
// Solidity: function injectDependencies(string name_) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactorSession) InjectDependencies(name_ string) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.InjectDependencies(&_RoleManagedRegistry.TransactOpts, name_)
}

// JustAddProxyContract is a paid mutator transaction binding the contract method 0x51dad82c.
//
// Solidity: function justAddProxyContract(string name_, address contractAddress_) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactor) JustAddProxyContract(opts *bind.TransactOpts, name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.contract.Transact(opts, "justAddProxyContract", name_, contractAddress_)
}

// JustAddProxyContract is a paid mutator transaction binding the contract method 0x51dad82c.
//
// Solidity: function justAddProxyContract(string name_, address contractAddress_) returns()
func (_RoleManagedRegistry *RoleManagedRegistrySession) JustAddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.JustAddProxyContract(&_RoleManagedRegistry.TransactOpts, name_, contractAddress_)
}

// JustAddProxyContract is a paid mutator transaction binding the contract method 0x51dad82c.
//
// Solidity: function justAddProxyContract(string name_, address contractAddress_) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactorSession) JustAddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.JustAddProxyContract(&_RoleManagedRegistry.TransactOpts, name_, contractAddress_)
}

// RemoveContract is a paid mutator transaction binding the contract method 0x97623b58.
//
// Solidity: function removeContract(string name_) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactor) RemoveContract(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _RoleManagedRegistry.contract.Transact(opts, "removeContract", name_)
}

// RemoveContract is a paid mutator transaction binding the contract method 0x97623b58.
//
// Solidity: function removeContract(string name_) returns()
func (_RoleManagedRegistry *RoleManagedRegistrySession) RemoveContract(name_ string) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.RemoveContract(&_RoleManagedRegistry.TransactOpts, name_)
}

// RemoveContract is a paid mutator transaction binding the contract method 0x97623b58.
//
// Solidity: function removeContract(string name_) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactorSession) RemoveContract(name_ string) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.RemoveContract(&_RoleManagedRegistry.TransactOpts, name_)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x1271bd53.
//
// Solidity: function upgradeContract(string name_, address newImplementation_) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactor) UpgradeContract(opts *bind.TransactOpts, name_ string, newImplementation_ common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.contract.Transact(opts, "upgradeContract", name_, newImplementation_)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x1271bd53.
//
// Solidity: function upgradeContract(string name_, address newImplementation_) returns()
func (_RoleManagedRegistry *RoleManagedRegistrySession) UpgradeContract(name_ string, newImplementation_ common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.UpgradeContract(&_RoleManagedRegistry.TransactOpts, name_, newImplementation_)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x1271bd53.
//
// Solidity: function upgradeContract(string name_, address newImplementation_) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactorSession) UpgradeContract(name_ string, newImplementation_ common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.UpgradeContract(&_RoleManagedRegistry.TransactOpts, name_, newImplementation_)
}

// UpgradeContractAndCall is a paid mutator transaction binding the contract method 0x6bbe8694.
//
// Solidity: function upgradeContractAndCall(string name_, address newImplementation_, bytes data_) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactor) UpgradeContractAndCall(opts *bind.TransactOpts, name_ string, newImplementation_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _RoleManagedRegistry.contract.Transact(opts, "upgradeContractAndCall", name_, newImplementation_, data_)
}

// UpgradeContractAndCall is a paid mutator transaction binding the contract method 0x6bbe8694.
//
// Solidity: function upgradeContractAndCall(string name_, address newImplementation_, bytes data_) returns()
func (_RoleManagedRegistry *RoleManagedRegistrySession) UpgradeContractAndCall(name_ string, newImplementation_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.UpgradeContractAndCall(&_RoleManagedRegistry.TransactOpts, name_, newImplementation_, data_)
}

// UpgradeContractAndCall is a paid mutator transaction binding the contract method 0x6bbe8694.
//
// Solidity: function upgradeContractAndCall(string name_, address newImplementation_, bytes data_) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactorSession) UpgradeContractAndCall(name_ string, newImplementation_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.UpgradeContractAndCall(&_RoleManagedRegistry.TransactOpts, name_, newImplementation_, data_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RoleManagedRegistry *RoleManagedRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.UpgradeTo(&_RoleManagedRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.UpgradeTo(&_RoleManagedRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RoleManagedRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RoleManagedRegistry *RoleManagedRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.UpgradeToAndCall(&_RoleManagedRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RoleManagedRegistry *RoleManagedRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RoleManagedRegistry.Contract.UpgradeToAndCall(&_RoleManagedRegistry.TransactOpts, newImplementation, data)
}

// RoleManagedRegistryAddedContractIterator is returned from FilterAddedContract and is used to iterate over the raw logs and unpacked data for AddedContract events raised by the RoleManagedRegistry contract.
type RoleManagedRegistryAddedContractIterator struct {
	Event *RoleManagedRegistryAddedContract // Event containing the contract specifics and raw log

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
func (it *RoleManagedRegistryAddedContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleManagedRegistryAddedContract)
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
		it.Event = new(RoleManagedRegistryAddedContract)
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
func (it *RoleManagedRegistryAddedContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleManagedRegistryAddedContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleManagedRegistryAddedContract represents a AddedContract event raised by the RoleManagedRegistry contract.
type RoleManagedRegistryAddedContract struct {
	Name            string
	ContractAddress common.Address
	IsProxy         bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddedContract is a free log retrieval operation binding the contract event 0xd8607592920f7ac4d1e9d198396c7f7cdaee8e675bc3ffb12924a71979867625.
//
// Solidity: event AddedContract(string name, address contractAddress, bool isProxy)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) FilterAddedContract(opts *bind.FilterOpts) (*RoleManagedRegistryAddedContractIterator, error) {

	logs, sub, err := _RoleManagedRegistry.contract.FilterLogs(opts, "AddedContract")
	if err != nil {
		return nil, err
	}
	return &RoleManagedRegistryAddedContractIterator{contract: _RoleManagedRegistry.contract, event: "AddedContract", logs: logs, sub: sub}, nil
}

// WatchAddedContract is a free log subscription operation binding the contract event 0xd8607592920f7ac4d1e9d198396c7f7cdaee8e675bc3ffb12924a71979867625.
//
// Solidity: event AddedContract(string name, address contractAddress, bool isProxy)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) WatchAddedContract(opts *bind.WatchOpts, sink chan<- *RoleManagedRegistryAddedContract) (event.Subscription, error) {

	logs, sub, err := _RoleManagedRegistry.contract.WatchLogs(opts, "AddedContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleManagedRegistryAddedContract)
				if err := _RoleManagedRegistry.contract.UnpackLog(event, "AddedContract", log); err != nil {
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

// ParseAddedContract is a log parse operation binding the contract event 0xd8607592920f7ac4d1e9d198396c7f7cdaee8e675bc3ffb12924a71979867625.
//
// Solidity: event AddedContract(string name, address contractAddress, bool isProxy)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) ParseAddedContract(log types.Log) (*RoleManagedRegistryAddedContract, error) {
	event := new(RoleManagedRegistryAddedContract)
	if err := _RoleManagedRegistry.contract.UnpackLog(event, "AddedContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleManagedRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the RoleManagedRegistry contract.
type RoleManagedRegistryAdminChangedIterator struct {
	Event *RoleManagedRegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *RoleManagedRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleManagedRegistryAdminChanged)
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
		it.Event = new(RoleManagedRegistryAdminChanged)
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
func (it *RoleManagedRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleManagedRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleManagedRegistryAdminChanged represents a AdminChanged event raised by the RoleManagedRegistry contract.
type RoleManagedRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*RoleManagedRegistryAdminChangedIterator, error) {

	logs, sub, err := _RoleManagedRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &RoleManagedRegistryAdminChangedIterator{contract: _RoleManagedRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *RoleManagedRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _RoleManagedRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleManagedRegistryAdminChanged)
				if err := _RoleManagedRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) ParseAdminChanged(log types.Log) (*RoleManagedRegistryAdminChanged, error) {
	event := new(RoleManagedRegistryAdminChanged)
	if err := _RoleManagedRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleManagedRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the RoleManagedRegistry contract.
type RoleManagedRegistryBeaconUpgradedIterator struct {
	Event *RoleManagedRegistryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *RoleManagedRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleManagedRegistryBeaconUpgraded)
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
		it.Event = new(RoleManagedRegistryBeaconUpgraded)
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
func (it *RoleManagedRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleManagedRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleManagedRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the RoleManagedRegistry contract.
type RoleManagedRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*RoleManagedRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _RoleManagedRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &RoleManagedRegistryBeaconUpgradedIterator{contract: _RoleManagedRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *RoleManagedRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _RoleManagedRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleManagedRegistryBeaconUpgraded)
				if err := _RoleManagedRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*RoleManagedRegistryBeaconUpgraded, error) {
	event := new(RoleManagedRegistryBeaconUpgraded)
	if err := _RoleManagedRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleManagedRegistryRemovedContractIterator is returned from FilterRemovedContract and is used to iterate over the raw logs and unpacked data for RemovedContract events raised by the RoleManagedRegistry contract.
type RoleManagedRegistryRemovedContractIterator struct {
	Event *RoleManagedRegistryRemovedContract // Event containing the contract specifics and raw log

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
func (it *RoleManagedRegistryRemovedContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleManagedRegistryRemovedContract)
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
		it.Event = new(RoleManagedRegistryRemovedContract)
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
func (it *RoleManagedRegistryRemovedContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleManagedRegistryRemovedContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleManagedRegistryRemovedContract represents a RemovedContract event raised by the RoleManagedRegistry contract.
type RoleManagedRegistryRemovedContract struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedContract is a free log retrieval operation binding the contract event 0x2216f135d7c337bc5b25f8bf96227730d40241c50adf6fde4d180deca4b25664.
//
// Solidity: event RemovedContract(string name)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) FilterRemovedContract(opts *bind.FilterOpts) (*RoleManagedRegistryRemovedContractIterator, error) {

	logs, sub, err := _RoleManagedRegistry.contract.FilterLogs(opts, "RemovedContract")
	if err != nil {
		return nil, err
	}
	return &RoleManagedRegistryRemovedContractIterator{contract: _RoleManagedRegistry.contract, event: "RemovedContract", logs: logs, sub: sub}, nil
}

// WatchRemovedContract is a free log subscription operation binding the contract event 0x2216f135d7c337bc5b25f8bf96227730d40241c50adf6fde4d180deca4b25664.
//
// Solidity: event RemovedContract(string name)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) WatchRemovedContract(opts *bind.WatchOpts, sink chan<- *RoleManagedRegistryRemovedContract) (event.Subscription, error) {

	logs, sub, err := _RoleManagedRegistry.contract.WatchLogs(opts, "RemovedContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleManagedRegistryRemovedContract)
				if err := _RoleManagedRegistry.contract.UnpackLog(event, "RemovedContract", log); err != nil {
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

// ParseRemovedContract is a log parse operation binding the contract event 0x2216f135d7c337bc5b25f8bf96227730d40241c50adf6fde4d180deca4b25664.
//
// Solidity: event RemovedContract(string name)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) ParseRemovedContract(log types.Log) (*RoleManagedRegistryRemovedContract, error) {
	event := new(RoleManagedRegistryRemovedContract)
	if err := _RoleManagedRegistry.contract.UnpackLog(event, "RemovedContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleManagedRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the RoleManagedRegistry contract.
type RoleManagedRegistryUpgradedIterator struct {
	Event *RoleManagedRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *RoleManagedRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleManagedRegistryUpgraded)
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
		it.Event = new(RoleManagedRegistryUpgraded)
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
func (it *RoleManagedRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleManagedRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleManagedRegistryUpgraded represents a Upgraded event raised by the RoleManagedRegistry contract.
type RoleManagedRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RoleManagedRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RoleManagedRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RoleManagedRegistryUpgradedIterator{contract: _RoleManagedRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RoleManagedRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RoleManagedRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleManagedRegistryUpgraded)
				if err := _RoleManagedRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RoleManagedRegistry *RoleManagedRegistryFilterer) ParseUpgraded(log types.Log) (*RoleManagedRegistryUpgraded, error) {
	event := new(RoleManagedRegistryUpgraded)
	if err := _RoleManagedRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
