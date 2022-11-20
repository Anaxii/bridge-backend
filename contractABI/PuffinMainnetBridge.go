// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// PuffinBridgeBridgeRequest is an auto generated low-level Go binding around an user-defined struct.
type PuffinBridgeBridgeRequest struct {
	Id     [32]byte
	User   common.Address
	Asset  common.Address
	Amount *big.Int
	Expiry *big.Int
}

// PuffinMainnetBridgeMetaData contains all meta data concerning the PuffinMainnetBridge contract.
var PuffinMainnetBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"bridgeIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"BridgeIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"BridgeOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"BridgeOutCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"BridgeOutWarm\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"cancelOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"markInComplete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"proposeOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"removeVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setKYC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setWarm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeInComplete\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeOutComplete\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"getRequestInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structPuffinBridge.BridgeRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"puffinAssets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"puffinKYC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"puffinWarmWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PuffinMainnetBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use PuffinMainnetBridgeMetaData.ABI instead.
var PuffinMainnetBridgeABI = PuffinMainnetBridgeMetaData.ABI

// PuffinMainnetBridge is an auto generated Go binding around an Ethereum contract.
type PuffinMainnetBridge struct {
	PuffinMainnetBridgeCaller     // Read-only binding to the contract
	PuffinMainnetBridgeTransactor // Write-only binding to the contract
	PuffinMainnetBridgeFilterer   // Log filterer for contract events
}

// PuffinMainnetBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PuffinMainnetBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinMainnetBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PuffinMainnetBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinMainnetBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PuffinMainnetBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinMainnetBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PuffinMainnetBridgeSession struct {
	Contract     *PuffinMainnetBridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PuffinMainnetBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PuffinMainnetBridgeCallerSession struct {
	Contract *PuffinMainnetBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PuffinMainnetBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PuffinMainnetBridgeTransactorSession struct {
	Contract     *PuffinMainnetBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PuffinMainnetBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PuffinMainnetBridgeRaw struct {
	Contract *PuffinMainnetBridge // Generic contract binding to access the raw methods on
}

// PuffinMainnetBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PuffinMainnetBridgeCallerRaw struct {
	Contract *PuffinMainnetBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// PuffinMainnetBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PuffinMainnetBridgeTransactorRaw struct {
	Contract *PuffinMainnetBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPuffinMainnetBridge creates a new instance of PuffinMainnetBridge, bound to a specific deployed contract.
func NewPuffinMainnetBridge(address common.Address, backend bind.ContractBackend) (*PuffinMainnetBridge, error) {
	contract, err := bindPuffinMainnetBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PuffinMainnetBridge{PuffinMainnetBridgeCaller: PuffinMainnetBridgeCaller{contract: contract}, PuffinMainnetBridgeTransactor: PuffinMainnetBridgeTransactor{contract: contract}, PuffinMainnetBridgeFilterer: PuffinMainnetBridgeFilterer{contract: contract}}, nil
}

// NewPuffinMainnetBridgeCaller creates a new read-only instance of PuffinMainnetBridge, bound to a specific deployed contract.
func NewPuffinMainnetBridgeCaller(address common.Address, caller bind.ContractCaller) (*PuffinMainnetBridgeCaller, error) {
	contract, err := bindPuffinMainnetBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinMainnetBridgeCaller{contract: contract}, nil
}

// NewPuffinMainnetBridgeTransactor creates a new write-only instance of PuffinMainnetBridge, bound to a specific deployed contract.
func NewPuffinMainnetBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*PuffinMainnetBridgeTransactor, error) {
	contract, err := bindPuffinMainnetBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinMainnetBridgeTransactor{contract: contract}, nil
}

// NewPuffinMainnetBridgeFilterer creates a new log filterer instance of PuffinMainnetBridge, bound to a specific deployed contract.
func NewPuffinMainnetBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*PuffinMainnetBridgeFilterer, error) {
	contract, err := bindPuffinMainnetBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PuffinMainnetBridgeFilterer{contract: contract}, nil
}

// bindPuffinMainnetBridge binds a generic wrapper to an already deployed contract.
func bindPuffinMainnetBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PuffinMainnetBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinMainnetBridge *PuffinMainnetBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinMainnetBridge.Contract.PuffinMainnetBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinMainnetBridge *PuffinMainnetBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.PuffinMainnetBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinMainnetBridge *PuffinMainnetBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.PuffinMainnetBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinMainnetBridge *PuffinMainnetBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinMainnetBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.contract.Transact(opts, method, params...)
}

