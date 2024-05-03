package generator

import (
	"strings"

	"github.com/devlongs/vanity-address-generator/util"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateVanityAddress(prefix string) (string, string) {
    for {
        privateKey, err := crypto.GenerateKey()
        if err != nil {
            panic(err)
        }

        address := util.PublicKeyToAddress(privateKey.PublicKey)
        if strings.HasPrefix(address.Hex()[2:], prefix) {
            return address.Hex(), util.PrivateKeyToString(privateKey)
        }
    }
}