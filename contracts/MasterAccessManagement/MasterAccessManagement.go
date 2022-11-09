// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MasterAccessManagement

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

// IRBACResourceWithPermissions is an auto generated low-level Go binding around an user-defined struct.
type IRBACResourceWithPermissions struct {
	Resource    string
	Permissions []string
}

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"permissionsToAdd\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AddedPermissions\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"rolesToGrant\",\"type\":\"string[]\"}],\"name\":\"GrantedRoles\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"permissionsToRemove\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"RemovedPermissions\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"rolesToRevoke\",\"type\":\"string[]\"}],\"name\":\"RevokedRoles\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ALL_PERMISSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ALL_RESOURCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSTANTS_REGISTRY_RESOURCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CREATE_PERMISSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELETE_PERMISSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXECUTE_PERMISSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_REGISTRY_RESOURCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_ROLE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RBAC_RESOURCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"READ_PERMISSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REVIEWABLE_REQUESTS_RESOURCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_PERMISSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"master_\",\"type\":\"address\"}],\"name\":\"__MasterAccessManagement_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"permissions\",\"type\":\"string[]\"}],\"internalType\":\"structIRBAC.ResourceWithPermissions[]\",\"name\":\"permissionsToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"addPermissionsToRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"}],\"name\":\"getRolePermissions\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"permissions\",\"type\":\"string[]\"}],\"internalType\":\"structIRBAC.ResourceWithPermissions[]\",\"name\":\"allowed\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"permissions\",\"type\":\"string[]\"}],\"internalType\":\"structIRBAC.ResourceWithPermissions[]\",\"name\":\"disallowed\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"getUserRoles\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"roles\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"rolesToGrant\",\"type\":\"string[]\"}],\"name\":\"grantRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasConstantsRegistryCreatePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasConstantsRegistryDeletePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasMasterContractsRegistryCreatePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasMasterContractsRegistryDeletePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasMasterContractsRegistryUpdatePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"permission\",\"type\":\"string\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasReviewableRequestsCreatePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasReviewableRequestsDeletePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasReviewableRequestsExecutePermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"resource\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"permissions\",\"type\":\"string[]\"}],\"internalType\":\"structIRBAC.ResourceWithPermissions[]\",\"name\":\"permissionsToRemove\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"removePermissionsFromRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"rolesToRevoke\",\"type\":\"string[]\"}],\"name\":\"revokeRoles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// ALLPERMISSION is a free data retrieval call binding the contract method 0xff846fb9.
//
// Solidity: function ALL_PERMISSION() view returns(string)
func (_Main *MainCaller) ALLPERMISSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "ALL_PERMISSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ALLPERMISSION is a free data retrieval call binding the contract method 0xff846fb9.
//
// Solidity: function ALL_PERMISSION() view returns(string)
func (_Main *MainSession) ALLPERMISSION() (string, error) {
	return _Main.Contract.ALLPERMISSION(&_Main.CallOpts)
}

// ALLPERMISSION is a free data retrieval call binding the contract method 0xff846fb9.
//
// Solidity: function ALL_PERMISSION() view returns(string)
func (_Main *MainCallerSession) ALLPERMISSION() (string, error) {
	return _Main.Contract.ALLPERMISSION(&_Main.CallOpts)
}

// ALLRESOURCE is a free data retrieval call binding the contract method 0x4a9e12c5.
//
// Solidity: function ALL_RESOURCE() view returns(string)
func (_Main *MainCaller) ALLRESOURCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "ALL_RESOURCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ALLRESOURCE is a free data retrieval call binding the contract method 0x4a9e12c5.
//
// Solidity: function ALL_RESOURCE() view returns(string)
func (_Main *MainSession) ALLRESOURCE() (string, error) {
	return _Main.Contract.ALLRESOURCE(&_Main.CallOpts)
}

// ALLRESOURCE is a free data retrieval call binding the contract method 0x4a9e12c5.
//
// Solidity: function ALL_RESOURCE() view returns(string)
func (_Main *MainCallerSession) ALLRESOURCE() (string, error) {
	return _Main.Contract.ALLRESOURCE(&_Main.CallOpts)
}

// CONSTANTSREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x6a5efd26.
//
// Solidity: function CONSTANTS_REGISTRY_RESOURCE() view returns(string)
func (_Main *MainCaller) CONSTANTSREGISTRYRESOURCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "CONSTANTS_REGISTRY_RESOURCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CONSTANTSREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x6a5efd26.
//
// Solidity: function CONSTANTS_REGISTRY_RESOURCE() view returns(string)
func (_Main *MainSession) CONSTANTSREGISTRYRESOURCE() (string, error) {
	return _Main.Contract.CONSTANTSREGISTRYRESOURCE(&_Main.CallOpts)
}

// CONSTANTSREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x6a5efd26.
//
// Solidity: function CONSTANTS_REGISTRY_RESOURCE() view returns(string)
func (_Main *MainCallerSession) CONSTANTSREGISTRYRESOURCE() (string, error) {
	return _Main.Contract.CONSTANTSREGISTRYRESOURCE(&_Main.CallOpts)
}

