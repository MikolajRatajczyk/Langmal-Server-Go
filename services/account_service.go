package services

import (
	"errors"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/google/uuid"
)

var ErrAccountAlreadyExists = errors.New("account already exists")
var ErrNoAccount = errors.New("account does not exist")
var ErrFailedToCreateAccount = errors.New("failed to create an account")
var ErrNotMatchingPasswords = errors.New("passwords don't match")
var ErrFailedToGenerateJwt = errors.New("failed to generate a JWT")

type AccountServiceInterface interface {
	Register(credentialsDto models.CredentialsDto) error
	Login(credentialsDto models.CredentialsDto) (models.JwtDto, error)
}

func NewAccountService(
	accountRepo repositories.AccountRepoInterface,
	cryptoUtil utils.CryptoUtilInterface,
	jwtUtil utils.JWTUtilInterface) AccountServiceInterface {
	return &accountService{
		accountRepo: accountRepo,
		cryptoUtil:  cryptoUtil,
		jwtUtil:     jwtUtil,
	}
}

type accountService struct {
	accountRepo repositories.AccountRepoInterface
	cryptoUtil  utils.CryptoUtilInterface
	jwtUtil     utils.JWTUtilInterface
}

func (as *accountService) Register(credentialsDto models.CredentialsDto) error {
	email := credentialsDto.Email
	_, accountExist := as.accountRepo.Find(email)
	if accountExist {
		return ErrAccountAlreadyExists
	}

	hashedPassword, err := as.cryptoUtil.HashPassword(credentialsDto.Password)
	if err != nil {
		return err
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	account := models.Account{
		Id:           uuid.String(),
		Email:        email,
		PasswordHash: hashedPassword,
	}

	success := as.accountRepo.Create(account)
	if !success {
		return ErrFailedToCreateAccount
	}

	return nil
}

func (as *accountService) Login(credentialsDto models.CredentialsDto) (models.JwtDto, error) {
	account, found := as.accountRepo.Find(credentialsDto.Email)
	if !found {
		return models.JwtDto{}, ErrNoAccount
	}

	isAuthenticated := as.cryptoUtil.ComparePassword(credentialsDto.Password, account.PasswordHash)
	if !isAuthenticated {
		return models.JwtDto{}, ErrNotMatchingPasswords
	}

	token, err := as.jwtUtil.Generate(account.Id)
	if err != nil {
		return models.JwtDto{}, ErrFailedToGenerateJwt
	}

	tokenDto := models.JwtDto{
		Token: token,
	}

	return tokenDto, nil
}
