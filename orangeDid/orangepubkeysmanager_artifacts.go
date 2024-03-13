package orangeDID

import (
	"encoding/hex"
	"fmt"

	"github.com/laizy/web3/abi"
)

var abiOrangePubkeysManager *abi.ABI

// OrangePubkeysManagerAbi returns the abi of the OrangePubkeysManager contract
func OrangePubkeysManagerAbi() *abi.ABI {
	return abiOrangePubkeysManager
}

var binOrangePubkeysManager []byte

// OrangePubkeysManagerBin returns the bin of the OrangePubkeysManager contract
func OrangePubkeysManagerBin() []byte {
	return binOrangePubkeysManager
}

var binRuntimeOrangePubkeysManager []byte

// OrangePubkeysManagerBinRuntime returns the runtime bin of the OrangePubkeysManager contract
func OrangePubkeysManagerBinRuntime() []byte {
	return binRuntimeOrangePubkeysManager
}

func init() {
	var err error
	abiOrangePubkeysManager, err = abi.NewABI(abiOrangePubkeysManagerStr)
	if err != nil {
		panic(fmt.Errorf("cannot parse OrangePubkeysManager abi: %v", err))
	}
	if len(binOrangePubkeysManagerStr) != 0 {
		binOrangePubkeysManager, err = hex.DecodeString(binOrangePubkeysManagerStr[2:])
		if err != nil {
			panic(fmt.Errorf("cannot parse OrangePubkeysManager bin: %v", err))
		}
	}
	if len(binRuntimeOrangePubkeysManagerStr) != 0 {
		binRuntimeOrangePubkeysManager, err = hex.DecodeString(binRuntimeOrangePubkeysManagerStr[2:])
		if err != nil {
			panic(fmt.Errorf("cannot parse OrangePubkeysManager bin runtime: %v", err))
		}
	}
}

