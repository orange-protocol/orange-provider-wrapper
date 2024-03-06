package service

import (
	"orange-provider-wrapper/orangeDid"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DidService struct {
	client *ethclient.Client
	// rpc             string
	// contractAddress string
	orangeDid *orangeDid.OrangeDID
}

var GlobalDidService *DidService

func InitDidService(rpc string, contractAddress string) error {
	eclient, err := ethclient.Dial(rpc)
	if err != nil {
		return err
	}
	od, err := orangeDid.NewOrangeDID(common.HexToAddress(contractAddress), eclient)
	if err != nil {
		return err
	}

	GlobalDidService = &DidService{client: eclient, orangeDid: od}
	return nil
}

func (d *DidService) GetDidPublicKey(did string) (string, error) {
	return d.orangeDid.GetDIDPublick(nil, did)
}
