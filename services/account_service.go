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
	Register(accountDto models.AccountDto) error
	Login(loginRequestDto models.LoginRequestDto) (models.JwtDto, error)
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

func (as *accountService) Register(accountDto models.AccountDto) error {
	email := accountDto.Email
	_, accountExist := as.accountRepo.Find(email)
	if accountExist {
		return ErrAccountAlreadyExists
	}

	hashedPassword, err := as.cryptoUtil.HashPassword(accountDto.Password)
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

func (as *accountService) Login(loginRequestDto models.LoginRequestDto) (models.JwtDto, error) {
	account, found := as.accountRepo.Find(loginRequestDto.Email)
	if !found {
		return models.JwtDto{}, ErrNoAccount
	}

	isAuthenticated := as.cryptoUtil.ComparePassword(loginRequestDto.Password, account.PasswordHash)
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