// BridgeInComplete is a free data retrieval call binding the contract method 0xdd68a1c9.
//
// Solidity: function bridgeInComplete(bytes32 ) view returns(bool)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCaller) BridgeInComplete(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _PuffinMainnetBridge.contract.Call(opts, &out, "bridgeInComplete", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeInComplete is a free data retrieval call binding the contract method 0xdd68a1c9.
//
// Solidity: function bridgeInComplete(bytes32 ) view returns(bool)
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) BridgeInComplete(arg0 [32]byte) (bool, error) {
	return _PuffinMainnetBridge.Contract.BridgeInComplete(&_PuffinMainnetBridge.CallOpts, arg0)
}

// BridgeInComplete is a free data retrieval call binding the contract method 0xdd68a1c9.
//
// Solidity: function bridgeInComplete(bytes32 ) view returns(bool)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCallerSession) BridgeInComplete(arg0 [32]byte) (bool, error) {
	return _PuffinMainnetBridge.Contract.BridgeInComplete(&_PuffinMainnetBridge.CallOpts, arg0)
}

// BridgeOutComplete is a free data retrieval call binding the contract method 0x4f9850ce.
//
// Solidity: function bridgeOutComplete(bytes32 ) view returns(bool)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCaller) BridgeOutComplete(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _PuffinMainnetBridge.contract.Call(opts, &out, "bridgeOutComplete", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeOutComplete is a free data retrieval call binding the contract method 0x4f9850ce.
//
// Solidity: function bridgeOutComplete(bytes32 ) view returns(bool)
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) BridgeOutComplete(arg0 [32]byte) (bool, error) {
	return _PuffinMainnetBridge.Contract.BridgeOutComplete(&_PuffinMainnetBridge.CallOpts, arg0)
}

// BridgeOutComplete is a free data retrieval call binding the contract method 0x4f9850ce.
//
// Solidity: function bridgeOutComplete(bytes32 ) view returns(bool)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCallerSession) BridgeOutComplete(arg0 [32]byte) (bool, error) {
	return _PuffinMainnetBridge.Contract.BridgeOutComplete(&_PuffinMainnetBridge.CallOpts, arg0)
}

// GetRequestInfo is a free data retrieval call binding the contract method 0x82d67e5a.
//
// Solidity: function getRequestInfo(bytes32 requestId) view returns((bytes32,address,address,uint256,uint256))
func (_PuffinMainnetBridge *PuffinMainnetBridgeCaller) GetRequestInfo(opts *bind.CallOpts, requestId [32]byte) (PuffinBridgeBridgeRequest, error) {
	var out []interface{}
	err := _PuffinMainnetBridge.contract.Call(opts, &out, "getRequestInfo", requestId)

	if err != nil {
		return *new(PuffinBridgeBridgeRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(PuffinBridgeBridgeRequest)).(*PuffinBridgeBridgeRequest)

	return out0, err

}

// GetRequestInfo is a free data retrieval call binding the contract method 0x82d67e5a.
//
// Solidity: function getRequestInfo(bytes32 requestId) view returns((bytes32,address,address,uint256,uint256))
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) GetRequestInfo(requestId [32]byte) (PuffinBridgeBridgeRequest, error) {
	return _PuffinMainnetBridge.Contract.GetRequestInfo(&_PuffinMainnetBridge.CallOpts, requestId)
}

