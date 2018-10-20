// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eths

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// PxcABI is the input ABI used to generate the binding from.
const PxcABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"foundation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// Pxc is an auto generated Go binding around an Ethereum contract.
type Pxc struct {
	PxcCaller     // Read-only binding to the contract
	PxcTransactor // Write-only binding to the contract
	PxcFilterer   // Log filterer for contract events
}

// PxcCaller is an auto generated read-only Go binding around an Ethereum contract.
type PxcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PxcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PxcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PxcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PxcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PxcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PxcSession struct {
	Contract     *Pxc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PxcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PxcCallerSession struct {
	Contract *PxcCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PxcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PxcTransactorSession struct {
	Contract     *PxcTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PxcRaw is an auto generated low-level Go binding around an Ethereum contract.
type PxcRaw struct {
	Contract *Pxc // Generic contract binding to access the raw methods on
}

// PxcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PxcCallerRaw struct {
	Contract *PxcCaller // Generic read-only contract binding to access the raw methods on
}

// PxcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PxcTransactorRaw struct {
	Contract *PxcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPxc creates a new instance of Pxc, bound to a specific deployed contract.
func NewPxc(address common.Address, backend bind.ContractBackend) (*Pxc, error) {
	contract, err := bindPxc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pxc{PxcCaller: PxcCaller{contract: contract}, PxcTransactor: PxcTransactor{contract: contract}, PxcFilterer: PxcFilterer{contract: contract}}, nil
}

// NewPxcCaller creates a new read-only instance of Pxc, bound to a specific deployed contract.
func NewPxcCaller(address common.Address, caller bind.ContractCaller) (*PxcCaller, error) {
	contract, err := bindPxc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PxcCaller{contract: contract}, nil
}

// NewPxcTransactor creates a new write-only instance of Pxc, bound to a specific deployed contract.
func NewPxcTransactor(address common.Address, transactor bind.ContractTransactor) (*PxcTransactor, error) {
	contract, err := bindPxc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PxcTransactor{contract: contract}, nil
}

// NewPxcFilterer creates a new log filterer instance of Pxc, bound to a specific deployed contract.
func NewPxcFilterer(address common.Address, filterer bind.ContractFilterer) (*PxcFilterer, error) {
	contract, err := bindPxc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PxcFilterer{contract: contract}, nil
}

// bindPxc binds a generic wrapper to an already deployed contract.
func bindPxc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PxcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pxc *PxcRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pxc.Contract.PxcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pxc *PxcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pxc.Contract.PxcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pxc *PxcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pxc.Contract.PxcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pxc *PxcCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pxc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pxc *PxcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pxc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pxc *PxcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pxc.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Pxc *PxcCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pxc.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Pxc *PxcSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Pxc.Contract.Allowance(&_Pxc.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Pxc *PxcCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Pxc.Contract.Allowance(&_Pxc.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Pxc *PxcCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pxc.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Pxc *PxcSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Pxc.Contract.BalanceOf(&_Pxc.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Pxc *PxcCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Pxc.Contract.BalanceOf(&_Pxc.CallOpts, _owner)
}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() constant returns(address)
func (_Pxc *PxcCaller) Foundation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pxc.contract.Call(opts, out, "foundation")
	return *ret0, err
}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() constant returns(address)
func (_Pxc *PxcSession) Foundation() (common.Address, error) {
	return _Pxc.Contract.Foundation(&_Pxc.CallOpts)
}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() constant returns(address)
func (_Pxc *PxcCallerSession) Foundation() (common.Address, error) {
	return _Pxc.Contract.Foundation(&_Pxc.CallOpts)
}

// GetAddr is a free data retrieval call binding the contract method 0xa74c2bb6.
//
// Solidity: function getAddr() constant returns(address)
func (_Pxc *PxcCaller) GetAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pxc.contract.Call(opts, out, "getAddr")
	return *ret0, err
}

// GetAddr is a free data retrieval call binding the contract method 0xa74c2bb6.
//
// Solidity: function getAddr() constant returns(address)
func (_Pxc *PxcSession) GetAddr() (common.Address, error) {
	return _Pxc.Contract.GetAddr(&_Pxc.CallOpts)
}

// GetAddr is a free data retrieval call binding the contract method 0xa74c2bb6.
//
// Solidity: function getAddr() constant returns(address)
func (_Pxc *PxcCallerSession) GetAddr() (common.Address, error) {
	return _Pxc.Contract.GetAddr(&_Pxc.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_Pxc *PxcCaller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pxc.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_Pxc *PxcSession) Issuer() (common.Address, error) {
	return _Pxc.Contract.Issuer(&_Pxc.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_Pxc *PxcCallerSession) Issuer() (common.Address, error) {
	return _Pxc.Contract.Issuer(&_Pxc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Pxc *PxcCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pxc.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Pxc *PxcSession) Name() (string, error) {
	return _Pxc.Contract.Name(&_Pxc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Pxc *PxcCallerSession) Name() (string, error) {
	return _Pxc.Contract.Name(&_Pxc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Pxc *PxcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pxc.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Pxc *PxcSession) Symbol() (string, error) {
	return _Pxc.Contract.Symbol(&_Pxc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Pxc *PxcCallerSession) Symbol() (string, error) {
	return _Pxc.Contract.Symbol(&_Pxc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(totalSupply uint256)
func (_Pxc *PxcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pxc.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(totalSupply uint256)
func (_Pxc *PxcSession) TotalSupply() (*big.Int, error) {
	return _Pxc.Contract.TotalSupply(&_Pxc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(totalSupply uint256)
func (_Pxc *PxcCallerSession) TotalSupply() (*big.Int, error) {
	return _Pxc.Contract.TotalSupply(&_Pxc.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Pxc *PxcTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Pxc *PxcSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.Contract.Approve(&_Pxc.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Pxc *PxcTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.Contract.Approve(&_Pxc.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Pxc *PxcTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Pxc *PxcSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.Contract.Transfer(&_Pxc.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Pxc *PxcTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.Contract.Transfer(&_Pxc.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Pxc *PxcTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Pxc *PxcSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.Contract.TransferFrom(&_Pxc.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Pxc *PxcTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.Contract.TransferFrom(&_Pxc.TransactOpts, _from, _to, _value)
}

// PxcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pxc contract.
type PxcApprovalIterator struct {
	Event *PxcApproval // Event containing the contract specifics and raw log

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
func (it *PxcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PxcApproval)
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
		it.Event = new(PxcApproval)
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
func (it *PxcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PxcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PxcApproval represents a Approval event raised by the Pxc contract.
type PxcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_Pxc *PxcFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*PxcApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Pxc.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &PxcApprovalIterator{contract: _Pxc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_Pxc *PxcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PxcApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Pxc.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PxcApproval)
				if err := _Pxc.contract.UnpackLog(event, "Approval", log); err != nil {
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

// PxcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pxc contract.
type PxcTransferIterator struct {
	Event *PxcTransfer // Event containing the contract specifics and raw log

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
func (it *PxcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PxcTransfer)
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
		it.Event = new(PxcTransfer)
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
func (it *PxcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PxcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PxcTransfer represents a Transfer event raised by the Pxc contract.
type PxcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_Pxc *PxcFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*PxcTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Pxc.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &PxcTransferIterator{contract: _Pxc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_Pxc *PxcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PxcTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Pxc.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PxcTransfer)
				if err := _Pxc.contract.UnpackLog(event, "Transfer", log); err != nil {
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
