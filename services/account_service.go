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

type AccountServiceInterface interface {
	Register(accountDto models.AccountDto) error
	Login(accountDto models.AccountDto) (jwtToken string, err error)
}

func NewAccountService(accountRepo repositories.AccountRepoInterface) AccountServiceInterface {
	return &accountService{
		accountRepo: accountRepo,
	}
}

type accountService struct {
	accountRepo repositories.AccountRepoInterface
}

func (as *accountService) Register(accountDto models.AccountDto) error {
	_, accountExist := as.accountRepo.Find(accountDto.Email)
	if accountExist {
		return ErrAccountAlreadyExists
	}

	password := accountDto.Password
	cryptoUtil := utils.NewCryptoUtil()
	hashedPassword, err := cryptoUtil.Hash(password)
	if err != nil {
		return err
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	account := models.Account{
		Id:           uuid.String(),
		Email:        accountDto.Email,
		PasswordHash: hashedPassword,
	}

	success := as.accountRepo.Create(account)
	if !success {
		return ErrFailedToCreateAccount
	}

	return nil
}

func (as *accountService) Login(accountDto models.AccountDto) (string, error) {
	email := accountDto.Email
	account, ok := as.accountRepo.Find(email)
	if !ok {
		return "", ErrNoAccount
	}

	cryptoUtil := utils.NewCryptoUtil()
	isAuthenticated := cryptoUtil.Compare(accountDto.Password, account.PasswordHash)
	if isAuthenticated {
		id := account.Id
		jwtUtil := utils.NewJWTUtil()
		jwtToken := jwtUtil.GenerateToken(id)
		return jwtToken, nil
	} else {
		return "", ErrNotMatchingPasswords
	}
}