// GetRequestInfo is a free data retrieval call binding the contract method 0x82d67e5a.
//
// Solidity: function getRequestInfo(bytes32 requestId) view returns((bytes32,address,address,uint256,uint256))
func (_PuffinMainnetBridge *PuffinMainnetBridgeCallerSession) GetRequestInfo(requestId [32]byte) (PuffinBridgeBridgeRequest, error) {
	return _PuffinMainnetBridge.Contract.GetRequestInfo(&_PuffinMainnetBridge.CallOpts, requestId)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address ) view returns(bool)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCaller) IsVoter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PuffinMainnetBridge.contract.Call(opts, &out, "isVoter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address ) view returns(bool)
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) IsVoter(arg0 common.Address) (bool, error) {
	return _PuffinMainnetBridge.Contract.IsVoter(&_PuffinMainnetBridge.CallOpts, arg0)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address ) view returns(bool)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCallerSession) IsVoter(arg0 common.Address) (bool, error) {
	return _PuffinMainnetBridge.Contract.IsVoter(&_PuffinMainnetBridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinMainnetBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) Owner() (common.Address, error) {
	return _PuffinMainnetBridge.Contract.Owner(&_PuffinMainnetBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCallerSession) Owner() (common.Address, error) {
	return _PuffinMainnetBridge.Contract.Owner(&_PuffinMainnetBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PuffinMainnetBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) Paused() (bool, error) {
	return _PuffinMainnetBridge.Contract.Paused(&_PuffinMainnetBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCallerSession) Paused() (bool, error) {
	return _PuffinMainnetBridge.Contract.Paused(&_PuffinMainnetBridge.CallOpts)
}

// PuffinAssets is a free data retrieval call binding the contract method 0x318f2e3e.
//
// Solidity: function puffinAssets() view returns(address)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCaller) PuffinAssets(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinMainnetBridge.contract.Call(opts, &out, "puffinAssets")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PuffinAssets is a free data retrieval call binding the contract method 0x318f2e3e.
//
// Solidity: function puffinAssets() view returns(address)
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) PuffinAssets() (common.Address, error) {
	return _PuffinMainnetBridge.Contract.PuffinAssets(&_PuffinMainnetBridge.CallOpts)
}

// PuffinAssets is a free data retrieval call binding the contract method 0x318f2e3e.
//
// Solidity: function puffinAssets() view returns(address)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCallerSession) PuffinAssets() (common.Address, error) {
	return _PuffinMainnetBridge.Contract.PuffinAssets(&_PuffinMainnetBridge.CallOpts)
}

// PuffinKYC is a free data retrieval call binding the contract method 0x717642d9.
//
// Solidity: function puffinKYC() view returns(address)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCaller) PuffinKYC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinMainnetBridge.contract.Call(opts, &out, "puffinKYC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PuffinKYC is a free data retrieval call binding the contract method 0x717642d9.
//
// Solidity: function puffinKYC() view returns(address)
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) PuffinKYC() (common.Address, error) {
	return _PuffinMainnetBridge.Contract.PuffinKYC(&_PuffinMainnetBridge.CallOpts)
}

// PuffinKYC is a free data retrieval call binding the contract method 0x717642d9.
//
// Solidity: function puffinKYC() view returns(address)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCallerSession) PuffinKYC() (common.Address, error) {
	return _PuffinMainnetBridge.Contract.PuffinKYC(&_PuffinMainnetBridge.CallOpts)
}

// PuffinWarmWallet is a free data retrieval call binding the contract method 0xd5a0ccf8.
//
// Solidity: function puffinWarmWallet() view returns(address)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCaller) PuffinWarmWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinMainnetBridge.contract.Call(opts, &out, "puffinWarmWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PuffinWarmWallet is a free data retrieval call binding the contract method 0xd5a0ccf8.
//
// Solidity: function puffinWarmWallet() view returns(address)
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) PuffinWarmWallet() (common.Address, error) {
	return _PuffinMainnetBridge.Contract.PuffinWarmWallet(&_PuffinMainnetBridge.CallOpts)
}

// PuffinWarmWallet is a free data retrieval call binding the contract method 0xd5a0ccf8.
//
// Solidity: function puffinWarmWallet() view returns(address)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCallerSession) PuffinWarmWallet() (common.Address, error) {
	return _PuffinMainnetBridge.Contract.PuffinWarmWallet(&_PuffinMainnetBridge.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x0526d3fe.
//
// Solidity: function requestCount(bytes32 ) view returns(uint256)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCaller) RequestCount(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PuffinMainnetBridge.contract.Call(opts, &out, "requestCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestCount is a free data retrieval call binding the contract method 0x0526d3fe.
//
// Solidity: function requestCount(bytes32 ) view returns(uint256)
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) RequestCount(arg0 [32]byte) (*big.Int, error) {
	return _PuffinMainnetBridge.Contract.RequestCount(&_PuffinMainnetBridge.CallOpts, arg0)
}

