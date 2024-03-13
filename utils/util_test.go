package utils

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/test-go/testify/assert"
)

func TestDecryptMessage(t *testing.T) {
	msg := "0x041813810d16c639f94be4eea05f192f9627ee4b5716a4ce37547fa795fac4a1837770b7328d07772cbbabbf6e35b036f38b99e48057080bd12dcd0d48340865dae5273c93feb9541a4df8f5041d38e0baf30845072ea3fb4e280cb46ca61f31be93812b1bdcc3cf97465155d4dcd3687a18b1912160a3b1627d7c85dfead20ff7c675ee5b9f2c84abac7c3169c3f00724076167d80675c5f782085b9e8164534e1d3367dbc001ae268df52652db26aa820f90bb603b917ad86e554bc2ec89dc09f338f7405fd62e5fe5909c7e0b6f86ead5ee69a44a5d016874a6e263c1d6722effcbb6ec125683c5e0b94e3edad2114aab28ade154ce79884ee85afb3b35f160b330dbd5f3ca972eeaab4658580d514ac337ac96bceb178597e44778893af3f9359a2516f811cb3787f3a1a4e46274f5e4be38b4af8728bdb9e5fcbe29955103d26a183ef1cbd7b1bb2cc71a316f98d4e9c3b1c7f1ada407e964e7784cfa59d2e9b4470aeb91a0ff534fc24ad7a139ba1a54a13479bf8b80744254249e6fccb8a498ac4291aab3c0d90d1f678876b5727d3cb398972672fb21dc4111add4362e0db349a3b9a9aa8f6d86959c14cbfebce1bb8961200a95dfff9326"
	privateKey := "0x"
	msgBytes, _ := hexutil.Decode(msg)
	privateKeyBytes, _ := hexutil.Decode(privateKey)
	prk, _ := crypto.ToECDSA(privateKeyBytes)
	bts, err := DecryptMessage(msgBytes, prk)
	assert.Nil(t, err)
	fmt.Printf("%s\n", string(bts))
}