var binOrangePubkeysManagerStr = "0x608060405234801561001057600080fd5b5033600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100845760006040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161007b919061019e565b60405180910390fd5b6100938161009960201b60201c565b506101b9565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101888261015d565b9050919050565b6101988161017d565b82525050565b60006020820190506101b3600083018461018f565b92915050565b610ca3806101c86000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80639e50f99c1161005b5780639e50f99c146100c6578063c8ec764b146100f6578063db36f83714610112578063f2fde38b146101425761007d565b80631570782014610082578063715018a61461009e5780638da5cb5b146100a8575b600080fd5b61009c600480360381019061009791906107b8565b61015e565b005b6100a66101f0565b005b6100b0610204565b6040516100bd9190610823565b60405180910390f35b6100e060048036038101906100db919061083e565b61022d565b6040516100ed91906108ea565b60405180910390f35b610110600480360381019061010b919061090c565b6102cd565b005b61012c6004803603810190610127919061083e565b610356565b60405161013991906108ea565b60405180910390f35b61015c6004803603810190610157919061083e565b610427565b005b6101666104ad565b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090816101b29190610b6b565b507fae07ad10c06d8cbc762909e88a3421963a0a1ec3fc0854acf03fc676a43e364f33826040516101e4929190610c3d565b60405180910390a15050565b6101f86104ad565b6102026000610534565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6001602052806000526040600020600091509050805461024c90610984565b80601f016020809104026020016040519081016040528092919081815260200182805461027890610984565b80156102c55780601f1061029a576101008083540402835291602001916102c5565b820191906000526020600020905b8154815290600101906020018083116102a857829003601f168201915b505050505081565b80600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090816103199190610b6b565b507fae07ad10c06d8cbc762909e88a3421963a0a1ec3fc0854acf03fc676a43e364f338260405161034b929190610c3d565b60405180910390a150565b6060600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080546103a290610984565b80601f01602080910402602001604051908101604052809291908181526020018280546103ce90610984565b801561041b5780601f106103f05761010080835404028352916020019161041b565b820191906000526020600020905b8154815290600101906020018083116103fe57829003601f168201915b50505050509050919050565b61042f6104ad565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036104a15760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016104989190610823565b60405180910390fd5b6104aa81610534565b50565b6104b56105f8565b73ffffffffffffffffffffffffffffffffffffffff166104d3610204565b73ffffffffffffffffffffffffffffffffffffffff1614610532576104f66105f8565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016105299190610823565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061063f82610614565b9050919050565b61064f81610634565b811461065a57600080fd5b50565b60008135905061066c81610646565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6106c58261067c565b810181811067ffffffffffffffff821117156106e4576106e361068d565b5b80604052505050565b60006106f7610600565b905061070382826106bc565b919050565b600067ffffffffffffffff8211156107235761072261068d565b5b61072c8261067c565b9050602081019050919050565b82818337600083830152505050565b600061075b61075684610708565b6106ed565b90508281526020810184848401111561077757610776610677565b5b610782848285610739565b509392505050565b600082601f83011261079f5761079e610672565b5b81356107af848260208601610748565b91505092915050565b600080604083850312156107cf576107ce61060a565b5b60006107dd8582860161065d565b925050602083013567ffffffffffffffff8111156107fe576107fd61060f565b5b61080a8582860161078a565b9150509250929050565b61081d81610634565b82525050565b60006020820190506108386000830184610814565b92915050565b6000602082840312156108545761085361060a565b5b60006108628482850161065d565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156108a557808201518184015260208101905061088a565b60008484015250505050565b60006108bc8261086b565b6108c68185610876565b93506108d6818560208601610887565b6108df8161067c565b840191505092915050565b6000602082019050818103600083015261090481846108b1565b905092915050565b6000602082840312156109225761092161060a565b5b600082013567ffffffffffffffff8111156109405761093f61060f565b5b61094c8482850161078a565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061099c57607f821691505b6020821081036109af576109ae610955565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610a177fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826109da565b610a2186836109da565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000610a68610a63610a5e84610a39565b610a43565b610a39565b9050919050565b6000819050919050565b610a8283610a4d565b610a96610a8e82610a6f565b8484546109e7565b825550505050565b600090565b610aab610a9e565b610ab6818484610a79565b505050565b5b81811015610ada57610acf600082610aa3565b600181019050610abc565b5050565b601f821115610b1f57610af0816109b5565b610af9846109ca565b81016020851015610b08578190505b610b1c610b14856109ca565b830182610abb565b50505b505050565b600082821c905092915050565b6000610b4260001984600802610b24565b1980831691505092915050565b6000610b5b8383610b31565b9150826002028217905092915050565b610b748261086b565b67ffffffffffffffff811115610b8d57610b8c61068d565b5b610b978254610984565b610ba2828285610ade565b600060209050601f831160018114610bd55760008415610bc3578287015190505b610bcd8582610b4f565b865550610c35565b601f198416610be3866109b5565b60005b82811015610c0b57848901518255600182019150602085019450602081019050610be6565b86831015610c285784890151610c24601f891682610b31565b8355505b6001600288020188555050505b505050505050565b6000604082019050610c526000830185610814565b8181036020830152610c6481846108b1565b9050939250505056fea26469706673582212201f8755a7eae6047901139f4ffcc15480915c4431c4e9efb91f2a5a32a6c9098964736f6c63430008140033"

