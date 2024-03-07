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
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getDIDPublick\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pubkeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"registerDID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"registerDIDByOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetDIDPublick is a free data retrieval call binding the contract method 0xdb36f837.
//
// Solidity: function getDIDPublick(address addr) view returns(bytes)
func (_OrangeDID *OrangeDIDCaller) GetDIDPublick(opts *bind.CallOpts, addr common.Address) ([]byte, error) {
	var out []interface{}
	err := _OrangeDID.contract.Call(opts, &out, "getDIDPublick", addr)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDIDPublick is a free data retrieval call binding the contract method 0xdb36f837.
//
// Solidity: function getDIDPublick(address addr) view returns(bytes)
func (_OrangeDID *OrangeDIDSession) GetDIDPublick(addr common.Address) ([]byte, error) {
	return _OrangeDID.Contract.GetDIDPublick(&_OrangeDID.CallOpts, addr)
}

// GetDIDPublick is a free data retrieval call binding the contract method 0xdb36f837.
//
// Solidity: function getDIDPublick(address addr) view returns(bytes)
func (_OrangeDID *OrangeDIDCallerSession) GetDIDPublick(addr common.Address) ([]byte, error) {
	return _OrangeDID.Contract.GetDIDPublick(&_OrangeDID.CallOpts, addr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrangeDID *OrangeDIDCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrangeDID.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrangeDID *OrangeDIDSession) Owner() (common.Address, error) {
	return _OrangeDID.Contract.Owner(&_OrangeDID.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OrangeDID *OrangeDIDCallerSession) Owner() (common.Address, error) {
	return _OrangeDID.Contract.Owner(&_OrangeDID.CallOpts)
}

// Pubkeys is a free data retrieval call binding the contract method 0x9e50f99c.
//
// Solidity: function pubkeys(address ) view returns(bytes)
func (_OrangeDID *OrangeDIDCaller) Pubkeys(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var out []interface{}
	err := _OrangeDID.contract.Call(opts, &out, "pubkeys", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Pubkeys is a free data retrieval call binding the contract method 0x9e50f99c.
//
// Solidity: function pubkeys(address ) view returns(bytes)
func (_OrangeDID *OrangeDIDSession) Pubkeys(arg0 common.Address) ([]byte, error) {
	return _OrangeDID.Contract.Pubkeys(&_OrangeDID.CallOpts, arg0)
}

// Pubkeys is a free data retrieval call binding the contract method 0x9e50f99c.
//
// Solidity: function pubkeys(address ) view returns(bytes)
func (_OrangeDID *OrangeDIDCallerSession) Pubkeys(arg0 common.Address) ([]byte, error) {
	return _OrangeDID.Contract.Pubkeys(&_OrangeDID.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_OrangeDID *OrangeDIDTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _OrangeDID.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_OrangeDID *OrangeDIDSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _OrangeDID.Contract.Initialize(&_OrangeDID.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_OrangeDID *OrangeDIDTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _OrangeDID.Contract.Initialize(&_OrangeDID.TransactOpts, _owner)
}

// RegisterDID is a paid mutator transaction binding the contract method 0xc8ec764b.
//
// Solidity: function registerDID(bytes pubkey) returns()
func (_OrangeDID *OrangeDIDTransactor) RegisterDID(opts *bind.TransactOpts, pubkey []byte) (*types.Transaction, error) {
	return _OrangeDID.contract.Transact(opts, "registerDID", pubkey)
}

// RegisterDID is a paid mutator transaction binding the contract method 0xc8ec764b.
//
// Solidity: function registerDID(bytes pubkey) returns()
func (_OrangeDID *OrangeDIDSession) RegisterDID(pubkey []byte) (*types.Transaction, error) {
	return _OrangeDID.Contract.RegisterDID(&_OrangeDID.TransactOpts, pubkey)
}

// RegisterDID is a paid mutator transaction binding the contract method 0xc8ec764b.
//
// Solidity: function registerDID(bytes pubkey) returns()
func (_OrangeDID *OrangeDIDTransactorSession) RegisterDID(pubkey []byte) (*types.Transaction, error) {
	return _OrangeDID.Contract.RegisterDID(&_OrangeDID.TransactOpts, pubkey)
}

// RegisterDIDByOwner is a paid mutator transaction binding the contract method 0x15707820.
//
// Solidity: function registerDIDByOwner(address wallet, bytes pubkey) returns()
func (_OrangeDID *OrangeDIDTransactor) RegisterDIDByOwner(opts *bind.TransactOpts, wallet common.Address, pubkey []byte) (*types.Transaction, error) {
	return _OrangeDID.contract.Transact(opts, "registerDIDByOwner", wallet, pubkey)
}

// RegisterDIDByOwner is a paid mutator transaction binding the contract method 0x15707820.
//
// Solidity: function registerDIDByOwner(address wallet, bytes pubkey) returns()
func (_OrangeDID *OrangeDIDSession) RegisterDIDByOwner(wallet common.Address, pubkey []byte) (*types.Transaction, error) {
	return _OrangeDID.Contract.RegisterDIDByOwner(&_OrangeDID.TransactOpts, wallet, pubkey)
}

// RegisterDIDByOwner is a paid mutator transaction binding the contract method 0x15707820.
//
// Solidity: function registerDIDByOwner(address wallet, bytes pubkey) returns()
func (_OrangeDID *OrangeDIDTransactorSession) RegisterDIDByOwner(wallet common.Address, pubkey []byte) (*types.Transaction, error) {
	return _OrangeDID.Contract.RegisterDIDByOwner(&_OrangeDID.TransactOpts, wallet, pubkey)
}

// OrangeDIDInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OrangeDID contract.
type OrangeDIDInitializedIterator struct {
	Event *OrangeDIDInitialized // Event containing the contract specifics and raw log

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
func (it *OrangeDIDInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrangeDIDInitialized)
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
		it.Event = new(OrangeDIDInitialized)
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
func (it *OrangeDIDInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrangeDIDInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrangeDIDInitialized represents a Initialized event raised by the OrangeDID contract.
type OrangeDIDInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OrangeDID *OrangeDIDFilterer) FilterInitialized(opts *bind.FilterOpts) (*OrangeDIDInitializedIterator, error) {

	logs, sub, err := _OrangeDID.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OrangeDIDInitializedIterator{contract: _OrangeDID.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OrangeDID *OrangeDIDFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OrangeDIDInitialized) (event.Subscription, error) {

	logs, sub, err := _OrangeDID.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrangeDIDInitialized)
				if err := _OrangeDID.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OrangeDID *OrangeDIDFilterer) ParseInitialized(log types.Log) (*OrangeDIDInitialized, error) {
	event := new(OrangeDIDInitialized)
	if err := _OrangeDID.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
