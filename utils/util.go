package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

func ETHVerifySig(from string, sig []byte, msg []byte) bool {
	fromAddr := common.HexToAddress(from)

	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if sig[64] != 0 && sig[64] != 1 {
		if sig[64] != 27 && sig[64] != 28 {
			return false
		}
		sig[64] -= 27
	}
	pubKey, err := crypto.SigToPub(crypto.Keccak256(msg), sig)
	if err != nil {
		return false
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return strings.EqualFold(fromAddr.Hex(), recoveredAddr.Hex())
}

func DIDToEthAddress(did string) (string, error) {
	arr := strings.Split(did, ":")
	if len(arr) != 2 {
		return "", fmt.Errorf("not a valid did")
	}
	if strings.EqualFold(arr[1], "etho") {
		return "", fmt.Errorf("not a etho did")
	}
	return "0x" + arr[2], nil
}

func EncryptMessageWithPubkey(msg []byte, pubkey []byte) ([]byte, error) {

	ecdsaPubkey, err := crypto.UnmarshalPubkey(pubkey)
	if err != nil {
		return nil, err
	}
	pk := ecies.ImportECDSAPublic(ecdsaPubkey)
	return ecies.Encrypt(rand.Reader, pk, msg, nil, nil)
}

func DecryptMessage(secret []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	privkey := ecies.ImportECDSA(privateKey)
	return privkey.Decrypt(secret, nil, nil)
}