// CREATEPERMISSION is a free data retrieval call binding the contract method 0xb3e657fb.
//
// Solidity: function CREATE_PERMISSION() view returns(string)
func (_Main *MainCaller) CREATEPERMISSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "CREATE_PERMISSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CREATEPERMISSION is a free data retrieval call binding the contract method 0xb3e657fb.
//
// Solidity: function CREATE_PERMISSION() view returns(string)
func (_Main *MainSession) CREATEPERMISSION() (string, error) {
	return _Main.Contract.CREATEPERMISSION(&_Main.CallOpts)
}

// CREATEPERMISSION is a free data retrieval call binding the contract method 0xb3e657fb.
//
// Solidity: function CREATE_PERMISSION() view returns(string)
func (_Main *MainCallerSession) CREATEPERMISSION() (string, error) {
	return _Main.Contract.CREATEPERMISSION(&_Main.CallOpts)
}

// DELETEPERMISSION is a free data retrieval call binding the contract method 0xb832a5a2.
//
// Solidity: function DELETE_PERMISSION() view returns(string)
func (_Main *MainCaller) DELETEPERMISSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "DELETE_PERMISSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DELETEPERMISSION is a free data retrieval call binding the contract method 0xb832a5a2.
//
// Solidity: function DELETE_PERMISSION() view returns(string)
func (_Main *MainSession) DELETEPERMISSION() (string, error) {
	return _Main.Contract.DELETEPERMISSION(&_Main.CallOpts)
}

// DELETEPERMISSION is a free data retrieval call binding the contract method 0xb832a5a2.
//
// Solidity: function DELETE_PERMISSION() view returns(string)
func (_Main *MainCallerSession) DELETEPERMISSION() (string, error) {
	return _Main.Contract.DELETEPERMISSION(&_Main.CallOpts)
}

// EXECUTEPERMISSION is a free data retrieval call binding the contract method 0xbcfa4784.
//
// Solidity: function EXECUTE_PERMISSION() view returns(string)
func (_Main *MainCaller) EXECUTEPERMISSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "EXECUTE_PERMISSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EXECUTEPERMISSION is a free data retrieval call binding the contract method 0xbcfa4784.
//
// Solidity: function EXECUTE_PERMISSION() view returns(string)
func (_Main *MainSession) EXECUTEPERMISSION() (string, error) {
	return _Main.Contract.EXECUTEPERMISSION(&_Main.CallOpts)
}

// EXECUTEPERMISSION is a free data retrieval call binding the contract method 0xbcfa4784.
//
// Solidity: function EXECUTE_PERMISSION() view returns(string)
func (_Main *MainCallerSession) EXECUTEPERMISSION() (string, error) {
	return _Main.Contract.EXECUTEPERMISSION(&_Main.CallOpts)
}

// MASTERREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x5db2bad0.
//
// Solidity: function MASTER_REGISTRY_RESOURCE() view returns(string)
func (_Main *MainCaller) MASTERREGISTRYRESOURCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "MASTER_REGISTRY_RESOURCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MASTERREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x5db2bad0.
//
// Solidity: function MASTER_REGISTRY_RESOURCE() view returns(string)
func (_Main *MainSession) MASTERREGISTRYRESOURCE() (string, error) {
	return _Main.Contract.MASTERREGISTRYRESOURCE(&_Main.CallOpts)
}

// MASTERREGISTRYRESOURCE is a free data retrieval call binding the contract method 0x5db2bad0.
//
// Solidity: function MASTER_REGISTRY_RESOURCE() view returns(string)
func (_Main *MainCallerSession) MASTERREGISTRYRESOURCE() (string, error) {
	return _Main.Contract.MASTERREGISTRYRESOURCE(&_Main.CallOpts)
}

// MASTERROLE is a free data retrieval call binding the contract method 0xdc224863.
//
// Solidity: function MASTER_ROLE() view returns(string)
func (_Main *MainCaller) MASTERROLE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "MASTER_ROLE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MASTERROLE is a free data retrieval call binding the contract method 0xdc224863.
//
// Solidity: function MASTER_ROLE() view returns(string)
func (_Main *MainSession) MASTERROLE() (string, error) {
	return _Main.Contract.MASTERROLE(&_Main.CallOpts)
}

// MASTERROLE is a free data retrieval call binding the contract method 0xdc224863.
//
// Solidity: function MASTER_ROLE() view returns(string)
func (_Main *MainCallerSession) MASTERROLE() (string, error) {
	return _Main.Contract.MASTERROLE(&_Main.CallOpts)
}

// RBACRESOURCE is a free data retrieval call binding the contract method 0x733352b3.
//
// Solidity: function RBAC_RESOURCE() view returns(string)
func (_Main *MainCaller) RBACRESOURCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "RBAC_RESOURCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RBACRESOURCE is a free data retrieval call binding the contract method 0x733352b3.
//
// Solidity: function RBAC_RESOURCE() view returns(string)
func (_Main *MainSession) RBACRESOURCE() (string, error) {
	return _Main.Contract.RBACRESOURCE(&_Main.CallOpts)
}

