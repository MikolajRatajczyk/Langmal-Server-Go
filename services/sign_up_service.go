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
	}
}

type signUpService struct {
	credentialsRepository repositories.CredentialsRepositoryInterface
}

func (sus *signUpService) SignUp(credentialsDto entities.CredentialsDto) bool {
	password := credentialsDto.Password
	cryptoUtil := utils.NewCryptoUtil()
	hashedPassword, err := cryptoUtil.Hash(password)
	if err != nil {
		return false
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return false
	}

	credentials := entities.Credentials{
		Id:           uuid.String(),
		Email:        credentialsDto.Email,
		PasswordHash: hashedPassword,
	}

	success := sus.credentialsRepository.Create(credentials)
	return success
}
