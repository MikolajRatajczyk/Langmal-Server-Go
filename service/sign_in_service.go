package service

import (
	"errors"
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entity"
	"github.com/MikolajRatajczyk/Langmal-Server/repository"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
)

type SignInServiceInterface interface {
	SignIn(credentials entity.Credentials) (string, error)
}

func NewSingInService(credentialsRepository repository.CredentialsRepositoryInterface,
	cryptoUtil utils.CryptoUtilInterface,
	jwtService JWTServiceInterface) SignInServiceInterface {
	return &signInService{
		credentialsRepository: credentialsRepository,
		cryptoUtil:            cryptoUtil,
		jwtService:            jwtService,
	}
}

type signInService struct {
	credentialsRepository repository.CredentialsRepositoryInterface
	cryptoUtil            utils.CryptoUtilInterface
	jwtService            JWTServiceInterface
}

func (sis *signInService) SignIn(credentials entity.Credentials) (string, error) {
	username := credentials.Username
	hashedCredentials := sis.credentialsRepository.Find(username)
	isAuthenticated := sis.cryptoUtil.Compare(credentials.Password, hashedCredentials.PasswordHash)
	if isAuthenticated {
		jwtToken := sis.jwtService.GenerateToken(username)
		return jwtToken, nil
	} else {
		log.Println("User does not exists!")
		return "", errors.New("User does not exist!")
	}
}
