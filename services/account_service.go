package services

import (
	"errors"
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/google/uuid"
)

type AccountServiceInterface interface {
	Register(credentialsDto entities.CredentialsDto) bool
	//	Returns JWT token
	Login(credentialsDto entities.CredentialsDto) (string, error)
}

func NewAccountService(credentialsRepository repositories.CredentialsRepositoryInterface) AccountServiceInterface {
	return &accountService{
		credentialsRepository: credentialsRepository,
	}
}

type accountService struct {
	credentialsRepository repositories.CredentialsRepositoryInterface
}

func (as *accountService) Register(credentialsDto entities.CredentialsDto) bool {
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

	success := as.credentialsRepository.Create(credentials)
	return success
}

func (as *accountService) Login(credentialsDto entities.CredentialsDto) (string, error) {
	email := credentialsDto.Email
	credentials := as.credentialsRepository.Find(email)

	cryptoUtil := utils.NewCryptoUtil()
	isAuthenticated := cryptoUtil.Compare(credentialsDto.Password, credentials.PasswordHash)
	if isAuthenticated {
		id := credentials.Id
		jwtUtil := utils.NewJWTUtil()
		jwtToken := jwtUtil.GenerateToken(id)
		return jwtToken, nil
	} else {
		log.Println("User does not exists!")
		return "", errors.New("user does not exist")
	}
}
