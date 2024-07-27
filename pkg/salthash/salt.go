package salthash

import (
	"crypto/rand"
	"encoding/hex"
)

// Generate a random salt of the given length
func GenerateSalt(length uint8) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	saltString := hex.EncodeToString(bytes)
	return saltString, nil
}
