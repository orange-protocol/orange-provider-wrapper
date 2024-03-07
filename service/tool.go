package service

import (
	"context"
	"fmt"
	"math/big"
	"orange-provider-wrapper/config"
	"orange-provider-wrapper/log"
	"orange-provider-wrapper/orangeDid"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

	client, err := ethclient.Dial(config.GlobalConfig.ChainRpc)
	if err != nil {
		return err
	}

	didContract, err := orangeDid.NewOrangeDID(common.HexToAddress(config.GlobalConfig.ContractAddress), client)
	if err != nil {
		return err
	}

	pubkey, err := didContract.GetDIDPublick(nil, common.HexToAddress("0x3F3407b63cF82f781B45274f71620357629eB24a"))
	if err != nil {
		log.Errorf("GetDIDPublick failed: %v", err)
		return err
	}
	fmt.Printf("pubkey: %v\n", hexutil.Encode(pubkey))

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, big.NewInt(config.GlobalConfig.ChainId))
	if err != nil {
		return err
	}
	nonce, err := client.PendingNonceAt(context.Background(), key.Address)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = 300000
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("gas price:%d\n", gasPrice)
	auth.GasPrice = gasPrice

	pubkeyBytes := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Printf("pubkey is %s\n", hexutil.Encode(pubkeyBytes))
	tx, err := didContract.RegisterDID(auth, pubkeyBytes)
	if err != nil {
		return err
	}
	fmt.Printf("txhash is %s\n", tx.Hash().Hex())
	return nil
}
