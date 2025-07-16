package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func CalculateHash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
