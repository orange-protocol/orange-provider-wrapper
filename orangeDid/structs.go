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

var OwnershipTransferredEventID = crypto.Keccak256Hash([]byte("OwnershipTransferred(address,address)"))

type OwnershipTransferredEvent struct {
	PreviousOwner web3.Address
	NewOwner      web3.Address

	Raw *web3.Log
}

var RegisterPubkeyEventID = crypto.Keccak256Hash([]byte("registerPubkey(address,bytes)"))

type RegisterPubkeyEvent struct {
	Arg0 web3.Address
	Arg1 []byte

	Raw *web3.Log
}
