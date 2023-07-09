package services

import (
	"errors"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
)

var (
	cryptoUtil = utils.NewCryptoUtil()
	jwtUtil    = utils.NewJWTUtil()

	credentialsDto = models.CredentialsDto{
		Email:    "foo@foo.com",
		Password: "foo",
	}
)

func TestAccountService_RegisterIfRepoSucceeds(t *testing.T) {
	accountRepoFake := AccountRepoFake{
		isCreateAlwaysSuccess: true,
	}
	sut := NewAccountService(
		&accountRepoFake,
		cryptoUtil,
		jwtUtil)

	err := sut.Register(credentialsDto)

	if err != nil {
		t.Error("Should not fail for successful repo")
	}
}

func TestAccountService_RegisterIfRepoFails(t *testing.T) {
	accountRepoFake := AccountRepoFake{
		isCreateAlwaysSuccess: false,
	}
	sut := NewAccountService(
		&accountRepoFake,
		cryptoUtil,
		jwtUtil)

	err := sut.Register(credentialsDto)

	if err == nil {
		t.Error("Should fail for failing repo")
	}
}

func TestAccountService_LoginIfRepoFails(t *testing.T) {
	accountRepoFake := AccountRepoFake{
		isCreateAlwaysSuccess: false,
		accountToFind:         nil,
	}
	sut := NewAccountService(
		&accountRepoFake,
		cryptoUtil,
		jwtUtil)

	jwt, err := sut.Login(credentialsDto)

	if err == nil {
		t.Error("Should fail for failing repo")
	}

	if len(jwt.Token) > 0 {
		t.Error("JWT should be empty for failing repo")
	}
}

func TestAccountService_LoginIfPasswordsDontMatch(t *testing.T) {
	account := models.Account{
		Id:           "123",
		Email:        "foo@foo.com",
		PasswordHash: []byte{1},
	}
	accountRepoFake := AccountRepoFake{
		isCreateAlwaysSuccess: false,
		accountToFind:         &account,
	}
	sut := NewAccountService(
		&accountRepoFake,
		cryptoUtil,
		jwtUtil)

	jwt, err := sut.Login(credentialsDto)

	if !errors.Is(err, ErrNotMatchingPasswords) {
		t.Error("Expected not matching passwords error")
	}

	if len(jwt.Token) > 0 {
		t.Error("JWT should be empty for not matching passwords")
	}
}

type AccountRepoFake struct {
	isCreateAlwaysSuccess bool
	accountToFind         *models.Account
}

func (arf *AccountRepoFake) Create(account models.Account) bool {
	return arf.isCreateAlwaysSuccess
}

func (arf *AccountRepoFake) Find(email string) (models.Account, bool) {
	if arf.accountToFind != nil {
		return *arf.accountToFind, true
	} else {
		return models.Account{}, false
	}
}

func (*AccountRepoFake) CloseDB() {}
