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

// PuffinBridgeMetaData contains all meta data concerning the PuffinBridge contract.
var PuffinBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"BridgeIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"BridgeOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"BridgeOutCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"BridgeOutWarm\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeIds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"bridgeIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeInComplete\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeOutComplete\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"cancelOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"getRequestInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structPuffinBridge.BridgeRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"markInComplete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"}],\"name\":\"proposeOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"puffinAssets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"puffinKYC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"puffinWarmWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"removeVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requiresWarmWallet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setKYC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setSubnetTokenDeployer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setWarm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subnetTokenDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferWarm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PuffinBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use PuffinBridgeMetaData.ABI instead.
var PuffinBridgeABI = PuffinBridgeMetaData.ABI

// PuffinBridge is an auto generated Go binding around an Ethereum contract.
type PuffinBridge struct {
	PuffinBridgeCaller     // Read-only binding to the contract
	PuffinBridgeTransactor // Write-only binding to the contract
	PuffinBridgeFilterer   // Log filterer for contract events
}

// PuffinBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PuffinBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PuffinBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PuffinBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PuffinBridgeSession struct {
	Contract     *PuffinBridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PuffinBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PuffinBridgeCallerSession struct {
	Contract *PuffinBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PuffinBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PuffinBridgeTransactorSession struct {
	Contract     *PuffinBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PuffinBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PuffinBridgeRaw struct {
	Contract *PuffinBridge // Generic contract binding to access the raw methods on
}

// PuffinBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PuffinBridgeCallerRaw struct {
	Contract *PuffinBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// PuffinBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PuffinBridgeTransactorRaw struct {
	Contract *PuffinBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPuffinBridge creates a new instance of PuffinBridge, bound to a specific deployed contract.
func NewPuffinBridge(address common.Address, backend bind.ContractBackend) (*PuffinBridge, error) {
	contract, err := bindPuffinBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PuffinBridge{PuffinBridgeCaller: PuffinBridgeCaller{contract: contract}, PuffinBridgeTransactor: PuffinBridgeTransactor{contract: contract}, PuffinBridgeFilterer: PuffinBridgeFilterer{contract: contract}}, nil
}

// NewPuffinBridgeCaller creates a new read-only instance of PuffinBridge, bound to a specific deployed contract.
func NewPuffinBridgeCaller(address common.Address, caller bind.ContractCaller) (*PuffinBridgeCaller, error) {
	contract, err := bindPuffinBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinBridgeCaller{contract: contract}, nil
}

// NewPuffinBridgeTransactor creates a new write-only instance of PuffinBridge, bound to a specific deployed contract.
func NewPuffinBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*PuffinBridgeTransactor, error) {
	contract, err := bindPuffinBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinBridgeTransactor{contract: contract}, nil
}

// NewPuffinBridgeFilterer creates a new log filterer instance of PuffinBridge, bound to a specific deployed contract.
func NewPuffinBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*PuffinBridgeFilterer, error) {
	contract, err := bindPuffinBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PuffinBridgeFilterer{contract: contract}, nil
}

// bindPuffinBridge binds a generic wrapper to an already deployed contract.
func bindPuffinBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PuffinBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinBridge *PuffinBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinBridge.Contract.PuffinBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinBridge *PuffinBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinBridge.Contract.PuffinBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinBridge *PuffinBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinBridge.Contract.PuffinBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinBridge *PuffinBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinBridge *PuffinBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinBridge *PuffinBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinBridge.Contract.contract.Transact(opts, method, params...)
}

// BridgeIds is a free data retrieval call binding the contract method 0xad05e8eb.
//
// Solidity: function bridgeIds(bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeCaller) BridgeIds(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "bridgeIds", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeIds is a free data retrieval call binding the contract method 0xad05e8eb.
//
// Solidity: function bridgeIds(bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeSession) BridgeIds(arg0 [32]byte) (bool, error) {
	return _PuffinBridge.Contract.BridgeIds(&_PuffinBridge.CallOpts, arg0)
}

// BridgeIds is a free data retrieval call binding the contract method 0xad05e8eb.
//
// Solidity: function bridgeIds(bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeCallerSession) BridgeIds(arg0 [32]byte) (bool, error) {
	return _PuffinBridge.Contract.BridgeIds(&_PuffinBridge.CallOpts, arg0)
}

// BridgeInComplete is a free data retrieval call binding the contract method 0xdd68a1c9.
//
// Solidity: function bridgeInComplete(bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeCaller) BridgeInComplete(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "bridgeInComplete", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeInComplete is a free data retrieval call binding the contract method 0xdd68a1c9.
//
// Solidity: function bridgeInComplete(bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeSession) BridgeInComplete(arg0 [32]byte) (bool, error) {
	return _PuffinBridge.Contract.BridgeInComplete(&_PuffinBridge.CallOpts, arg0)
}

// BridgeInComplete is a free data retrieval call binding the contract method 0xdd68a1c9.
//
// Solidity: function bridgeInComplete(bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeCallerSession) BridgeInComplete(arg0 [32]byte) (bool, error) {
	return _PuffinBridge.Contract.BridgeInComplete(&_PuffinBridge.CallOpts, arg0)
}

