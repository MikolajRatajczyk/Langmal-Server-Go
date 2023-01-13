package services

import (
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
)

var accountDto = models.AccountDto{
	Email:    "foo@foo.com",
	Password: "foo",
}

var account = models.Account{
	Id:           "123",
	Email:        "foo@foo.com",
	PasswordHash: []byte{1},
}

func TestAccountService_RegisterIfRepoSucceeds(t *testing.T) {
	fakeRepo := &AccountRepoFake{
		isCreateAlwaysSuccess: true,
	}
	sut := NewAccountService(fakeRepo)

	success := sut.Register(accountDto)

	if !success {
		t.Error("Should not fail for successful repo")
	}

	usedAccount := fakeRepo.usedAccountInCreate

	emailsMatch := usedAccount.Email == accountDto.Email
	if !emailsMatch {
		t.Error("Emails should match")
	}

	isIdEmpty := len(usedAccount.Id) == 0
	if isIdEmpty {
		t.Error("Should not try to register an account with empty id")
	}

	isPasswordHashEmpty := len(usedAccount.PasswordHash) == 0
	if isPasswordHashEmpty {
		t.Error("Should not try to register an account with empty password hash")
	}
}

func TestAccountService_RegisterIfRepoFails(t *testing.T) {
	fakeRepo := &AccountRepoFake{
		isCreateAlwaysSuccess: false,
	}
	sut := NewAccountService(fakeRepo)

	success := sut.Register(accountDto)

	if success {
		t.Error("Reported success despite the repo failing")
	}
}

func TestAccountService_LoginIfRepoFails(t *testing.T) {
	fakeRepo := AccountRepoFake{
		isCreateAlwaysSuccess: false,
		usedAccountInCreate:   nil,
		accountToFind:         nil,
	}
	sut := NewAccountService(&fakeRepo)

	jwt, err := sut.Login(accountDto)

	if err == nil {
		t.Error("Error should be filled for failing repo")
	}

	if len(jwt) > 0 {
		t.Error("JWT should be empty for failing repo")
	}
}

func TestAccountService_LoginIfPasswordsDontMatch(t *testing.T) {
	fakeRepo := AccountRepoFake{
		isCreateAlwaysSuccess: false,
		usedAccountInCreate:   nil,
		accountToFind:         &account,
	}
	sut := NewAccountService(&fakeRepo)

	jwt, err := sut.Login(accountDto)

	if err == nil {
		t.Error("Error should be filled for not matching passwords")
	}

	if len(jwt) > 0 {
		t.Error("JWT should be empty for not matching passwords")
	}
}

type AccountRepoFake struct {
	isCreateAlwaysSuccess bool
	// Enables spying on what was passed when calling Create
	usedAccountInCreate *models.Account
	accountToFind       *models.Account
}

func (arf *AccountRepoFake) Create(account models.Account) bool {
	arf.usedAccountInCreate = &account
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
