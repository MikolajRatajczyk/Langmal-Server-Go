package service

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entity"
	"github.com/MikolajRatajczyk/Langmal-Server/repository"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
)

type SignUpServiceInterface interface {
	SignUp(credentials entity.Credentials) bool
}

func NewSignUpService(credentialsRepository repository.CredentialsRepositoryInterface,
	cryptoUtil utils.CryptoUtilInterface) SignUpServiceInterface {
	return &signUpService{
		credentialsRepository: credentialsRepository,
		cryptoUtil:            cryptoUtil,
	}
}

type signUpService struct {
	credentialsRepository repository.CredentialsRepositoryInterface
	cryptoUtil            utils.CryptoUtilInterface
}

func (sus *signUpService) SignUp(credentials entity.Credentials) bool {
	password := credentials.Password
	hashedPassword, err := sus.cryptoUtil.Hash(password)
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
