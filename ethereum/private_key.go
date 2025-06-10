package ethereum

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func NewRandomPrivateKey() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

func PrivateKeyToAddress(key *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(key.PublicKey)
}

func PrivateKeyToPublicKey(key *ecdsa.PrivateKey) ecdsa.PublicKey {
	return key.PublicKey
}

func PrivateKeyToHex(key *ecdsa.PrivateKey, withPrefix bool) (h string) {
	h = common.Bytes2Hex(crypto.FromECDSA(key))
	if withPrefix {
		h = "0x" + h
	}

	return h
}

func HexToPrivateKey(keyHex string) (*ecdsa.PrivateKey, error) {
	return crypto.ToECDSA(common.FromHex(keyHex))
}
