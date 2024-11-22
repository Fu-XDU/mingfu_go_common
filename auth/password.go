package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/labstack/gommon/random"
	"golang.org/x/crypto/pbkdf2"
)

func PasswordHash(password string, salt string) (string, string) {
	// random salt
	if salt == "" {
		salt = random.String(32)
	}

	dk := pbkdf2.Key([]byte(password), []byte(salt), 4096, 32, sha256.New)
	hash := hex.EncodeToString(dk)
	return salt, hash
}
