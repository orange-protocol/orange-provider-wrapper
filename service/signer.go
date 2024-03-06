package service

import (
	"fmt"
	"orange-provider-wrapper/log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

type SignerService struct {
	ks  *keystore.KeyStore
	pwd string
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

	GlobalSignerService = &SignerService{ks: capitalKeyStore, pwd: walletPwd}
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
