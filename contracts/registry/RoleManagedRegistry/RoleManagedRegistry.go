// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RoleManagedRegistry

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isProxy\",\"type\":\"bool\"}],\"name\":\"AddedContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"RemovedContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MASTER_ACCESS_MANAGEMENT_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"addContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"addProxyContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyUpgrader\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"hasContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"injectDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"contractAddress_\",\"type\":\"address\"}],\"name\":\"justAddProxyContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"removeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"}],\"name\":\"upgradeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"upgradeContractAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

// MASTERACCESSMANAGEMENTNAME is a free data retrieval call binding the contract method 0xa4ce3439.
//
// Solidity: function MASTER_ACCESS_MANAGEMENT_NAME() view returns(string)
func (_Main *MainCaller) MASTERACCESSMANAGEMENTNAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "MASTER_ACCESS_MANAGEMENT_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MASTERACCESSMANAGEMENTNAME is a free data retrieval call binding the contract method 0xa4ce3439.
//
// Solidity: function MASTER_ACCESS_MANAGEMENT_NAME() view returns(string)
func (_Main *MainSession) MASTERACCESSMANAGEMENTNAME() (string, error) {
	return _Main.Contract.MASTERACCESSMANAGEMENTNAME(&_Main.CallOpts)
}

// MASTERACCESSMANAGEMENTNAME is a free data retrieval call binding the contract method 0xa4ce3439.
//
// Solidity: function MASTER_ACCESS_MANAGEMENT_NAME() view returns(string)
func (_Main *MainCallerSession) MASTERACCESSMANAGEMENTNAME() (string, error) {
	return _Main.Contract.MASTERACCESSMANAGEMENTNAME(&_Main.CallOpts)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string name) view returns(address)
