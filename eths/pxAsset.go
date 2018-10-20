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

// PxaABI is the input ABI used to generate the binding from.
const PxaABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_approved\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"NewAsset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\"},{\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"splitAsset\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"name\":\"contentHash\",\"type\":\"bytes32\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\"},{\"name\":\"metaData\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"foundation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPXCAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getPXCBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Pxa is an auto generated Go binding around an Ethereum contract.
type Pxa struct {
	PxaCaller     // Read-only binding to the contract
	PxaTransactor // Write-only binding to the contract
	PxaFilterer   // Log filterer for contract events
}

// PxaCaller is an auto generated read-only Go binding around an Ethereum contract.
type PxaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PxaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PxaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PxaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PxaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PxaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PxaSession struct {
	Contract     *Pxa              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PxaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PxaCallerSession struct {
	Contract *PxaCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PxaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PxaTransactorSession struct {
	Contract     *PxaTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PxaRaw is an auto generated low-level Go binding around an Ethereum contract.
type PxaRaw struct {
	Contract *Pxa // Generic contract binding to access the raw methods on
}

// PxaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PxaCallerRaw struct {
	Contract *PxaCaller // Generic read-only contract binding to access the raw methods on
}

// PxaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PxaTransactorRaw struct {
	Contract *PxaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPxa creates a new instance of Pxa, bound to a specific deployed contract.
func NewPxa(address common.Address, backend bind.ContractBackend) (*Pxa, error) {
	contract, err := bindPxa(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pxa{PxaCaller: PxaCaller{contract: contract}, PxaTransactor: PxaTransactor{contract: contract}, PxaFilterer: PxaFilterer{contract: contract}}, nil
}

// NewPxaCaller creates a new read-only instance of Pxa, bound to a specific deployed contract.
func NewPxaCaller(address common.Address, caller bind.ContractCaller) (*PxaCaller, error) {
	contract, err := bindPxa(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PxaCaller{contract: contract}, nil
}

// NewPxaTransactor creates a new write-only instance of Pxa, bound to a specific deployed contract.
func NewPxaTransactor(address common.Address, transactor bind.ContractTransactor) (*PxaTransactor, error) {
	contract, err := bindPxa(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PxaTransactor{contract: contract}, nil
}

// NewPxaFilterer creates a new log filterer instance of Pxa, bound to a specific deployed contract.
func NewPxaFilterer(address common.Address, filterer bind.ContractFilterer) (*PxaFilterer, error) {
	contract, err := bindPxa(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PxaFilterer{contract: contract}, nil
}

// bindPxa binds a generic wrapper to an already deployed contract.
func bindPxa(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PxaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pxa *PxaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pxa.Contract.PxaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pxa *PxaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pxa.Contract.PxaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pxa *PxaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pxa.Contract.PxaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pxa *PxaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pxa.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pxa *PxaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pxa.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pxa *PxaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pxa.Contract.contract.Transact(opts, method, params...)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(contentHash bytes32, price uint256, weight uint256, metaData string)
func (_Pxa *PxaCaller) Assets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ContentHash [32]byte
	Price       *big.Int
	Weight      *big.Int
	MetaData    string
}, error) {
	ret := new(struct {
		ContentHash [32]byte
		Price       *big.Int
		Weight      *big.Int
		MetaData    string
	})
	out := ret
	err := _Pxa.contract.Call(opts, out, "assets", arg0)
	return *ret, err
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(contentHash bytes32, price uint256, weight uint256, metaData string)
func (_Pxa *PxaSession) Assets(arg0 *big.Int) (struct {
	ContentHash [32]byte
	Price       *big.Int
	Weight      *big.Int
	MetaData    string
}, error) {
	return _Pxa.Contract.Assets(&_Pxa.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(contentHash bytes32, price uint256, weight uint256, metaData string)
func (_Pxa *PxaCallerSession) Assets(arg0 *big.Int) (struct {
	ContentHash [32]byte
	Price       *big.Int
	Weight      *big.Int
	MetaData    string
}, error) {
	return _Pxa.Contract.Assets(&_Pxa.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Pxa *PxaCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pxa.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Pxa *PxaSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Pxa.Contract.BalanceOf(&_Pxa.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Pxa *PxaCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Pxa.Contract.BalanceOf(&_Pxa.CallOpts, _owner)
}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() constant returns(address)
func (_Pxa *PxaCaller) Foundation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pxa.contract.Call(opts, out, "foundation")
	return *ret0, err
}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() constant returns(address)
func (_Pxa *PxaSession) Foundation() (common.Address, error) {
	return _Pxa.Contract.Foundation(&_Pxa.CallOpts)
}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() constant returns(address)
func (_Pxa *PxaCallerSession) Foundation() (common.Address, error) {
	return _Pxa.Contract.Foundation(&_Pxa.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_Pxa *PxaCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pxa.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_Pxa *PxaSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Pxa.Contract.GetApproved(&_Pxa.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_Pxa *PxaCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Pxa.Contract.GetApproved(&_Pxa.CallOpts, _tokenId)
}

// GetPXCAddr is a free data retrieval call binding the contract method 0xf593870d.
//
// Solidity: function getPXCAddr() constant returns(address)
func (_Pxa *PxaCaller) GetPXCAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pxa.contract.Call(opts, out, "getPXCAddr")
	return *ret0, err
}

// GetPXCAddr is a free data retrieval call binding the contract method 0xf593870d.
//
// Solidity: function getPXCAddr() constant returns(address)
func (_Pxa *PxaSession) GetPXCAddr() (common.Address, error) {
	return _Pxa.Contract.GetPXCAddr(&_Pxa.CallOpts)
}

// GetPXCAddr is a free data retrieval call binding the contract method 0xf593870d.
//
// Solidity: function getPXCAddr() constant returns(address)
func (_Pxa *PxaCallerSession) GetPXCAddr() (common.Address, error) {
	return _Pxa.Contract.GetPXCAddr(&_Pxa.CallOpts)
}

// GetPXCBalance is a free data retrieval call binding the contract method 0x7197d0bc.
//
// Solidity: function getPXCBalance(_owner address) constant returns(uint256)
func (_Pxa *PxaCaller) GetPXCBalance(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pxa.contract.Call(opts, out, "getPXCBalance", _owner)
	return *ret0, err
}

// GetPXCBalance is a free data retrieval call binding the contract method 0x7197d0bc.
//
// Solidity: function getPXCBalance(_owner address) constant returns(uint256)
func (_Pxa *PxaSession) GetPXCBalance(_owner common.Address) (*big.Int, error) {
	return _Pxa.Contract.GetPXCBalance(&_Pxa.CallOpts, _owner)
}

// GetPXCBalance is a free data retrieval call binding the contract method 0x7197d0bc.
//
// Solidity: function getPXCBalance(_owner address) constant returns(uint256)
func (_Pxa *PxaCallerSession) GetPXCBalance(_owner common.Address) (*big.Int, error) {
	return _Pxa.Contract.GetPXCBalance(&_Pxa.CallOpts, _owner)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_Pxa *PxaCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pxa.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_Pxa *PxaSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Pxa.Contract.IsApprovedForAll(&_Pxa.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_Pxa *PxaCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Pxa.Contract.IsApprovedForAll(&_Pxa.CallOpts, _owner, _operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_Pxa *PxaCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pxa.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_Pxa *PxaSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Pxa.Contract.OwnerOf(&_Pxa.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_Pxa *PxaCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Pxa.Contract.OwnerOf(&_Pxa.CallOpts, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_approved address, _tokenId uint256) returns()
func (_Pxa *PxaTransactor) Approve(opts *bind.TransactOpts, _approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pxa.contract.Transact(opts, "approve", _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_approved address, _tokenId uint256) returns()
func (_Pxa *PxaSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pxa.Contract.Approve(&_Pxa.TransactOpts, _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_approved address, _tokenId uint256) returns()
func (_Pxa *PxaTransactorSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pxa.Contract.Approve(&_Pxa.TransactOpts, _approved, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x5fd8cb69.
//
// Solidity: function mint(_hash bytes32, _price uint256, _weight uint256, _data string) returns()
func (_Pxa *PxaTransactor) Mint(opts *bind.TransactOpts, _hash [32]byte, _price *big.Int, _weight *big.Int, _data string) (*types.Transaction, error) {
	return _Pxa.contract.Transact(opts, "mint", _hash, _price, _weight, _data)
}

// Mint is a paid mutator transaction binding the contract method 0x5fd8cb69.
//
// Solidity: function mint(_hash bytes32, _price uint256, _weight uint256, _data string) returns()
func (_Pxa *PxaSession) Mint(_hash [32]byte, _price *big.Int, _weight *big.Int, _data string) (*types.Transaction, error) {
	return _Pxa.Contract.Mint(&_Pxa.TransactOpts, _hash, _price, _weight, _data)
}

// Mint is a paid mutator transaction binding the contract method 0x5fd8cb69.
//
// Solidity: function mint(_hash bytes32, _price uint256, _weight uint256, _data string) returns()
func (_Pxa *PxaTransactorSession) Mint(_hash [32]byte, _price *big.Int, _weight *big.Int, _data string) (*types.Transaction, error) {
	return _Pxa.Contract.Mint(&_Pxa.TransactOpts, _hash, _price, _weight, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_Pxa *PxaTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Pxa.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_Pxa *PxaSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Pxa.Contract.SafeTransferFrom(&_Pxa.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_Pxa *PxaTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Pxa.Contract.SafeTransferFrom(&_Pxa.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_Pxa *PxaTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Pxa.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_Pxa *PxaSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Pxa.Contract.SetApprovalForAll(&_Pxa.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_Pxa *PxaTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Pxa.Contract.SetApprovalForAll(&_Pxa.TransactOpts, _operator, _approved)
}

// SplitAsset is a paid mutator transaction binding the contract method 0x7532e0f6.
//
// Solidity: function splitAsset(_tokenId uint256, _weight uint256, _buyer address) returns(uint256)
func (_Pxa *PxaTransactor) SplitAsset(opts *bind.TransactOpts, _tokenId *big.Int, _weight *big.Int, _buyer common.Address) (*types.Transaction, error) {
	return _Pxa.contract.Transact(opts, "splitAsset", _tokenId, _weight, _buyer)
}

// SplitAsset is a paid mutator transaction binding the contract method 0x7532e0f6.
//
// Solidity: function splitAsset(_tokenId uint256, _weight uint256, _buyer address) returns(uint256)
func (_Pxa *PxaSession) SplitAsset(_tokenId *big.Int, _weight *big.Int, _buyer common.Address) (*types.Transaction, error) {
	return _Pxa.Contract.SplitAsset(&_Pxa.TransactOpts, _tokenId, _weight, _buyer)
}

// SplitAsset is a paid mutator transaction binding the contract method 0x7532e0f6.
//
// Solidity: function splitAsset(_tokenId uint256, _weight uint256, _buyer address) returns(uint256)
func (_Pxa *PxaTransactorSession) SplitAsset(_tokenId *big.Int, _weight *big.Int, _buyer common.Address) (*types.Transaction, error) {
	return _Pxa.Contract.SplitAsset(&_Pxa.TransactOpts, _tokenId, _weight, _buyer)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_Pxa *PxaTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pxa.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_Pxa *PxaSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pxa.Contract.TransferFrom(&_Pxa.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_Pxa *PxaTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Pxa.Contract.TransferFrom(&_Pxa.TransactOpts, _from, _to, _tokenId)
}

// PxaApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pxa contract.
type PxaApprovalIterator struct {
	Event *PxaApproval // Event containing the contract specifics and raw log

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
func (it *PxaApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PxaApproval)
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
		it.Event = new(PxaApproval)
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
func (it *PxaApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PxaApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PxaApproval represents a Approval event raised by the Pxa contract.
type PxaApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_Pxa *PxaFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*PxaApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Pxa.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PxaApprovalIterator{contract: _Pxa.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_Pxa *PxaFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PxaApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Pxa.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PxaApproval)
				if err := _Pxa.contract.UnpackLog(event, "Approval", log); err != nil {
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

// PxaApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Pxa contract.
type PxaApprovalForAllIterator struct {
	Event *PxaApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PxaApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PxaApprovalForAll)
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
		it.Event = new(PxaApprovalForAll)
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
func (it *PxaApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PxaApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PxaApprovalForAll represents a ApprovalForAll event raised by the Pxa contract.
type PxaApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_Pxa *PxaFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*PxaApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Pxa.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &PxaApprovalForAllIterator{contract: _Pxa.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_Pxa *PxaFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PxaApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Pxa.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PxaApprovalForAll)
				if err := _Pxa.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// PxaNewAssetIterator is returned from FilterNewAsset and is used to iterate over the raw logs and unpacked data for NewAsset events raised by the Pxa contract.
type PxaNewAssetIterator struct {
	Event *PxaNewAsset // Event containing the contract specifics and raw log

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
func (it *PxaNewAssetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PxaNewAsset)
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
		it.Event = new(PxaNewAsset)
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
func (it *PxaNewAssetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PxaNewAssetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PxaNewAsset represents a NewAsset event raised by the Pxa contract.
type PxaNewAsset struct {
	Hash    [32]byte
	Owner   common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewAsset is a free log retrieval operation binding the contract event 0xc3a70059a71ca182bc9c23ad7e76e1d0e59af9a6250879f5fb33ba449165a8f6.
//
// Solidity: e NewAsset(_hash bytes32, _owner address, _tokenId uint256)
func (_Pxa *PxaFilterer) FilterNewAsset(opts *bind.FilterOpts) (*PxaNewAssetIterator, error) {

	logs, sub, err := _Pxa.contract.FilterLogs(opts, "NewAsset")
	if err != nil {
		return nil, err
	}
	return &PxaNewAssetIterator{contract: _Pxa.contract, event: "NewAsset", logs: logs, sub: sub}, nil
}

// WatchNewAsset is a free log subscription operation binding the contract event 0xc3a70059a71ca182bc9c23ad7e76e1d0e59af9a6250879f5fb33ba449165a8f6.
//
// Solidity: e NewAsset(_hash bytes32, _owner address, _tokenId uint256)
func (_Pxa *PxaFilterer) WatchNewAsset(opts *bind.WatchOpts, sink chan<- *PxaNewAsset) (event.Subscription, error) {

	logs, sub, err := _Pxa.contract.WatchLogs(opts, "NewAsset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PxaNewAsset)
				if err := _Pxa.contract.UnpackLog(event, "NewAsset", log); err != nil {
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

// PxaTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pxa contract.
type PxaTransferIterator struct {
	Event *PxaTransfer // Event containing the contract specifics and raw log

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
func (it *PxaTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PxaTransfer)
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
		it.Event = new(PxaTransfer)
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
func (it *PxaTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PxaTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PxaTransfer represents a Transfer event raised by the Pxa contract.
type PxaTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_Pxa *PxaFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*PxaTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Pxa.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PxaTransferIterator{contract: _Pxa.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_Pxa *PxaFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PxaTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Pxa.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PxaTransfer)
				if err := _Pxa.contract.UnpackLog(event, "Transfer", log); err != nil {
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
