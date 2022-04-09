package service

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entity"
	"github.com/MikolajRatajczyk/Langmal-Server/repository"
)

type SignUpServiceInterface interface {
	SignUp(credentials entity.Credentials) bool
}

func NewSignUpService(credentialsRepository repository.CredentialsRepositoryInterface) SignUpServiceInterface {
	return &signUpService{
		credentialsRepository: credentialsRepository,
	}
}

type signUpService struct {
	credentialsRepository repository.CredentialsRepositoryInterface
}

func (sus *signUpService) SignUp(credentials entity.Credentials) bool {
	success := sus.credentialsRepository.Create(credentials)
	return success
}