// BridgeOutComplete is a free data retrieval call binding the contract method 0x4f9850ce.
//
// Solidity: function bridgeOutComplete(bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeCaller) BridgeOutComplete(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "bridgeOutComplete", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeOutComplete is a free data retrieval call binding the contract method 0x4f9850ce.
//
// Solidity: function bridgeOutComplete(bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeSession) BridgeOutComplete(arg0 [32]byte) (bool, error) {
	return _PuffinBridge.Contract.BridgeOutComplete(&_PuffinBridge.CallOpts, arg0)
}

// BridgeOutComplete is a free data retrieval call binding the contract method 0x4f9850ce.
//
// Solidity: function bridgeOutComplete(bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeCallerSession) BridgeOutComplete(arg0 [32]byte) (bool, error) {
	return _PuffinBridge.Contract.BridgeOutComplete(&_PuffinBridge.CallOpts, arg0)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_PuffinBridge *PuffinBridgeCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_PuffinBridge *PuffinBridgeSession) ChainId() (*big.Int, error) {
	return _PuffinBridge.Contract.ChainId(&_PuffinBridge.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_PuffinBridge *PuffinBridgeCallerSession) ChainId() (*big.Int, error) {
	return _PuffinBridge.Contract.ChainId(&_PuffinBridge.CallOpts)
}

// GetRequestInfo is a free data retrieval call binding the contract method 0x82d67e5a.
//
// Solidity: function getRequestInfo(bytes32 requestId) view returns((bytes32,address,address,uint256,uint256))
func (_PuffinBridge *PuffinBridgeCaller) GetRequestInfo(opts *bind.CallOpts, requestId [32]byte) (PuffinBridgeBridgeRequest, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "getRequestInfo", requestId)

	if err != nil {
		return *new(PuffinBridgeBridgeRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(PuffinBridgeBridgeRequest)).(*PuffinBridgeBridgeRequest)

	return out0, err

}

// GetRequestInfo is a free data retrieval call binding the contract method 0x82d67e5a.
//
// Solidity: function getRequestInfo(bytes32 requestId) view returns((bytes32,address,address,uint256,uint256))
func (_PuffinBridge *PuffinBridgeSession) GetRequestInfo(requestId [32]byte) (PuffinBridgeBridgeRequest, error) {
	return _PuffinBridge.Contract.GetRequestInfo(&_PuffinBridge.CallOpts, requestId)
}

// GetRequestInfo is a free data retrieval call binding the contract method 0x82d67e5a.
//
// Solidity: function getRequestInfo(bytes32 requestId) view returns((bytes32,address,address,uint256,uint256))
func (_PuffinBridge *PuffinBridgeCallerSession) GetRequestInfo(requestId [32]byte) (PuffinBridgeBridgeRequest, error) {
	return _PuffinBridge.Contract.GetRequestInfo(&_PuffinBridge.CallOpts, requestId)
}

// HasVoted is a free data retrieval call binding the contract method 0x4a1dc5cc.
//
// Solidity: function hasVoted(address , bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeCaller) HasVoted(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "hasVoted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x4a1dc5cc.
//
// Solidity: function hasVoted(address , bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeSession) HasVoted(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _PuffinBridge.Contract.HasVoted(&_PuffinBridge.CallOpts, arg0, arg1)
}

