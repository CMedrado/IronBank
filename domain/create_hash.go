package domain

import (
	"crypto/md5"
	"encoding/hex"
)

// CreateHash returns the secret received as a hash
func CreateHash(secret string) string {
	secretHash := md5.New()
	secretHash.Write([]byte(secret))
	secretHashFinal := hex.EncodeToString(secretHash.Sum(nil))
	return secretHashFinal
}
