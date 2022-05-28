package services

import (
	"errors"
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
)

type SignInServiceInterface interface {
	SignIn(credentialsDto entities.CredentialsDto) (string, error)
}

func NewSingInService(credentialsRepository repositories.CredentialsRepositoryInterface) SignInServiceInterface {
	return &signInService{
		credentialsRepository: credentialsRepository,
		cryptoUtil:            utils.NewCryptoUtil(),
		jwtUtil:               utils.NewJWTUtil(),
	}
}

type signInService struct {
	credentialsRepository repositories.CredentialsRepositoryInterface
	cryptoUtil            utils.CryptoUtilInterface
	jwtUtil               utils.JWTUtilInterface
}

func (sis *signInService) SignIn(credentialsDto entities.CredentialsDto) (string, error) {
	username := credentialsDto.Username
	hashedCredentials := sis.credentialsRepository.Find(username)
	isAuthenticated := sis.cryptoUtil.Compare(credentialsDto.Password, hashedCredentials.PasswordHash)
	if isAuthenticated {
		jwtToken := sis.jwtUtil.GenerateToken(username)
		return jwtToken, nil
	} else {
		log.Println("User does not exists!")
		return "", errors.New("User does not exist!")
	}
}
