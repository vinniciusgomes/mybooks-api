package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateSecureToken generates a secure token by creating 32 random bytes and encoding them to a hexadecimal string.
//
// Returns the generated secure token as a string and any error encountered during the process.
func GenerateSecureToken() (string, error) {
	bytes := make([]byte, 32)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
