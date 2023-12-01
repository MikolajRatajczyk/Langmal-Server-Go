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
	Register(email string, password string) error
	Login(email string, password string) (token string, err error)
}

func NewAccountService(
	accountRepo repositories.AccountRepoInterface,
	cryptoUtil utils.CryptoUtil,
	jwtUtil utils.JwtUtilInterface) AccountServiceInterface {
	return &accountService{
		accountRepo: accountRepo,
		cryptoUtil:  cryptoUtil,
		jwtUtil:     jwtUtil,
	}
}

type accountService struct {
	accountRepo repositories.AccountRepoInterface
	cryptoUtil  utils.CryptoUtil
	jwtUtil     utils.JwtUtilInterface
}

func (as *accountService) Register(email string, password string) error {
	_, accountExist := as.accountRepo.Find(email)
	if accountExist {
		return ErrAccountAlreadyExists
	}

	hashedPassword, err := as.cryptoUtil.HashPassword(password)
	if err != nil {
		return err
	}

	account := models.AccountEntity{
		Id:           uuid.New().String(),
		Email:        email,
		PasswordHash: hashedPassword,
	}

	success := as.accountRepo.Create(account)
	if !success {
		return ErrFailedToCreateAccount
	}

	return nil
}

func (as *accountService) Login(email string, password string) (string, error) {
	account, found := as.accountRepo.Find(email)
	if !found {
		return "", ErrNoAccount
	}

	isAuthenticated := as.cryptoUtil.ComparePassword(password, account.PasswordHash)
	if !isAuthenticated {
		return "", ErrNotMatchingPasswords
	}

	token, err := as.jwtUtil.Generate(account.Id)
	if err != nil {
		return "", ErrFailedToGenerateJwt
	}

	return token, nil
}
