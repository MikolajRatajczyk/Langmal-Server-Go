package services

import (
	"errors"
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/google/uuid"
)

type AccountServiceInterface interface {
	Register(accountDto entities.AccountDto) bool
	//	Returns JWT token
	Login(accountDto entities.AccountDto) (string, error)
}

func NewAccountService(accountRepo repositories.AccountRepoInterface) AccountServiceInterface {
	return &accountService{
		accountRepo: accountRepo,
	}
}

type accountService struct {
	accountRepo repositories.AccountRepoInterface
}

func (as *accountService) Register(accountDto entities.AccountDto) bool {
	password := accountDto.Password
	cryptoUtil := utils.NewCryptoUtil()
	hashedPassword, err := cryptoUtil.Hash(password)
	if err != nil {
		return false
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return false
	}

	account := entities.Account{
		Id:           uuid.String(),
		Email:        accountDto.Email,
		PasswordHash: hashedPassword,
	}

	success := as.accountRepo.Create(account)
	return success
}

func (as *accountService) Login(accountDto entities.AccountDto) (string, error) {
	email := accountDto.Email
	account, ok := as.accountRepo.Find(email)
	if !ok {
		return "", errors.New("account does not exist")
	}

	cryptoUtil := utils.NewCryptoUtil()
	isAuthenticated := cryptoUtil.Compare(accountDto.Password, account.PasswordHash)
	if isAuthenticated {
		id := account.Id
		jwtUtil := utils.NewJWTUtil()
		jwtToken := jwtUtil.GenerateToken(id)
		return jwtToken, nil
	} else {
		log.Println("Account does not exists!")
		return "", errors.New("account does not exist")
	}
}