var binRuntimeOrangePubkeysManagerStr = "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c80639e50f99c1161005b5780639e50f99c146100c6578063c8ec764b146100f6578063db36f83714610112578063f2fde38b146101425761007d565b80631570782014610082578063715018a61461009e5780638da5cb5b146100a8575b600080fd5b61009c600480360381019061009791906107b8565b61015e565b005b6100a66101f0565b005b6100b0610204565b6040516100bd9190610823565b60405180910390f35b6100e060048036038101906100db919061083e565b61022d565b6040516100ed91906108ea565b60405180910390f35b610110600480360381019061010b919061090c565b6102cd565b005b61012c6004803603810190610127919061083e565b610356565b60405161013991906108ea565b60405180910390f35b61015c6004803603810190610157919061083e565b610427565b005b6101666104ad565b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090816101b29190610b6b565b507fae07ad10c06d8cbc762909e88a3421963a0a1ec3fc0854acf03fc676a43e364f33826040516101e4929190610c3d565b60405180910390a15050565b6101f86104ad565b6102026000610534565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6001602052806000526040600020600091509050805461024c90610984565b80601f016020809104026020016040519081016040528092919081815260200182805461027890610984565b80156102c55780601f1061029a576101008083540402835291602001916102c5565b820191906000526020600020905b8154815290600101906020018083116102a857829003601f168201915b505050505081565b80600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090816103199190610b6b565b507fae07ad10c06d8cbc762909e88a3421963a0a1ec3fc0854acf03fc676a43e364f338260405161034b929190610c3d565b60405180910390a150565b6060600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080546103a290610984565b80601f01602080910402602001604051908101604052809291908181526020018280546103ce90610984565b801561041b5780601f106103f05761010080835404028352916020019161041b565b820191906000526020600020905b8154815290600101906020018083116103fe57829003601f168201915b50505050509050919050565b61042f6104ad565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036104a15760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016104989190610823565b60405180910390fd5b6104aa81610534565b50565b6104b56105f8565b73ffffffffffffffffffffffffffffffffffffffff166104d3610204565b73ffffffffffffffffffffffffffffffffffffffff1614610532576104f66105f8565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016105299190610823565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061063f82610614565b9050919050565b61064f81610634565b811461065a57600080fd5b50565b60008135905061066c81610646565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6106c58261067c565b810181811067ffffffffffffffff821117156106e4576106e361068d565b5b80604052505050565b60006106f7610600565b905061070382826106bc565b919050565b600067ffffffffffffffff8211156107235761072261068d565b5b61072c8261067c565b9050602081019050919050565b82818337600083830152505050565b600061075b61075684610708565b6106ed565b90508281526020810184848401111561077757610776610677565b5b610782848285610739565b509392505050565b600082601f83011261079f5761079e610672565b5b81356107af848260208601610748565b91505092915050565b600080604083850312156107cf576107ce61060a565b5b60006107dd8582860161065d565b925050602083013567ffffffffffffffff8111156107fe576107fd61060f565b5b61080a8582860161078a565b9150509250929050565b61081d81610634565b82525050565b60006020820190506108386000830184610814565b92915050565b6000602082840312156108545761085361060a565b5b60006108628482850161065d565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156108a557808201518184015260208101905061088a565b60008484015250505050565b60006108bc8261086b565b6108c68185610876565b93506108d6818560208601610887565b6108df8161067c565b840191505092915050565b6000602082019050818103600083015261090481846108b1565b905092915050565b6000602082840312156109225761092161060a565b5b600082013567ffffffffffffffff8111156109405761093f61060f565b5b61094c8482850161078a565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061099c57607f821691505b6020821081036109af576109ae610955565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610a177fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826109da565b610a2186836109da565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000610a68610a63610a5e84610a39565b610a43565b610a39565b9050919050565b6000819050919050565b610a8283610a4d565b610a96610a8e82610a6f565b8484546109e7565b825550505050565b600090565b610aab610a9e565b610ab6818484610a79565b505050565b5b81811015610ada57610acf600082610aa3565b600181019050610abc565b5050565b601f821115610b1f57610af0816109b5565b610af9846109ca565b81016020851015610b08578190505b610b1c610b14856109ca565b830182610abb565b50505b505050565b600082821c905092915050565b6000610b4260001984600802610b24565b1980831691505092915050565b6000610b5b8383610b31565b9150826002028217905092915050565b610b748261086b565b67ffffffffffffffff811115610b8d57610b8c61068d565b5b610b978254610984565b610ba2828285610ade565b600060209050601f831160018114610bd55760008415610bc3578287015190505b610bcd8582610b4f565b865550610c35565b601f198416610be3866109b5565b60005b82811015610c0b57848901518255600182019150602085019450602081019050610be6565b86831015610c285784890151610c24601f891682610b31565b8355505b6001600288020188555050505b505050505050565b6000604082019050610c526000830185610814565b8181036020830152610c6481846108b1565b9050939250505056fea26469706673582212201f8755a7eae6047901139f4ffcc15480915c4431c4e9efb91f2a5a32a6c9098964736f6c63430008140033"

var abiOrangePubkeysManagerStr = `[{"inputs":[],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"OwnableInvalidOwner","type":"error"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"OwnableUnauthorizedAccount","type":"error"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"","type":"address"},{"indexed":false,"internalType":"bytes","name":"","type":"bytes"}],"name":"registerPubkey","type":"event"},{"inputs":[{"internalType":"address","name":"addr","type":"address"}],"name":"getDIDPublick","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"pubkeys","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes","name":"pubkey","type":"bytes"}],"name":"registerDID","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"wallet","type":"address"},{"internalType":"bytes","name":"pubkey","type":"bytes"}],"name":"registerDIDByOwner","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
