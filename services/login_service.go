package services

import (
	"errors"
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
)

type LoginServiceInterface interface {
	//	Returns JWT token
	Login(credentialsDto entities.CredentialsDto) (string, error)
}

func NewLoginService(credentialsRepository repositories.CredentialsRepositoryInterface) LoginServiceInterface {
	return &loginService{
		credentialsRepository: credentialsRepository,
	}
}

type loginService struct {
	credentialsRepository repositories.CredentialsRepositoryInterface
}

func (ls *loginService) Login(credentialsDto entities.CredentialsDto) (string, error) {
	email := credentialsDto.Email
	credentials := ls.credentialsRepository.Find(email)

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