// RBACRESOURCE is a free data retrieval call binding the contract method 0x733352b3.
//
// Solidity: function RBAC_RESOURCE() view returns(string)
func (_Main *MainCallerSession) RBACRESOURCE() (string, error) {
	return _Main.Contract.RBACRESOURCE(&_Main.CallOpts)
}

// READPERMISSION is a free data retrieval call binding the contract method 0x03bc0b3e.
//
// Solidity: function READ_PERMISSION() view returns(string)
func (_Main *MainCaller) READPERMISSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "READ_PERMISSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// READPERMISSION is a free data retrieval call binding the contract method 0x03bc0b3e.
//
// Solidity: function READ_PERMISSION() view returns(string)
func (_Main *MainSession) READPERMISSION() (string, error) {
	return _Main.Contract.READPERMISSION(&_Main.CallOpts)
}

// READPERMISSION is a free data retrieval call binding the contract method 0x03bc0b3e.
//
// Solidity: function READ_PERMISSION() view returns(string)
func (_Main *MainCallerSession) READPERMISSION() (string, error) {
	return _Main.Contract.READPERMISSION(&_Main.CallOpts)
}

// REVIEWABLEREQUESTSRESOURCE is a free data retrieval call binding the contract method 0xf55303bd.
//
// Solidity: function REVIEWABLE_REQUESTS_RESOURCE() view returns(string)
func (_Main *MainCaller) REVIEWABLEREQUESTSRESOURCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "REVIEWABLE_REQUESTS_RESOURCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// REVIEWABLEREQUESTSRESOURCE is a free data retrieval call binding the contract method 0xf55303bd.
//
// Solidity: function REVIEWABLE_REQUESTS_RESOURCE() view returns(string)
func (_Main *MainSession) REVIEWABLEREQUESTSRESOURCE() (string, error) {
	return _Main.Contract.REVIEWABLEREQUESTSRESOURCE(&_Main.CallOpts)
}

// REVIEWABLEREQUESTSRESOURCE is a free data retrieval call binding the contract method 0xf55303bd.
//
// Solidity: function REVIEWABLE_REQUESTS_RESOURCE() view returns(string)
func (_Main *MainCallerSession) REVIEWABLEREQUESTSRESOURCE() (string, error) {
	return _Main.Contract.REVIEWABLEREQUESTSRESOURCE(&_Main.CallOpts)
}

// UPDATEPERMISSION is a free data retrieval call binding the contract method 0x0ead6f1e.
//
// Solidity: function UPDATE_PERMISSION() view returns(string)
func (_Main *MainCaller) UPDATEPERMISSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "UPDATE_PERMISSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPDATEPERMISSION is a free data retrieval call binding the contract method 0x0ead6f1e.
//
// Solidity: function UPDATE_PERMISSION() view returns(string)
func (_Main *MainSession) UPDATEPERMISSION() (string, error) {
	return _Main.Contract.UPDATEPERMISSION(&_Main.CallOpts)
}

// UPDATEPERMISSION is a free data retrieval call binding the contract method 0x0ead6f1e.
//
// Solidity: function UPDATE_PERMISSION() view returns(string)
func (_Main *MainCallerSession) UPDATEPERMISSION() (string, error) {
	return _Main.Contract.UPDATEPERMISSION(&_Main.CallOpts)
}

