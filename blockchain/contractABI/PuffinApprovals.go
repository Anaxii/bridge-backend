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

// PuffinApprovalsMetaData contains all meta data concerning the PuffinApprovals contract.
var PuffinApprovalsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"allowApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"ApprovalUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"removeApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"canApprove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PuffinApprovalsABI is the input ABI used to generate the binding from.
// Deprecated: Use PuffinApprovalsMetaData.ABI instead.
var PuffinApprovalsABI = PuffinApprovalsMetaData.ABI

// PuffinApprovals is an auto generated Go binding around an Ethereum contract.
type PuffinApprovals struct {
	PuffinApprovalsCaller     // Read-only binding to the contract
	PuffinApprovalsTransactor // Write-only binding to the contract
	PuffinApprovalsFilterer   // Log filterer for contract events
}

// PuffinApprovalsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PuffinApprovalsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinApprovalsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PuffinApprovalsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinApprovalsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PuffinApprovalsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinApprovalsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PuffinApprovalsSession struct {
	Contract     *PuffinApprovals  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PuffinApprovalsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PuffinApprovalsCallerSession struct {
	Contract *PuffinApprovalsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PuffinApprovalsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PuffinApprovalsTransactorSession struct {
	Contract     *PuffinApprovalsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PuffinApprovalsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PuffinApprovalsRaw struct {
	Contract *PuffinApprovals // Generic contract binding to access the raw methods on
}

// PuffinApprovalsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PuffinApprovalsCallerRaw struct {
	Contract *PuffinApprovalsCaller // Generic read-only contract binding to access the raw methods on
}

// PuffinApprovalsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PuffinApprovalsTransactorRaw struct {
	Contract *PuffinApprovalsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPuffinApprovals creates a new instance of PuffinApprovals, bound to a specific deployed contract.
func NewPuffinApprovals(address common.Address, backend bind.ContractBackend) (*PuffinApprovals, error) {
	contract, err := bindPuffinApprovals(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PuffinApprovals{PuffinApprovalsCaller: PuffinApprovalsCaller{contract: contract}, PuffinApprovalsTransactor: PuffinApprovalsTransactor{contract: contract}, PuffinApprovalsFilterer: PuffinApprovalsFilterer{contract: contract}}, nil
}

// NewPuffinApprovalsCaller creates a new read-only instance of PuffinApprovals, bound to a specific deployed contract.
func NewPuffinApprovalsCaller(address common.Address, caller bind.ContractCaller) (*PuffinApprovalsCaller, error) {
	contract, err := bindPuffinApprovals(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinApprovalsCaller{contract: contract}, nil
}

// NewPuffinApprovalsTransactor creates a new write-only instance of PuffinApprovals, bound to a specific deployed contract.
func NewPuffinApprovalsTransactor(address common.Address, transactor bind.ContractTransactor) (*PuffinApprovalsTransactor, error) {
	contract, err := bindPuffinApprovals(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinApprovalsTransactor{contract: contract}, nil
}

// NewPuffinApprovalsFilterer creates a new log filterer instance of PuffinApprovals, bound to a specific deployed contract.
func NewPuffinApprovalsFilterer(address common.Address, filterer bind.ContractFilterer) (*PuffinApprovalsFilterer, error) {
	contract, err := bindPuffinApprovals(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PuffinApprovalsFilterer{contract: contract}, nil
}

// bindPuffinApprovals binds a generic wrapper to an already deployed contract.
func bindPuffinApprovals(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PuffinApprovalsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinApprovals *PuffinApprovalsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinApprovals.Contract.PuffinApprovalsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinApprovals *PuffinApprovalsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinApprovals.Contract.PuffinApprovalsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinApprovals *PuffinApprovalsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinApprovals.Contract.PuffinApprovalsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinApprovals *PuffinApprovalsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinApprovals.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinApprovals *PuffinApprovalsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinApprovals.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinApprovals *PuffinApprovalsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinApprovals.Contract.contract.Transact(opts, method, params...)
}

// CanApprove is a free data retrieval call binding the contract method 0x74743444.
//
// Solidity: function canApprove(address ) view returns(bool)
func (_PuffinApprovals *PuffinApprovalsCaller) CanApprove(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PuffinApprovals.contract.Call(opts, &out, "canApprove", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanApprove is a free data retrieval call binding the contract method 0x74743444.
//
// Solidity: function canApprove(address ) view returns(bool)
func (_PuffinApprovals *PuffinApprovalsSession) CanApprove(arg0 common.Address) (bool, error) {
	return _PuffinApprovals.Contract.CanApprove(&_PuffinApprovals.CallOpts, arg0)
}

// CanApprove is a free data retrieval call binding the contract method 0x74743444.
//
// Solidity: function canApprove(address ) view returns(bool)
func (_PuffinApprovals *PuffinApprovalsCallerSession) CanApprove(arg0 common.Address) (bool, error) {
	return _PuffinApprovals.Contract.CanApprove(&_PuffinApprovals.CallOpts, arg0)
}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(address ) view returns(bool)
func (_PuffinApprovals *PuffinApprovalsCaller) IsApproved(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PuffinApprovals.contract.Call(opts, &out, "isApproved", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(address ) view returns(bool)
func (_PuffinApprovals *PuffinApprovalsSession) IsApproved(arg0 common.Address) (bool, error) {
	return _PuffinApprovals.Contract.IsApproved(&_PuffinApprovals.CallOpts, arg0)
}

// IsApproved is a free data retrieval call binding the contract method 0x673448dd.
//
// Solidity: function isApproved(address ) view returns(bool)
func (_PuffinApprovals *PuffinApprovalsCallerSession) IsApproved(arg0 common.Address) (bool, error) {
	return _PuffinApprovals.Contract.IsApproved(&_PuffinApprovals.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinApprovals *PuffinApprovalsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinApprovals.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinApprovals *PuffinApprovalsSession) Owner() (common.Address, error) {
	return _PuffinApprovals.Contract.Owner(&_PuffinApprovals.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinApprovals *PuffinApprovalsCallerSession) Owner() (common.Address, error) {
	return _PuffinApprovals.Contract.Owner(&_PuffinApprovals.CallOpts)
}

// AllowApprove is a paid mutator transaction binding the contract method 0xedd1cf9a.
//
// Solidity: function allowApprove(address user) returns()
func (_PuffinApprovals *PuffinApprovalsTransactor) AllowApprove(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.contract.Transact(opts, "allowApprove", user)
}

// AllowApprove is a paid mutator transaction binding the contract method 0xedd1cf9a.
//
// Solidity: function allowApprove(address user) returns()
func (_PuffinApprovals *PuffinApprovalsSession) AllowApprove(user common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.Contract.AllowApprove(&_PuffinApprovals.TransactOpts, user)
}

// AllowApprove is a paid mutator transaction binding the contract method 0xedd1cf9a.
//
// Solidity: function allowApprove(address user) returns()
func (_PuffinApprovals *PuffinApprovalsTransactorSession) AllowApprove(user common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.Contract.AllowApprove(&_PuffinApprovals.TransactOpts, user)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address user) returns()
func (_PuffinApprovals *PuffinApprovalsTransactor) Approve(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.contract.Transact(opts, "approve", user)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address user) returns()
func (_PuffinApprovals *PuffinApprovalsSession) Approve(user common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.Contract.Approve(&_PuffinApprovals.TransactOpts, user)
}

// Approve is a paid mutator transaction binding the contract method 0xdaea85c5.
//
// Solidity: function approve(address user) returns()
func (_PuffinApprovals *PuffinApprovalsTransactorSession) Approve(user common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.Contract.Approve(&_PuffinApprovals.TransactOpts, user)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address user) returns()
func (_PuffinApprovals *PuffinApprovalsTransactor) Remove(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.contract.Transact(opts, "remove", user)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address user) returns()
func (_PuffinApprovals *PuffinApprovalsSession) Remove(user common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.Contract.Remove(&_PuffinApprovals.TransactOpts, user)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(address user) returns()
func (_PuffinApprovals *PuffinApprovalsTransactorSession) Remove(user common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.Contract.Remove(&_PuffinApprovals.TransactOpts, user)
}

// RemoveApprove is a paid mutator transaction binding the contract method 0xd5a0d76f.
//
// Solidity: function removeApprove(address user) returns()
func (_PuffinApprovals *PuffinApprovalsTransactor) RemoveApprove(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.contract.Transact(opts, "removeApprove", user)
}

// RemoveApprove is a paid mutator transaction binding the contract method 0xd5a0d76f.
//
// Solidity: function removeApprove(address user) returns()
func (_PuffinApprovals *PuffinApprovalsSession) RemoveApprove(user common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.Contract.RemoveApprove(&_PuffinApprovals.TransactOpts, user)
}

// RemoveApprove is a paid mutator transaction binding the contract method 0xd5a0d76f.
//
// Solidity: function removeApprove(address user) returns()
func (_PuffinApprovals *PuffinApprovalsTransactorSession) RemoveApprove(user common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.Contract.RemoveApprove(&_PuffinApprovals.TransactOpts, user)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinApprovals *PuffinApprovalsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinApprovals.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinApprovals *PuffinApprovalsSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinApprovals.Contract.RenounceOwnership(&_PuffinApprovals.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinApprovals *PuffinApprovalsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinApprovals.Contract.RenounceOwnership(&_PuffinApprovals.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinApprovals *PuffinApprovalsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinApprovals *PuffinApprovalsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.Contract.TransferOwnership(&_PuffinApprovals.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinApprovals *PuffinApprovalsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinApprovals.Contract.TransferOwnership(&_PuffinApprovals.TransactOpts, newOwner)
}

// PuffinApprovalsApprovalUpdateIterator is returned from FilterApprovalUpdate and is used to iterate over the raw logs and unpacked data for ApprovalUpdate events raised by the PuffinApprovals contract.
type PuffinApprovalsApprovalUpdateIterator struct {
	Event *PuffinApprovalsApprovalUpdate // Event containing the contract specifics and raw log

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
func (it *PuffinApprovalsApprovalUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinApprovalsApprovalUpdate)
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
		it.Event = new(PuffinApprovalsApprovalUpdate)
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
func (it *PuffinApprovalsApprovalUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinApprovalsApprovalUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinApprovalsApprovalUpdate represents a ApprovalUpdate event raised by the PuffinApprovals contract.
type PuffinApprovalsApprovalUpdate struct {
	User   common.Address
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApprovalUpdate is a free log retrieval operation binding the contract event 0x6f5912b9ff37dc15253ef37ffa5348b313bc329fc4e600d2ea98334b6509a5ff.
//
// Solidity: event ApprovalUpdate(address indexed user, bool indexed status)
func (_PuffinApprovals *PuffinApprovalsFilterer) FilterApprovalUpdate(opts *bind.FilterOpts, user []common.Address, status []bool) (*PuffinApprovalsApprovalUpdateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _PuffinApprovals.contract.FilterLogs(opts, "ApprovalUpdate", userRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &PuffinApprovalsApprovalUpdateIterator{contract: _PuffinApprovals.contract, event: "ApprovalUpdate", logs: logs, sub: sub}, nil
}

// WatchApprovalUpdate is a free log subscription operation binding the contract event 0x6f5912b9ff37dc15253ef37ffa5348b313bc329fc4e600d2ea98334b6509a5ff.
//
// Solidity: event ApprovalUpdate(address indexed user, bool indexed status)
func (_PuffinApprovals *PuffinApprovalsFilterer) WatchApprovalUpdate(opts *bind.WatchOpts, sink chan<- *PuffinApprovalsApprovalUpdate, user []common.Address, status []bool) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _PuffinApprovals.contract.WatchLogs(opts, "ApprovalUpdate", userRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinApprovalsApprovalUpdate)
				if err := _PuffinApprovals.contract.UnpackLog(event, "ApprovalUpdate", log); err != nil {
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

// ParseApprovalUpdate is a log parse operation binding the contract event 0x6f5912b9ff37dc15253ef37ffa5348b313bc329fc4e600d2ea98334b6509a5ff.
//
// Solidity: event ApprovalUpdate(address indexed user, bool indexed status)
func (_PuffinApprovals *PuffinApprovalsFilterer) ParseApprovalUpdate(log types.Log) (*PuffinApprovalsApprovalUpdate, error) {
	event := new(PuffinApprovalsApprovalUpdate)
	if err := _PuffinApprovals.contract.UnpackLog(event, "ApprovalUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinApprovalsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PuffinApprovals contract.
type PuffinApprovalsOwnershipTransferredIterator struct {
	Event *PuffinApprovalsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PuffinApprovalsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinApprovalsOwnershipTransferred)
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
		it.Event = new(PuffinApprovalsOwnershipTransferred)
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
func (it *PuffinApprovalsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinApprovalsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinApprovalsOwnershipTransferred represents a OwnershipTransferred event raised by the PuffinApprovals contract.
type PuffinApprovalsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinApprovals *PuffinApprovalsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PuffinApprovalsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinApprovals.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PuffinApprovalsOwnershipTransferredIterator{contract: _PuffinApprovals.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinApprovals *PuffinApprovalsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PuffinApprovalsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinApprovals.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinApprovalsOwnershipTransferred)
				if err := _PuffinApprovals.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PuffinApprovals *PuffinApprovalsFilterer) ParseOwnershipTransferred(log types.Log) (*PuffinApprovalsOwnershipTransferred, error) {
	event := new(PuffinApprovalsOwnershipTransferred)
	if err := _PuffinApprovals.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