func (_Main *MainCaller) GetContract(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getContract", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string name) view returns(address)
func (_Main *MainSession) GetContract(name string) (common.Address, error) {
	return _Main.Contract.GetContract(&_Main.CallOpts, name)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string name) view returns(address)
func (_Main *MainCallerSession) GetContract(name string) (common.Address, error) {
	return _Main.Contract.GetContract(&_Main.CallOpts, name)
}

// GetImplementation is a free data retrieval call binding the contract method 0x6b683896.
//
// Solidity: function getImplementation(string name) view returns(address)
func (_Main *MainCaller) GetImplementation(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getImplementation", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0x6b683896.
//
// Solidity: function getImplementation(string name) view returns(address)
func (_Main *MainSession) GetImplementation(name string) (common.Address, error) {
	return _Main.Contract.GetImplementation(&_Main.CallOpts, name)
}

// GetImplementation is a free data retrieval call binding the contract method 0x6b683896.
//
// Solidity: function getImplementation(string name) view returns(address)
func (_Main *MainCallerSession) GetImplementation(name string) (common.Address, error) {
	return _Main.Contract.GetImplementation(&_Main.CallOpts, name)
}

// GetProxyUpgrader is a free data retrieval call binding the contract method 0xd10611fc.
//
// Solidity: function getProxyUpgrader() view returns(address)
func (_Main *MainCaller) GetProxyUpgrader(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getProxyUpgrader")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyUpgrader is a free data retrieval call binding the contract method 0xd10611fc.
//
// Solidity: function getProxyUpgrader() view returns(address)
func (_Main *MainSession) GetProxyUpgrader() (common.Address, error) {
	return _Main.Contract.GetProxyUpgrader(&_Main.CallOpts)
}

// GetProxyUpgrader is a free data retrieval call binding the contract method 0xd10611fc.
//
// Solidity: function getProxyUpgrader() view returns(address)
func (_Main *MainCallerSession) GetProxyUpgrader() (common.Address, error) {
	return _Main.Contract.GetProxyUpgrader(&_Main.CallOpts)
}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(string name) view returns(bool)
func (_Main *MainCaller) HasContract(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "hasContract", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(string name) view returns(bool)
func (_Main *MainSession) HasContract(name string) (bool, error) {
	return _Main.Contract.HasContract(&_Main.CallOpts, name)
}

// HasContract is a free data retrieval call binding the contract method 0x8c223601.
//
// Solidity: function hasContract(string name) view returns(bool)
func (_Main *MainCallerSession) HasContract(name string) (bool, error) {
	return _Main.Contract.HasContract(&_Main.CallOpts, name)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Main *MainCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Main *MainSession) ProxiableUUID() ([32]byte, error) {
	return _Main.Contract.ProxiableUUID(&_Main.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Main *MainCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Main.Contract.ProxiableUUID(&_Main.CallOpts)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string name_, address contractAddress_) returns()
func (_Main *MainTransactor) AddContract(opts *bind.TransactOpts, name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addContract", name_, contractAddress_)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string name_, address contractAddress_) returns()
func (_Main *MainSession) AddContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddContract(&_Main.TransactOpts, name_, contractAddress_)
}

// AddContract is a paid mutator transaction binding the contract method 0xbf5b6016.
//
// Solidity: function addContract(string name_, address contractAddress_) returns()
func (_Main *MainTransactorSession) AddContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddContract(&_Main.TransactOpts, name_, contractAddress_)
}

// AddProxyContract is a paid mutator transaction binding the contract method 0xe0e084f8.
//
// Solidity: function addProxyContract(string name_, address contractAddress_) returns()
func (_Main *MainTransactor) AddProxyContract(opts *bind.TransactOpts, name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addProxyContract", name_, contractAddress_)
}

// AddProxyContract is a paid mutator transaction binding the contract method 0xe0e084f8.
//
// Solidity: function addProxyContract(string name_, address contractAddress_) returns()
func (_Main *MainSession) AddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddProxyContract(&_Main.TransactOpts, name_, contractAddress_)
}

// AddProxyContract is a paid mutator transaction binding the contract method 0xe0e084f8.
//
// Solidity: function addProxyContract(string name_, address contractAddress_) returns()
func (_Main *MainTransactorSession) AddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Main.Contract.AddProxyContract(&_Main.TransactOpts, name_, contractAddress_)
}

// InjectDependencies is a paid mutator transaction binding the contract method 0x1adad8cf.
//
// Solidity: function injectDependencies(string name_) returns()
func (_Main *MainTransactor) InjectDependencies(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "injectDependencies", name_)
}

// InjectDependencies is a paid mutator transaction binding the contract method 0x1adad8cf.
//
// Solidity: function injectDependencies(string name_) returns()
func (_Main *MainSession) InjectDependencies(name_ string) (*types.Transaction, error) {
	return _Main.Contract.InjectDependencies(&_Main.TransactOpts, name_)
}

// InjectDependencies is a paid mutator transaction binding the contract method 0x1adad8cf.
//
// Solidity: function injectDependencies(string name_) returns()
func (_Main *MainTransactorSession) InjectDependencies(name_ string) (*types.Transaction, error) {
	return _Main.Contract.InjectDependencies(&_Main.TransactOpts, name_)
}

// JustAddProxyContract is a paid mutator transaction binding the contract method 0x51dad82c.
//
// Solidity: function justAddProxyContract(string name_, address contractAddress_) returns()
func (_Main *MainTransactor) JustAddProxyContract(opts *bind.TransactOpts, name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "justAddProxyContract", name_, contractAddress_)
}

// JustAddProxyContract is a paid mutator transaction binding the contract method 0x51dad82c.
//
// Solidity: function justAddProxyContract(string name_, address contractAddress_) returns()
func (_Main *MainSession) JustAddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Main.Contract.JustAddProxyContract(&_Main.TransactOpts, name_, contractAddress_)
}

// JustAddProxyContract is a paid mutator transaction binding the contract method 0x51dad82c.
//
// Solidity: function justAddProxyContract(string name_, address contractAddress_) returns()
func (_Main *MainTransactorSession) JustAddProxyContract(name_ string, contractAddress_ common.Address) (*types.Transaction, error) {
	return _Main.Contract.JustAddProxyContract(&_Main.TransactOpts, name_, contractAddress_)
}

// RemoveContract is a paid mutator transaction binding the contract method 0x97623b58.
//
// Solidity: function removeContract(string name_) returns()
func (_Main *MainTransactor) RemoveContract(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeContract", name_)
}

// RemoveContract is a paid mutator transaction binding the contract method 0x97623b58.
//
// Solidity: function removeContract(string name_) returns()
func (_Main *MainSession) RemoveContract(name_ string) (*types.Transaction, error) {
	return _Main.Contract.RemoveContract(&_Main.TransactOpts, name_)
}