// GetRolePermissions is a free data retrieval call binding the contract method 0x002f5bc0.
//
// Solidity: function getRolePermissions(string role) view returns((string,string[])[] allowed, (string,string[])[] disallowed)
func (_Main *MainCaller) GetRolePermissions(opts *bind.CallOpts, role string) (struct {
	Allowed    []IRBACResourceWithPermissions
	Disallowed []IRBACResourceWithPermissions
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getRolePermissions", role)

	outstruct := new(struct {
		Allowed    []IRBACResourceWithPermissions
		Disallowed []IRBACResourceWithPermissions
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Allowed = *abi.ConvertType(out[0], new([]IRBACResourceWithPermissions)).(*[]IRBACResourceWithPermissions)
	outstruct.Disallowed = *abi.ConvertType(out[1], new([]IRBACResourceWithPermissions)).(*[]IRBACResourceWithPermissions)

	return *outstruct, err

}

// GetRolePermissions is a free data retrieval call binding the contract method 0x002f5bc0.
//
// Solidity: function getRolePermissions(string role) view returns((string,string[])[] allowed, (string,string[])[] disallowed)
func (_Main *MainSession) GetRolePermissions(role string) (struct {
	Allowed    []IRBACResourceWithPermissions
	Disallowed []IRBACResourceWithPermissions
}, error) {
	return _Main.Contract.GetRolePermissions(&_Main.CallOpts, role)
}

// GetRolePermissions is a free data retrieval call binding the contract method 0x002f5bc0.
//
// Solidity: function getRolePermissions(string role) view returns((string,string[])[] allowed, (string,string[])[] disallowed)
func (_Main *MainCallerSession) GetRolePermissions(role string) (struct {
	Allowed    []IRBACResourceWithPermissions
	Disallowed []IRBACResourceWithPermissions
}, error) {
	return _Main.Contract.GetRolePermissions(&_Main.CallOpts, role)
}

// GetUserRoles is a free data retrieval call binding the contract method 0x06a36aee.
//
// Solidity: function getUserRoles(address who) view returns(string[] roles)
func (_Main *MainCaller) GetUserRoles(opts *bind.CallOpts, who common.Address) ([]string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getUserRoles", who)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetUserRoles is a free data retrieval call binding the contract method 0x06a36aee.
//
// Solidity: function getUserRoles(address who) view returns(string[] roles)
func (_Main *MainSession) GetUserRoles(who common.Address) ([]string, error) {
	return _Main.Contract.GetUserRoles(&_Main.CallOpts, who)
}

// GetUserRoles is a free data retrieval call binding the contract method 0x06a36aee.
//
// Solidity: function getUserRoles(address who) view returns(string[] roles)
func (_Main *MainCallerSession) GetUserRoles(who common.Address) ([]string, error) {
	return _Main.Contract.GetUserRoles(&_Main.CallOpts, who)
}

// HasConstantsRegistryCreatePermission is a free data retrieval call binding the contract method 0x5b44a636.
//
// Solidity: function hasConstantsRegistryCreatePermission(address account_) view returns(bool)
func (_Main *MainCaller) HasConstantsRegistryCreatePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "hasConstantsRegistryCreatePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasConstantsRegistryCreatePermission is a free data retrieval call binding the contract method 0x5b44a636.
//
// Solidity: function hasConstantsRegistryCreatePermission(address account_) view returns(bool)
func (_Main *MainSession) HasConstantsRegistryCreatePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasConstantsRegistryCreatePermission(&_Main.CallOpts, account_)
}

// HasConstantsRegistryCreatePermission is a free data retrieval call binding the contract method 0x5b44a636.
//
// Solidity: function hasConstantsRegistryCreatePermission(address account_) view returns(bool)
func (_Main *MainCallerSession) HasConstantsRegistryCreatePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasConstantsRegistryCreatePermission(&_Main.CallOpts, account_)
}

// HasConstantsRegistryDeletePermission is a free data retrieval call binding the contract method 0x05e10d32.
//
// Solidity: function hasConstantsRegistryDeletePermission(address account_) view returns(bool)
func (_Main *MainCaller) HasConstantsRegistryDeletePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "hasConstantsRegistryDeletePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasConstantsRegistryDeletePermission is a free data retrieval call binding the contract method 0x05e10d32.
//
// Solidity: function hasConstantsRegistryDeletePermission(address account_) view returns(bool)
func (_Main *MainSession) HasConstantsRegistryDeletePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasConstantsRegistryDeletePermission(&_Main.CallOpts, account_)
}

// HasConstantsRegistryDeletePermission is a free data retrieval call binding the contract method 0x05e10d32.
//
// Solidity: function hasConstantsRegistryDeletePermission(address account_) view returns(bool)
func (_Main *MainCallerSession) HasConstantsRegistryDeletePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasConstantsRegistryDeletePermission(&_Main.CallOpts, account_)
}

// HasMasterContractsRegistryCreatePermission is a free data retrieval call binding the contract method 0x25536582.
//
// Solidity: function hasMasterContractsRegistryCreatePermission(address account_) view returns(bool)
func (_Main *MainCaller) HasMasterContractsRegistryCreatePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "hasMasterContractsRegistryCreatePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMasterContractsRegistryCreatePermission is a free data retrieval call binding the contract method 0x25536582.
//
// Solidity: function hasMasterContractsRegistryCreatePermission(address account_) view returns(bool)
func (_Main *MainSession) HasMasterContractsRegistryCreatePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasMasterContractsRegistryCreatePermission(&_Main.CallOpts, account_)
}

// HasMasterContractsRegistryCreatePermission is a free data retrieval call binding the contract method 0x25536582.
//
// Solidity: function hasMasterContractsRegistryCreatePermission(address account_) view returns(bool)
func (_Main *MainCallerSession) HasMasterContractsRegistryCreatePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasMasterContractsRegistryCreatePermission(&_Main.CallOpts, account_)
}

// HasMasterContractsRegistryDeletePermission is a free data retrieval call binding the contract method 0x7043f0d8.
//
// Solidity: function hasMasterContractsRegistryDeletePermission(address account_) view returns(bool)
func (_Main *MainCaller) HasMasterContractsRegistryDeletePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "hasMasterContractsRegistryDeletePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMasterContractsRegistryDeletePermission is a free data retrieval call binding the contract method 0x7043f0d8.
//
// Solidity: function hasMasterContractsRegistryDeletePermission(address account_) view returns(bool)
func (_Main *MainSession) HasMasterContractsRegistryDeletePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasMasterContractsRegistryDeletePermission(&_Main.CallOpts, account_)
}

// HasMasterContractsRegistryDeletePermission is a free data retrieval call binding the contract method 0x7043f0d8.
//
// Solidity: function hasMasterContractsRegistryDeletePermission(address account_) view returns(bool)
func (_Main *MainCallerSession) HasMasterContractsRegistryDeletePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasMasterContractsRegistryDeletePermission(&_Main.CallOpts, account_)
}

