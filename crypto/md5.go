package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(data []byte) (md5Hex string) {
	hash := md5.Sum(data)
	md5Hex = hex.EncodeToString(hash[:])
	return
}
