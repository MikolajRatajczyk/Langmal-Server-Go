package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type CryptoUtilInterface interface {
	Hash(password string) ([]byte, error)
	Compare(password string, hash []byte) bool
}

func NewCryptoUtil() CryptoUtilInterface {
	return &cryptoUtil{}
}

type cryptoUtil struct{}

func (cu *cryptoUtil) Hash(password string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return hashed, err
}

func (cu *cryptoUtil) Compare(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	success := err == nil
	return success
}
