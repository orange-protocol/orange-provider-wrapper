package orangeDID

import (
	"fmt"
	"math/big"

	"github.com/laizy/web3"
	"github.com/laizy/web3/crypto"
)

var (
	_ = big.NewInt
	_ = fmt.Printf
	_ = web3.HexToAddress
	_ = crypto.Keccak256Hash
)

var InitializedEventID = crypto.Keccak256Hash([]byte("Initialized(uint64)"))

type InitializedEvent struct {
	Version uint64

	Raw *web3.Log
}

var RegisterCredentialEventEventID = crypto.Keccak256Hash([]byte("registerCredentialEvent(string,address,uint256,address)"))

type RegisterCredentialEventEvent struct {
	Arg0 string
	Arg1 web3.Address
	Arg2 *big.Int
	Arg3 web3.Address

	Raw *web3.Log
}

var RegisterPubkeyEventEventID = crypto.Keccak256Hash([]byte("registerPubkeyEvent(address,bytes)"))

type RegisterPubkeyEventEvent struct {
	Arg0 web3.Address
	Arg1 []byte

	Raw *web3.Log
}

var RevokeCredentialEventEventID = crypto.Keccak256Hash([]byte("revokeCredentialEvent(string)"))

type RevokeCredentialEventEvent struct {
	Arg0 string

	Raw *web3.Log
}
