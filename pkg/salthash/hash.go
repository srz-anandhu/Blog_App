package salthash

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hashes the password with given salt (SHA-256)
func HashPassword(password, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashBytes := hash.Sum(nil)
	hashedPass := hex.EncodeToString(hashBytes)
	return hashedPass
}
