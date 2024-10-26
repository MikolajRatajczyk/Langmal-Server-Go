package services

import (
	"errors"

	"github.com/MikolajRatajczyk/Langmal-Server-Go/models"
	"github.com/MikolajRatajczyk/Langmal-Server-Go/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server-Go/utils"
	"github.com/google/uuid"
)

var ErrUserAlreadyExists = errors.New("user already exists")
var ErrNoUser = errors.New("user does not exist")
var ErrFailedToCreateUser = errors.New("failed to create a user")
var ErrNotMatchingPasswords = errors.New("passwords don't match")
var ErrFailedToGenerateJwt = errors.New("failed to generate a JWT")

type UserServiceInterface interface {
	Register(email string, password string) error
	Login(email string, password string) (token string, err error)
}

func NewUserService(
	userRepo repositories.UserRepoInterface,
	cryptoUtil utils.CryptoUtil,
	jwtUtil utils.JwtUtilInterface) UserServiceInterface {
	return &userService{
		userRepo:   userRepo,
		cryptoUtil: cryptoUtil,
		jwtUtil:    jwtUtil,
	}
}

type userService struct {
	userRepo   repositories.UserRepoInterface
	cryptoUtil utils.CryptoUtil
	jwtUtil    utils.JwtUtilInterface
}

func (as *userService) Register(email string, password string) error {
	_, userExist := as.userRepo.Find(email)
	if userExist {
		return ErrUserAlreadyExists
	}

	hashedPassword, err := as.cryptoUtil.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.UserEntity{
		Id:           uuid.New().String(),
		Email:        email,
		PasswordHash: hashedPassword,
	}

	success := as.userRepo.Create(user)
	if !success {
		return ErrFailedToCreateUser
	}

	return nil
}

func (as *userService) Login(email string, password string) (string, error) {
	user, found := as.userRepo.Find(email)
	if !found {
		return "", ErrNoUser
	}

	isAuthenticated := as.cryptoUtil.ComparePassword(password, user.PasswordHash)
	if !isAuthenticated {
		return "", ErrNotMatchingPasswords
	}

	token, err := as.jwtUtil.Generate(user.Id)
	if err != nil {
		return "", ErrFailedToGenerateJwt
	}

	return token, nil
}