// HasVoted is a free data retrieval call binding the contract method 0x4a1dc5cc.
//
// Solidity: function hasVoted(address , bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeCallerSession) HasVoted(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _PuffinBridge.Contract.HasVoted(&_PuffinBridge.CallOpts, arg0, arg1)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address ) view returns(bool)
func (_PuffinBridge *PuffinBridgeCaller) IsVoter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "isVoter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address ) view returns(bool)
func (_PuffinBridge *PuffinBridgeSession) IsVoter(arg0 common.Address) (bool, error) {
	return _PuffinBridge.Contract.IsVoter(&_PuffinBridge.CallOpts, arg0)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(address ) view returns(bool)
func (_PuffinBridge *PuffinBridgeCallerSession) IsVoter(arg0 common.Address) (bool, error) {
	return _PuffinBridge.Contract.IsVoter(&_PuffinBridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinBridge *PuffinBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinBridge *PuffinBridgeSession) Owner() (common.Address, error) {
	return _PuffinBridge.Contract.Owner(&_PuffinBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinBridge *PuffinBridgeCallerSession) Owner() (common.Address, error) {
	return _PuffinBridge.Contract.Owner(&_PuffinBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PuffinBridge *PuffinBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PuffinBridge *PuffinBridgeSession) Paused() (bool, error) {
	return _PuffinBridge.Contract.Paused(&_PuffinBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PuffinBridge *PuffinBridgeCallerSession) Paused() (bool, error) {
	return _PuffinBridge.Contract.Paused(&_PuffinBridge.CallOpts)
}

// PuffinAssets is a free data retrieval call binding the contract method 0x318f2e3e.
//
// Solidity: function puffinAssets() view returns(address)
func (_PuffinBridge *PuffinBridgeCaller) PuffinAssets(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "puffinAssets")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PuffinAssets is a free data retrieval call binding the contract method 0x318f2e3e.
//
// Solidity: function puffinAssets() view returns(address)
func (_PuffinBridge *PuffinBridgeSession) PuffinAssets() (common.Address, error) {
	return _PuffinBridge.Contract.PuffinAssets(&_PuffinBridge.CallOpts)
}

// PuffinAssets is a free data retrieval call binding the contract method 0x318f2e3e.
//
// Solidity: function puffinAssets() view returns(address)
func (_PuffinBridge *PuffinBridgeCallerSession) PuffinAssets() (common.Address, error) {
	return _PuffinBridge.Contract.PuffinAssets(&_PuffinBridge.CallOpts)
}

// PuffinKYC is a free data retrieval call binding the contract method 0x717642d9.
//
// Solidity: function puffinKYC() view returns(address)
func (_PuffinBridge *PuffinBridgeCaller) PuffinKYC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "puffinKYC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PuffinKYC is a free data retrieval call binding the contract method 0x717642d9.
//
// Solidity: function puffinKYC() view returns(address)
func (_PuffinBridge *PuffinBridgeSession) PuffinKYC() (common.Address, error) {
	return _PuffinBridge.Contract.PuffinKYC(&_PuffinBridge.CallOpts)
}

// PuffinKYC is a free data retrieval call binding the contract method 0x717642d9.
//
// Solidity: function puffinKYC() view returns(address)
func (_PuffinBridge *PuffinBridgeCallerSession) PuffinKYC() (common.Address, error) {
	return _PuffinBridge.Contract.PuffinKYC(&_PuffinBridge.CallOpts)
}

// PuffinWarmWallet is a free data retrieval call binding the contract method 0xd5a0ccf8.
//
// Solidity: function puffinWarmWallet() view returns(address)
func (_PuffinBridge *PuffinBridgeCaller) PuffinWarmWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "puffinWarmWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PuffinWarmWallet is a free data retrieval call binding the contract method 0xd5a0ccf8.
//
// Solidity: function puffinWarmWallet() view returns(address)
func (_PuffinBridge *PuffinBridgeSession) PuffinWarmWallet() (common.Address, error) {
	return _PuffinBridge.Contract.PuffinWarmWallet(&_PuffinBridge.CallOpts)
}

// PuffinWarmWallet is a free data retrieval call binding the contract method 0xd5a0ccf8.
//
// Solidity: function puffinWarmWallet() view returns(address)
func (_PuffinBridge *PuffinBridgeCallerSession) PuffinWarmWallet() (common.Address, error) {
	return _PuffinBridge.Contract.PuffinWarmWallet(&_PuffinBridge.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x0526d3fe.
//
// Solidity: function requestCount(bytes32 ) view returns(uint256)
func (_PuffinBridge *PuffinBridgeCaller) RequestCount(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "requestCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestCount is a free data retrieval call binding the contract method 0x0526d3fe.
//
// Solidity: function requestCount(bytes32 ) view returns(uint256)
func (_PuffinBridge *PuffinBridgeSession) RequestCount(arg0 [32]byte) (*big.Int, error) {
	return _PuffinBridge.Contract.RequestCount(&_PuffinBridge.CallOpts, arg0)
}

// RequestCount is a free data retrieval call binding the contract method 0x0526d3fe.
//
// Solidity: function requestCount(bytes32 ) view returns(uint256)
func (_PuffinBridge *PuffinBridgeCallerSession) RequestCount(arg0 [32]byte) (*big.Int, error) {
	return _PuffinBridge.Contract.RequestCount(&_PuffinBridge.CallOpts, arg0)
}

// RequestInfo is a free data retrieval call binding the contract method 0x361e4d28.
//
// Solidity: function requestInfo(bytes32 ) view returns(bytes32 id, address user, address asset, uint256 amount, uint256 expiry)
func (_PuffinBridge *PuffinBridgeCaller) RequestInfo(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Id     [32]byte
	User   common.Address
	Asset  common.Address
	Amount *big.Int
	Expiry *big.Int
}, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "requestInfo", arg0)

	outstruct := new(struct {
		Id     [32]byte
		User   common.Address
		Asset  common.Address
		Amount *big.Int
		Expiry *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.User = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Asset = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Expiry = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RequestInfo is a free data retrieval call binding the contract method 0x361e4d28.
//
// Solidity: function requestInfo(bytes32 ) view returns(bytes32 id, address user, address asset, uint256 amount, uint256 expiry)
func (_PuffinBridge *PuffinBridgeSession) RequestInfo(arg0 [32]byte) (struct {
	Id     [32]byte
	User   common.Address
	Asset  common.Address
	Amount *big.Int
	Expiry *big.Int
}, error) {
	return _PuffinBridge.Contract.RequestInfo(&_PuffinBridge.CallOpts, arg0)
}

// RequestInfo is a free data retrieval call binding the contract method 0x361e4d28.
//
// Solidity: function requestInfo(bytes32 ) view returns(bytes32 id, address user, address asset, uint256 amount, uint256 expiry)
func (_PuffinBridge *PuffinBridgeCallerSession) RequestInfo(arg0 [32]byte) (struct {
	Id     [32]byte
	User   common.Address
	Asset  common.Address
	Amount *big.Int
	Expiry *big.Int
}, error) {
	return _PuffinBridge.Contract.RequestInfo(&_PuffinBridge.CallOpts, arg0)
}

// RequiresWarmWallet is a free data retrieval call binding the contract method 0x576fd9a8.
//
// Solidity: function requiresWarmWallet(bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeCaller) RequiresWarmWallet(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "requiresWarmWallet", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequiresWarmWallet is a free data retrieval call binding the contract method 0x576fd9a8.
//
// Solidity: function requiresWarmWallet(bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeSession) RequiresWarmWallet(arg0 [32]byte) (bool, error) {
	return _PuffinBridge.Contract.RequiresWarmWallet(&_PuffinBridge.CallOpts, arg0)
}

// RequiresWarmWallet is a free data retrieval call binding the contract method 0x576fd9a8.
//
// Solidity: function requiresWarmWallet(bytes32 ) view returns(bool)
func (_PuffinBridge *PuffinBridgeCallerSession) RequiresWarmWallet(arg0 [32]byte) (bool, error) {
	return _PuffinBridge.Contract.RequiresWarmWallet(&_PuffinBridge.CallOpts, arg0)
}

// SubnetTokenDeployer is a free data retrieval call binding the contract method 0x41c4ed79.
//
// Solidity: function subnetTokenDeployer() view returns(address)
func (_PuffinBridge *PuffinBridgeCaller) SubnetTokenDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "subnetTokenDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SubnetTokenDeployer is a free data retrieval call binding the contract method 0x41c4ed79.
//
// Solidity: function subnetTokenDeployer() view returns(address)
func (_PuffinBridge *PuffinBridgeSession) SubnetTokenDeployer() (common.Address, error) {
	return _PuffinBridge.Contract.SubnetTokenDeployer(&_PuffinBridge.CallOpts)
}

// SubnetTokenDeployer is a free data retrieval call binding the contract method 0x41c4ed79.
//
// Solidity: function subnetTokenDeployer() view returns(address)
func (_PuffinBridge *PuffinBridgeCallerSession) SubnetTokenDeployer() (common.Address, error) {
	return _PuffinBridge.Contract.SubnetTokenDeployer(&_PuffinBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_PuffinBridge *PuffinBridgeCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PuffinBridge.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_PuffinBridge *PuffinBridgeSession) Threshold() (*big.Int, error) {
	return _PuffinBridge.Contract.Threshold(&_PuffinBridge.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_PuffinBridge *PuffinBridgeCallerSession) Threshold() (*big.Int, error) {
	return _PuffinBridge.Contract.Threshold(&_PuffinBridge.CallOpts)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(address user) returns()
func (_PuffinBridge *PuffinBridgeTransactor) AddVoter(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "addVoter", user)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(address user) returns()
func (_PuffinBridge *PuffinBridgeSession) AddVoter(user common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.AddVoter(&_PuffinBridge.TransactOpts, user)
}

// AddVoter is a paid mutator transaction binding the contract method 0xf4ab9adf.
//
// Solidity: function addVoter(address user) returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) AddVoter(user common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.AddVoter(&_PuffinBridge.TransactOpts, user)
}

// BridgeIn is a paid mutator transaction binding the contract method 0x7107fa82.
//
// Solidity: function bridgeIn(uint256 amount, address asset) returns()
func (_PuffinBridge *PuffinBridgeTransactor) BridgeIn(opts *bind.TransactOpts, amount *big.Int, asset common.Address) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "bridgeIn", amount, asset)
}

// BridgeIn is a paid mutator transaction binding the contract method 0x7107fa82.
//
// Solidity: function bridgeIn(uint256 amount, address asset) returns()
func (_PuffinBridge *PuffinBridgeSession) BridgeIn(amount *big.Int, asset common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.BridgeIn(&_PuffinBridge.TransactOpts, amount, asset)
}

// BridgeIn is a paid mutator transaction binding the contract method 0x7107fa82.
//
// Solidity: function bridgeIn(uint256 amount, address asset) returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) BridgeIn(amount *big.Int, asset common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.BridgeIn(&_PuffinBridge.TransactOpts, amount, asset)
}

// CancelOut is a paid mutator transaction binding the contract method 0x1533c400.
//
// Solidity: function cancelOut(bytes32 requestId) returns()
func (_PuffinBridge *PuffinBridgeTransactor) CancelOut(opts *bind.TransactOpts, requestId [32]byte) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "cancelOut", requestId)
}

// CancelOut is a paid mutator transaction binding the contract method 0x1533c400.
//
// Solidity: function cancelOut(bytes32 requestId) returns()
func (_PuffinBridge *PuffinBridgeSession) CancelOut(requestId [32]byte) (*types.Transaction, error) {
	return _PuffinBridge.Contract.CancelOut(&_PuffinBridge.TransactOpts, requestId)
}

// CancelOut is a paid mutator transaction binding the contract method 0x1533c400.
//
// Solidity: function cancelOut(bytes32 requestId) returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) CancelOut(requestId [32]byte) (*types.Transaction, error) {
	return _PuffinBridge.Contract.CancelOut(&_PuffinBridge.TransactOpts, requestId)
}

// MarkInComplete is a paid mutator transaction binding the contract method 0x49049bb1.
//
// Solidity: function markInComplete(bytes32 requestId) returns()
func (_PuffinBridge *PuffinBridgeTransactor) MarkInComplete(opts *bind.TransactOpts, requestId [32]byte) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "markInComplete", requestId)
}

// MarkInComplete is a paid mutator transaction binding the contract method 0x49049bb1.
//
// Solidity: function markInComplete(bytes32 requestId) returns()
func (_PuffinBridge *PuffinBridgeSession) MarkInComplete(requestId [32]byte) (*types.Transaction, error) {
	return _PuffinBridge.Contract.MarkInComplete(&_PuffinBridge.TransactOpts, requestId)
}

// MarkInComplete is a paid mutator transaction binding the contract method 0x49049bb1.
//
// Solidity: function markInComplete(bytes32 requestId) returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) MarkInComplete(requestId [32]byte) (*types.Transaction, error) {
	return _PuffinBridge.Contract.MarkInComplete(&_PuffinBridge.TransactOpts, requestId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PuffinBridge *PuffinBridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PuffinBridge *PuffinBridgeSession) Pause() (*types.Transaction, error) {
	return _PuffinBridge.Contract.Pause(&_PuffinBridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _PuffinBridge.Contract.Pause(&_PuffinBridge.TransactOpts)
}

// ProposeOut is a paid mutator transaction binding the contract method 0x53a3a09c.
//
// Solidity: function proposeOut(address asset, address user, uint256 amount, bytes32 requestId, uint256 fromChainId) returns()
func (_PuffinBridge *PuffinBridgeTransactor) ProposeOut(opts *bind.TransactOpts, asset common.Address, user common.Address, amount *big.Int, requestId [32]byte, fromChainId *big.Int) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "proposeOut", asset, user, amount, requestId, fromChainId)
}

// ProposeOut is a paid mutator transaction binding the contract method 0x53a3a09c.
//
// Solidity: function proposeOut(address asset, address user, uint256 amount, bytes32 requestId, uint256 fromChainId) returns()
func (_PuffinBridge *PuffinBridgeSession) ProposeOut(asset common.Address, user common.Address, amount *big.Int, requestId [32]byte, fromChainId *big.Int) (*types.Transaction, error) {
	return _PuffinBridge.Contract.ProposeOut(&_PuffinBridge.TransactOpts, asset, user, amount, requestId, fromChainId)
}

// ProposeOut is a paid mutator transaction binding the contract method 0x53a3a09c.
//
// Solidity: function proposeOut(address asset, address user, uint256 amount, bytes32 requestId, uint256 fromChainId) returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) ProposeOut(asset common.Address, user common.Address, amount *big.Int, requestId [32]byte, fromChainId *big.Int) (*types.Transaction, error) {
	return _PuffinBridge.Contract.ProposeOut(&_PuffinBridge.TransactOpts, asset, user, amount, requestId, fromChainId)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(address user) returns()
func (_PuffinBridge *PuffinBridgeTransactor) RemoveVoter(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "removeVoter", user)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(address user) returns()
func (_PuffinBridge *PuffinBridgeSession) RemoveVoter(user common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.RemoveVoter(&_PuffinBridge.TransactOpts, user)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0x86c1ff68.
//
// Solidity: function removeVoter(address user) returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) RemoveVoter(user common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.RemoveVoter(&_PuffinBridge.TransactOpts, user)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinBridge *PuffinBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinBridge *PuffinBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinBridge.Contract.RenounceOwnership(&_PuffinBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinBridge.Contract.RenounceOwnership(&_PuffinBridge.TransactOpts)
}

// SetAssets is a paid mutator transaction binding the contract method 0xee1559da.
//
// Solidity: function setAssets(address _contract) returns()
func (_PuffinBridge *PuffinBridgeTransactor) SetAssets(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "setAssets", _contract)
}