// RequestCount is a free data retrieval call binding the contract method 0x0526d3fe.
//
// Solidity: function requestCount(bytes32 ) view returns(uint256)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCallerSession) RequestCount(arg0 [32]byte) (*big.Int, error) {
	return _PuffinMainnetBridge.Contract.RequestCount(&_PuffinMainnetBridge.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PuffinMainnetBridge.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) Threshold() (*big.Int, error) {
	return _PuffinMainnetBridge.Contract.Threshold(&_PuffinMainnetBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_PuffinMainnetBridge *PuffinMainnetBridgeCallerSession) Threshold() (*big.Int, error) {
	return _PuffinMainnetBridge.Contract.Threshold(&_PuffinMainnetBridge.CallOpts)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(address user) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactor) AddVoter(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.contract.Transact(opts, "addVoter", user)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(address user) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) AddVoter(user common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.AddVoter(&_PuffinMainnetBridge.TransactOpts, user)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(address user) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorSession) AddVoter(user common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.AddVoter(&_PuffinMainnetBridge.TransactOpts, user)
}

// BridgeIn is a paid mutator transaction binding the contract method 0x7107fa82.
//
// Solidity: function bridgeIn(uint256 amount, address asset) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactor) BridgeIn(opts *bind.TransactOpts, amount *big.Int, asset common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.contract.Transact(opts, "bridgeIn", amount, asset)
}

// BridgeIn is a paid mutator transaction binding the contract method 0x7107fa82.
//
// Solidity: function bridgeIn(uint256 amount, address asset) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) BridgeIn(amount *big.Int, asset common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.BridgeIn(&_PuffinMainnetBridge.TransactOpts, amount, asset)
}

// BridgeIn is a paid mutator transaction binding the contract method 0x7107fa82.
//
// Solidity: function bridgeIn(uint256 amount, address asset) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorSession) BridgeIn(amount *big.Int, asset common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.BridgeIn(&_PuffinMainnetBridge.TransactOpts, amount, asset)
}

// CancelOut is a paid mutator transaction binding the contract method 0x1533c400.
//
// Solidity: function cancelOut(bytes32 requestId) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactor) CancelOut(opts *bind.TransactOpts, requestId [32]byte) (*types.Transaction, error) {
	return _PuffinMainnetBridge.contract.Transact(opts, "cancelOut", requestId)
}

// CancelOut is a paid mutator transaction binding the contract method 0x1533c400.
//
// Solidity: function cancelOut(bytes32 requestId) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) CancelOut(requestId [32]byte) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.CancelOut(&_PuffinMainnetBridge.TransactOpts, requestId)
}

// CancelOut is a paid mutator transaction binding the contract method 0x1533c400.
//
// Solidity: function cancelOut(bytes32 requestId) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorSession) CancelOut(requestId [32]byte) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.CancelOut(&_PuffinMainnetBridge.TransactOpts, requestId)
}

// MarkInComplete is a paid mutator transaction binding the contract method 0x49049bb1.
//
// Solidity: function markInComplete(bytes32 requestId) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactor) MarkInComplete(opts *bind.TransactOpts, requestId [32]byte) (*types.Transaction, error) {
	return _PuffinMainnetBridge.contract.Transact(opts, "markInComplete", requestId)
}

// MarkInComplete is a paid mutator transaction binding the contract method 0x49049bb1.
//
// Solidity: function markInComplete(bytes32 requestId) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) MarkInComplete(requestId [32]byte) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.MarkInComplete(&_PuffinMainnetBridge.TransactOpts, requestId)
}

// MarkInComplete is a paid mutator transaction binding the contract method 0x49049bb1.
//
// Solidity: function markInComplete(bytes32 requestId) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorSession) MarkInComplete(requestId [32]byte) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.MarkInComplete(&_PuffinMainnetBridge.TransactOpts, requestId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinMainnetBridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) Pause() (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.Pause(&_PuffinMainnetBridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.Pause(&_PuffinMainnetBridge.TransactOpts)
}

// ProposeOut is a paid mutator transaction binding the contract method 0x809a7022.
//
// Solidity: function proposeOut(address asset, address user, uint256 amount, bytes32 requestId) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactor) ProposeOut(opts *bind.TransactOpts, asset common.Address, user common.Address, amount *big.Int, requestId [32]byte) (*types.Transaction, error) {
	return _PuffinMainnetBridge.contract.Transact(opts, "proposeOut", asset, user, amount, requestId)
}

// ProposeOut is a paid mutator transaction binding the contract method 0x809a7022.
//
// Solidity: function proposeOut(address asset, address user, uint256 amount, bytes32 requestId) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) ProposeOut(asset common.Address, user common.Address, amount *big.Int, requestId [32]byte) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.ProposeOut(&_PuffinMainnetBridge.TransactOpts, asset, user, amount, requestId)
}

