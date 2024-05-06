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

// Credentials calls the credentials method in the solidity contract
func (_a *OrangePubkeysManager) Credentials(val0 string, block ...web3.BlockNumber) (retval0 web3.Address, retval1 *big.Int, retval2 web3.Address, retval3 bool, retval4 [32]byte, err error) {
	var out map[string]interface{}
	_ = out // avoid not used compiler error

	out, err = _a.c.Call("credentials", web3.EncodeBlock(block...), val0)
	if err != nil {
		return
	}

	// decode outputs

	if err = mapstructure.Decode(out["issuer"], &retval0); err != nil {
		err = fmt.Errorf("failed to encode output at index 0")
	}
	if err = mapstructure.Decode(out["idx"], &retval1); err != nil {
		err = fmt.Errorf("failed to encode output at index 1")
	}
	if err = mapstructure.Decode(out["holder"], &retval2); err != nil {
		err = fmt.Errorf("failed to encode output at index 2")
	}
	if err = mapstructure.Decode(out["isValid"], &retval3); err != nil {
		err = fmt.Errorf("failed to encode output at index 3")
	}
	if err = mapstructure.Decode(out["subjectHash"], &retval4); err != nil {
		err = fmt.Errorf("failed to encode output at index 4")
	}

	return
}

// GetDIDPublickey calls the getDIDPublickey method in the solidity contract
func (_a *OrangePubkeysManager) GetDIDPublickey(addr web3.Address, block ...web3.BlockNumber) (retval0 []byte, err error) {
	var out map[string]interface{}
	_ = out // avoid not used compiler error

	out, err = _a.c.Call("getDIDPublickey", web3.EncodeBlock(block...), addr)
	if err != nil {
		return
	}

	// decode outputs

	if err = mapstructure.Decode(out["0"], &retval0); err != nil {
		err = fmt.Errorf("failed to encode output at index 0")
	}

	return
}