// SetAssets is a paid mutator transaction binding the contract method 0xee1559da.
//
// Solidity: function setAssets(address _contract) returns()
func (_PuffinBridge *PuffinBridgeSession) SetAssets(_contract common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.SetAssets(&_PuffinBridge.TransactOpts, _contract)
}

// SetAssets is a paid mutator transaction binding the contract method 0xee1559da.
//
// Solidity: function setAssets(address _contract) returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) SetAssets(_contract common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.SetAssets(&_PuffinBridge.TransactOpts, _contract)
}

// SetKYC is a paid mutator transaction binding the contract method 0x483a83df.
//
// Solidity: function setKYC(address _contract) returns()
func (_PuffinBridge *PuffinBridgeTransactor) SetKYC(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "setKYC", _contract)
}

// SetKYC is a paid mutator transaction binding the contract method 0x483a83df.
//
// Solidity: function setKYC(address _contract) returns()
func (_PuffinBridge *PuffinBridgeSession) SetKYC(_contract common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.SetKYC(&_PuffinBridge.TransactOpts, _contract)
}

// SetKYC is a paid mutator transaction binding the contract method 0x483a83df.
//
// Solidity: function setKYC(address _contract) returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) SetKYC(_contract common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.SetKYC(&_PuffinBridge.TransactOpts, _contract)
}