// ProposeOut is a paid mutator transaction binding the contract method 0x809a7022.
//
// Solidity: function proposeOut(address asset, address user, uint256 amount, bytes32 requestId) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorSession) ProposeOut(asset common.Address, user common.Address, amount *big.Int, requestId [32]byte) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.ProposeOut(&_PuffinMainnetBridge.TransactOpts, asset, user, amount, requestId)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(address user) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactor) RemoveVoter(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.contract.Transact(opts, "removeVoter", user)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(address user) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) RemoveVoter(user common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.RemoveVoter(&_PuffinMainnetBridge.TransactOpts, user)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(address user) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorSession) RemoveVoter(user common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.RemoveVoter(&_PuffinMainnetBridge.TransactOpts, user)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinMainnetBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.RenounceOwnership(&_PuffinMainnetBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.RenounceOwnership(&_PuffinMainnetBridge.TransactOpts)
}

// SetAssets is a paid mutator transaction binding the contract method 0xee1559da.
//
// Solidity: function setAssets(address _contract) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactor) SetAssets(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.contract.Transact(opts, "setAssets", _contract)
}

// SetAssets is a paid mutator transaction binding the contract method 0xee1559da.
//
// Solidity: function setAssets(address _contract) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) SetAssets(_contract common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.SetAssets(&_PuffinMainnetBridge.TransactOpts, _contract)
}

// SetAssets is a paid mutator transaction binding the contract method 0xee1559da.
//
// Solidity: function setAssets(address _contract) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorSession) SetAssets(_contract common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.SetAssets(&_PuffinMainnetBridge.TransactOpts, _contract)
}

// SetKYC is a paid mutator transaction binding the contract method 0x483a83df.
//
// Solidity: function setKYC(address _contract) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactor) SetKYC(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.contract.Transact(opts, "setKYC", _contract)
}

// SetKYC is a paid mutator transaction binding the contract method 0x483a83df.
//
// Solidity: function setKYC(address _contract) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) SetKYC(_contract common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.SetKYC(&_PuffinMainnetBridge.TransactOpts, _contract)
}

// SetKYC is a paid mutator transaction binding the contract method 0x483a83df.
//
// Solidity: function setKYC(address _contract) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorSession) SetKYC(_contract common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.SetKYC(&_PuffinMainnetBridge.TransactOpts, _contract)
}

// SetWarm is a paid mutator transaction binding the contract method 0x7eea8c34.
//
// Solidity: function setWarm(address _contract) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactor) SetWarm(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.contract.Transact(opts, "setWarm", _contract)
}

// SetWarm is a paid mutator transaction binding the contract method 0x7eea8c34.
//
// Solidity: function setWarm(address _contract) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) SetWarm(_contract common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.SetWarm(&_PuffinMainnetBridge.TransactOpts, _contract)
}

// SetWarm is a paid mutator transaction binding the contract method 0x7eea8c34.
//
// Solidity: function setWarm(address _contract) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorSession) SetWarm(_contract common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.SetWarm(&_PuffinMainnetBridge.TransactOpts, _contract)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.TransferOwnership(&_PuffinMainnetBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.TransferOwnership(&_PuffinMainnetBridge.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinMainnetBridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeSession) Unpause() (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.Unpause(&_PuffinMainnetBridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PuffinMainnetBridge *PuffinMainnetBridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _PuffinMainnetBridge.Contract.Unpause(&_PuffinMainnetBridge.TransactOpts)
}

