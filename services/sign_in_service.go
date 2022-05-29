package services

import (
	"errors"
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
)

type SignInServiceInterface interface {
	//	Returns JWT token
	SignIn(credentialsDto entities.CredentialsDto) (string, error)
}

func NewSingInService(credentialsRepository repositories.CredentialsRepositoryInterface) SignInServiceInterface {
	return &signInService{
		credentialsRepository: credentialsRepository,
	}
}

type signInService struct {
	credentialsRepository repositories.CredentialsRepositoryInterface
}

func (sis *signInService) SignIn(credentialsDto entities.CredentialsDto) (string, error) {
	email := credentialsDto.Email
	credentials := sis.credentialsRepository.Find(email)

	cryptoUtil := utils.NewCryptoUtil()
	isAuthenticated := cryptoUtil.Compare(credentialsDto.Password, credentials.PasswordHash)
	if isAuthenticated {
		id := credentials.Id
		jwtUtil := utils.NewJWTUtil()
		jwtToken := jwtUtil.GenerateToken(id)
		return jwtToken, nil
	} else {
		log.Println("User does not exists!")
		return "", errors.New("User does not exist!")
	}
}