// SetSubnetTokenDeployer is a paid mutator transaction binding the contract method 0x1b399680.
//
// Solidity: function setSubnetTokenDeployer(address _contract) returns()
func (_PuffinBridge *PuffinBridgeTransactor) SetSubnetTokenDeployer(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "setSubnetTokenDeployer", _contract)
}

// SetSubnetTokenDeployer is a paid mutator transaction binding the contract method 0x1b399680.
//
// Solidity: function setSubnetTokenDeployer(address _contract) returns()
func (_PuffinBridge *PuffinBridgeSession) SetSubnetTokenDeployer(_contract common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.SetSubnetTokenDeployer(&_PuffinBridge.TransactOpts, _contract)
}

// SetSubnetTokenDeployer is a paid mutator transaction binding the contract method 0x1b399680.
//
// Solidity: function setSubnetTokenDeployer(address _contract) returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) SetSubnetTokenDeployer(_contract common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.SetSubnetTokenDeployer(&_PuffinBridge.TransactOpts, _contract)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 _threshold) returns()
func (_PuffinBridge *PuffinBridgeTransactor) SetThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "setThreshold", _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 _threshold) returns()
func (_PuffinBridge *PuffinBridgeSession) SetThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _PuffinBridge.Contract.SetThreshold(&_PuffinBridge.TransactOpts, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 _threshold) returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) SetThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _PuffinBridge.Contract.SetThreshold(&_PuffinBridge.TransactOpts, _threshold)
}

