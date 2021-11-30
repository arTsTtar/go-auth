package utils

import (
	"crypto/rand"
	"encoding/base32"
	"golang.org/x/crypto/bcrypt"
)

func GenerateRandomPasswd() *string {
	return getToken(8)
}

func getToken(length int) *string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	token := base32.StdEncoding.EncodeToString(randomBytes)[:length]
	return &token
}

func CompareHashAndPassword(hash []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)

	if err != nil {
		return err
	}
	return nil
}
