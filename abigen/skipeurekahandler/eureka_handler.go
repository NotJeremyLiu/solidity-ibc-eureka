// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package skipeurekahandler

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

// IEurekaHandlerFees is an auto generated low-level Go binding around an user-defined struct.
type IEurekaHandlerFees struct {
	RelayFee          *big.Int
	RelayFeeRecipient common.Address
	QuoteExpiry       uint64
}

// IEurekaHandlerTransferParams is an auto generated low-level Go binding around an user-defined struct.
type IEurekaHandlerTransferParams struct {
	Token            common.Address
	Recipient        string
	SourceClient     string
	DestPort         string
	TimeoutTimestamp uint64
	Memo             string
}

// EurekaHandlerMetaData contains all meta data concerning the EurekaHandler contract.
var EurekaHandlerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ics20Transfer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ics20Transfer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_swapRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lbtcVoucher\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_lbtc\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lbtc\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lbtcVoucher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lombardTransfer\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transferParams\",\"type\":\"tuple\",\"internalType\":\"structIEurekaHandler.TransferParams\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"fees\",\"type\":\"tuple\",\"internalType\":\"structIEurekaHandler.Fees\",\"components\":[{\"name\":\"relayFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"relayFeeRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quoteExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapAndTransfer\",\"inputs\":[{\"name\":\"swapInputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"swapInputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"swapCalldata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"transferParams\",\"type\":\"tuple\",\"internalType\":\"structIEurekaHandler.TransferParams\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"fees\",\"type\":\"tuple\",\"internalType\":\"structIEurekaHandler.Fees\",\"components\":[{\"name\":\"relayFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"relayFeeRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quoteExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transferParams\",\"type\":\"tuple\",\"internalType\":\"structIEurekaHandler.TransferParams\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sourceClient\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"destPort\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"memo\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"fees\",\"type\":\"tuple\",\"internalType\":\"structIEurekaHandler.Fees\",\"components\":[{\"name\":\"relayFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"relayFeeRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quoteExpiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"relayFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// EurekaHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use EurekaHandlerMetaData.ABI instead.
var EurekaHandlerABI = EurekaHandlerMetaData.ABI

// EurekaHandler is an auto generated Go binding around an Ethereum contract.
type EurekaHandler struct {
	EurekaHandlerCaller     // Read-only binding to the contract
	EurekaHandlerTransactor // Write-only binding to the contract
	EurekaHandlerFilterer   // Log filterer for contract events
}

// EurekaHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EurekaHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EurekaHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EurekaHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EurekaHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EurekaHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EurekaHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EurekaHandlerSession struct {
	Contract     *EurekaHandler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EurekaHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EurekaHandlerCallerSession struct {
	Contract *EurekaHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EurekaHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EurekaHandlerTransactorSession struct {
	Contract     *EurekaHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EurekaHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EurekaHandlerRaw struct {
	Contract *EurekaHandler // Generic contract binding to access the raw methods on
}

// EurekaHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EurekaHandlerCallerRaw struct {
	Contract *EurekaHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// EurekaHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EurekaHandlerTransactorRaw struct {
	Contract *EurekaHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEurekaHandler creates a new instance of EurekaHandler, bound to a specific deployed contract.
func NewEurekaHandler(address common.Address, backend bind.ContractBackend) (*EurekaHandler, error) {
	contract, err := bindEurekaHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EurekaHandler{EurekaHandlerCaller: EurekaHandlerCaller{contract: contract}, EurekaHandlerTransactor: EurekaHandlerTransactor{contract: contract}, EurekaHandlerFilterer: EurekaHandlerFilterer{contract: contract}}, nil
}

// NewEurekaHandlerCaller creates a new read-only instance of EurekaHandler, bound to a specific deployed contract.
func NewEurekaHandlerCaller(address common.Address, caller bind.ContractCaller) (*EurekaHandlerCaller, error) {
	contract, err := bindEurekaHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EurekaHandlerCaller{contract: contract}, nil
}

// NewEurekaHandlerTransactor creates a new write-only instance of EurekaHandler, bound to a specific deployed contract.
func NewEurekaHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*EurekaHandlerTransactor, error) {
	contract, err := bindEurekaHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EurekaHandlerTransactor{contract: contract}, nil
}

// NewEurekaHandlerFilterer creates a new log filterer instance of EurekaHandler, bound to a specific deployed contract.
func NewEurekaHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*EurekaHandlerFilterer, error) {
	contract, err := bindEurekaHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EurekaHandlerFilterer{contract: contract}, nil
}