// SetWarm is a paid mutator transaction binding the contract method 0x7eea8c34.
//
// Solidity: function setWarm(address _contract) returns()
func (_PuffinBridge *PuffinBridgeTransactor) SetWarm(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "setWarm", _contract)
}

// SetWarm is a paid mutator transaction binding the contract method 0x7eea8c34.
//
// Solidity: function setWarm(address _contract) returns()
func (_PuffinBridge *PuffinBridgeSession) SetWarm(_contract common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.SetWarm(&_PuffinBridge.TransactOpts, _contract)
}

// SetWarm is a paid mutator transaction binding the contract method 0x7eea8c34.
//
// Solidity: function setWarm(address _contract) returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) SetWarm(_contract common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.SetWarm(&_PuffinBridge.TransactOpts, _contract)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinBridge *PuffinBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinBridge *PuffinBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.TransferOwnership(&_PuffinBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinBridge.Contract.TransferOwnership(&_PuffinBridge.TransactOpts, newOwner)
}

// TransferWarm is a paid mutator transaction binding the contract method 0x817747c1.
//
// Solidity: function transferWarm(address asset, uint256 amount) returns()
func (_PuffinBridge *PuffinBridgeTransactor) TransferWarm(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "transferWarm", asset, amount)
}

// TransferWarm is a paid mutator transaction binding the contract method 0x817747c1.
//
// Solidity: function transferWarm(address asset, uint256 amount) returns()
func (_PuffinBridge *PuffinBridgeSession) TransferWarm(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PuffinBridge.Contract.TransferWarm(&_PuffinBridge.TransactOpts, asset, amount)
}

// TransferWarm is a paid mutator transaction binding the contract method 0x817747c1.
//
// Solidity: function transferWarm(address asset, uint256 amount) returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) TransferWarm(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PuffinBridge.Contract.TransferWarm(&_PuffinBridge.TransactOpts, asset, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PuffinBridge *PuffinBridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinBridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PuffinBridge *PuffinBridgeSession) Unpause() (*types.Transaction, error) {
	return _PuffinBridge.Contract.Unpause(&_PuffinBridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PuffinBridge *PuffinBridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _PuffinBridge.Contract.Unpause(&_PuffinBridge.TransactOpts)
}

