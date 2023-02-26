package utils

import (
	"crypto/sha256"

	"golang.org/x/crypto/bcrypt"
)

type CryptoUtilInterface interface {
	HashPassword(password string) ([]byte, error)
	HashToken(tokenString string) ([]byte, error)

	ComparePassword(password string, hash []byte) bool
	CompareToken(tokenString string, hash []byte) bool
}

func NewCryptoUtil() CryptoUtilInterface {
	return &cryptoUtil{}
}

type cryptoUtil struct{}

func (cu *cryptoUtil) HashPassword(password string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return hashed, err
}

func (cu *cryptoUtil) HashToken(tokenString string) ([]byte, error) {
	shorter := cu.shorten(tokenString)
	hashed, err := bcrypt.GenerateFromPassword(shorter, 8)
	return hashed, err
}

func (cu *cryptoUtil) ComparePassword(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	success := err == nil
	return success
}

func (cu *cryptoUtil) CompareToken(tokenString string, hash []byte) bool {
	shorter := cu.shorten(tokenString)
	err := bcrypt.CompareHashAndPassword(hash, shorter)
	success := err == nil
	return success
}

// Can be used when a value is too long.
// For example bcrypt.GenerateFromPassword accepts max 72 bytes.
// This shortens the value to 32 bytes.
func (*cryptoUtil) shorten(value string) []byte {
	sha256hash := sha256.Sum256([]byte(value))
	return sha256hash[:]
}