// PuffinMainnetBridgeBridgeInIterator is returned from FilterBridgeIn and is used to iterate over the raw logs and unpacked data for BridgeIn events raised by the PuffinMainnetBridge contract.
type PuffinMainnetBridgeBridgeInIterator struct {
	Event *PuffinMainnetBridgeBridgeIn // Event containing the contract specifics and raw log

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
func (it *PuffinMainnetBridgeBridgeInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinMainnetBridgeBridgeIn)
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
		it.Event = new(PuffinMainnetBridgeBridgeIn)
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
func (it *PuffinMainnetBridgeBridgeInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinMainnetBridgeBridgeInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinMainnetBridgeBridgeIn represents a BridgeIn event raised by the PuffinMainnetBridge contract.
type PuffinMainnetBridgeBridgeIn struct {
	User   common.Address
	Asset  common.Address
	Amount *big.Int
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeIn is a free log retrieval operation binding the contract event 0xdb3da595892fff4fc96023c53890339901f50dcce8530e8765febdd1b5809ead.
//
// Solidity: event BridgeIn(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) FilterBridgeIn(opts *bind.FilterOpts, user []common.Address, asset []common.Address, amount []*big.Int) (*PuffinMainnetBridgeBridgeInIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _PuffinMainnetBridge.contract.FilterLogs(opts, "BridgeIn", userRule, assetRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &PuffinMainnetBridgeBridgeInIterator{contract: _PuffinMainnetBridge.contract, event: "BridgeIn", logs: logs, sub: sub}, nil
}

// WatchBridgeIn is a free log subscription operation binding the contract event 0xdb3da595892fff4fc96023c53890339901f50dcce8530e8765febdd1b5809ead.
//
// Solidity: event BridgeIn(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) WatchBridgeIn(opts *bind.WatchOpts, sink chan<- *PuffinMainnetBridgeBridgeIn, user []common.Address, asset []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _PuffinMainnetBridge.contract.WatchLogs(opts, "BridgeIn", userRule, assetRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinMainnetBridgeBridgeIn)
				if err := _PuffinMainnetBridge.contract.UnpackLog(event, "BridgeIn", log); err != nil {
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

// ParseBridgeIn is a log parse operation binding the contract event 0xdb3da595892fff4fc96023c53890339901f50dcce8530e8765febdd1b5809ead.
//
// Solidity: event BridgeIn(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) ParseBridgeIn(log types.Log) (*PuffinMainnetBridgeBridgeIn, error) {
	event := new(PuffinMainnetBridgeBridgeIn)
	if err := _PuffinMainnetBridge.contract.UnpackLog(event, "BridgeIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinMainnetBridgeBridgeOutIterator is returned from FilterBridgeOut and is used to iterate over the raw logs and unpacked data for BridgeOut events raised by the PuffinMainnetBridge contract.
type PuffinMainnetBridgeBridgeOutIterator struct {
	Event *PuffinMainnetBridgeBridgeOut // Event containing the contract specifics and raw log

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
func (it *PuffinMainnetBridgeBridgeOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinMainnetBridgeBridgeOut)
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
		it.Event = new(PuffinMainnetBridgeBridgeOut)
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
func (it *PuffinMainnetBridgeBridgeOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinMainnetBridgeBridgeOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinMainnetBridgeBridgeOut represents a BridgeOut event raised by the PuffinMainnetBridge contract.
type PuffinMainnetBridgeBridgeOut struct {
	User   common.Address
	Asset  common.Address
	Amount *big.Int
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeOut is a free log retrieval operation binding the contract event 0x4905daf08853d889d890e35036f113fe7028be15b84a95ecee546c11f0d7c529.
//
// Solidity: event BridgeOut(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) FilterBridgeOut(opts *bind.FilterOpts, user []common.Address, asset []common.Address, amount []*big.Int) (*PuffinMainnetBridgeBridgeOutIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _PuffinMainnetBridge.contract.FilterLogs(opts, "BridgeOut", userRule, assetRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &PuffinMainnetBridgeBridgeOutIterator{contract: _PuffinMainnetBridge.contract, event: "BridgeOut", logs: logs, sub: sub}, nil
}

// WatchBridgeOut is a free log subscription operation binding the contract event 0x4905daf08853d889d890e35036f113fe7028be15b84a95ecee546c11f0d7c529.
//
// Solidity: event BridgeOut(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) WatchBridgeOut(opts *bind.WatchOpts, sink chan<- *PuffinMainnetBridgeBridgeOut, user []common.Address, asset []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _PuffinMainnetBridge.contract.WatchLogs(opts, "BridgeOut", userRule, assetRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinMainnetBridgeBridgeOut)
				if err := _PuffinMainnetBridge.contract.UnpackLog(event, "BridgeOut", log); err != nil {
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

// ParseBridgeOut is a log parse operation binding the contract event 0x4905daf08853d889d890e35036f113fe7028be15b84a95ecee546c11f0d7c529.
//
// Solidity: event BridgeOut(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) ParseBridgeOut(log types.Log) (*PuffinMainnetBridgeBridgeOut, error) {
	event := new(PuffinMainnetBridgeBridgeOut)
	if err := _PuffinMainnetBridge.contract.UnpackLog(event, "BridgeOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinMainnetBridgeBridgeOutCanceledIterator is returned from FilterBridgeOutCanceled and is used to iterate over the raw logs and unpacked data for BridgeOutCanceled events raised by the PuffinMainnetBridge contract.
type PuffinMainnetBridgeBridgeOutCanceledIterator struct {
	Event *PuffinMainnetBridgeBridgeOutCanceled // Event containing the contract specifics and raw log

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
func (it *PuffinMainnetBridgeBridgeOutCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinMainnetBridgeBridgeOutCanceled)
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
		it.Event = new(PuffinMainnetBridgeBridgeOutCanceled)
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
func (it *PuffinMainnetBridgeBridgeOutCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinMainnetBridgeBridgeOutCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinMainnetBridgeBridgeOutCanceled represents a BridgeOutCanceled event raised by the PuffinMainnetBridge contract.
type PuffinMainnetBridgeBridgeOutCanceled struct {
	User common.Address
	Id   [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBridgeOutCanceled is a free log retrieval operation binding the contract event 0x40278f7011e047d826661d20ccba81e3851b1837ed425b63a9a55167efabeb30.
//
// Solidity: event BridgeOutCanceled(address indexed user, bytes32 indexed id)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) FilterBridgeOutCanceled(opts *bind.FilterOpts, user []common.Address, id [][32]byte) (*PuffinMainnetBridgeBridgeOutCanceledIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PuffinMainnetBridge.contract.FilterLogs(opts, "BridgeOutCanceled", userRule, idRule)
	if err != nil {
		return nil, err
	}
	return &PuffinMainnetBridgeBridgeOutCanceledIterator{contract: _PuffinMainnetBridge.contract, event: "BridgeOutCanceled", logs: logs, sub: sub}, nil
}

// WatchBridgeOutCanceled is a free log subscription operation binding the contract event 0x40278f7011e047d826661d20ccba81e3851b1837ed425b63a9a55167efabeb30.
//
// Solidity: event BridgeOutCanceled(address indexed user, bytes32 indexed id)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) WatchBridgeOutCanceled(opts *bind.WatchOpts, sink chan<- *PuffinMainnetBridgeBridgeOutCanceled, user []common.Address, id [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PuffinMainnetBridge.contract.WatchLogs(opts, "BridgeOutCanceled", userRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinMainnetBridgeBridgeOutCanceled)
				if err := _PuffinMainnetBridge.contract.UnpackLog(event, "BridgeOutCanceled", log); err != nil {
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

// ParseBridgeOutCanceled is a log parse operation binding the contract event 0x40278f7011e047d826661d20ccba81e3851b1837ed425b63a9a55167efabeb30.
//
// Solidity: event BridgeOutCanceled(address indexed user, bytes32 indexed id)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) ParseBridgeOutCanceled(log types.Log) (*PuffinMainnetBridgeBridgeOutCanceled, error) {
	event := new(PuffinMainnetBridgeBridgeOutCanceled)
	if err := _PuffinMainnetBridge.contract.UnpackLog(event, "BridgeOutCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinMainnetBridgeBridgeOutWarmIterator is returned from FilterBridgeOutWarm and is used to iterate over the raw logs and unpacked data for BridgeOutWarm events raised by the PuffinMainnetBridge contract.
type PuffinMainnetBridgeBridgeOutWarmIterator struct {
	Event *PuffinMainnetBridgeBridgeOutWarm // Event containing the contract specifics and raw log

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
func (it *PuffinMainnetBridgeBridgeOutWarmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinMainnetBridgeBridgeOutWarm)
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
		it.Event = new(PuffinMainnetBridgeBridgeOutWarm)
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
func (it *PuffinMainnetBridgeBridgeOutWarmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinMainnetBridgeBridgeOutWarmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinMainnetBridgeBridgeOutWarm represents a BridgeOutWarm event raised by the PuffinMainnetBridge contract.
type PuffinMainnetBridgeBridgeOutWarm struct {
	User   common.Address
	Asset  common.Address
	Amount *big.Int
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeOutWarm is a free log retrieval operation binding the contract event 0xcdf43e020a0b86c3f9312c3a32a8db60cddc764644a74a77294bf38d5a6265bc.
//
// Solidity: event BridgeOutWarm(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) FilterBridgeOutWarm(opts *bind.FilterOpts, user []common.Address, asset []common.Address, amount []*big.Int) (*PuffinMainnetBridgeBridgeOutWarmIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _PuffinMainnetBridge.contract.FilterLogs(opts, "BridgeOutWarm", userRule, assetRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &PuffinMainnetBridgeBridgeOutWarmIterator{contract: _PuffinMainnetBridge.contract, event: "BridgeOutWarm", logs: logs, sub: sub}, nil
}

// WatchBridgeOutWarm is a free log subscription operation binding the contract event 0xcdf43e020a0b86c3f9312c3a32a8db60cddc764644a74a77294bf38d5a6265bc.
//
// Solidity: event BridgeOutWarm(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) WatchBridgeOutWarm(opts *bind.WatchOpts, sink chan<- *PuffinMainnetBridgeBridgeOutWarm, user []common.Address, asset []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _PuffinMainnetBridge.contract.WatchLogs(opts, "BridgeOutWarm", userRule, assetRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinMainnetBridgeBridgeOutWarm)
				if err := _PuffinMainnetBridge.contract.UnpackLog(event, "BridgeOutWarm", log); err != nil {
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

// ParseBridgeOutWarm is a log parse operation binding the contract event 0xcdf43e020a0b86c3f9312c3a32a8db60cddc764644a74a77294bf38d5a6265bc.
//
// Solidity: event BridgeOutWarm(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) ParseBridgeOutWarm(log types.Log) (*PuffinMainnetBridgeBridgeOutWarm, error) {
	event := new(PuffinMainnetBridgeBridgeOutWarm)
	if err := _PuffinMainnetBridge.contract.UnpackLog(event, "BridgeOutWarm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinMainnetBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PuffinMainnetBridge contract.
type PuffinMainnetBridgeOwnershipTransferredIterator struct {
	Event *PuffinMainnetBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PuffinMainnetBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinMainnetBridgeOwnershipTransferred)
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
		it.Event = new(PuffinMainnetBridgeOwnershipTransferred)
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
func (it *PuffinMainnetBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinMainnetBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinMainnetBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the PuffinMainnetBridge contract.
type PuffinMainnetBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PuffinMainnetBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinMainnetBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PuffinMainnetBridgeOwnershipTransferredIterator{contract: _PuffinMainnetBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PuffinMainnetBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinMainnetBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinMainnetBridgeOwnershipTransferred)
				if err := _PuffinMainnetBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*PuffinMainnetBridgeOwnershipTransferred, error) {
	event := new(PuffinMainnetBridgeOwnershipTransferred)
	if err := _PuffinMainnetBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinMainnetBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PuffinMainnetBridge contract.
type PuffinMainnetBridgePausedIterator struct {
	Event *PuffinMainnetBridgePaused // Event containing the contract specifics and raw log

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
func (it *PuffinMainnetBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinMainnetBridgePaused)
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
		it.Event = new(PuffinMainnetBridgePaused)
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
func (it *PuffinMainnetBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinMainnetBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinMainnetBridgePaused represents a Paused event raised by the PuffinMainnetBridge contract.
type PuffinMainnetBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*PuffinMainnetBridgePausedIterator, error) {

	logs, sub, err := _PuffinMainnetBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PuffinMainnetBridgePausedIterator{contract: _PuffinMainnetBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PuffinMainnetBridgePaused) (event.Subscription, error) {

	logs, sub, err := _PuffinMainnetBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinMainnetBridgePaused)
				if err := _PuffinMainnetBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) ParsePaused(log types.Log) (*PuffinMainnetBridgePaused, error) {
	event := new(PuffinMainnetBridgePaused)
	if err := _PuffinMainnetBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinMainnetBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PuffinMainnetBridge contract.
type PuffinMainnetBridgeUnpausedIterator struct {
	Event *PuffinMainnetBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *PuffinMainnetBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinMainnetBridgeUnpaused)
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
		it.Event = new(PuffinMainnetBridgeUnpaused)
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
func (it *PuffinMainnetBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinMainnetBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinMainnetBridgeUnpaused represents a Unpaused event raised by the PuffinMainnetBridge contract.
type PuffinMainnetBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PuffinMainnetBridgeUnpausedIterator, error) {

	logs, sub, err := _PuffinMainnetBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PuffinMainnetBridgeUnpausedIterator{contract: _PuffinMainnetBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PuffinMainnetBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _PuffinMainnetBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinMainnetBridgeUnpaused)
				if err := _PuffinMainnetBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PuffinMainnetBridge *PuffinMainnetBridgeFilterer) ParseUnpaused(log types.Log) (*PuffinMainnetBridgeUnpaused, error) {
	event := new(PuffinMainnetBridgeUnpaused)
	if err := _PuffinMainnetBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