// HasMasterContractsRegistryUpdatePermission is a free data retrieval call binding the contract method 0x0aefca57.
//
// Solidity: function hasMasterContractsRegistryUpdatePermission(address account_) view returns(bool)
func (_Main *MainCaller) HasMasterContractsRegistryUpdatePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "hasMasterContractsRegistryUpdatePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMasterContractsRegistryUpdatePermission is a free data retrieval call binding the contract method 0x0aefca57.
//
// Solidity: function hasMasterContractsRegistryUpdatePermission(address account_) view returns(bool)
func (_Main *MainSession) HasMasterContractsRegistryUpdatePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasMasterContractsRegistryUpdatePermission(&_Main.CallOpts, account_)
}

// HasMasterContractsRegistryUpdatePermission is a free data retrieval call binding the contract method 0x0aefca57.
//
// Solidity: function hasMasterContractsRegistryUpdatePermission(address account_) view returns(bool)
func (_Main *MainCallerSession) HasMasterContractsRegistryUpdatePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasMasterContractsRegistryUpdatePermission(&_Main.CallOpts, account_)
}

// HasPermission is a free data retrieval call binding the contract method 0x7951c6da.
//
// Solidity: function hasPermission(address who, string resource, string permission) view returns(bool)
func (_Main *MainCaller) HasPermission(opts *bind.CallOpts, who common.Address, resource string, permission string) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "hasPermission", who, resource, permission)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0x7951c6da.
//
// Solidity: function hasPermission(address who, string resource, string permission) view returns(bool)
func (_Main *MainSession) HasPermission(who common.Address, resource string, permission string) (bool, error) {
	return _Main.Contract.HasPermission(&_Main.CallOpts, who, resource, permission)
}

// HasPermission is a free data retrieval call binding the contract method 0x7951c6da.
//
// Solidity: function hasPermission(address who, string resource, string permission) view returns(bool)
func (_Main *MainCallerSession) HasPermission(who common.Address, resource string, permission string) (bool, error) {
	return _Main.Contract.HasPermission(&_Main.CallOpts, who, resource, permission)
}

// HasReviewableRequestsCreatePermission is a free data retrieval call binding the contract method 0xe6cdf5d0.
//
// Solidity: function hasReviewableRequestsCreatePermission(address account_) view returns(bool)
func (_Main *MainCaller) HasReviewableRequestsCreatePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "hasReviewableRequestsCreatePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasReviewableRequestsCreatePermission is a free data retrieval call binding the contract method 0xe6cdf5d0.
//
// Solidity: function hasReviewableRequestsCreatePermission(address account_) view returns(bool)
func (_Main *MainSession) HasReviewableRequestsCreatePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasReviewableRequestsCreatePermission(&_Main.CallOpts, account_)
}

// HasReviewableRequestsCreatePermission is a free data retrieval call binding the contract method 0xe6cdf5d0.
//
// Solidity: function hasReviewableRequestsCreatePermission(address account_) view returns(bool)
func (_Main *MainCallerSession) HasReviewableRequestsCreatePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasReviewableRequestsCreatePermission(&_Main.CallOpts, account_)
}

// HasReviewableRequestsDeletePermission is a free data retrieval call binding the contract method 0x6d19225e.
//
// Solidity: function hasReviewableRequestsDeletePermission(address account_) view returns(bool)
func (_Main *MainCaller) HasReviewableRequestsDeletePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "hasReviewableRequestsDeletePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasReviewableRequestsDeletePermission is a free data retrieval call binding the contract method 0x6d19225e.
//
// Solidity: function hasReviewableRequestsDeletePermission(address account_) view returns(bool)
func (_Main *MainSession) HasReviewableRequestsDeletePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasReviewableRequestsDeletePermission(&_Main.CallOpts, account_)
}

// HasReviewableRequestsDeletePermission is a free data retrieval call binding the contract method 0x6d19225e.
//
// Solidity: function hasReviewableRequestsDeletePermission(address account_) view returns(bool)
func (_Main *MainCallerSession) HasReviewableRequestsDeletePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasReviewableRequestsDeletePermission(&_Main.CallOpts, account_)
}

// HasReviewableRequestsExecutePermission is a free data retrieval call binding the contract method 0x20b55d81.
//
// Solidity: function hasReviewableRequestsExecutePermission(address account_) view returns(bool)
func (_Main *MainCaller) HasReviewableRequestsExecutePermission(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "hasReviewableRequestsExecutePermission", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasReviewableRequestsExecutePermission is a free data retrieval call binding the contract method 0x20b55d81.
//
// Solidity: function hasReviewableRequestsExecutePermission(address account_) view returns(bool)
func (_Main *MainSession) HasReviewableRequestsExecutePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasReviewableRequestsExecutePermission(&_Main.CallOpts, account_)
}

// HasReviewableRequestsExecutePermission is a free data retrieval call binding the contract method 0x20b55d81.
//
// Solidity: function hasReviewableRequestsExecutePermission(address account_) view returns(bool)
func (_Main *MainCallerSession) HasReviewableRequestsExecutePermission(account_ common.Address) (bool, error) {
	return _Main.Contract.HasReviewableRequestsExecutePermission(&_Main.CallOpts, account_)
}