// PuffinBridgeBridgeInIterator is returned from FilterBridgeIn and is used to iterate over the raw logs and unpacked data for BridgeIn events raised by the PuffinBridge contract.
type PuffinBridgeBridgeInIterator struct {
	Event *PuffinBridgeBridgeIn // Event containing the contract specifics and raw log

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
func (it *PuffinBridgeBridgeInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinBridgeBridgeIn)
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
		it.Event = new(PuffinBridgeBridgeIn)
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
func (it *PuffinBridgeBridgeInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinBridgeBridgeInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinBridgeBridgeIn represents a BridgeIn event raised by the PuffinBridge contract.
type PuffinBridgeBridgeIn struct {
	User    common.Address
	Asset   common.Address
	Amount  *big.Int
	Id      [32]byte
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgeIn is a free log retrieval operation binding the contract event 0x59606946ecc613ff002f1195416fae74a6f48dec462cb8f858b5f733a5872798.
//
// Solidity: event BridgeIn(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id, uint256 chainId)
func (_PuffinBridge *PuffinBridgeFilterer) FilterBridgeIn(opts *bind.FilterOpts, user []common.Address, asset []common.Address, amount []*big.Int) (*PuffinBridgeBridgeInIterator, error) {

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

	logs, sub, err := _PuffinBridge.contract.FilterLogs(opts, "BridgeIn", userRule, assetRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &PuffinBridgeBridgeInIterator{contract: _PuffinBridge.contract, event: "BridgeIn", logs: logs, sub: sub}, nil
}

// WatchBridgeIn is a free log subscription operation binding the contract event 0x59606946ecc613ff002f1195416fae74a6f48dec462cb8f858b5f733a5872798.
//
// Solidity: event BridgeIn(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id, uint256 chainId)
func (_PuffinBridge *PuffinBridgeFilterer) WatchBridgeIn(opts *bind.WatchOpts, sink chan<- *PuffinBridgeBridgeIn, user []common.Address, asset []common.Address, amount []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _PuffinBridge.contract.WatchLogs(opts, "BridgeIn", userRule, assetRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinBridgeBridgeIn)
				if err := _PuffinBridge.contract.UnpackLog(event, "BridgeIn", log); err != nil {
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

// ParseBridgeIn is a log parse operation binding the contract event 0x59606946ecc613ff002f1195416fae74a6f48dec462cb8f858b5f733a5872798.
//
// Solidity: event BridgeIn(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id, uint256 chainId)
func (_PuffinBridge *PuffinBridgeFilterer) ParseBridgeIn(log types.Log) (*PuffinBridgeBridgeIn, error) {
	event := new(PuffinBridgeBridgeIn)
	if err := _PuffinBridge.contract.UnpackLog(event, "BridgeIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinBridgeBridgeOutIterator is returned from FilterBridgeOut and is used to iterate over the raw logs and unpacked data for BridgeOut events raised by the PuffinBridge contract.
type PuffinBridgeBridgeOutIterator struct {
	Event *PuffinBridgeBridgeOut // Event containing the contract specifics and raw log

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
func (it *PuffinBridgeBridgeOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinBridgeBridgeOut)
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
		it.Event = new(PuffinBridgeBridgeOut)
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
func (it *PuffinBridgeBridgeOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinBridgeBridgeOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinBridgeBridgeOut represents a BridgeOut event raised by the PuffinBridge contract.
type PuffinBridgeBridgeOut struct {
	User    common.Address
	Asset   common.Address
	Amount  *big.Int
	Id      [32]byte
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgeOut is a free log retrieval operation binding the contract event 0x202d63c32c205a84ab4059059b8c74e3ba1d071bac7bb773b569103fc6285330.
//
// Solidity: event BridgeOut(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id, uint256 chainId)
func (_PuffinBridge *PuffinBridgeFilterer) FilterBridgeOut(opts *bind.FilterOpts, user []common.Address, asset []common.Address, amount []*big.Int) (*PuffinBridgeBridgeOutIterator, error) {

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

	logs, sub, err := _PuffinBridge.contract.FilterLogs(opts, "BridgeOut", userRule, assetRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &PuffinBridgeBridgeOutIterator{contract: _PuffinBridge.contract, event: "BridgeOut", logs: logs, sub: sub}, nil
}

// WatchBridgeOut is a free log subscription operation binding the contract event 0x202d63c32c205a84ab4059059b8c74e3ba1d071bac7bb773b569103fc6285330.
//
// Solidity: event BridgeOut(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id, uint256 chainId)
func (_PuffinBridge *PuffinBridgeFilterer) WatchBridgeOut(opts *bind.WatchOpts, sink chan<- *PuffinBridgeBridgeOut, user []common.Address, asset []common.Address, amount []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _PuffinBridge.contract.WatchLogs(opts, "BridgeOut", userRule, assetRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinBridgeBridgeOut)
				if err := _PuffinBridge.contract.UnpackLog(event, "BridgeOut", log); err != nil {
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

// ParseBridgeOut is a log parse operation binding the contract event 0x202d63c32c205a84ab4059059b8c74e3ba1d071bac7bb773b569103fc6285330.
//
// Solidity: event BridgeOut(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id, uint256 chainId)
func (_PuffinBridge *PuffinBridgeFilterer) ParseBridgeOut(log types.Log) (*PuffinBridgeBridgeOut, error) {
	event := new(PuffinBridgeBridgeOut)
	if err := _PuffinBridge.contract.UnpackLog(event, "BridgeOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinBridgeBridgeOutCanceledIterator is returned from FilterBridgeOutCanceled and is used to iterate over the raw logs and unpacked data for BridgeOutCanceled events raised by the PuffinBridge contract.
type PuffinBridgeBridgeOutCanceledIterator struct {
	Event *PuffinBridgeBridgeOutCanceled // Event containing the contract specifics and raw log

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
func (it *PuffinBridgeBridgeOutCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinBridgeBridgeOutCanceled)
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
		it.Event = new(PuffinBridgeBridgeOutCanceled)
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
func (it *PuffinBridgeBridgeOutCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinBridgeBridgeOutCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinBridgeBridgeOutCanceled represents a BridgeOutCanceled event raised by the PuffinBridge contract.
type PuffinBridgeBridgeOutCanceled struct {
	User common.Address
	Id   [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBridgeOutCanceled is a free log retrieval operation binding the contract event 0x40278f7011e047d826661d20ccba81e3851b1837ed425b63a9a55167efabeb30.
//
// Solidity: event BridgeOutCanceled(address indexed user, bytes32 indexed id)
func (_PuffinBridge *PuffinBridgeFilterer) FilterBridgeOutCanceled(opts *bind.FilterOpts, user []common.Address, id [][32]byte) (*PuffinBridgeBridgeOutCanceledIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PuffinBridge.contract.FilterLogs(opts, "BridgeOutCanceled", userRule, idRule)
	if err != nil {
		return nil, err
	}
	return &PuffinBridgeBridgeOutCanceledIterator{contract: _PuffinBridge.contract, event: "BridgeOutCanceled", logs: logs, sub: sub}, nil
}

// WatchBridgeOutCanceled is a free log subscription operation binding the contract event 0x40278f7011e047d826661d20ccba81e3851b1837ed425b63a9a55167efabeb30.
//
// Solidity: event BridgeOutCanceled(address indexed user, bytes32 indexed id)
func (_PuffinBridge *PuffinBridgeFilterer) WatchBridgeOutCanceled(opts *bind.WatchOpts, sink chan<- *PuffinBridgeBridgeOutCanceled, user []common.Address, id [][32]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PuffinBridge.contract.WatchLogs(opts, "BridgeOutCanceled", userRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinBridgeBridgeOutCanceled)
				if err := _PuffinBridge.contract.UnpackLog(event, "BridgeOutCanceled", log); err != nil {
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
func (_PuffinBridge *PuffinBridgeFilterer) ParseBridgeOutCanceled(log types.Log) (*PuffinBridgeBridgeOutCanceled, error) {
	event := new(PuffinBridgeBridgeOutCanceled)
	if err := _PuffinBridge.contract.UnpackLog(event, "BridgeOutCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinBridgeBridgeOutWarmIterator is returned from FilterBridgeOutWarm and is used to iterate over the raw logs and unpacked data for BridgeOutWarm events raised by the PuffinBridge contract.
type PuffinBridgeBridgeOutWarmIterator struct {
	Event *PuffinBridgeBridgeOutWarm // Event containing the contract specifics and raw log

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
func (it *PuffinBridgeBridgeOutWarmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinBridgeBridgeOutWarm)
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
		it.Event = new(PuffinBridgeBridgeOutWarm)
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
func (it *PuffinBridgeBridgeOutWarmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinBridgeBridgeOutWarmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinBridgeBridgeOutWarm represents a BridgeOutWarm event raised by the PuffinBridge contract.
type PuffinBridgeBridgeOutWarm struct {
	User    common.Address
	Asset   common.Address
	Amount  *big.Int
	Id      [32]byte
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgeOutWarm is a free log retrieval operation binding the contract event 0x2187160e892170ad60c67b7a14abf71140572bc939953835d5712b88566b8136.
//
// Solidity: event BridgeOutWarm(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id, uint256 chainId)
func (_PuffinBridge *PuffinBridgeFilterer) FilterBridgeOutWarm(opts *bind.FilterOpts, user []common.Address, asset []common.Address, amount []*big.Int) (*PuffinBridgeBridgeOutWarmIterator, error) {

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

	logs, sub, err := _PuffinBridge.contract.FilterLogs(opts, "BridgeOutWarm", userRule, assetRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &PuffinBridgeBridgeOutWarmIterator{contract: _PuffinBridge.contract, event: "BridgeOutWarm", logs: logs, sub: sub}, nil
}

// WatchBridgeOutWarm is a free log subscription operation binding the contract event 0x2187160e892170ad60c67b7a14abf71140572bc939953835d5712b88566b8136.
//
// Solidity: event BridgeOutWarm(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id, uint256 chainId)
func (_PuffinBridge *PuffinBridgeFilterer) WatchBridgeOutWarm(opts *bind.WatchOpts, sink chan<- *PuffinBridgeBridgeOutWarm, user []common.Address, asset []common.Address, amount []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _PuffinBridge.contract.WatchLogs(opts, "BridgeOutWarm", userRule, assetRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinBridgeBridgeOutWarm)
				if err := _PuffinBridge.contract.UnpackLog(event, "BridgeOutWarm", log); err != nil {
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

// ParseBridgeOutWarm is a log parse operation binding the contract event 0x2187160e892170ad60c67b7a14abf71140572bc939953835d5712b88566b8136.
//
// Solidity: event BridgeOutWarm(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id, uint256 chainId)
func (_PuffinBridge *PuffinBridgeFilterer) ParseBridgeOutWarm(log types.Log) (*PuffinBridgeBridgeOutWarm, error) {
	event := new(PuffinBridgeBridgeOutWarm)
	if err := _PuffinBridge.contract.UnpackLog(event, "BridgeOutWarm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PuffinBridge contract.
type PuffinBridgeOwnershipTransferredIterator struct {
	Event *PuffinBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PuffinBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinBridgeOwnershipTransferred)
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
		it.Event = new(PuffinBridgeOwnershipTransferred)
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
func (it *PuffinBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the PuffinBridge contract.
type PuffinBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinBridge *PuffinBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PuffinBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PuffinBridgeOwnershipTransferredIterator{contract: _PuffinBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinBridge *PuffinBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PuffinBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinBridgeOwnershipTransferred)
				if err := _PuffinBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PuffinBridge *PuffinBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*PuffinBridgeOwnershipTransferred, error) {
	event := new(PuffinBridgeOwnershipTransferred)
	if err := _PuffinBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PuffinBridge contract.
type PuffinBridgePausedIterator struct {
	Event *PuffinBridgePaused // Event containing the contract specifics and raw log

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
func (it *PuffinBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinBridgePaused)
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
		it.Event = new(PuffinBridgePaused)
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
func (it *PuffinBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinBridgePaused represents a Paused event raised by the PuffinBridge contract.
type PuffinBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PuffinBridge *PuffinBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*PuffinBridgePausedIterator, error) {

	logs, sub, err := _PuffinBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PuffinBridgePausedIterator{contract: _PuffinBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PuffinBridge *PuffinBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PuffinBridgePaused) (event.Subscription, error) {

	logs, sub, err := _PuffinBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinBridgePaused)
				if err := _PuffinBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PuffinBridge *PuffinBridgeFilterer) ParsePaused(log types.Log) (*PuffinBridgePaused, error) {
	event := new(PuffinBridgePaused)
	if err := _PuffinBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PuffinBridge contract.
type PuffinBridgeUnpausedIterator struct {
	Event *PuffinBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *PuffinBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinBridgeUnpaused)
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
		it.Event = new(PuffinBridgeUnpaused)
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
func (it *PuffinBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinBridgeUnpaused represents a Unpaused event raised by the PuffinBridge contract.
type PuffinBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PuffinBridge *PuffinBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PuffinBridgeUnpausedIterator, error) {

	logs, sub, err := _PuffinBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PuffinBridgeUnpausedIterator{contract: _PuffinBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PuffinBridge *PuffinBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PuffinBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _PuffinBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinBridgeUnpaused)
				if err := _PuffinBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PuffinBridge *PuffinBridgeFilterer) ParseUnpaused(log types.Log) (*PuffinBridgeUnpaused, error) {
	event := new(PuffinBridgeUnpaused)
	if err := _PuffinBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
