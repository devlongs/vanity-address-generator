package util

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func PublicKeyToAddress(publicKey ecdsa.PublicKey) common.Address {
    return crypto.PubkeyToAddress(publicKey)
}

func PrivateKeyToString(privateKey *ecdsa.PrivateKey) string {
    return hexutil.Encode(crypto.FromECDSA(privateKey))
}