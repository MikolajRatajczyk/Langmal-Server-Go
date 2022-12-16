package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/google/uuid"
)

type RegisterServiceInterface interface {
	Register(credentialsDto entities.CredentialsDto) bool
}

func NewRegisterService(credentialsRepository repositories.CredentialsRepositoryInterface) RegisterServiceInterface {
	return &registerService{
		credentialsRepository: credentialsRepository,
	}
}

type registerService struct {
	credentialsRepository repositories.CredentialsRepositoryInterface
}

func (rs *registerService) Register(credentialsDto entities.CredentialsDto) bool {
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

	success := rs.credentialsRepository.Create(credentials)
	return success
}
