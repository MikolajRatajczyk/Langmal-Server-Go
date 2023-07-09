package services

import (
	"errors"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
)

var accountDto = models.AccountDto{
	Email:    "foo@foo.com",
	Password: "foo",
}

var loginRequestDto = models.LoginRequestDto{
	Email:    "foo@foo.com",
	Password: "foo",
	DeviceId: "123",
}

var cryptoUtil = utils.NewCryptoUtil()
var jwtUtil = utils.NewJWTUtil()

func TestAccountService_RegisterIfRepoSucceeds(t *testing.T) {
	accountRepoFake := AccountRepoFake{
		isCreateAlwaysSuccess: true,
	}
	refreshRepoFake := RefreshRepoFake{
		isCreateAlwaysSuccess: true,
	}
	sut := NewAccountService(
		&accountRepoFake,
		&refreshRepoFake,
		cryptoUtil,
		jwtUtil)

	err := sut.Register(accountDto)

	if err != nil {
		t.Error("Should not fail for successful repos")
	}
}

func TestAccountService_RegisterIfRepoFails(t *testing.T) {
	accountRepoFake := AccountRepoFake{
		isCreateAlwaysSuccess: false,
	}
	refreshRepoFake := RefreshRepoFake{
		isCreateAlwaysSuccess: false,
	}
	sut := NewAccountService(
		&accountRepoFake,
		&refreshRepoFake,
		cryptoUtil,
		jwtUtil)

	err := sut.Register(accountDto)

	if err == nil {
		t.Error("Should fail for failing repos")
	}
}

func TestAccountService_LoginIfRepoFails(t *testing.T) {
	accountRepoFake := AccountRepoFake{
		isCreateAlwaysSuccess: false,
		accountToFind:         nil,
	}
	refreshRepoFake := RefreshRepoFake{
		isCreateAlwaysSuccess: false,
	}
	sut := NewAccountService(
		&accountRepoFake,
		&refreshRepoFake,
		cryptoUtil,
		jwtUtil)

	jwt, err := sut.Login(loginRequestDto)

	if err == nil {
		t.Error("Should fail for failing repos")
	}

	if len(jwt.Refresh) > 0 && len(jwt.Access) > 0 {
		t.Error("JWTs should be empty for failing repos")
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
	refreshRepoFake := RefreshRepoFake{
		isCreateAlwaysSuccess: false,
	}
	sut := NewAccountService(
		&accountRepoFake,
		&refreshRepoFake,
		cryptoUtil,
		jwtUtil)

	jwt, err := sut.Login(loginRequestDto)

	if !errors.Is(err, ErrNotMatchingPasswords) {
		t.Error("Expected not matching passwords error")
	}

	if len(jwt.Refresh) > 0 && len(jwt.Access) > 0 {
		t.Error("JWTs should be empty for not matching passwords")
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

type RefreshRepoFake struct {
	isCreateAlwaysSuccess bool
}

func (rrf *RefreshRepoFake) Create(tokenHash []byte, accountId string, deviceId string) bool {
	return rrf.isCreateAlwaysSuccess
}

func (rrf *RefreshRepoFake) Find(accountId string) ([]models.AssociatedToken, bool) {
	panic("unimplemented")
}
