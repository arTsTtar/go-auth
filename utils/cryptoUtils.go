package utils

import (
	"crypto/rand"
	"encoding/base32"
)

func GenerateRandomPasswd() string {
	return getToken(8)
}

func getToken(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}
