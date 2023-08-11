package utils

import (
	"crypto/sha256"

	"golang.org/x/crypto/bcrypt"
)

type CryptoUtil struct{}

func (cu *CryptoUtil) HashPassword(password string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return hashed, err
}

func (cu *CryptoUtil) HashToken(tokenString string) ([]byte, error) {
	shorter := cu.shorten(tokenString)
	hashed, err := bcrypt.GenerateFromPassword(shorter, 8)
	return hashed, err
}

func (cu *CryptoUtil) ComparePassword(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	success := err == nil
	return success
}

func (cu *CryptoUtil) CompareToken(tokenString string, hash []byte) bool {
	shorter := cu.shorten(tokenString)
	err := bcrypt.CompareHashAndPassword(hash, shorter)
	success := err == nil
	return success
}

// Can be used when a value is too long.
// For example bcrypt.GenerateFromPassword accepts max 72 bytes.
// This shortens the value to 32 bytes.
func (*CryptoUtil) shorten(value string) []byte {
	sha256hash := sha256.Sum256([]byte(value))
	return sha256hash[:]
}
