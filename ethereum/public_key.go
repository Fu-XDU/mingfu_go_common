package ethereum

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func PublicKeyToAddress(key ecdsa.PublicKey) common.Address {
	return crypto.PubkeyToAddress(key)
}

func PublicKeyToHex(key *ecdsa.PublicKey, withPrefix bool) (h string) {
	h = common.Bytes2Hex(crypto.FromECDSAPub(key))
	if withPrefix {
		h = "0x" + h
	}

	return h
}

func HexToPublicKey(keyHex string) (*ecdsa.PublicKey, error) {
	return crypto.UnmarshalPubkey(common.FromHex(keyHex))
}
