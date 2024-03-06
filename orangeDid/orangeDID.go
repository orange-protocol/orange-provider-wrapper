// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package orangeDid

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

// OrangeDIDMetaData contains all meta data concerning the OrangeDID contract.
var OrangeDIDMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getDIDPublick\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OrangeDIDABI is the input ABI used to generate the binding from.
// Deprecated: Use OrangeDIDMetaData.ABI instead.
var OrangeDIDABI = OrangeDIDMetaData.ABI

// OrangeDID is an auto generated Go binding around an Ethereum contract.
type OrangeDID struct {
	OrangeDIDCaller     // Read-only binding to the contract
	OrangeDIDTransactor // Write-only binding to the contract
	OrangeDIDFilterer   // Log filterer for contract events
}

// OrangeDIDCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrangeDIDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrangeDIDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrangeDIDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrangeDIDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrangeDIDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrangeDIDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrangeDIDSession struct {
	Contract     *OrangeDID        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrangeDIDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrangeDIDCallerSession struct {
	Contract *OrangeDIDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OrangeDIDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrangeDIDTransactorSession struct {
	Contract     *OrangeDIDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OrangeDIDRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrangeDIDRaw struct {
	Contract *OrangeDID // Generic contract binding to access the raw methods on
}

// OrangeDIDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrangeDIDCallerRaw struct {
	Contract *OrangeDIDCaller // Generic read-only contract binding to access the raw methods on
}

// OrangeDIDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrangeDIDTransactorRaw struct {
	Contract *OrangeDIDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrangeDID creates a new instance of OrangeDID, bound to a specific deployed contract.
func NewOrangeDID(address common.Address, backend bind.ContractBackend) (*OrangeDID, error) {
	contract, err := bindOrangeDID(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrangeDID{OrangeDIDCaller: OrangeDIDCaller{contract: contract}, OrangeDIDTransactor: OrangeDIDTransactor{contract: contract}, OrangeDIDFilterer: OrangeDIDFilterer{contract: contract}}, nil
}

// NewOrangeDIDCaller creates a new read-only instance of OrangeDID, bound to a specific deployed contract.
func NewOrangeDIDCaller(address common.Address, caller bind.ContractCaller) (*OrangeDIDCaller, error) {
	contract, err := bindOrangeDID(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrangeDIDCaller{contract: contract}, nil
}

// NewOrangeDIDTransactor creates a new write-only instance of OrangeDID, bound to a specific deployed contract.
func NewOrangeDIDTransactor(address common.Address, transactor bind.ContractTransactor) (*OrangeDIDTransactor, error) {
	contract, err := bindOrangeDID(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrangeDIDTransactor{contract: contract}, nil
}

// NewOrangeDIDFilterer creates a new log filterer instance of OrangeDID, bound to a specific deployed contract.
func NewOrangeDIDFilterer(address common.Address, filterer bind.ContractFilterer) (*OrangeDIDFilterer, error) {
	contract, err := bindOrangeDID(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrangeDIDFilterer{contract: contract}, nil
}

// bindOrangeDID binds a generic wrapper to an already deployed contract.
func bindOrangeDID(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrangeDIDMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrangeDID *OrangeDIDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrangeDID.Contract.OrangeDIDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrangeDID *OrangeDIDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrangeDID.Contract.OrangeDIDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrangeDID *OrangeDIDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrangeDID.Contract.OrangeDIDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrangeDID *OrangeDIDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrangeDID.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrangeDID *OrangeDIDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrangeDID.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrangeDID *OrangeDIDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrangeDID.Contract.contract.Transact(opts, method, params...)
}

// GetDIDPublick is a free data retrieval call binding the contract method 0x3f36aba8.
//
// Solidity: function getDIDPublick(string did) view returns(string)
func (_OrangeDID *OrangeDIDCaller) GetDIDPublick(opts *bind.CallOpts, did string) (string, error) {
	var out []interface{}
	err := _OrangeDID.contract.Call(opts, &out, "getDIDPublick", did)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDIDPublick is a free data retrieval call binding the contract method 0x3f36aba8.
//
// Solidity: function getDIDPublick(string did) view returns(string)
func (_OrangeDID *OrangeDIDSession) GetDIDPublick(did string) (string, error) {
	return _OrangeDID.Contract.GetDIDPublick(&_OrangeDID.CallOpts, did)
}

// GetDIDPublick is a free data retrieval call binding the contract method 0x3f36aba8.
//
// Solidity: function getDIDPublick(string did) view returns(string)
func (_OrangeDID *OrangeDIDCallerSession) GetDIDPublick(did string) (string, error) {
	return _OrangeDID.Contract.GetDIDPublick(&_OrangeDID.CallOpts, did)
}
