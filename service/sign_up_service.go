package service

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entity"
	"github.com/MikolajRatajczyk/Langmal-Server/repository"
)

type SignUpServiceInterface interface {
	SignUp(credentials entity.Credentials) bool
}

func NewSignUpService(credentialsRepository repository.CredentialsRepositoryInterface,
	cryptoService CryptoServiceInterface) SignUpServiceInterface {
	return &signUpService{
		credentialsRepository: credentialsRepository,
		cryptoService:         cryptoService,
	}
}

type signUpService struct {
	credentialsRepository repository.CredentialsRepositoryInterface
	cryptoService         CryptoServiceInterface
}

func (sus *signUpService) SignUp(credentials entity.Credentials) bool {
	password := credentials.Password
	hashedPassword, err := sus.cryptoService.Hash(password)
	if err != nil {
		return false
	}

	hashedCredentials := entity.HashedCredentials{
		Username:     credentials.Username,
		PasswordHash: hashedPassword,
	}

	success := sus.credentialsRepository.Create(hashedCredentials)
	return success
}
