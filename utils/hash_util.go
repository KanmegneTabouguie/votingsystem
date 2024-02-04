package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashString hashes a string using SHA-256 and returns the hex-encoded hash
func HashString(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
