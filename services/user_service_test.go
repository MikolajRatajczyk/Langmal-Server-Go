package services

import (
	"errors"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server-Go/models"
	"github.com/MikolajRatajczyk/Langmal-Server-Go/utils"
)

var (
	cryptoUtil = utils.CryptoUtil{}
	jwtUtil    = utils.NewJWTUtil("")
)

const (
	email    = "foo@foo.com"
	password = "foo"
)

func TestUserService_RegisterIfRepoSucceeds(t *testing.T) {
	userRepoFake := UserRepoFake{
		isCreateAlwaysSuccess: true,
	}
	sut := NewUserService(
		&userRepoFake,
		cryptoUtil,
		jwtUtil)

	err := sut.Register(email, password)

	if err != nil {
		t.Error("Should not fail for successful repo")
	}
}

func TestUserService_RegisterIfRepoFails(t *testing.T) {
	userRepoFake := UserRepoFake{
		isCreateAlwaysSuccess: false,
	}
	sut := NewUserService(
		&userRepoFake,
		cryptoUtil,
		jwtUtil)

	err := sut.Register(email, password)

	if err == nil {
		t.Error("Should fail for failing repo")
	}
}

func TestUserService_LoginIfRepoFails(t *testing.T) {
	userRepoFake := UserRepoFake{
		isCreateAlwaysSuccess: false,
		userToFind:            nil,
	}
	sut := NewUserService(
		&userRepoFake,
		cryptoUtil,
		jwtUtil)

	token, err := sut.Login(email, password)

	if err == nil {
		t.Error("Should fail for failing repo")
	}

	if len(token) > 0 {
		t.Error("JWT should be empty for failing repo")
	}
}

func TestUserService_LoginIfPasswordsDontMatch(t *testing.T) {
	user := models.UserEntity{
		Id:           "123",
		Email:        "foo@foo.com",
		PasswordHash: []byte{1},
	}
	userRepoFake := UserRepoFake{
		isCreateAlwaysSuccess: false,
		userToFind:            &user,
	}
	sut := NewUserService(
		&userRepoFake,
		cryptoUtil,
		jwtUtil)

	token, err := sut.Login(email, password)

	if !errors.Is(err, ErrNotMatchingPasswords) {
		t.Error("Expected not matching passwords error")
	}

	if len(token) > 0 {
		t.Error("JWT should be empty for not matching passwords")
	}
}

type UserRepoFake struct {
	isCreateAlwaysSuccess bool
	userToFind            *models.UserEntity
}

func (arf *UserRepoFake) Create(user models.UserEntity) bool {
	return arf.isCreateAlwaysSuccess
}

func (arf *UserRepoFake) Find(email string) (models.UserEntity, bool) {
	if arf.userToFind != nil {
		return *arf.userToFind, true
	} else {
		return models.UserEntity{}, false
	}
}
