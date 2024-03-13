package service

import (
	orangeDID "orange-provider-wrapper/orangeDid"
	"orange-provider-wrapper/utils"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/laizy/web3"
	"github.com/laizy/web3/jsonrpc"
)

type DidService struct {
	client *jsonrpc.Client
	// rpc             string
	// contractAddress string
	orangeDid *orangeDID.OrangePubkeysManager
}

var GlobalDidService *DidService

func InitDidService(rpc string, contractAddress string) error {
	client, err := jsonrpc.NewClient(rpc)
	if err != nil {
		return err
	}
	od := orangeDID.NewOrangePubkeysManager(web3.HexToAddress(contractAddress), client)

	GlobalDidService = &DidService{client: client, orangeDid: od}
	return nil
}

func (d *DidService) GetDidPublicKey(did string) (string, error) {
	addr, err := utils.DIDToEthAddress(did)
	if err != nil {
		return "", err
	}
	btes, err := d.orangeDid.GetDIDPublick(web3.HexToAddress(addr))
	return hexutil.Encode(btes), err
}
