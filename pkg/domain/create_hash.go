package domain

import (
	"crypto/sha256"
	"encoding/hex"
)

// CreateHash returns the secret received as a hash
func CreateHash(secret string) string {
	secretHash := sha256.New()
	secretHash.Write([]byte(secret))
	secretHashFinal := hex.EncodeToString(secretHash.Sum(nil))
	return secretHashFinal
}
