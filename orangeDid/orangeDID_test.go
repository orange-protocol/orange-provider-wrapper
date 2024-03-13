package orangeDID

import (
	"fmt"
	"testing"

	"github.com/laizy/web3"
	"github.com/laizy/web3/jsonrpc"
	"github.com/test-go/testify/assert"
)

func TestGetPubkeyByDid2(t *testing.T) {

	contractAddr := web3.HexToAddress("0xed371215358cea76F2B0aA0dAA258670BcB89558")
	client, err := jsonrpc.NewClient("https://polaris2.ont.io:10339")
	assert.Nil(t, err)
	c := NewOrangePubkeysManager(contractAddr, client)
	assert.Nil(t, err)
	addr, err := c.Owner()
	assert.Nil(t, err)
	fmt.Println(addr)

}

func TestInt(t *testing.T) {
	fmt.Println(1)
}

func TestGetPubkeyByDid(t *testing.T) {
	// eclient, err := ethclient.Dial("https://polygon-mumbai-pokt.nodies.app")
	// eclient, err := ethclient.Dial("http://polaris1.ont.io:20339")
	// assert.Nil(t, err)

	// height, err := eclient.BlockNumber(context.Background())
	// assert.Nil(t, err)
	// fmt.Printf("block number: %d\n", height)
	// didcontract, err := NewOrangePubkeysManager(common.HexToAddress("0xed371215358cea76F2B0aA0dAA258670BcB89558"), eclient)
	// assert.Nil(t, err)
	// key, err := didcontract.Pubkeys(nil, common.HexToAddress("0x3F3407b63cF82f781B45274f71620357629eB24a"))
	// assert.Nil(t, err)
	// fmt.Printf("%s\n", key)

	// odid, err := NewOrangeDID(common.HexToAddress("0x5E777f913d73Bb9BD07e370168e13d9E6bE6Cd3C"), eclient)
	// odid, err := NewOrangeDID(common.HexToAddress("0xed371215358cea76F2B0aA0dAA258670BcB89558"), eclient)
	// assert.Nil(t, err)
	// odid.

	// owner, err := odid.Owner(nil)
	// assert.Nil(t, err)
	// fmt.Printf("owner: %v\n", owner.String())
	// pkbts, err := odid.Pubkeys(nil, common.HexToAddress("0x3F3407b63cF82f781B45274f71620357629eB24a"))
	// assert.Nil(t, err)
	// fmt.Printf("pubkeys: %v\n", hexutil.Encode(pkbts))
}