// MasterAccessManagementInit is a paid mutator transaction binding the contract method 0x305e4c6c.
//
// Solidity: function __MasterAccessManagement_init(address master_) returns()
func (_Main *MainTransactor) MasterAccessManagementInit(opts *bind.TransactOpts, master_ common.Address) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "__MasterAccessManagement_init", master_)
}

// MasterAccessManagementInit is a paid mutator transaction binding the contract method 0x305e4c6c.
//
// Solidity: function __MasterAccessManagement_init(address master_) returns()
func (_Main *MainSession) MasterAccessManagementInit(master_ common.Address) (*types.Transaction, error) {
	return _Main.Contract.MasterAccessManagementInit(&_Main.TransactOpts, master_)
}

// MasterAccessManagementInit is a paid mutator transaction binding the contract method 0x305e4c6c.
//
// Solidity: function __MasterAccessManagement_init(address master_) returns()
func (_Main *MainTransactorSession) MasterAccessManagementInit(master_ common.Address) (*types.Transaction, error) {
	return _Main.Contract.MasterAccessManagementInit(&_Main.TransactOpts, master_)
}

// AddPermissionsToRole is a paid mutator transaction binding the contract method 0x37ff630d.
//
// Solidity: function addPermissionsToRole(string role, (string,string[])[] permissionsToAdd, bool allowed) returns()
func (_Main *MainTransactor) AddPermissionsToRole(opts *bind.TransactOpts, role string, permissionsToAdd []IRBACResourceWithPermissions, allowed bool) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "addPermissionsToRole", role, permissionsToAdd, allowed)
}

// AddPermissionsToRole is a paid mutator transaction binding the contract method 0x37ff630d.
//
// Solidity: function addPermissionsToRole(string role, (string,string[])[] permissionsToAdd, bool allowed) returns()
func (_Main *MainSession) AddPermissionsToRole(role string, permissionsToAdd []IRBACResourceWithPermissions, allowed bool) (*types.Transaction, error) {
	return _Main.Contract.AddPermissionsToRole(&_Main.TransactOpts, role, permissionsToAdd, allowed)
}

// AddPermissionsToRole is a paid mutator transaction binding the contract method 0x37ff630d.
//
// Solidity: function addPermissionsToRole(string role, (string,string[])[] permissionsToAdd, bool allowed) returns()
func (_Main *MainTransactorSession) AddPermissionsToRole(role string, permissionsToAdd []IRBACResourceWithPermissions, allowed bool) (*types.Transaction, error) {
	return _Main.Contract.AddPermissionsToRole(&_Main.TransactOpts, role, permissionsToAdd, allowed)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xee2f6ce5.
//
// Solidity: function grantRoles(address to, string[] rolesToGrant) returns()
func (_Main *MainTransactor) GrantRoles(opts *bind.TransactOpts, to common.Address, rolesToGrant []string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "grantRoles", to, rolesToGrant)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xee2f6ce5.
//
// Solidity: function grantRoles(address to, string[] rolesToGrant) returns()
func (_Main *MainSession) GrantRoles(to common.Address, rolesToGrant []string) (*types.Transaction, error) {
	return _Main.Contract.GrantRoles(&_Main.TransactOpts, to, rolesToGrant)
}

// GrantRoles is a paid mutator transaction binding the contract method 0xee2f6ce5.
//
// Solidity: function grantRoles(address to, string[] rolesToGrant) returns()
func (_Main *MainTransactorSession) GrantRoles(to common.Address, rolesToGrant []string) (*types.Transaction, error) {
	return _Main.Contract.GrantRoles(&_Main.TransactOpts, to, rolesToGrant)
}

// RemovePermissionsFromRole is a paid mutator transaction binding the contract method 0x75e025e7.
//
// Solidity: function removePermissionsFromRole(string role, (string,string[])[] permissionsToRemove, bool allowed) returns()
func (_Main *MainTransactor) RemovePermissionsFromRole(opts *bind.TransactOpts, role string, permissionsToRemove []IRBACResourceWithPermissions, allowed bool) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removePermissionsFromRole", role, permissionsToRemove, allowed)
}

// RemovePermissionsFromRole is a paid mutator transaction binding the contract method 0x75e025e7.
//
// Solidity: function removePermissionsFromRole(string role, (string,string[])[] permissionsToRemove, bool allowed) returns()
func (_Main *MainSession) RemovePermissionsFromRole(role string, permissionsToRemove []IRBACResourceWithPermissions, allowed bool) (*types.Transaction, error) {
	return _Main.Contract.RemovePermissionsFromRole(&_Main.TransactOpts, role, permissionsToRemove, allowed)
}

