package service

import (
	"crypto/ecdsa"
	"fmt"
	"orange-provider-wrapper/config"
	"orange-provider-wrapper/log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

type SignerService struct {
	ks            *keystore.KeyStore
	pwd           string
	privateKey    *ecdsa.PrivateKey
	WalletAddress string
}

var GlobalSignerService *SignerService

func InitSignerService(walletFile, walletPwd string) error {
	capitalKeyStore := keystore.NewKeyStore(walletFile, keystore.StandardScryptN,
		keystore.StandardScryptP)
	acctArr := capitalKeyStore.Accounts()
	if len(acctArr) == 0 {
		return fmt.Errorf("no account found")
	}
	str := ""
	for i, v := range acctArr {
		str += fmt.Sprintf("(no.%d acc: %s), ", i+1, v.Address.String())
	}
	log.Infof("server are using accounts: [ %s ]", str)

	//get private key
	folderPath := walletFile
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
	fmt.Printf("private key:%s\n", key.PrivateKey)

	GlobalSignerService = &SignerService{ks: capitalKeyStore, pwd: walletPwd, privateKey: key.PrivateKey, WalletAddress: acctArr[0].Address.Hex()}
	return nil
}

func (ss *SignerService) SignMsg(msgBytes []byte) ([]byte, error) {

	msghash := crypto.Keccak256Hash(msgBytes)
	sig, err := ss.ks.SignHashWithPassphrase(ss.ks.Accounts()[0], ss.pwd, msghash.Bytes())
	if err != nil {
		return nil, err
	}
	if len(sig) != 65 {
		return nil, fmt.Errorf("sig length is not 65")
	}
	sig[64] += 27

	return sig, nil
}

func (ss *SignerService) GetPrivateKey() *ecdsa.PrivateKey {
	return ss.privateKey
}
