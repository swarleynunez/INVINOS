package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateAPIToken() string {

	t := make([]byte, 32)
	_, _ = rand.Read(t)

	return hex.EncodeToString(t)
}