// Manager calls the manager method in the solidity contract
func (_a *OrangePubkeysManager) Manager(block ...web3.BlockNumber) (retval0 web3.Address, err error) {
	var out map[string]interface{}
	_ = out // avoid not used compiler error

	out, err = _a.c.Call("manager", web3.EncodeBlock(block...))
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

// Initialize sends a initialize transaction in the solidity contract
func (_a *OrangePubkeysManager) Initialize(manager web3.Address) *contract.Txn {
	return _a.c.Txn("initialize", manager)
}

// RegisterCredential sends a registerCredential transaction in the solidity contract
func (_a *OrangePubkeysManager) RegisterCredential(credentialId string, issuer web3.Address, idx *big.Int, sig []byte, subjectHash [32]byte) *contract.Txn {
	return _a.c.Txn("registerCredential", credentialId, issuer, idx, sig, subjectHash)
}

// RegisterCredentialByManager sends a registerCredentialByManager transaction in the solidity contract
func (_a *OrangePubkeysManager) RegisterCredentialByManager(credentialId string, issuer web3.Address, idx *big.Int, holder web3.Address, subjectHash [32]byte) *contract.Txn {
	return _a.c.Txn("registerCredentialByManager", credentialId, issuer, idx, holder, subjectHash)
}

// RegisterDID sends a registerDID transaction in the solidity contract
func (_a *OrangePubkeysManager) RegisterDID(pubkey []byte) *contract.Txn {
	return _a.c.Txn("registerDID", pubkey)
}

// RegisterDIDByOwner sends a registerDIDByOwner transaction in the solidity contract
func (_a *OrangePubkeysManager) RegisterDIDByOwner(wallet web3.Address, pubkey []byte) *contract.Txn {
	return _a.c.Txn("registerDIDByOwner", wallet, pubkey)
}

// RevokeCredentialByManager sends a revokeCredentialByManager transaction in the solidity contract
func (_a *OrangePubkeysManager) RevokeCredentialByManager(credentialId string) *contract.Txn {
	return _a.c.Txn("revokeCredentialByManager", credentialId)
}

// RevokeCredentialByOwner sends a revokeCredentialByOwner transaction in the solidity contract
func (_a *OrangePubkeysManager) RevokeCredentialByOwner(credentialId string) *contract.Txn {
	return _a.c.Txn("revokeCredentialByOwner", credentialId)
}

// SetManager sends a setManager transaction in the solidity contract
func (_a *OrangePubkeysManager) SetManager(newManager web3.Address) *contract.Txn {
	return _a.c.Txn("setManager", newManager)
}

// TransferOwner sends a transferOwner transaction in the solidity contract
func (_a *OrangePubkeysManager) TransferOwner(newOwner web3.Address) *contract.Txn {
	return _a.c.Txn("transferOwner", newOwner)
}

// events

func (_a *OrangePubkeysManager) InitializedTopicFilter() [][]web3.Hash {

	var query [][]interface{}
	query = append(query, []interface{}{InitializedEventID})

	topics, err := contract.MakeTopics(query...)
	utils.Ensure(err)

	return topics
}

func (_a *OrangePubkeysManager) FilterInitializedEvent(startBlock uint64, endBlock ...uint64) ([]*InitializedEvent, error) {
	topic := _a.InitializedTopicFilter()

	logs, err := _a.c.FilterLogsWithTopic(topic, startBlock, endBlock...)
	if err != nil {
		return nil, err
	}
	res := make([]*InitializedEvent, 0)
	evts := _a.c.Abi.Events["Initialized"]
	for _, log := range logs {
		args, err := evts.ParseLog(log)
		if err != nil {
			return nil, err
		}
		var evtItem InitializedEvent
		err = json.Unmarshal([]byte(utils.JsonStr(args)), &evtItem)
		if err != nil {
			return nil, err
		}
		evtItem.Raw = log
		res = append(res, &evtItem)
	}
	return res, nil
}

func (_a *OrangePubkeysManager) RegisterCredentialEventTopicFilter() [][]web3.Hash {

	var query [][]interface{}
	query = append(query, []interface{}{RegisterCredentialEventEventID})

	topics, err := contract.MakeTopics(query...)
	utils.Ensure(err)

	return topics
}

func (_a *OrangePubkeysManager) FilterRegisterCredentialEventEvent(startBlock uint64, endBlock ...uint64) ([]*RegisterCredentialEventEvent, error) {
	topic := _a.RegisterCredentialEventTopicFilter()

	logs, err := _a.c.FilterLogsWithTopic(topic, startBlock, endBlock...)
	if err != nil {
		return nil, err
	}
	res := make([]*RegisterCredentialEventEvent, 0)
	evts := _a.c.Abi.Events["registerCredentialEvent"]
	for _, log := range logs {
		args, err := evts.ParseLog(log)
		if err != nil {
			return nil, err
		}
		var evtItem RegisterCredentialEventEvent
		err = json.Unmarshal([]byte(utils.JsonStr(args)), &evtItem)
		if err != nil {
			return nil, err
		}
		evtItem.Raw = log
		res = append(res, &evtItem)
	}
	return res, nil
}

func (_a *OrangePubkeysManager) RegisterPubkeyEventTopicFilter() [][]web3.Hash {

	var query [][]interface{}
	query = append(query, []interface{}{RegisterPubkeyEventEventID})

	topics, err := contract.MakeTopics(query...)
	utils.Ensure(err)

	return topics
}

func (_a *OrangePubkeysManager) FilterRegisterPubkeyEventEvent(startBlock uint64, endBlock ...uint64) ([]*RegisterPubkeyEventEvent, error) {
	topic := _a.RegisterPubkeyEventTopicFilter()

	logs, err := _a.c.FilterLogsWithTopic(topic, startBlock, endBlock...)
	if err != nil {
		return nil, err
	}
	res := make([]*RegisterPubkeyEventEvent, 0)
	evts := _a.c.Abi.Events["registerPubkeyEvent"]
	for _, log := range logs {
		args, err := evts.ParseLog(log)
		if err != nil {
			return nil, err
		}
		var evtItem RegisterPubkeyEventEvent
		err = json.Unmarshal([]byte(utils.JsonStr(args)), &evtItem)
		if err != nil {
			return nil, err
		}
		evtItem.Raw = log
		res = append(res, &evtItem)
	}
	return res, nil
}

func (_a *OrangePubkeysManager) RevokeCredentialEventTopicFilter() [][]web3.Hash {

	var query [][]interface{}
	query = append(query, []interface{}{RevokeCredentialEventEventID})

	topics, err := contract.MakeTopics(query...)
	utils.Ensure(err)

	return topics
}

func (_a *OrangePubkeysManager) FilterRevokeCredentialEventEvent(startBlock uint64, endBlock ...uint64) ([]*RevokeCredentialEventEvent, error) {
	topic := _a.RevokeCredentialEventTopicFilter()

	logs, err := _a.c.FilterLogsWithTopic(topic, startBlock, endBlock...)
	if err != nil {
		return nil, err
	}
	res := make([]*RevokeCredentialEventEvent, 0)
	evts := _a.c.Abi.Events["revokeCredentialEvent"]
	for _, log := range logs {
		args, err := evts.ParseLog(log)
		if err != nil {
			return nil, err
		}
		var evtItem RevokeCredentialEventEvent
		err = json.Unmarshal([]byte(utils.JsonStr(args)), &evtItem)
		if err != nil {
			return nil, err
		}
		evtItem.Raw = log
		res = append(res, &evtItem)
	}
	return res, nil
}
