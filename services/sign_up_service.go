package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/google/uuid"
)

type SignUpServiceInterface interface {
	SignUp(credentialsDto entities.CredentialsDto) bool
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

func (sus *signUpService) SignUp(credentialsDto entities.CredentialsDto) bool {
	password := credentialsDto.Password
	hashedPassword, err := sus.cryptoUtil.Hash(password)
	if err != nil {
		return false
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return false
	}

	hashedCredentials := entities.HashedCredentials{
		Id:           uuid.String(),
		Email:        credentialsDto.Email,
		PasswordHash: hashedPassword,
	}

	success := sus.credentialsRepository.Create(hashedCredentials)
	return success
}