// RemoveContract is a paid mutator transaction binding the contract method 0x97623b58.
//
// Solidity: function removeContract(string name_) returns()
func (_Main *MainTransactorSession) RemoveContract(name_ string) (*types.Transaction, error) {
	return _Main.Contract.RemoveContract(&_Main.TransactOpts, name_)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x1271bd53.
//
// Solidity: function upgradeContract(string name_, address newImplementation_) returns()
func (_Main *MainTransactor) UpgradeContract(opts *bind.TransactOpts, name_ string, newImplementation_ common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "upgradeContract", name_, newImplementation_)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x1271bd53.
//
// Solidity: function upgradeContract(string name_, address newImplementation_) returns()
func (_Main *MainSession) UpgradeContract(name_ string, newImplementation_ common.Address) (*types.Transaction, error) {
	return _Main.Contract.UpgradeContract(&_Main.TransactOpts, name_, newImplementation_)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0x1271bd53.
//
// Solidity: function upgradeContract(string name_, address newImplementation_) returns()
func (_Main *MainTransactorSession) UpgradeContract(name_ string, newImplementation_ common.Address) (*types.Transaction, error) {
	return _Main.Contract.UpgradeContract(&_Main.TransactOpts, name_, newImplementation_)
}

// UpgradeContractAndCall is a paid mutator transaction binding the contract method 0x6bbe8694.
//
// Solidity: function upgradeContractAndCall(string name_, address newImplementation_, bytes data_) returns()
func (_Main *MainTransactor) UpgradeContractAndCall(opts *bind.TransactOpts, name_ string, newImplementation_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "upgradeContractAndCall", name_, newImplementation_, data_)
}

// UpgradeContractAndCall is a paid mutator transaction binding the contract method 0x6bbe8694.
//
// Solidity: function upgradeContractAndCall(string name_, address newImplementation_, bytes data_) returns()
func (_Main *MainSession) UpgradeContractAndCall(name_ string, newImplementation_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _Main.Contract.UpgradeContractAndCall(&_Main.TransactOpts, name_, newImplementation_, data_)
}

// UpgradeContractAndCall is a paid mutator transaction binding the contract method 0x6bbe8694.
//
// Solidity: function upgradeContractAndCall(string name_, address newImplementation_, bytes data_) returns()
func (_Main *MainTransactorSession) UpgradeContractAndCall(name_ string, newImplementation_ common.Address, data_ []byte) (*types.Transaction, error) {
	return _Main.Contract.UpgradeContractAndCall(&_Main.TransactOpts, name_, newImplementation_, data_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Main *MainTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Main *MainSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Main.Contract.UpgradeTo(&_Main.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Main *MainTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Main.Contract.UpgradeTo(&_Main.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Main *MainTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Main *MainSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Main.Contract.UpgradeToAndCall(&_Main.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Main *MainTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Main.Contract.UpgradeToAndCall(&_Main.TransactOpts, newImplementation, data)
}

// MainAddedContractIterator is returned from FilterAddedContract and is used to iterate over the raw logs and unpacked data for AddedContract events raised by the Main contract.
type MainAddedContractIterator struct {
	Event *MainAddedContract // Event containing the contract specifics and raw log

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
func (it *MainAddedContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainAddedContract)
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
		it.Event = new(MainAddedContract)
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
func (it *MainAddedContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainAddedContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainAddedContract represents a AddedContract event raised by the Main contract.
type MainAddedContract struct {
	Name            string
	ContractAddress common.Address
	IsProxy         bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddedContract is a free log retrieval operation binding the contract event 0xd8607592920f7ac4d1e9d198396c7f7cdaee8e675bc3ffb12924a71979867625.
//
// Solidity: event AddedContract(string name, address contractAddress, bool isProxy)
func (_Main *MainFilterer) FilterAddedContract(opts *bind.FilterOpts) (*MainAddedContractIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "AddedContract")
	if err != nil {
		return nil, err
	}
	return &MainAddedContractIterator{contract: _Main.contract, event: "AddedContract", logs: logs, sub: sub}, nil
}

// WatchAddedContract is a free log subscription operation binding the contract event 0xd8607592920f7ac4d1e9d198396c7f7cdaee8e675bc3ffb12924a71979867625.
//
// Solidity: event AddedContract(string name, address contractAddress, bool isProxy)
func (_Main *MainFilterer) WatchAddedContract(opts *bind.WatchOpts, sink chan<- *MainAddedContract) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "AddedContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainAddedContract)
				if err := _Main.contract.UnpackLog(event, "AddedContract", log); err != nil {
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
func (_Main *MainFilterer) ParseAddedContract(log types.Log) (*MainAddedContract, error) {
	event := new(MainAddedContract)
	if err := _Main.contract.UnpackLog(event, "AddedContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Main contract.
type MainAdminChangedIterator struct {
	Event *MainAdminChanged // Event containing the contract specifics and raw log

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
func (it *MainAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainAdminChanged)
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
		it.Event = new(MainAdminChanged)
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
func (it *MainAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainAdminChanged represents a AdminChanged event raised by the Main contract.
type MainAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Main *MainFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*MainAdminChangedIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &MainAdminChangedIterator{contract: _Main.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Main *MainFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *MainAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainAdminChanged)
				if err := _Main.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Main *MainFilterer) ParseAdminChanged(log types.Log) (*MainAdminChanged, error) {
	event := new(MainAdminChanged)
	if err := _Main.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Main contract.
type MainBeaconUpgradedIterator struct {
	Event *MainBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *MainBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainBeaconUpgraded)
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
		it.Event = new(MainBeaconUpgraded)
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
func (it *MainBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainBeaconUpgraded represents a BeaconUpgraded event raised by the Main contract.
type MainBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Main *MainFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*MainBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &MainBeaconUpgradedIterator{contract: _Main.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Main *MainFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *MainBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainBeaconUpgraded)
				if err := _Main.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Main *MainFilterer) ParseBeaconUpgraded(log types.Log) (*MainBeaconUpgraded, error) {
	event := new(MainBeaconUpgraded)
	if err := _Main.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRemovedContractIterator is returned from FilterRemovedContract and is used to iterate over the raw logs and unpacked data for RemovedContract events raised by the Main contract.
type MainRemovedContractIterator struct {
	Event *MainRemovedContract // Event containing the contract specifics and raw log

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
func (it *MainRemovedContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRemovedContract)
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
		it.Event = new(MainRemovedContract)
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
func (it *MainRemovedContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRemovedContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRemovedContract represents a RemovedContract event raised by the Main contract.
type MainRemovedContract struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedContract is a free log retrieval operation binding the contract event 0x2216f135d7c337bc5b25f8bf96227730d40241c50adf6fde4d180deca4b25664.
//
// Solidity: event RemovedContract(string name)
func (_Main *MainFilterer) FilterRemovedContract(opts *bind.FilterOpts) (*MainRemovedContractIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RemovedContract")
	if err != nil {
		return nil, err
	}
	return &MainRemovedContractIterator{contract: _Main.contract, event: "RemovedContract", logs: logs, sub: sub}, nil
}

// WatchRemovedContract is a free log subscription operation binding the contract event 0x2216f135d7c337bc5b25f8bf96227730d40241c50adf6fde4d180deca4b25664.
//
// Solidity: event RemovedContract(string name)
func (_Main *MainFilterer) WatchRemovedContract(opts *bind.WatchOpts, sink chan<- *MainRemovedContract) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RemovedContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRemovedContract)
				if err := _Main.contract.UnpackLog(event, "RemovedContract", log); err != nil {
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
func (_Main *MainFilterer) ParseRemovedContract(log types.Log) (*MainRemovedContract, error) {
	event := new(MainRemovedContract)
	if err := _Main.contract.UnpackLog(event, "RemovedContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Main contract.
type MainUpgradedIterator struct {
	Event *MainUpgraded // Event containing the contract specifics and raw log

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
func (it *MainUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainUpgraded)
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
		it.Event = new(MainUpgraded)
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
func (it *MainUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainUpgraded represents a Upgraded event raised by the Main contract.
type MainUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Main *MainFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MainUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MainUpgradedIterator{contract: _Main.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Main *MainFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MainUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainUpgraded)
				if err := _Main.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Main *MainFilterer) ParseUpgraded(log types.Log) (*MainUpgraded, error) {
	event := new(MainUpgraded)
	if err := _Main.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