// RemovePermissionsFromRole is a paid mutator transaction binding the contract method 0x75e025e7.
//
// Solidity: function removePermissionsFromRole(string role, (string,string[])[] permissionsToRemove, bool allowed) returns()
func (_Main *MainTransactorSession) RemovePermissionsFromRole(role string, permissionsToRemove []IRBACResourceWithPermissions, allowed bool) (*types.Transaction, error) {
	return _Main.Contract.RemovePermissionsFromRole(&_Main.TransactOpts, role, permissionsToRemove, allowed)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4f0d84e3.
//
// Solidity: function revokeRoles(address from, string[] rolesToRevoke) returns()
func (_Main *MainTransactor) RevokeRoles(opts *bind.TransactOpts, from common.Address, rolesToRevoke []string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "revokeRoles", from, rolesToRevoke)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4f0d84e3.
//
// Solidity: function revokeRoles(address from, string[] rolesToRevoke) returns()
func (_Main *MainSession) RevokeRoles(from common.Address, rolesToRevoke []string) (*types.Transaction, error) {
	return _Main.Contract.RevokeRoles(&_Main.TransactOpts, from, rolesToRevoke)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4f0d84e3.
//
// Solidity: function revokeRoles(address from, string[] rolesToRevoke) returns()
func (_Main *MainTransactorSession) RevokeRoles(from common.Address, rolesToRevoke []string) (*types.Transaction, error) {
	return _Main.Contract.RevokeRoles(&_Main.TransactOpts, from, rolesToRevoke)
}

// MainAddedPermissionsIterator is returned from FilterAddedPermissions and is used to iterate over the raw logs and unpacked data for AddedPermissions events raised by the Main contract.
type MainAddedPermissionsIterator struct {
	Event *MainAddedPermissions // Event containing the contract specifics and raw log

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
func (it *MainAddedPermissionsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainAddedPermissions)
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
		it.Event = new(MainAddedPermissions)
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
func (it *MainAddedPermissionsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainAddedPermissionsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainAddedPermissions represents a AddedPermissions event raised by the Main contract.
type MainAddedPermissions struct {
	Role             string
	Resource         string
	PermissionsToAdd []string
	Allowed          bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAddedPermissions is a free log retrieval operation binding the contract event 0x06ea5bb05be257dd30f31cc76f4077a3df230e09f7e33e9b52fb09c40a8f695e.
//
// Solidity: event AddedPermissions(string role, string resource, string[] permissionsToAdd, bool allowed)
func (_Main *MainFilterer) FilterAddedPermissions(opts *bind.FilterOpts) (*MainAddedPermissionsIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "AddedPermissions")
	if err != nil {
		return nil, err
	}
	return &MainAddedPermissionsIterator{contract: _Main.contract, event: "AddedPermissions", logs: logs, sub: sub}, nil
}

// WatchAddedPermissions is a free log subscription operation binding the contract event 0x06ea5bb05be257dd30f31cc76f4077a3df230e09f7e33e9b52fb09c40a8f695e.
//
// Solidity: event AddedPermissions(string role, string resource, string[] permissionsToAdd, bool allowed)
func (_Main *MainFilterer) WatchAddedPermissions(opts *bind.WatchOpts, sink chan<- *MainAddedPermissions) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "AddedPermissions")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainAddedPermissions)
				if err := _Main.contract.UnpackLog(event, "AddedPermissions", log); err != nil {
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

// ParseAddedPermissions is a log parse operation binding the contract event 0x06ea5bb05be257dd30f31cc76f4077a3df230e09f7e33e9b52fb09c40a8f695e.
//
// Solidity: event AddedPermissions(string role, string resource, string[] permissionsToAdd, bool allowed)
func (_Main *MainFilterer) ParseAddedPermissions(log types.Log) (*MainAddedPermissions, error) {
	event := new(MainAddedPermissions)
	if err := _Main.contract.UnpackLog(event, "AddedPermissions", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainGrantedRolesIterator is returned from FilterGrantedRoles and is used to iterate over the raw logs and unpacked data for GrantedRoles events raised by the Main contract.
type MainGrantedRolesIterator struct {
	Event *MainGrantedRoles // Event containing the contract specifics and raw log

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
func (it *MainGrantedRolesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainGrantedRoles)
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
		it.Event = new(MainGrantedRoles)
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
func (it *MainGrantedRolesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainGrantedRolesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainGrantedRoles represents a GrantedRoles event raised by the Main contract.
type MainGrantedRoles struct {
	To           common.Address
	RolesToGrant []string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGrantedRoles is a free log retrieval operation binding the contract event 0x44240f5b60cedf44a65c3717503d91c46a899ef33c5348880e4c29131ac87311.
//
// Solidity: event GrantedRoles(address to, string[] rolesToGrant)
func (_Main *MainFilterer) FilterGrantedRoles(opts *bind.FilterOpts) (*MainGrantedRolesIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "GrantedRoles")
	if err != nil {
		return nil, err
	}
	return &MainGrantedRolesIterator{contract: _Main.contract, event: "GrantedRoles", logs: logs, sub: sub}, nil
}

// WatchGrantedRoles is a free log subscription operation binding the contract event 0x44240f5b60cedf44a65c3717503d91c46a899ef33c5348880e4c29131ac87311.
//
// Solidity: event GrantedRoles(address to, string[] rolesToGrant)
func (_Main *MainFilterer) WatchGrantedRoles(opts *bind.WatchOpts, sink chan<- *MainGrantedRoles) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "GrantedRoles")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainGrantedRoles)
				if err := _Main.contract.UnpackLog(event, "GrantedRoles", log); err != nil {
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

// ParseGrantedRoles is a log parse operation binding the contract event 0x44240f5b60cedf44a65c3717503d91c46a899ef33c5348880e4c29131ac87311.
//
// Solidity: event GrantedRoles(address to, string[] rolesToGrant)
func (_Main *MainFilterer) ParseGrantedRoles(log types.Log) (*MainGrantedRoles, error) {
	event := new(MainGrantedRoles)
	if err := _Main.contract.UnpackLog(event, "GrantedRoles", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRemovedPermissionsIterator is returned from FilterRemovedPermissions and is used to iterate over the raw logs and unpacked data for RemovedPermissions events raised by the Main contract.
type MainRemovedPermissionsIterator struct {
	Event *MainRemovedPermissions // Event containing the contract specifics and raw log

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
func (it *MainRemovedPermissionsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRemovedPermissions)
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
		it.Event = new(MainRemovedPermissions)
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
func (it *MainRemovedPermissionsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRemovedPermissionsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRemovedPermissions represents a RemovedPermissions event raised by the Main contract.
type MainRemovedPermissions struct {
	Role                string
	Resource            string
	PermissionsToRemove []string
	Allowed             bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRemovedPermissions is a free log retrieval operation binding the contract event 0x27f963dfdedb973db8acb94be0be26fbd4c55498bc54cfd0733b9da9c0ab8296.
//
// Solidity: event RemovedPermissions(string role, string resource, string[] permissionsToRemove, bool allowed)
func (_Main *MainFilterer) FilterRemovedPermissions(opts *bind.FilterOpts) (*MainRemovedPermissionsIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RemovedPermissions")
	if err != nil {
		return nil, err
	}
	return &MainRemovedPermissionsIterator{contract: _Main.contract, event: "RemovedPermissions", logs: logs, sub: sub}, nil
}

// WatchRemovedPermissions is a free log subscription operation binding the contract event 0x27f963dfdedb973db8acb94be0be26fbd4c55498bc54cfd0733b9da9c0ab8296.
//
// Solidity: event RemovedPermissions(string role, string resource, string[] permissionsToRemove, bool allowed)
func (_Main *MainFilterer) WatchRemovedPermissions(opts *bind.WatchOpts, sink chan<- *MainRemovedPermissions) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RemovedPermissions")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRemovedPermissions)
				if err := _Main.contract.UnpackLog(event, "RemovedPermissions", log); err != nil {
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

// ParseRemovedPermissions is a log parse operation binding the contract event 0x27f963dfdedb973db8acb94be0be26fbd4c55498bc54cfd0733b9da9c0ab8296.
//
// Solidity: event RemovedPermissions(string role, string resource, string[] permissionsToRemove, bool allowed)
func (_Main *MainFilterer) ParseRemovedPermissions(log types.Log) (*MainRemovedPermissions, error) {
	event := new(MainRemovedPermissions)
	if err := _Main.contract.UnpackLog(event, "RemovedPermissions", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainRevokedRolesIterator is returned from FilterRevokedRoles and is used to iterate over the raw logs and unpacked data for RevokedRoles events raised by the Main contract.
type MainRevokedRolesIterator struct {
	Event *MainRevokedRoles // Event containing the contract specifics and raw log

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
func (it *MainRevokedRolesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainRevokedRoles)
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
		it.Event = new(MainRevokedRoles)
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
func (it *MainRevokedRolesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainRevokedRolesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainRevokedRoles represents a RevokedRoles event raised by the Main contract.
type MainRevokedRoles struct {
	From          common.Address
	RolesToRevoke []string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevokedRoles is a free log retrieval operation binding the contract event 0x037c273ac6ee1105154063d4b014a5afeec9981076c152f47a0899c6e2854740.
//
// Solidity: event RevokedRoles(address from, string[] rolesToRevoke)
func (_Main *MainFilterer) FilterRevokedRoles(opts *bind.FilterOpts) (*MainRevokedRolesIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "RevokedRoles")
	if err != nil {
		return nil, err
	}
	return &MainRevokedRolesIterator{contract: _Main.contract, event: "RevokedRoles", logs: logs, sub: sub}, nil
}

// WatchRevokedRoles is a free log subscription operation binding the contract event 0x037c273ac6ee1105154063d4b014a5afeec9981076c152f47a0899c6e2854740.
//
// Solidity: event RevokedRoles(address from, string[] rolesToRevoke)
func (_Main *MainFilterer) WatchRevokedRoles(opts *bind.WatchOpts, sink chan<- *MainRevokedRoles) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "RevokedRoles")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainRevokedRoles)
				if err := _Main.contract.UnpackLog(event, "RevokedRoles", log); err != nil {
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

// ParseRevokedRoles is a log parse operation binding the contract event 0x037c273ac6ee1105154063d4b014a5afeec9981076c152f47a0899c6e2854740.
//
// Solidity: event RevokedRoles(address from, string[] rolesToRevoke)
func (_Main *MainFilterer) ParseRevokedRoles(log types.Log) (*MainRevokedRoles, error) {
	event := new(MainRevokedRoles)
	if err := _Main.contract.UnpackLog(event, "RevokedRoles", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
