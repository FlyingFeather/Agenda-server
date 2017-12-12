package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha(password string) string {
	r := sha1.Sum([]byte(password))
	return hex.EncodeToString(r[:])
}
