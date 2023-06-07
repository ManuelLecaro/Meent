// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package meent

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

// MeentMetaData contains all meta data concerning the Meent contract.
var MeentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"}],\"name\":\"EventCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"apiresponse\",\"type\":\"uint256\"}],\"name\":\"RequestTicketCreation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"_setTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apiresponse\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"}],\"name\":\"buyTicket\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTickets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"realOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"}],\"name\":\"createEvent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"eventOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"events\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTickets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ticketsSold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_volume\",\"type\":\"uint256\"}],\"name\":\"fulfill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ticketId\",\"type\":\"uint256\"}],\"name\":\"refundTicket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferEventOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"eventName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalTickets\",\"type\":\"uint256\"}],\"name\":\"updateEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MeentABI is the input ABI used to generate the binding from.
// Deprecated: Use MeentMetaData.ABI instead.
var MeentABI = MeentMetaData.ABI

// Meent is an auto generated Go binding around an Ethereum contract.
type Meent struct {
	MeentCaller     // Read-only binding to the contract
	MeentTransactor // Write-only binding to the contract
	MeentFilterer   // Log filterer for contract events
}

// MeentCaller is an auto generated read-only Go binding around an Ethereum contract.
type MeentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MeentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MeentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MeentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MeentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MeentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MeentSession struct {
	Contract     *Meent            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MeentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MeentCallerSession struct {
	Contract *MeentCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MeentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MeentTransactorSession struct {
	Contract     *MeentTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MeentRaw is an auto generated low-level Go binding around an Ethereum contract.
type MeentRaw struct {
	Contract *Meent // Generic contract binding to access the raw methods on
}

// MeentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MeentCallerRaw struct {
	Contract *MeentCaller // Generic read-only contract binding to access the raw methods on
}

// MeentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MeentTransactorRaw struct {
	Contract *MeentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMeent creates a new instance of Meent, bound to a specific deployed contract.
func NewMeent(address common.Address, backend bind.ContractBackend) (*Meent, error) {
	contract, err := bindMeent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Meent{MeentCaller: MeentCaller{contract: contract}, MeentTransactor: MeentTransactor{contract: contract}, MeentFilterer: MeentFilterer{contract: contract}}, nil
}

// NewMeentCaller creates a new read-only instance of Meent, bound to a specific deployed contract.
func NewMeentCaller(address common.Address, caller bind.ContractCaller) (*MeentCaller, error) {
	contract, err := bindMeent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MeentCaller{contract: contract}, nil
}

// NewMeentTransactor creates a new write-only instance of Meent, bound to a specific deployed contract.
func NewMeentTransactor(address common.Address, transactor bind.ContractTransactor) (*MeentTransactor, error) {
	contract, err := bindMeent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MeentTransactor{contract: contract}, nil
}

// NewMeentFilterer creates a new log filterer instance of Meent, bound to a specific deployed contract.
func NewMeentFilterer(address common.Address, filterer bind.ContractFilterer) (*MeentFilterer, error) {
	contract, err := bindMeent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MeentFilterer{contract: contract}, nil
}

// bindMeent binds a generic wrapper to an already deployed contract.
func bindMeent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MeentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Meent *MeentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Meent.Contract.MeentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Meent *MeentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meent.Contract.MeentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Meent *MeentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Meent.Contract.MeentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Meent *MeentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Meent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Meent *MeentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Meent *MeentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Meent.Contract.contract.Transact(opts, method, params...)
}

// Apiresponse is a free data retrieval call binding the contract method 0xcb7cc4e9.
//
// Solidity: function apiresponse() view returns(uint256)
func (_Meent *MeentCaller) Apiresponse(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Meent.contract.Call(opts, &out, "apiresponse")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Apiresponse is a free data retrieval call binding the contract method 0xcb7cc4e9.
//
// Solidity: function apiresponse() view returns(uint256)
func (_Meent *MeentSession) Apiresponse() (*big.Int, error) {
	return _Meent.Contract.Apiresponse(&_Meent.CallOpts)
}

// Apiresponse is a free data retrieval call binding the contract method 0xcb7cc4e9.
//
// Solidity: function apiresponse() view returns(uint256)
func (_Meent *MeentCallerSession) Apiresponse() (*big.Int, error) {
	return _Meent.Contract.Apiresponse(&_Meent.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Meent *MeentCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Meent.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Meent *MeentSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Meent.Contract.BalanceOf(&_Meent.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Meent *MeentCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Meent.Contract.BalanceOf(&_Meent.CallOpts, owner)
}

// EventOwners is a free data retrieval call binding the contract method 0x336bca15.
//
// Solidity: function eventOwners(uint256 ) view returns(address)
func (_Meent *MeentCaller) EventOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Meent.contract.Call(opts, &out, "eventOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EventOwners is a free data retrieval call binding the contract method 0x336bca15.
//
// Solidity: function eventOwners(uint256 ) view returns(address)
func (_Meent *MeentSession) EventOwners(arg0 *big.Int) (common.Address, error) {
	return _Meent.Contract.EventOwners(&_Meent.CallOpts, arg0)
}

// EventOwners is a free data retrieval call binding the contract method 0x336bca15.
//
// Solidity: function eventOwners(uint256 ) view returns(address)
func (_Meent *MeentCallerSession) EventOwners(arg0 *big.Int) (common.Address, error) {
	return _Meent.Contract.EventOwners(&_Meent.CallOpts, arg0)
}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint256 eventId, string eventName, uint256 price, uint256 totalTickets, uint256 ticketsSold, uint256 totalRewards)
func (_Meent *MeentCaller) Events(opts *bind.CallOpts, arg0 *big.Int) (struct {
	EventId      *big.Int
	EventName    string
	Price        *big.Int
	TotalTickets *big.Int
	TicketsSold  *big.Int
	TotalRewards *big.Int
}, error) {
	var out []interface{}
	err := _Meent.contract.Call(opts, &out, "events", arg0)

	outstruct := new(struct {
		EventId      *big.Int
		EventName    string
		Price        *big.Int
		TotalTickets *big.Int
		TicketsSold  *big.Int
		TotalRewards *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EventId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EventName = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalTickets = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TicketsSold = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalRewards = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint256 eventId, string eventName, uint256 price, uint256 totalTickets, uint256 ticketsSold, uint256 totalRewards)
func (_Meent *MeentSession) Events(arg0 *big.Int) (struct {
	EventId      *big.Int
	EventName    string
	Price        *big.Int
	TotalTickets *big.Int
	TicketsSold  *big.Int
	TotalRewards *big.Int
}, error) {
	return _Meent.Contract.Events(&_Meent.CallOpts, arg0)
}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint256 eventId, string eventName, uint256 price, uint256 totalTickets, uint256 ticketsSold, uint256 totalRewards)
func (_Meent *MeentCallerSession) Events(arg0 *big.Int) (struct {
	EventId      *big.Int
	EventName    string
	Price        *big.Int
	TotalTickets *big.Int
	TicketsSold  *big.Int
	TotalRewards *big.Int
}, error) {
	return _Meent.Contract.Events(&_Meent.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Meent *MeentCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Meent.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Meent *MeentSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Meent.Contract.GetApproved(&_Meent.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Meent *MeentCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Meent.Contract.GetApproved(&_Meent.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Meent *MeentCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Meent.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Meent *MeentSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Meent.Contract.IsApprovedForAll(&_Meent.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Meent *MeentCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Meent.Contract.IsApprovedForAll(&_Meent.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Meent *MeentCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Meent.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Meent *MeentSession) Name() (string, error) {
	return _Meent.Contract.Name(&_Meent.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Meent *MeentCallerSession) Name() (string, error) {
	return _Meent.Contract.Name(&_Meent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Meent *MeentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Meent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Meent *MeentSession) Owner() (common.Address, error) {
	return _Meent.Contract.Owner(&_Meent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Meent *MeentCallerSession) Owner() (common.Address, error) {
	return _Meent.Contract.Owner(&_Meent.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Meent *MeentCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Meent.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Meent *MeentSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Meent.Contract.OwnerOf(&_Meent.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Meent *MeentCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Meent.Contract.OwnerOf(&_Meent.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Meent *MeentCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Meent.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Meent *MeentSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Meent.Contract.SupportsInterface(&_Meent.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Meent *MeentCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Meent.Contract.SupportsInterface(&_Meent.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Meent *MeentCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Meent.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Meent *MeentSession) Symbol() (string, error) {
	return _Meent.Contract.Symbol(&_Meent.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Meent *MeentCallerSession) Symbol() (string, error) {
	return _Meent.Contract.Symbol(&_Meent.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Meent *MeentCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Meent.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Meent *MeentSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Meent.Contract.TokenURI(&_Meent.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Meent *MeentCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Meent.Contract.TokenURI(&_Meent.CallOpts, tokenId)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x01538868.
//
// Solidity: function _setTokenURI(uint256 tokenId, string _tokenURI) returns()
func (_Meent *MeentTransactor) SetTokenURI(opts *bind.TransactOpts, tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "_setTokenURI", tokenId, _tokenURI)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x01538868.
//
// Solidity: function _setTokenURI(uint256 tokenId, string _tokenURI) returns()
func (_Meent *MeentSession) SetTokenURI(tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Meent.Contract.SetTokenURI(&_Meent.TransactOpts, tokenId, _tokenURI)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x01538868.
//
// Solidity: function _setTokenURI(uint256 tokenId, string _tokenURI) returns()
func (_Meent *MeentTransactorSession) SetTokenURI(tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Meent.Contract.SetTokenURI(&_Meent.TransactOpts, tokenId, _tokenURI)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Meent *MeentTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Meent *MeentSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.Approve(&_Meent.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Meent *MeentTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.Approve(&_Meent.TransactOpts, to, tokenId)
}

// BuyTicket is a paid mutator transaction binding the contract method 0x67dd74ca.
//
// Solidity: function buyTicket(uint256 eventId) payable returns()
func (_Meent *MeentTransactor) BuyTicket(opts *bind.TransactOpts, eventId *big.Int) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "buyTicket", eventId)
}

// BuyTicket is a paid mutator transaction binding the contract method 0x67dd74ca.
//
// Solidity: function buyTicket(uint256 eventId) payable returns()
func (_Meent *MeentSession) BuyTicket(eventId *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.BuyTicket(&_Meent.TransactOpts, eventId)
}

// BuyTicket is a paid mutator transaction binding the contract method 0x67dd74ca.
//
// Solidity: function buyTicket(uint256 eventId) payable returns()
func (_Meent *MeentTransactorSession) BuyTicket(eventId *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.BuyTicket(&_Meent.TransactOpts, eventId)
}

// CreateEvent is a paid mutator transaction binding the contract method 0x4eecf711.
//
// Solidity: function createEvent(string eventName, uint256 price, uint256 totalTickets, address realOwner, uint256 totalRewards) returns(uint256)
func (_Meent *MeentTransactor) CreateEvent(opts *bind.TransactOpts, eventName string, price *big.Int, totalTickets *big.Int, realOwner common.Address, totalRewards *big.Int) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "createEvent", eventName, price, totalTickets, realOwner, totalRewards)
}

// CreateEvent is a paid mutator transaction binding the contract method 0x4eecf711.
//
// Solidity: function createEvent(string eventName, uint256 price, uint256 totalTickets, address realOwner, uint256 totalRewards) returns(uint256)
func (_Meent *MeentSession) CreateEvent(eventName string, price *big.Int, totalTickets *big.Int, realOwner common.Address, totalRewards *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.CreateEvent(&_Meent.TransactOpts, eventName, price, totalTickets, realOwner, totalRewards)
}

// CreateEvent is a paid mutator transaction binding the contract method 0x4eecf711.
//
// Solidity: function createEvent(string eventName, uint256 price, uint256 totalTickets, address realOwner, uint256 totalRewards) returns(uint256)
func (_Meent *MeentTransactorSession) CreateEvent(eventName string, price *big.Int, totalTickets *big.Int, realOwner common.Address, totalRewards *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.CreateEvent(&_Meent.TransactOpts, eventName, price, totalTickets, realOwner, totalRewards)
}

// Fulfill is a paid mutator transaction binding the contract method 0x4357855e.
//
// Solidity: function fulfill(bytes32 _requestId, uint256 _volume) returns()
func (_Meent *MeentTransactor) Fulfill(opts *bind.TransactOpts, _requestId [32]byte, _volume *big.Int) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "fulfill", _requestId, _volume)
}

// Fulfill is a paid mutator transaction binding the contract method 0x4357855e.
//
// Solidity: function fulfill(bytes32 _requestId, uint256 _volume) returns()
func (_Meent *MeentSession) Fulfill(_requestId [32]byte, _volume *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.Fulfill(&_Meent.TransactOpts, _requestId, _volume)
}

// Fulfill is a paid mutator transaction binding the contract method 0x4357855e.
//
// Solidity: function fulfill(bytes32 _requestId, uint256 _volume) returns()
func (_Meent *MeentTransactorSession) Fulfill(_requestId [32]byte, _volume *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.Fulfill(&_Meent.TransactOpts, _requestId, _volume)
}

// RefundTicket is a paid mutator transaction binding the contract method 0x1d196a7f.
//
// Solidity: function refundTicket(uint256 eventId, address buyer, uint256 ticketId) returns()
func (_Meent *MeentTransactor) RefundTicket(opts *bind.TransactOpts, eventId *big.Int, buyer common.Address, ticketId *big.Int) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "refundTicket", eventId, buyer, ticketId)
}

// RefundTicket is a paid mutator transaction binding the contract method 0x1d196a7f.
//
// Solidity: function refundTicket(uint256 eventId, address buyer, uint256 ticketId) returns()
func (_Meent *MeentSession) RefundTicket(eventId *big.Int, buyer common.Address, ticketId *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.RefundTicket(&_Meent.TransactOpts, eventId, buyer, ticketId)
}

// RefundTicket is a paid mutator transaction binding the contract method 0x1d196a7f.
//
// Solidity: function refundTicket(uint256 eventId, address buyer, uint256 ticketId) returns()
func (_Meent *MeentTransactorSession) RefundTicket(eventId *big.Int, buyer common.Address, ticketId *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.RefundTicket(&_Meent.TransactOpts, eventId, buyer, ticketId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Meent *MeentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Meent *MeentSession) RenounceOwnership() (*types.Transaction, error) {
	return _Meent.Contract.RenounceOwnership(&_Meent.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Meent *MeentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Meent.Contract.RenounceOwnership(&_Meent.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Meent *MeentTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Meent *MeentSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.SafeTransferFrom(&_Meent.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Meent *MeentTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.SafeTransferFrom(&_Meent.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Meent *MeentTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Meent *MeentSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Meent.Contract.SafeTransferFrom0(&_Meent.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Meent *MeentTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Meent.Contract.SafeTransferFrom0(&_Meent.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Meent *MeentTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Meent *MeentSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Meent.Contract.SetApprovalForAll(&_Meent.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Meent *MeentTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Meent.Contract.SetApprovalForAll(&_Meent.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Meent *MeentTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "setBaseURI", baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Meent *MeentSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _Meent.Contract.SetBaseURI(&_Meent.TransactOpts, baseURI_)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI_) returns()
func (_Meent *MeentTransactorSession) SetBaseURI(baseURI_ string) (*types.Transaction, error) {
	return _Meent.Contract.SetBaseURI(&_Meent.TransactOpts, baseURI_)
}

// TransferEventOwnership is a paid mutator transaction binding the contract method 0x86f4c772.
//
// Solidity: function transferEventOwnership(uint256 eventId, address newOwner) returns()
func (_Meent *MeentTransactor) TransferEventOwnership(opts *bind.TransactOpts, eventId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "transferEventOwnership", eventId, newOwner)
}

// TransferEventOwnership is a paid mutator transaction binding the contract method 0x86f4c772.
//
// Solidity: function transferEventOwnership(uint256 eventId, address newOwner) returns()
func (_Meent *MeentSession) TransferEventOwnership(eventId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _Meent.Contract.TransferEventOwnership(&_Meent.TransactOpts, eventId, newOwner)
}

// TransferEventOwnership is a paid mutator transaction binding the contract method 0x86f4c772.
//
// Solidity: function transferEventOwnership(uint256 eventId, address newOwner) returns()
func (_Meent *MeentTransactorSession) TransferEventOwnership(eventId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _Meent.Contract.TransferEventOwnership(&_Meent.TransactOpts, eventId, newOwner)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Meent *MeentTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Meent *MeentSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.TransferFrom(&_Meent.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Meent *MeentTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.TransferFrom(&_Meent.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Meent *MeentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Meent *MeentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Meent.Contract.TransferOwnership(&_Meent.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Meent *MeentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Meent.Contract.TransferOwnership(&_Meent.TransactOpts, newOwner)
}

// UpdateEvent is a paid mutator transaction binding the contract method 0xa77d5639.
//
// Solidity: function updateEvent(uint256 eventId, string eventName, uint256 price, uint256 totalTickets) returns()
func (_Meent *MeentTransactor) UpdateEvent(opts *bind.TransactOpts, eventId *big.Int, eventName string, price *big.Int, totalTickets *big.Int) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "updateEvent", eventId, eventName, price, totalTickets)
}

// UpdateEvent is a paid mutator transaction binding the contract method 0xa77d5639.
//
// Solidity: function updateEvent(uint256 eventId, string eventName, uint256 price, uint256 totalTickets) returns()
func (_Meent *MeentSession) UpdateEvent(eventId *big.Int, eventName string, price *big.Int, totalTickets *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.UpdateEvent(&_Meent.TransactOpts, eventId, eventName, price, totalTickets)
}

// UpdateEvent is a paid mutator transaction binding the contract method 0xa77d5639.
//
// Solidity: function updateEvent(uint256 eventId, string eventName, uint256 price, uint256 totalTickets) returns()
func (_Meent *MeentTransactorSession) UpdateEvent(eventId *big.Int, eventName string, price *big.Int, totalTickets *big.Int) (*types.Transaction, error) {
	return _Meent.Contract.UpdateEvent(&_Meent.TransactOpts, eventId, eventName, price, totalTickets)
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_Meent *MeentTransactor) WithdrawLink(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meent.contract.Transact(opts, "withdrawLink")
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_Meent *MeentSession) WithdrawLink() (*types.Transaction, error) {
	return _Meent.Contract.WithdrawLink(&_Meent.TransactOpts)
}

// WithdrawLink is a paid mutator transaction binding the contract method 0x8dc654a2.
//
// Solidity: function withdrawLink() returns()
func (_Meent *MeentTransactorSession) WithdrawLink() (*types.Transaction, error) {
	return _Meent.Contract.WithdrawLink(&_Meent.TransactOpts)
}

// MeentApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Meent contract.
type MeentApprovalIterator struct {
	Event *MeentApproval // Event containing the contract specifics and raw log

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
func (it *MeentApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MeentApproval)
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
		it.Event = new(MeentApproval)
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
func (it *MeentApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MeentApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MeentApproval represents a Approval event raised by the Meent contract.
type MeentApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Meent *MeentFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MeentApprovalIterator, error) {

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

	logs, sub, err := _Meent.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MeentApprovalIterator{contract: _Meent.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Meent *MeentFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MeentApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Meent.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MeentApproval)
				if err := _Meent.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Meent *MeentFilterer) ParseApproval(log types.Log) (*MeentApproval, error) {
	event := new(MeentApproval)
	if err := _Meent.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MeentApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Meent contract.
type MeentApprovalForAllIterator struct {
	Event *MeentApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MeentApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MeentApprovalForAll)
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
		it.Event = new(MeentApprovalForAll)
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
func (it *MeentApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MeentApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MeentApprovalForAll represents a ApprovalForAll event raised by the Meent contract.
type MeentApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Meent *MeentFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MeentApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Meent.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MeentApprovalForAllIterator{contract: _Meent.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Meent *MeentFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MeentApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Meent.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MeentApprovalForAll)
				if err := _Meent.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Meent *MeentFilterer) ParseApprovalForAll(log types.Log) (*MeentApprovalForAll, error) {
	event := new(MeentApprovalForAll)
	if err := _Meent.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MeentChainlinkCancelledIterator is returned from FilterChainlinkCancelled and is used to iterate over the raw logs and unpacked data for ChainlinkCancelled events raised by the Meent contract.
type MeentChainlinkCancelledIterator struct {
	Event *MeentChainlinkCancelled // Event containing the contract specifics and raw log

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
func (it *MeentChainlinkCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MeentChainlinkCancelled)
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
		it.Event = new(MeentChainlinkCancelled)
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
func (it *MeentChainlinkCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MeentChainlinkCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MeentChainlinkCancelled represents a ChainlinkCancelled event raised by the Meent contract.
type MeentChainlinkCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkCancelled is a free log retrieval operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_Meent *MeentFilterer) FilterChainlinkCancelled(opts *bind.FilterOpts, id [][32]byte) (*MeentChainlinkCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Meent.contract.FilterLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &MeentChainlinkCancelledIterator{contract: _Meent.contract, event: "ChainlinkCancelled", logs: logs, sub: sub}, nil
}

// WatchChainlinkCancelled is a free log subscription operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_Meent *MeentFilterer) WatchChainlinkCancelled(opts *bind.WatchOpts, sink chan<- *MeentChainlinkCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Meent.contract.WatchLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MeentChainlinkCancelled)
				if err := _Meent.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
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

// ParseChainlinkCancelled is a log parse operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_Meent *MeentFilterer) ParseChainlinkCancelled(log types.Log) (*MeentChainlinkCancelled, error) {
	event := new(MeentChainlinkCancelled)
	if err := _Meent.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MeentChainlinkFulfilledIterator is returned from FilterChainlinkFulfilled and is used to iterate over the raw logs and unpacked data for ChainlinkFulfilled events raised by the Meent contract.
type MeentChainlinkFulfilledIterator struct {
	Event *MeentChainlinkFulfilled // Event containing the contract specifics and raw log

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
func (it *MeentChainlinkFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MeentChainlinkFulfilled)
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
		it.Event = new(MeentChainlinkFulfilled)
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
func (it *MeentChainlinkFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MeentChainlinkFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MeentChainlinkFulfilled represents a ChainlinkFulfilled event raised by the Meent contract.
type MeentChainlinkFulfilled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkFulfilled is a free log retrieval operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_Meent *MeentFilterer) FilterChainlinkFulfilled(opts *bind.FilterOpts, id [][32]byte) (*MeentChainlinkFulfilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Meent.contract.FilterLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return &MeentChainlinkFulfilledIterator{contract: _Meent.contract, event: "ChainlinkFulfilled", logs: logs, sub: sub}, nil
}

// WatchChainlinkFulfilled is a free log subscription operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_Meent *MeentFilterer) WatchChainlinkFulfilled(opts *bind.WatchOpts, sink chan<- *MeentChainlinkFulfilled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Meent.contract.WatchLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MeentChainlinkFulfilled)
				if err := _Meent.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
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

// ParseChainlinkFulfilled is a log parse operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_Meent *MeentFilterer) ParseChainlinkFulfilled(log types.Log) (*MeentChainlinkFulfilled, error) {
	event := new(MeentChainlinkFulfilled)
	if err := _Meent.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MeentChainlinkRequestedIterator is returned from FilterChainlinkRequested and is used to iterate over the raw logs and unpacked data for ChainlinkRequested events raised by the Meent contract.
type MeentChainlinkRequestedIterator struct {
	Event *MeentChainlinkRequested // Event containing the contract specifics and raw log

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
func (it *MeentChainlinkRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MeentChainlinkRequested)
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
		it.Event = new(MeentChainlinkRequested)
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
func (it *MeentChainlinkRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MeentChainlinkRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MeentChainlinkRequested represents a ChainlinkRequested event raised by the Meent contract.
type MeentChainlinkRequested struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkRequested is a free log retrieval operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_Meent *MeentFilterer) FilterChainlinkRequested(opts *bind.FilterOpts, id [][32]byte) (*MeentChainlinkRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Meent.contract.FilterLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return &MeentChainlinkRequestedIterator{contract: _Meent.contract, event: "ChainlinkRequested", logs: logs, sub: sub}, nil
}

// WatchChainlinkRequested is a free log subscription operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_Meent *MeentFilterer) WatchChainlinkRequested(opts *bind.WatchOpts, sink chan<- *MeentChainlinkRequested, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Meent.contract.WatchLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MeentChainlinkRequested)
				if err := _Meent.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
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

// ParseChainlinkRequested is a log parse operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_Meent *MeentFilterer) ParseChainlinkRequested(log types.Log) (*MeentChainlinkRequested, error) {
	event := new(MeentChainlinkRequested)
	if err := _Meent.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MeentEventCreatedIterator is returned from FilterEventCreated and is used to iterate over the raw logs and unpacked data for EventCreated events raised by the Meent contract.
type MeentEventCreatedIterator struct {
	Event *MeentEventCreated // Event containing the contract specifics and raw log

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
func (it *MeentEventCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MeentEventCreated)
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
		it.Event = new(MeentEventCreated)
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
func (it *MeentEventCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MeentEventCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MeentEventCreated represents a EventCreated event raised by the Meent contract.
type MeentEventCreated struct {
	EventId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEventCreated is a free log retrieval operation binding the contract event 0x748ff5cfd7b66bdd8e21581c2a864756fe7085378c0003183e579bfafba28325.
//
// Solidity: event EventCreated(uint256 eventId)
func (_Meent *MeentFilterer) FilterEventCreated(opts *bind.FilterOpts) (*MeentEventCreatedIterator, error) {

	logs, sub, err := _Meent.contract.FilterLogs(opts, "EventCreated")
	if err != nil {
		return nil, err
	}
	return &MeentEventCreatedIterator{contract: _Meent.contract, event: "EventCreated", logs: logs, sub: sub}, nil
}

// WatchEventCreated is a free log subscription operation binding the contract event 0x748ff5cfd7b66bdd8e21581c2a864756fe7085378c0003183e579bfafba28325.
//
// Solidity: event EventCreated(uint256 eventId)
func (_Meent *MeentFilterer) WatchEventCreated(opts *bind.WatchOpts, sink chan<- *MeentEventCreated) (event.Subscription, error) {

	logs, sub, err := _Meent.contract.WatchLogs(opts, "EventCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MeentEventCreated)
				if err := _Meent.contract.UnpackLog(event, "EventCreated", log); err != nil {
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

// ParseEventCreated is a log parse operation binding the contract event 0x748ff5cfd7b66bdd8e21581c2a864756fe7085378c0003183e579bfafba28325.
//
// Solidity: event EventCreated(uint256 eventId)
func (_Meent *MeentFilterer) ParseEventCreated(log types.Log) (*MeentEventCreated, error) {
	event := new(MeentEventCreated)
	if err := _Meent.contract.UnpackLog(event, "EventCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MeentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Meent contract.
type MeentOwnershipTransferredIterator struct {
	Event *MeentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MeentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MeentOwnershipTransferred)
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
		it.Event = new(MeentOwnershipTransferred)
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
func (it *MeentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MeentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MeentOwnershipTransferred represents a OwnershipTransferred event raised by the Meent contract.
type MeentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Meent *MeentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MeentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Meent.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MeentOwnershipTransferredIterator{contract: _Meent.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Meent *MeentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MeentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Meent.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MeentOwnershipTransferred)
				if err := _Meent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Meent *MeentFilterer) ParseOwnershipTransferred(log types.Log) (*MeentOwnershipTransferred, error) {
	event := new(MeentOwnershipTransferred)
	if err := _Meent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MeentRequestTicketCreationIterator is returned from FilterRequestTicketCreation and is used to iterate over the raw logs and unpacked data for RequestTicketCreation events raised by the Meent contract.
type MeentRequestTicketCreationIterator struct {
	Event *MeentRequestTicketCreation // Event containing the contract specifics and raw log

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
func (it *MeentRequestTicketCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MeentRequestTicketCreation)
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
		it.Event = new(MeentRequestTicketCreation)
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
func (it *MeentRequestTicketCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MeentRequestTicketCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MeentRequestTicketCreation represents a RequestTicketCreation event raised by the Meent contract.
type MeentRequestTicketCreation struct {
	RequestId   [32]byte
	Apiresponse *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRequestTicketCreation is a free log retrieval operation binding the contract event 0x5738d88ae38b9b6206e6cbe055f26ee438a0e03653068c4b3a4d48b61f28ad44.
//
// Solidity: event RequestTicketCreation(bytes32 indexed requestId, uint256 apiresponse)
func (_Meent *MeentFilterer) FilterRequestTicketCreation(opts *bind.FilterOpts, requestId [][32]byte) (*MeentRequestTicketCreationIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Meent.contract.FilterLogs(opts, "RequestTicketCreation", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &MeentRequestTicketCreationIterator{contract: _Meent.contract, event: "RequestTicketCreation", logs: logs, sub: sub}, nil
}

// WatchRequestTicketCreation is a free log subscription operation binding the contract event 0x5738d88ae38b9b6206e6cbe055f26ee438a0e03653068c4b3a4d48b61f28ad44.
//
// Solidity: event RequestTicketCreation(bytes32 indexed requestId, uint256 apiresponse)
func (_Meent *MeentFilterer) WatchRequestTicketCreation(opts *bind.WatchOpts, sink chan<- *MeentRequestTicketCreation, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Meent.contract.WatchLogs(opts, "RequestTicketCreation", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MeentRequestTicketCreation)
				if err := _Meent.contract.UnpackLog(event, "RequestTicketCreation", log); err != nil {
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

// ParseRequestTicketCreation is a log parse operation binding the contract event 0x5738d88ae38b9b6206e6cbe055f26ee438a0e03653068c4b3a4d48b61f28ad44.
//
// Solidity: event RequestTicketCreation(bytes32 indexed requestId, uint256 apiresponse)
func (_Meent *MeentFilterer) ParseRequestTicketCreation(log types.Log) (*MeentRequestTicketCreation, error) {
	event := new(MeentRequestTicketCreation)
	if err := _Meent.contract.UnpackLog(event, "RequestTicketCreation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MeentTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Meent contract.
type MeentTransferIterator struct {
	Event *MeentTransfer // Event containing the contract specifics and raw log

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
func (it *MeentTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MeentTransfer)
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
		it.Event = new(MeentTransfer)
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
func (it *MeentTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MeentTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MeentTransfer represents a Transfer event raised by the Meent contract.
type MeentTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Meent *MeentFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MeentTransferIterator, error) {

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

	logs, sub, err := _Meent.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MeentTransferIterator{contract: _Meent.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Meent *MeentFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MeentTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Meent.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MeentTransfer)
				if err := _Meent.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Meent *MeentFilterer) ParseTransfer(log types.Log) (*MeentTransfer, error) {
	event := new(MeentTransfer)
	if err := _Meent.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
