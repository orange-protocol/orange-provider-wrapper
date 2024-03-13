package service

import (
	"fmt"
	"orange-provider-wrapper/config"
	"orange-provider-wrapper/log"
	orangeDID "orange-provider-wrapper/orangeDid"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/laizy/web3"
	"github.com/laizy/web3/contract"
	"github.com/laizy/web3/jsonrpc"
)

func NewWallet() error {
	folderPath := config.GlobalConfig.Keystore
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// Folder does not exist, attempt to create it
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create folder: %s, error: %v", folderPath, err)
		}
	}
	capitalKeyStore := keystore.NewKeyStore(folderPath, keystore.StandardScryptN,
		keystore.StandardScryptP)

	accArr := capitalKeyStore.Accounts()
	if len(accArr) > 0 {
		return fmt.Errorf("already has account: %s ", accArr[0].Address.String())
	}
	acct, err := capitalKeyStore.NewAccount(config.GlobalConfig.WalletPwd)
	if err != nil {
		return err
	}
	fmt.Printf("new account created: %s\n", acct.Address.String())
	return nil
}

func RegisterDID() error {
	folderPath := config.GlobalConfig.Keystore
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return fmt.Errorf("no wallet file found!")
	}

	//read file name under directory
	fileInfos, err := os.ReadDir(folderPath)
	if err != nil {
		log.Error("initSigAccount ReadDir failed", "err", err.Error())
		return err
	}

	if len(fileInfos) == 0 {
		log.Error("initSigAccount ReadDir failed", "err", "no files under the "+folderPath)
		return err
	}

	jsonFileName := ""
	for _, f := range fileInfos {
		if !f.IsDir() {
			jsonFileName = f.Name()
			break
		}
	}

	if jsonFileName == "" {
		log.Error("initSigAccount ReadDir failed", "err", "no files under the "+folderPath)
		return err
	}

	keystoreJSON, err := os.ReadFile(folderPath + "/" + jsonFileName)
	if err != nil {
		log.Error("initSigAccount ReadFile failed", "err", err.Error())
		return err
	}
	key, err := keystore.DecryptKey(keystoreJSON, config.GlobalConfig.WalletPwd)
	if err != nil {
		log.Error("initSigAccount DecryptKey failed", "err", err.Error())
		return err
	}
	fmt.Printf("wallet address: %v\n:", key.Address.Hex())

	client, err := jsonrpc.NewClient(config.GlobalConfig.ChainRpc)
	if err != nil {
		return err
	}

	didContract := orangeDID.NewOrangePubkeysManager(web3.HexToAddress(config.GlobalConfig.ContractAddress), client)
	if err != nil {
		return err
	}
	didContract.Contract().SetFrom(web3.Address(key.Address))

	pubkey, err := didContract.GetDIDPublick(web3.BytesToAddress(key.Address[:]))
	if err != nil {
		log.Errorf("GetDIDPublick failed: %v", err)
		return err
	}
	fmt.Printf("pubkey: %v\n", hexutil.Encode(pubkey))

	privkey := hexutil.Encode(crypto.FromECDSA(key.PrivateKey))
	signer := contract.NewSigner(privkey, client, uint64(config.GlobalConfig.ChainId))

	pubkeyBytes := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Printf("pubkey is %s\n", hexutil.Encode(pubkeyBytes))

	signer.Submit = true
	receipt := didContract.RegisterDID(pubkeyBytes).Sign(signer).SendTransaction(signer)
	fmt.Printf("txhash:%s\n", receipt.TransactionHash.String())
	return nil
}
