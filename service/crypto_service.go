package service

import (
	"golang.org/x/crypto/bcrypt"
)

type CryptoServiceInterface interface {
	Hash(password string) ([]byte, error)
	Compare(password string, hash []byte) bool
}

func NewCryptoService() CryptoServiceInterface {
	return &cryptoService{}
}

type cryptoService struct{}

func (cs *cryptoService) Hash(password string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return hashed, err
}

func (cs *cryptoService) Compare(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	success := err == nil
	return success
}
