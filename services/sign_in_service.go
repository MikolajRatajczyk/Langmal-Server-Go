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
	email := credentialsDto.Email
	credentials := sis.credentialsRepository.Find(email)

	isAuthenticated := sis.cryptoUtil.Compare(credentialsDto.Password, credentials.PasswordHash)
	if isAuthenticated {
		id := credentials.Id
		jwtToken := sis.jwtUtil.GenerateToken(id)
		return jwtToken, nil
	} else {
		log.Println("User does not exists!")
		return "", errors.New("User does not exist!")
	}
}