// bindEurekaHandler binds a generic wrapper to an already deployed contract.
func bindEurekaHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EurekaHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EurekaHandler *EurekaHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EurekaHandler.Contract.EurekaHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EurekaHandler *EurekaHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EurekaHandler.Contract.EurekaHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EurekaHandler *EurekaHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EurekaHandler.Contract.EurekaHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EurekaHandler *EurekaHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EurekaHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EurekaHandler *EurekaHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EurekaHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EurekaHandler *EurekaHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EurekaHandler.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EurekaHandler *EurekaHandlerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EurekaHandler.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EurekaHandler *EurekaHandlerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _EurekaHandler.Contract.UPGRADEINTERFACEVERSION(&_EurekaHandler.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_EurekaHandler *EurekaHandlerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _EurekaHandler.Contract.UPGRADEINTERFACEVERSION(&_EurekaHandler.CallOpts)
}

// Ics20Transfer is a free data retrieval call binding the contract method 0x25a2feb8.
//
// Solidity: function ics20Transfer() view returns(address)
func (_EurekaHandler *EurekaHandlerCaller) Ics20Transfer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EurekaHandler.contract.Call(opts, &out, "ics20Transfer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ics20Transfer is a free data retrieval call binding the contract method 0x25a2feb8.
//
// Solidity: function ics20Transfer() view returns(address)
func (_EurekaHandler *EurekaHandlerSession) Ics20Transfer() (common.Address, error) {
	return _EurekaHandler.Contract.Ics20Transfer(&_EurekaHandler.CallOpts)
}

// Ics20Transfer is a free data retrieval call binding the contract method 0x25a2feb8.
//
// Solidity: function ics20Transfer() view returns(address)
func (_EurekaHandler *EurekaHandlerCallerSession) Ics20Transfer() (common.Address, error) {
	return _EurekaHandler.Contract.Ics20Transfer(&_EurekaHandler.CallOpts)
}

// Lbtc is a free data retrieval call binding the contract method 0x68b3c910.
//
// Solidity: function lbtc() view returns(address)
func (_EurekaHandler *EurekaHandlerCaller) Lbtc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EurekaHandler.contract.Call(opts, &out, "lbtc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lbtc is a free data retrieval call binding the contract method 0x68b3c910.
//
// Solidity: function lbtc() view returns(address)
func (_EurekaHandler *EurekaHandlerSession) Lbtc() (common.Address, error) {
	return _EurekaHandler.Contract.Lbtc(&_EurekaHandler.CallOpts)
}

// Lbtc is a free data retrieval call binding the contract method 0x68b3c910.
//
// Solidity: function lbtc() view returns(address)
func (_EurekaHandler *EurekaHandlerCallerSession) Lbtc() (common.Address, error) {
	return _EurekaHandler.Contract.Lbtc(&_EurekaHandler.CallOpts)
}

// LbtcVoucher is a free data retrieval call binding the contract method 0xfcb49a3c.
//
// Solidity: function lbtcVoucher() view returns(address)
func (_EurekaHandler *EurekaHandlerCaller) LbtcVoucher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EurekaHandler.contract.Call(opts, &out, "lbtcVoucher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LbtcVoucher is a free data retrieval call binding the contract method 0xfcb49a3c.
//
// Solidity: function lbtcVoucher() view returns(address)
func (_EurekaHandler *EurekaHandlerSession) LbtcVoucher() (common.Address, error) {
	return _EurekaHandler.Contract.LbtcVoucher(&_EurekaHandler.CallOpts)
}

// LbtcVoucher is a free data retrieval call binding the contract method 0xfcb49a3c.
//
// Solidity: function lbtcVoucher() view returns(address)
func (_EurekaHandler *EurekaHandlerCallerSession) LbtcVoucher() (common.Address, error) {
	return _EurekaHandler.Contract.LbtcVoucher(&_EurekaHandler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EurekaHandler *EurekaHandlerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EurekaHandler.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EurekaHandler *EurekaHandlerSession) Owner() (common.Address, error) {
	return _EurekaHandler.Contract.Owner(&_EurekaHandler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EurekaHandler *EurekaHandlerCallerSession) Owner() (common.Address, error) {
	return _EurekaHandler.Contract.Owner(&_EurekaHandler.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EurekaHandler *EurekaHandlerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EurekaHandler.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EurekaHandler *EurekaHandlerSession) ProxiableUUID() ([32]byte, error) {
	return _EurekaHandler.Contract.ProxiableUUID(&_EurekaHandler.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EurekaHandler *EurekaHandlerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _EurekaHandler.Contract.ProxiableUUID(&_EurekaHandler.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_EurekaHandler *EurekaHandlerCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EurekaHandler.contract.Call(opts, &out, "swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_EurekaHandler *EurekaHandlerSession) SwapRouter() (common.Address, error) {
	return _EurekaHandler.Contract.SwapRouter(&_EurekaHandler.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_EurekaHandler *EurekaHandlerCallerSession) SwapRouter() (common.Address, error) {
	return _EurekaHandler.Contract.SwapRouter(&_EurekaHandler.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _ics20Transfer, address _swapRouter, address _lbtcVoucher, address _lbtc) returns()
func (_EurekaHandler *EurekaHandlerTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _ics20Transfer common.Address, _swapRouter common.Address, _lbtcVoucher common.Address, _lbtc common.Address) (*types.Transaction, error) {
	return _EurekaHandler.contract.Transact(opts, "initialize", _owner, _ics20Transfer, _swapRouter, _lbtcVoucher, _lbtc)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _ics20Transfer, address _swapRouter, address _lbtcVoucher, address _lbtc) returns()
func (_EurekaHandler *EurekaHandlerSession) Initialize(_owner common.Address, _ics20Transfer common.Address, _swapRouter common.Address, _lbtcVoucher common.Address, _lbtc common.Address) (*types.Transaction, error) {
	return _EurekaHandler.Contract.Initialize(&_EurekaHandler.TransactOpts, _owner, _ics20Transfer, _swapRouter, _lbtcVoucher, _lbtc)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _ics20Transfer, address _swapRouter, address _lbtcVoucher, address _lbtc) returns()
func (_EurekaHandler *EurekaHandlerTransactorSession) Initialize(_owner common.Address, _ics20Transfer common.Address, _swapRouter common.Address, _lbtcVoucher common.Address, _lbtc common.Address) (*types.Transaction, error) {
	return _EurekaHandler.Contract.Initialize(&_EurekaHandler.TransactOpts, _owner, _ics20Transfer, _swapRouter, _lbtcVoucher, _lbtc)
}

// LombardTransfer is a paid mutator transaction binding the contract method 0x9130d915.
//
// Solidity: function lombardTransfer(uint256 amount, (address,string,string,string,uint64,string) transferParams, (uint256,address,uint64) fees) returns()
func (_EurekaHandler *EurekaHandlerTransactor) LombardTransfer(opts *bind.TransactOpts, amount *big.Int, transferParams IEurekaHandlerTransferParams, fees IEurekaHandlerFees) (*types.Transaction, error) {
	return _EurekaHandler.contract.Transact(opts, "lombardTransfer", amount, transferParams, fees)
}

// LombardTransfer is a paid mutator transaction binding the contract method 0x9130d915.
//
// Solidity: function lombardTransfer(uint256 amount, (address,string,string,string,uint64,string) transferParams, (uint256,address,uint64) fees) returns()
func (_EurekaHandler *EurekaHandlerSession) LombardTransfer(amount *big.Int, transferParams IEurekaHandlerTransferParams, fees IEurekaHandlerFees) (*types.Transaction, error) {
	return _EurekaHandler.Contract.LombardTransfer(&_EurekaHandler.TransactOpts, amount, transferParams, fees)
}

// LombardTransfer is a paid mutator transaction binding the contract method 0x9130d915.
//
// Solidity: function lombardTransfer(uint256 amount, (address,string,string,string,uint64,string) transferParams, (uint256,address,uint64) fees) returns()
func (_EurekaHandler *EurekaHandlerTransactorSession) LombardTransfer(amount *big.Int, transferParams IEurekaHandlerTransferParams, fees IEurekaHandlerFees) (*types.Transaction, error) {
	return _EurekaHandler.Contract.LombardTransfer(&_EurekaHandler.TransactOpts, amount, transferParams, fees)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EurekaHandler *EurekaHandlerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EurekaHandler.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EurekaHandler *EurekaHandlerSession) RenounceOwnership() (*types.Transaction, error) {
	return _EurekaHandler.Contract.RenounceOwnership(&_EurekaHandler.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EurekaHandler *EurekaHandlerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EurekaHandler.Contract.RenounceOwnership(&_EurekaHandler.TransactOpts)
}

// SwapAndTransfer is a paid mutator transaction binding the contract method 0xb9b5974b.
//
// Solidity: function swapAndTransfer(address swapInputToken, uint256 swapInputAmount, bytes swapCalldata, (address,string,string,string,uint64,string) transferParams, (uint256,address,uint64) fees) returns()
func (_EurekaHandler *EurekaHandlerTransactor) SwapAndTransfer(opts *bind.TransactOpts, swapInputToken common.Address, swapInputAmount *big.Int, swapCalldata []byte, transferParams IEurekaHandlerTransferParams, fees IEurekaHandlerFees) (*types.Transaction, error) {
	return _EurekaHandler.contract.Transact(opts, "swapAndTransfer", swapInputToken, swapInputAmount, swapCalldata, transferParams, fees)
}

// SwapAndTransfer is a paid mutator transaction binding the contract method 0xb9b5974b.
//
// Solidity: function swapAndTransfer(address swapInputToken, uint256 swapInputAmount, bytes swapCalldata, (address,string,string,string,uint64,string) transferParams, (uint256,address,uint64) fees) returns()
func (_EurekaHandler *EurekaHandlerSession) SwapAndTransfer(swapInputToken common.Address, swapInputAmount *big.Int, swapCalldata []byte, transferParams IEurekaHandlerTransferParams, fees IEurekaHandlerFees) (*types.Transaction, error) {
	return _EurekaHandler.Contract.SwapAndTransfer(&_EurekaHandler.TransactOpts, swapInputToken, swapInputAmount, swapCalldata, transferParams, fees)
}

// SwapAndTransfer is a paid mutator transaction binding the contract method 0xb9b5974b.
//
// Solidity: function swapAndTransfer(address swapInputToken, uint256 swapInputAmount, bytes swapCalldata, (address,string,string,string,uint64,string) transferParams, (uint256,address,uint64) fees) returns()
func (_EurekaHandler *EurekaHandlerTransactorSession) SwapAndTransfer(swapInputToken common.Address, swapInputAmount *big.Int, swapCalldata []byte, transferParams IEurekaHandlerTransferParams, fees IEurekaHandlerFees) (*types.Transaction, error) {
	return _EurekaHandler.Contract.SwapAndTransfer(&_EurekaHandler.TransactOpts, swapInputToken, swapInputAmount, swapCalldata, transferParams, fees)
}

// Transfer is a paid mutator transaction binding the contract method 0xeecf31f9.
//
// Solidity: function transfer(uint256 amount, (address,string,string,string,uint64,string) transferParams, (uint256,address,uint64) fees) returns()
func (_EurekaHandler *EurekaHandlerTransactor) Transfer(opts *bind.TransactOpts, amount *big.Int, transferParams IEurekaHandlerTransferParams, fees IEurekaHandlerFees) (*types.Transaction, error) {
	return _EurekaHandler.contract.Transact(opts, "transfer", amount, transferParams, fees)
}

// Transfer is a paid mutator transaction binding the contract method 0xeecf31f9.
//
// Solidity: function transfer(uint256 amount, (address,string,string,string,uint64,string) transferParams, (uint256,address,uint64) fees) returns()
func (_EurekaHandler *EurekaHandlerSession) Transfer(amount *big.Int, transferParams IEurekaHandlerTransferParams, fees IEurekaHandlerFees) (*types.Transaction, error) {
	return _EurekaHandler.Contract.Transfer(&_EurekaHandler.TransactOpts, amount, transferParams, fees)
}

// Transfer is a paid mutator transaction binding the contract method 0xeecf31f9.
//
// Solidity: function transfer(uint256 amount, (address,string,string,string,uint64,string) transferParams, (uint256,address,uint64) fees) returns()
func (_EurekaHandler *EurekaHandlerTransactorSession) Transfer(amount *big.Int, transferParams IEurekaHandlerTransferParams, fees IEurekaHandlerFees) (*types.Transaction, error) {
	return _EurekaHandler.Contract.Transfer(&_EurekaHandler.TransactOpts, amount, transferParams, fees)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EurekaHandler *EurekaHandlerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EurekaHandler.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EurekaHandler *EurekaHandlerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EurekaHandler.Contract.TransferOwnership(&_EurekaHandler.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EurekaHandler *EurekaHandlerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EurekaHandler.Contract.TransferOwnership(&_EurekaHandler.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EurekaHandler *EurekaHandlerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EurekaHandler.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EurekaHandler *EurekaHandlerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EurekaHandler.Contract.UpgradeToAndCall(&_EurekaHandler.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EurekaHandler *EurekaHandlerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EurekaHandler.Contract.UpgradeToAndCall(&_EurekaHandler.TransactOpts, newImplementation, data)
}

// EurekaHandlerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EurekaHandler contract.
type EurekaHandlerInitializedIterator struct {
	Event *EurekaHandlerInitialized // Event containing the contract specifics and raw log

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
func (it *EurekaHandlerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EurekaHandlerInitialized)
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
		it.Event = new(EurekaHandlerInitialized)
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
func (it *EurekaHandlerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EurekaHandlerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EurekaHandlerInitialized represents a Initialized event raised by the EurekaHandler contract.
type EurekaHandlerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EurekaHandler *EurekaHandlerFilterer) FilterInitialized(opts *bind.FilterOpts) (*EurekaHandlerInitializedIterator, error) {

	logs, sub, err := _EurekaHandler.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EurekaHandlerInitializedIterator{contract: _EurekaHandler.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EurekaHandler *EurekaHandlerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EurekaHandlerInitialized) (event.Subscription, error) {

	logs, sub, err := _EurekaHandler.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EurekaHandlerInitialized)
				if err := _EurekaHandler.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_EurekaHandler *EurekaHandlerFilterer) ParseInitialized(log types.Log) (*EurekaHandlerInitialized, error) {
	event := new(EurekaHandlerInitialized)
	if err := _EurekaHandler.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EurekaHandlerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EurekaHandler contract.
type EurekaHandlerOwnershipTransferredIterator struct {
	Event *EurekaHandlerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EurekaHandlerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EurekaHandlerOwnershipTransferred)
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
		it.Event = new(EurekaHandlerOwnershipTransferred)
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
func (it *EurekaHandlerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EurekaHandlerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EurekaHandlerOwnershipTransferred represents a OwnershipTransferred event raised by the EurekaHandler contract.
type EurekaHandlerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EurekaHandler *EurekaHandlerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EurekaHandlerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EurekaHandler.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EurekaHandlerOwnershipTransferredIterator{contract: _EurekaHandler.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EurekaHandler *EurekaHandlerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EurekaHandlerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EurekaHandler.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EurekaHandlerOwnershipTransferred)
				if err := _EurekaHandler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EurekaHandler *EurekaHandlerFilterer) ParseOwnershipTransferred(log types.Log) (*EurekaHandlerOwnershipTransferred, error) {
	event := new(EurekaHandlerOwnershipTransferred)
	if err := _EurekaHandler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EurekaHandlerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EurekaHandler contract.
type EurekaHandlerTransferIterator struct {
	Event *EurekaHandlerTransfer // Event containing the contract specifics and raw log

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
func (it *EurekaHandlerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EurekaHandlerTransfer)
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
		it.Event = new(EurekaHandlerTransfer)
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
func (it *EurekaHandlerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EurekaHandlerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EurekaHandlerTransfer represents a Transfer event raised by the EurekaHandler contract.
type EurekaHandlerTransfer struct {
	Token    common.Address
	Amount   *big.Int
	RelayFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x7fa9aafeb8bb803d77de5d84bc2f2edbd842ca91b20cd5020aa21dfe26ab0be9.
//
// Solidity: event Transfer(address indexed token, uint256 amount, uint256 relayFee)
func (_EurekaHandler *EurekaHandlerFilterer) FilterTransfer(opts *bind.FilterOpts, token []common.Address) (*EurekaHandlerTransferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EurekaHandler.contract.FilterLogs(opts, "Transfer", tokenRule)
	if err != nil {
		return nil, err
	}
	return &EurekaHandlerTransferIterator{contract: _EurekaHandler.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x7fa9aafeb8bb803d77de5d84bc2f2edbd842ca91b20cd5020aa21dfe26ab0be9.
//
// Solidity: event Transfer(address indexed token, uint256 amount, uint256 relayFee)
func (_EurekaHandler *EurekaHandlerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EurekaHandlerTransfer, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EurekaHandler.contract.WatchLogs(opts, "Transfer", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EurekaHandlerTransfer)
				if err := _EurekaHandler.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x7fa9aafeb8bb803d77de5d84bc2f2edbd842ca91b20cd5020aa21dfe26ab0be9.
//
// Solidity: event Transfer(address indexed token, uint256 amount, uint256 relayFee)
func (_EurekaHandler *EurekaHandlerFilterer) ParseTransfer(log types.Log) (*EurekaHandlerTransfer, error) {
	event := new(EurekaHandlerTransfer)
	if err := _EurekaHandler.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EurekaHandlerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EurekaHandler contract.
type EurekaHandlerUpgradedIterator struct {
	Event *EurekaHandlerUpgraded // Event containing the contract specifics and raw log

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
func (it *EurekaHandlerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EurekaHandlerUpgraded)
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
		it.Event = new(EurekaHandlerUpgraded)
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
func (it *EurekaHandlerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EurekaHandlerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EurekaHandlerUpgraded represents a Upgraded event raised by the EurekaHandler contract.
type EurekaHandlerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EurekaHandler *EurekaHandlerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EurekaHandlerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EurekaHandler.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EurekaHandlerUpgradedIterator{contract: _EurekaHandler.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EurekaHandler *EurekaHandlerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EurekaHandlerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EurekaHandler.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EurekaHandlerUpgraded)
				if err := _EurekaHandler.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_EurekaHandler *EurekaHandlerFilterer) ParseUpgraded(log types.Log) (*EurekaHandlerUpgraded, error) {
	event := new(EurekaHandlerUpgraded)
	if err := _EurekaHandler.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
