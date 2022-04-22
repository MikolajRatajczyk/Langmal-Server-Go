package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
)

type SignUpServiceInterface interface {
	SignUp(credentials entities.Credentials) bool
}

func NewSignUpService(credentialsRepository repositories.CredentialsRepositoryInterface) SignUpServiceInterface {
	return &signUpService{
		credentialsRepository: credentialsRepository,
		cryptoUtil:            utils.NewCryptoUtil(),
	}
}

type signUpService struct {
	credentialsRepository repositories.CredentialsRepositoryInterface
	cryptoUtil            utils.CryptoUtilInterface
}

func (sus *signUpService) SignUp(credentials entities.Credentials) bool {
	password := credentials.Password
	hashedPassword, err := sus.cryptoUtil.Hash(password)
	if err != nil {
		return false
	}

	hashedCredentials := entities.HashedCredentials{
		Username:     credentials.Username,
		PasswordHash: hashedPassword,
	}

	success := sus.credentialsRepository.Create(hashedCredentials)
	return success
}
