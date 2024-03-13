package orangeDID

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/laizy/web3"
	"github.com/laizy/web3/contract"
	"github.com/laizy/web3/jsonrpc"
	"github.com/laizy/web3/utils"
	"github.com/mitchellh/mapstructure"
)

var (
	_ = json.Unmarshal
	_ = big.NewInt
	_ = fmt.Printf
	_ = utils.JsonStr
	_ = mapstructure.Decode
)

// OrangePubkeysManager is a solidity contract
type OrangePubkeysManager struct {
	c *contract.Contract
}

// DeployOrangePubkeysManager deploys a new OrangePubkeysManager contract
func DeployOrangePubkeysManager(provider *jsonrpc.Client, from web3.Address) *contract.Txn {
	return contract.DeployContract(provider, from, abiOrangePubkeysManager, binOrangePubkeysManager)
}

// NewOrangePubkeysManager creates a new instance of the contract at a specific address
func NewOrangePubkeysManager(addr web3.Address, provider *jsonrpc.Client) *OrangePubkeysManager {
	return &OrangePubkeysManager{c: contract.NewContract(addr, abiOrangePubkeysManager, provider)}
}

// Contract returns the contract object
func (_a *OrangePubkeysManager) Contract() *contract.Contract {
	return _a.c
}

// calls

// GetDIDPublick calls the getDIDPublick method in the solidity contract
func (_a *OrangePubkeysManager) GetDIDPublick(addr web3.Address, block ...web3.BlockNumber) (retval0 []byte, err error) {
	var out map[string]interface{}
	_ = out // avoid not used compiler error

	out, err = _a.c.Call("getDIDPublick", web3.EncodeBlock(block...), addr)
	if err != nil {
		return
	}

	// decode outputs

	if err = mapstructure.Decode(out["0"], &retval0); err != nil {
		err = fmt.Errorf("failed to encode output at index 0")
	}

	return
}

// Owner calls the owner method in the solidity contract
func (_a *OrangePubkeysManager) Owner(block ...web3.BlockNumber) (retval0 web3.Address, err error) {
	var out map[string]interface{}
	_ = out // avoid not used compiler error

	out, err = _a.c.Call("owner", web3.EncodeBlock(block...))
	if err != nil {
		return
	}

	// decode outputs

	if err = mapstructure.Decode(out["0"], &retval0); err != nil {
		err = fmt.Errorf("failed to encode output at index 0")
	}

	return
}

// Pubkeys calls the pubkeys method in the solidity contract
func (_a *OrangePubkeysManager) Pubkeys(val0 web3.Address, block ...web3.BlockNumber) (retval0 []byte, err error) {
	var out map[string]interface{}
	_ = out // avoid not used compiler error

	out, err = _a.c.Call("pubkeys", web3.EncodeBlock(block...), val0)
	if err != nil {
		return
	}

	// decode outputs

	if err = mapstructure.Decode(out["0"], &retval0); err != nil {
		err = fmt.Errorf("failed to encode output at index 0")
	}

	return
}

// txns

// RegisterDID sends a registerDID transaction in the solidity contract
func (_a *OrangePubkeysManager) RegisterDID(pubkey []byte) *contract.Txn {
	return _a.c.Txn("registerDID", pubkey)
}

// RegisterDIDByOwner sends a registerDIDByOwner transaction in the solidity contract
func (_a *OrangePubkeysManager) RegisterDIDByOwner(wallet web3.Address, pubkey []byte) *contract.Txn {
	return _a.c.Txn("registerDIDByOwner", wallet, pubkey)
}

// RenounceOwnership sends a renounceOwnership transaction in the solidity contract
func (_a *OrangePubkeysManager) RenounceOwnership() *contract.Txn {
	return _a.c.Txn("renounceOwnership")
}

// TransferOwnership sends a transferOwnership transaction in the solidity contract
func (_a *OrangePubkeysManager) TransferOwnership(newOwner web3.Address) *contract.Txn {
	return _a.c.Txn("transferOwnership", newOwner)
}

// events

func (_a *OrangePubkeysManager) OwnershipTransferredTopicFilter(previousOwner []web3.Address, newOwner []web3.Address) [][]web3.Hash {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	var query [][]interface{}
	query = append(query, []interface{}{OwnershipTransferredEventID}, previousOwnerRule, newOwnerRule)

	topics, err := contract.MakeTopics(query...)
	utils.Ensure(err)

	return topics
}

func (_a *OrangePubkeysManager) FilterOwnershipTransferredEvent(previousOwner []web3.Address, newOwner []web3.Address, startBlock uint64, endBlock ...uint64) ([]*OwnershipTransferredEvent, error) {
	topic := _a.OwnershipTransferredTopicFilter(previousOwner, newOwner)

	logs, err := _a.c.FilterLogsWithTopic(topic, startBlock, endBlock...)
	if err != nil {
		return nil, err
	}
	res := make([]*OwnershipTransferredEvent, 0)
	evts := _a.c.Abi.Events["OwnershipTransferred"]
	for _, log := range logs {
		args, err := evts.ParseLog(log)
		if err != nil {
			return nil, err
		}
		var evtItem OwnershipTransferredEvent
		err = json.Unmarshal([]byte(utils.JsonStr(args)), &evtItem)
		if err != nil {
			return nil, err
		}
		evtItem.Raw = log
		res = append(res, &evtItem)
	}
	return res, nil
}

func (_a *OrangePubkeysManager) RegisterPubkeyTopicFilter() [][]web3.Hash {

	var query [][]interface{}
	query = append(query, []interface{}{RegisterPubkeyEventID})

	topics, err := contract.MakeTopics(query...)
	utils.Ensure(err)

	return topics
}

func (_a *OrangePubkeysManager) FilterRegisterPubkeyEvent(startBlock uint64, endBlock ...uint64) ([]*RegisterPubkeyEvent, error) {
	topic := _a.RegisterPubkeyTopicFilter()

	logs, err := _a.c.FilterLogsWithTopic(topic, startBlock, endBlock...)
	if err != nil {
		return nil, err
	}
	res := make([]*RegisterPubkeyEvent, 0)
	evts := _a.c.Abi.Events["registerPubkey"]
	for _, log := range logs {
		args, err := evts.ParseLog(log)
		if err != nil {
			return nil, err
		}
		var evtItem RegisterPubkeyEvent
		err = json.Unmarshal([]byte(utils.JsonStr(args)), &evtItem)
		if err != nil {
			return nil, err
		}
		evtItem.Raw = log
		res = append(res, &evtItem)
	}
	return res, nil
}
