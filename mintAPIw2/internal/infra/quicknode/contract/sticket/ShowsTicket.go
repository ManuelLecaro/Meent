// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sticket

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
	_ = abi.ConvertType
)

// SticketMetaData contains all meta data concerning the Sticket contract.
var SticketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vrfCoordinator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"place\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"specialPrices\",\"type\":\"string[]\"}],\"name\":\"ShowCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"showId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"TicketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"place\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"specialPrices\",\"type\":\"string[]\"}],\"name\":\"createShow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"showId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"showId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"createTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ticketId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"showId\",\"type\":\"uint256\"}],\"name\":\"getShow\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"place\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"specialPrices\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ticketId\",\"type\":\"uint256\"}],\"name\":\"getTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"showId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextShowId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTicketId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"shows\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"place\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tickets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"showId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SticketABI is the input ABI used to generate the binding from.
// Deprecated: Use SticketMetaData.ABI instead.
var SticketABI = SticketMetaData.ABI

// Sticket is an auto generated Go binding around an Ethereum contract.
type Sticket struct {
	SticketCaller     // Read-only binding to the contract
	SticketTransactor // Write-only binding to the contract
	SticketFilterer   // Log filterer for contract events
}

// SticketCaller is an auto generated read-only Go binding around an Ethereum contract.
type SticketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SticketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SticketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SticketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SticketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SticketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SticketSession struct {
	Contract     *Sticket          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SticketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SticketCallerSession struct {
	Contract *SticketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SticketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SticketTransactorSession struct {
	Contract     *SticketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SticketRaw is an auto generated low-level Go binding around an Ethereum contract.
type SticketRaw struct {
	Contract *Sticket // Generic contract binding to access the raw methods on
}

// SticketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SticketCallerRaw struct {
	Contract *SticketCaller // Generic read-only contract binding to access the raw methods on
}

// SticketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SticketTransactorRaw struct {
	Contract *SticketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSticket creates a new instance of Sticket, bound to a specific deployed contract.
func NewSticket(address common.Address, backend bind.ContractBackend) (*Sticket, error) {
	contract, err := bindSticket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sticket{SticketCaller: SticketCaller{contract: contract}, SticketTransactor: SticketTransactor{contract: contract}, SticketFilterer: SticketFilterer{contract: contract}}, nil
}

// NewSticketCaller creates a new read-only instance of Sticket, bound to a specific deployed contract.
func NewSticketCaller(address common.Address, caller bind.ContractCaller) (*SticketCaller, error) {
	contract, err := bindSticket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SticketCaller{contract: contract}, nil
}

// NewSticketTransactor creates a new write-only instance of Sticket, bound to a specific deployed contract.
func NewSticketTransactor(address common.Address, transactor bind.ContractTransactor) (*SticketTransactor, error) {
	contract, err := bindSticket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SticketTransactor{contract: contract}, nil
}

// NewSticketFilterer creates a new log filterer instance of Sticket, bound to a specific deployed contract.
func NewSticketFilterer(address common.Address, filterer bind.ContractFilterer) (*SticketFilterer, error) {
	contract, err := bindSticket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SticketFilterer{contract: contract}, nil
}

// bindSticket binds a generic wrapper to an already deployed contract.
func bindSticket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SticketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sticket *SticketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sticket.Contract.SticketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sticket *SticketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sticket.Contract.SticketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sticket *SticketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sticket.Contract.SticketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sticket *SticketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sticket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sticket *SticketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sticket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sticket *SticketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sticket.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Sticket *SticketCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Sticket *SticketSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Sticket.Contract.BalanceOf(&_Sticket.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Sticket *SticketCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Sticket.Contract.BalanceOf(&_Sticket.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Sticket *SticketCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Sticket *SticketSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Sticket.Contract.GetApproved(&_Sticket.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Sticket *SticketCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Sticket.Contract.GetApproved(&_Sticket.CallOpts, tokenId)
}

// GetShow is a free data retrieval call binding the contract method 0x590f9ece.
//
// Solidity: function getShow(uint256 showId) view returns(string name, uint256 date, string place, string[] specialPrices)
func (_Sticket *SticketCaller) GetShow(opts *bind.CallOpts, showId *big.Int) (struct {
	Name          string
	Date          *big.Int
	Place         string
	SpecialPrices []string
}, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "getShow", showId)

	outstruct := new(struct {
		Name          string
		Date          *big.Int
		Place         string
		SpecialPrices []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Date = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Place = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.SpecialPrices = *abi.ConvertType(out[3], new([]string)).(*[]string)

	return *outstruct, err

}

// GetShow is a free data retrieval call binding the contract method 0x590f9ece.
//
// Solidity: function getShow(uint256 showId) view returns(string name, uint256 date, string place, string[] specialPrices)
func (_Sticket *SticketSession) GetShow(showId *big.Int) (struct {
	Name          string
	Date          *big.Int
	Place         string
	SpecialPrices []string
}, error) {
	return _Sticket.Contract.GetShow(&_Sticket.CallOpts, showId)
}

// GetShow is a free data retrieval call binding the contract method 0x590f9ece.
//
// Solidity: function getShow(uint256 showId) view returns(string name, uint256 date, string place, string[] specialPrices)
func (_Sticket *SticketCallerSession) GetShow(showId *big.Int) (struct {
	Name          string
	Date          *big.Int
	Place         string
	SpecialPrices []string
}, error) {
	return _Sticket.Contract.GetShow(&_Sticket.CallOpts, showId)
}

// GetTicket is a free data retrieval call binding the contract method 0x7dc379fa.
//
// Solidity: function getTicket(uint256 ticketId) view returns(uint256 showId, string metadata, uint256 price)
func (_Sticket *SticketCaller) GetTicket(opts *bind.CallOpts, ticketId *big.Int) (struct {
	ShowId   *big.Int
	Metadata string
	Price    *big.Int
}, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "getTicket", ticketId)

	outstruct := new(struct {
		ShowId   *big.Int
		Metadata string
		Price    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ShowId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Metadata = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTicket is a free data retrieval call binding the contract method 0x7dc379fa.
//
// Solidity: function getTicket(uint256 ticketId) view returns(uint256 showId, string metadata, uint256 price)
func (_Sticket *SticketSession) GetTicket(ticketId *big.Int) (struct {
	ShowId   *big.Int
	Metadata string
	Price    *big.Int
}, error) {
	return _Sticket.Contract.GetTicket(&_Sticket.CallOpts, ticketId)
}

// GetTicket is a free data retrieval call binding the contract method 0x7dc379fa.
//
// Solidity: function getTicket(uint256 ticketId) view returns(uint256 showId, string metadata, uint256 price)
func (_Sticket *SticketCallerSession) GetTicket(ticketId *big.Int) (struct {
	ShowId   *big.Int
	Metadata string
	Price    *big.Int
}, error) {
	return _Sticket.Contract.GetTicket(&_Sticket.CallOpts, ticketId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Sticket *SticketCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Sticket *SticketSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Sticket.Contract.IsApprovedForAll(&_Sticket.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Sticket *SticketCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Sticket.Contract.IsApprovedForAll(&_Sticket.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Sticket *SticketCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Sticket *SticketSession) Name() (string, error) {
	return _Sticket.Contract.Name(&_Sticket.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Sticket *SticketCallerSession) Name() (string, error) {
	return _Sticket.Contract.Name(&_Sticket.CallOpts)
}

// NextShowId is a free data retrieval call binding the contract method 0xaf61b2ca.
//
// Solidity: function nextShowId() view returns(uint256)
func (_Sticket *SticketCaller) NextShowId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "nextShowId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextShowId is a free data retrieval call binding the contract method 0xaf61b2ca.
//
// Solidity: function nextShowId() view returns(uint256)
func (_Sticket *SticketSession) NextShowId() (*big.Int, error) {
	return _Sticket.Contract.NextShowId(&_Sticket.CallOpts)
}

// NextShowId is a free data retrieval call binding the contract method 0xaf61b2ca.
//
// Solidity: function nextShowId() view returns(uint256)
func (_Sticket *SticketCallerSession) NextShowId() (*big.Int, error) {
	return _Sticket.Contract.NextShowId(&_Sticket.CallOpts)
}

// NextTicketId is a free data retrieval call binding the contract method 0x6031a52f.
//
// Solidity: function nextTicketId() view returns(uint256)
func (_Sticket *SticketCaller) NextTicketId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "nextTicketId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTicketId is a free data retrieval call binding the contract method 0x6031a52f.
//
// Solidity: function nextTicketId() view returns(uint256)
func (_Sticket *SticketSession) NextTicketId() (*big.Int, error) {
	return _Sticket.Contract.NextTicketId(&_Sticket.CallOpts)
}

// NextTicketId is a free data retrieval call binding the contract method 0x6031a52f.
//
// Solidity: function nextTicketId() view returns(uint256)
func (_Sticket *SticketCallerSession) NextTicketId() (*big.Int, error) {
	return _Sticket.Contract.NextTicketId(&_Sticket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sticket *SticketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sticket *SticketSession) Owner() (common.Address, error) {
	return _Sticket.Contract.Owner(&_Sticket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sticket *SticketCallerSession) Owner() (common.Address, error) {
	return _Sticket.Contract.Owner(&_Sticket.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Sticket *SticketCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Sticket *SticketSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Sticket.Contract.OwnerOf(&_Sticket.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Sticket *SticketCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Sticket.Contract.OwnerOf(&_Sticket.CallOpts, tokenId)
}

// Shows is a free data retrieval call binding the contract method 0x30f0931e.
//
// Solidity: function shows(uint256 ) view returns(string name, uint256 date, string place)
func (_Sticket *SticketCaller) Shows(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name  string
	Date  *big.Int
	Place string
}, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "shows", arg0)

	outstruct := new(struct {
		Name  string
		Date  *big.Int
		Place string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Date = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Place = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// Shows is a free data retrieval call binding the contract method 0x30f0931e.
//
// Solidity: function shows(uint256 ) view returns(string name, uint256 date, string place)
func (_Sticket *SticketSession) Shows(arg0 *big.Int) (struct {
	Name  string
	Date  *big.Int
	Place string
}, error) {
	return _Sticket.Contract.Shows(&_Sticket.CallOpts, arg0)
}

// Shows is a free data retrieval call binding the contract method 0x30f0931e.
//
// Solidity: function shows(uint256 ) view returns(string name, uint256 date, string place)
func (_Sticket *SticketCallerSession) Shows(arg0 *big.Int) (struct {
	Name  string
	Date  *big.Int
	Place string
}, error) {
	return _Sticket.Contract.Shows(&_Sticket.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sticket *SticketCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sticket *SticketSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Sticket.Contract.SupportsInterface(&_Sticket.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sticket *SticketCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Sticket.Contract.SupportsInterface(&_Sticket.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sticket *SticketCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sticket *SticketSession) Symbol() (string, error) {
	return _Sticket.Contract.Symbol(&_Sticket.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Sticket *SticketCallerSession) Symbol() (string, error) {
	return _Sticket.Contract.Symbol(&_Sticket.CallOpts)
}

// Tickets is a free data retrieval call binding the contract method 0x50b44712.
//
// Solidity: function tickets(uint256 ) view returns(uint256 showId, string metadata, uint256 price)
func (_Sticket *SticketCaller) Tickets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ShowId   *big.Int
	Metadata string
	Price    *big.Int
}, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "tickets", arg0)

	outstruct := new(struct {
		ShowId   *big.Int
		Metadata string
		Price    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ShowId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Metadata = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Tickets is a free data retrieval call binding the contract method 0x50b44712.
//
// Solidity: function tickets(uint256 ) view returns(uint256 showId, string metadata, uint256 price)
func (_Sticket *SticketSession) Tickets(arg0 *big.Int) (struct {
	ShowId   *big.Int
	Metadata string
	Price    *big.Int
}, error) {
	return _Sticket.Contract.Tickets(&_Sticket.CallOpts, arg0)
}

// Tickets is a free data retrieval call binding the contract method 0x50b44712.
//
// Solidity: function tickets(uint256 ) view returns(uint256 showId, string metadata, uint256 price)
func (_Sticket *SticketCallerSession) Tickets(arg0 *big.Int) (struct {
	ShowId   *big.Int
	Metadata string
	Price    *big.Int
}, error) {
	return _Sticket.Contract.Tickets(&_Sticket.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Sticket *SticketCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Sticket.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Sticket *SticketSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Sticket.Contract.TokenURI(&_Sticket.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Sticket *SticketCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Sticket.Contract.TokenURI(&_Sticket.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Sticket *SticketTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Sticket.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Sticket *SticketSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Sticket.Contract.Approve(&_Sticket.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Sticket *SticketTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Sticket.Contract.Approve(&_Sticket.TransactOpts, to, tokenId)
}

// CreateShow is a paid mutator transaction binding the contract method 0x8805d01e.
//
// Solidity: function createShow(string name, uint256 date, string place, string[] specialPrices) returns(uint256 showId)
func (_Sticket *SticketTransactor) CreateShow(opts *bind.TransactOpts, name string, date *big.Int, place string, specialPrices []string) (*types.Transaction, error) {
	return _Sticket.contract.Transact(opts, "createShow", name, date, place, specialPrices)
}

// CreateShow is a paid mutator transaction binding the contract method 0x8805d01e.
//
// Solidity: function createShow(string name, uint256 date, string place, string[] specialPrices) returns(uint256 showId)
func (_Sticket *SticketSession) CreateShow(name string, date *big.Int, place string, specialPrices []string) (*types.Transaction, error) {
	return _Sticket.Contract.CreateShow(&_Sticket.TransactOpts, name, date, place, specialPrices)
}

// CreateShow is a paid mutator transaction binding the contract method 0x8805d01e.
//
// Solidity: function createShow(string name, uint256 date, string place, string[] specialPrices) returns(uint256 showId)
func (_Sticket *SticketTransactorSession) CreateShow(name string, date *big.Int, place string, specialPrices []string) (*types.Transaction, error) {
	return _Sticket.Contract.CreateShow(&_Sticket.TransactOpts, name, date, place, specialPrices)
}

// CreateTicket is a paid mutator transaction binding the contract method 0x5a92035c.
//
// Solidity: function createTicket(uint256 showId, string metadata, uint256 price) returns(uint256 ticketId)
func (_Sticket *SticketTransactor) CreateTicket(opts *bind.TransactOpts, showId *big.Int, metadata string, price *big.Int) (*types.Transaction, error) {
	return _Sticket.contract.Transact(opts, "createTicket", showId, metadata, price)
}

// CreateTicket is a paid mutator transaction binding the contract method 0x5a92035c.
//
// Solidity: function createTicket(uint256 showId, string metadata, uint256 price) returns(uint256 ticketId)
func (_Sticket *SticketSession) CreateTicket(showId *big.Int, metadata string, price *big.Int) (*types.Transaction, error) {
	return _Sticket.Contract.CreateTicket(&_Sticket.TransactOpts, showId, metadata, price)
}

// CreateTicket is a paid mutator transaction binding the contract method 0x5a92035c.
//
// Solidity: function createTicket(uint256 showId, string metadata, uint256 price) returns(uint256 ticketId)
func (_Sticket *SticketTransactorSession) CreateTicket(showId *big.Int, metadata string, price *big.Int) (*types.Transaction, error) {
	return _Sticket.Contract.CreateTicket(&_Sticket.TransactOpts, showId, metadata, price)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Sticket *SticketTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Sticket.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Sticket *SticketSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Sticket.Contract.RawFulfillRandomWords(&_Sticket.TransactOpts, requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Sticket *SticketTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Sticket.Contract.RawFulfillRandomWords(&_Sticket.TransactOpts, requestId, randomWords)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sticket *SticketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sticket.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sticket *SticketSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sticket.Contract.RenounceOwnership(&_Sticket.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Sticket *SticketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Sticket.Contract.RenounceOwnership(&_Sticket.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Sticket *SticketTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Sticket.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Sticket *SticketSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Sticket.Contract.SafeTransferFrom(&_Sticket.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Sticket *SticketTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Sticket.Contract.SafeTransferFrom(&_Sticket.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Sticket *SticketTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Sticket.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Sticket *SticketSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Sticket.Contract.SafeTransferFrom0(&_Sticket.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Sticket *SticketTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Sticket.Contract.SafeTransferFrom0(&_Sticket.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Sticket *SticketTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Sticket.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Sticket *SticketSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Sticket.Contract.SetApprovalForAll(&_Sticket.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Sticket *SticketTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Sticket.Contract.SetApprovalForAll(&_Sticket.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Sticket *SticketTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Sticket.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Sticket *SticketSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Sticket.Contract.TransferFrom(&_Sticket.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Sticket *SticketTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Sticket.Contract.TransferFrom(&_Sticket.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sticket *SticketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Sticket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sticket *SticketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sticket.Contract.TransferOwnership(&_Sticket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Sticket *SticketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Sticket.Contract.TransferOwnership(&_Sticket.TransactOpts, newOwner)
}

// SticketApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Sticket contract.
type SticketApprovalIterator struct {
	Event *SticketApproval // Event containing the contract specifics and raw log

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
func (it *SticketApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SticketApproval)
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
		it.Event = new(SticketApproval)
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
func (it *SticketApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SticketApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SticketApproval represents a Approval event raised by the Sticket contract.
type SticketApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Sticket *SticketFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SticketApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Sticket.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SticketApprovalIterator{contract: _Sticket.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Sticket *SticketFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SticketApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Sticket.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SticketApproval)
				if err := _Sticket.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Sticket *SticketFilterer) ParseApproval(log types.Log) (*SticketApproval, error) {
	event := new(SticketApproval)
	if err := _Sticket.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SticketApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Sticket contract.
type SticketApprovalForAllIterator struct {
	Event *SticketApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SticketApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SticketApprovalForAll)
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
		it.Event = new(SticketApprovalForAll)
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
func (it *SticketApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SticketApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SticketApprovalForAll represents a ApprovalForAll event raised by the Sticket contract.
type SticketApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Sticket *SticketFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SticketApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Sticket.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SticketApprovalForAllIterator{contract: _Sticket.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Sticket *SticketFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SticketApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Sticket.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SticketApprovalForAll)
				if err := _Sticket.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Sticket *SticketFilterer) ParseApprovalForAll(log types.Log) (*SticketApprovalForAll, error) {
	event := new(SticketApprovalForAll)
	if err := _Sticket.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SticketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Sticket contract.
type SticketOwnershipTransferredIterator struct {
	Event *SticketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SticketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SticketOwnershipTransferred)
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
		it.Event = new(SticketOwnershipTransferred)
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
func (it *SticketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SticketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SticketOwnershipTransferred represents a OwnershipTransferred event raised by the Sticket contract.
type SticketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sticket *SticketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SticketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sticket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SticketOwnershipTransferredIterator{contract: _Sticket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Sticket *SticketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SticketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Sticket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SticketOwnershipTransferred)
				if err := _Sticket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Sticket *SticketFilterer) ParseOwnershipTransferred(log types.Log) (*SticketOwnershipTransferred, error) {
	event := new(SticketOwnershipTransferred)
	if err := _Sticket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SticketShowCreatedIterator is returned from FilterShowCreated and is used to iterate over the raw logs and unpacked data for ShowCreated events raised by the Sticket contract.
type SticketShowCreatedIterator struct {
	Event *SticketShowCreated // Event containing the contract specifics and raw log

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
func (it *SticketShowCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SticketShowCreated)
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
		it.Event = new(SticketShowCreated)
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
func (it *SticketShowCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SticketShowCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SticketShowCreated represents a ShowCreated event raised by the Sticket contract.
type SticketShowCreated struct {
	Id            *big.Int
	Name          string
	Date          *big.Int
	Place         string
	SpecialPrices []string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterShowCreated is a free log retrieval operation binding the contract event 0xb9ddbe6a8f97ae64a3e5114d9fbf2fb71700218f9e1fdc071e2f62c81a223f99.
//
// Solidity: event ShowCreated(uint256 indexed id, string name, uint256 date, string place, string[] specialPrices)
func (_Sticket *SticketFilterer) FilterShowCreated(opts *bind.FilterOpts, id []*big.Int) (*SticketShowCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Sticket.contract.FilterLogs(opts, "ShowCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &SticketShowCreatedIterator{contract: _Sticket.contract, event: "ShowCreated", logs: logs, sub: sub}, nil
}

// WatchShowCreated is a free log subscription operation binding the contract event 0xb9ddbe6a8f97ae64a3e5114d9fbf2fb71700218f9e1fdc071e2f62c81a223f99.
//
// Solidity: event ShowCreated(uint256 indexed id, string name, uint256 date, string place, string[] specialPrices)
func (_Sticket *SticketFilterer) WatchShowCreated(opts *bind.WatchOpts, sink chan<- *SticketShowCreated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Sticket.contract.WatchLogs(opts, "ShowCreated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SticketShowCreated)
				if err := _Sticket.contract.UnpackLog(event, "ShowCreated", log); err != nil {
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

// ParseShowCreated is a log parse operation binding the contract event 0xb9ddbe6a8f97ae64a3e5114d9fbf2fb71700218f9e1fdc071e2f62c81a223f99.
//
// Solidity: event ShowCreated(uint256 indexed id, string name, uint256 date, string place, string[] specialPrices)
func (_Sticket *SticketFilterer) ParseShowCreated(log types.Log) (*SticketShowCreated, error) {
	event := new(SticketShowCreated)
	if err := _Sticket.contract.UnpackLog(event, "ShowCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SticketTicketCreatedIterator is returned from FilterTicketCreated and is used to iterate over the raw logs and unpacked data for TicketCreated events raised by the Sticket contract.
type SticketTicketCreatedIterator struct {
	Event *SticketTicketCreated // Event containing the contract specifics and raw log

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
func (it *SticketTicketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SticketTicketCreated)
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
		it.Event = new(SticketTicketCreated)
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
func (it *SticketTicketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SticketTicketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SticketTicketCreated represents a TicketCreated event raised by the Sticket contract.
type SticketTicketCreated struct {
	Id       *big.Int
	ShowId   *big.Int
	Metadata string
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTicketCreated is a free log retrieval operation binding the contract event 0xe1240c65bce597d6828726b77db764ec72682f08444faf22990dbf1da190c0d3.
//
// Solidity: event TicketCreated(uint256 indexed id, uint256 showId, string metadata, uint256 price)
func (_Sticket *SticketFilterer) FilterTicketCreated(opts *bind.FilterOpts, id []*big.Int) (*SticketTicketCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Sticket.contract.FilterLogs(opts, "TicketCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &SticketTicketCreatedIterator{contract: _Sticket.contract, event: "TicketCreated", logs: logs, sub: sub}, nil
}

// WatchTicketCreated is a free log subscription operation binding the contract event 0xe1240c65bce597d6828726b77db764ec72682f08444faf22990dbf1da190c0d3.
//
// Solidity: event TicketCreated(uint256 indexed id, uint256 showId, string metadata, uint256 price)
func (_Sticket *SticketFilterer) WatchTicketCreated(opts *bind.WatchOpts, sink chan<- *SticketTicketCreated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Sticket.contract.WatchLogs(opts, "TicketCreated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SticketTicketCreated)
				if err := _Sticket.contract.UnpackLog(event, "TicketCreated", log); err != nil {
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

// ParseTicketCreated is a log parse operation binding the contract event 0xe1240c65bce597d6828726b77db764ec72682f08444faf22990dbf1da190c0d3.
//
// Solidity: event TicketCreated(uint256 indexed id, uint256 showId, string metadata, uint256 price)
func (_Sticket *SticketFilterer) ParseTicketCreated(log types.Log) (*SticketTicketCreated, error) {
	event := new(SticketTicketCreated)
	if err := _Sticket.contract.UnpackLog(event, "TicketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SticketTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Sticket contract.
type SticketTransferIterator struct {
	Event *SticketTransfer // Event containing the contract specifics and raw log

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
func (it *SticketTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SticketTransfer)
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
		it.Event = new(SticketTransfer)
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
func (it *SticketTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SticketTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SticketTransfer represents a Transfer event raised by the Sticket contract.
type SticketTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Sticket *SticketFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SticketTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Sticket.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SticketTransferIterator{contract: _Sticket.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Sticket *SticketFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SticketTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Sticket.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SticketTransfer)
				if err := _Sticket.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Sticket *SticketFilterer) ParseTransfer(log types.Log) (*SticketTransfer, error) {
	event := new(SticketTransfer)
	if err := _Sticket.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
