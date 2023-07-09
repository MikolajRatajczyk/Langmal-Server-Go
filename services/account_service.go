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
var ErrFailedToGenerateRefreshJwt = errors.New("failed to generate a refresh JWT")
var ErrFailedToGenerateAccessJwt = errors.New("failed to generate an access JWT")
var ErrFailedToStoreRefreshJwt = errors.New("failed to store a refresh JWT")
var ErrFailedToCreateNewAccessToken = errors.New("failed to create (refresh) a new access token")

type AccountServiceInterface interface {
	Register(accountDto models.AccountDto) error
	Login(loginRequestDto models.LoginRequestDto) (models.JwtTokenPairDto, error)
	NewAccessToken(refreshToken string) (accessToken string, err error)
}

func NewAccountService(
	accountRepo repositories.AccountRepoInterface,
	refreshTokenRepo repositories.RefreshTokenRepoInterface,
	cryptoUtil utils.CryptoUtilInterface,
	jwtUtil utils.JWTUtilInterface) AccountServiceInterface {
	return &accountService{
		accountRepo:      accountRepo,
		refreshTokenRepo: refreshTokenRepo,
		cryptoUtil:       cryptoUtil,
		jwtUtil:          jwtUtil,
	}
}

type accountService struct {
	accountRepo      repositories.AccountRepoInterface
	refreshTokenRepo repositories.RefreshTokenRepoInterface
	cryptoUtil       utils.CryptoUtilInterface
	jwtUtil          utils.JWTUtilInterface
}

func (as *accountService) Register(accountDto models.AccountDto) error {
	_, accountExist := as.accountRepo.Find(accountDto.Email)
	if accountExist {
		return ErrAccountAlreadyExists
	}

	password := accountDto.Password
	hashedPassword, err := as.cryptoUtil.HashPassword(password)
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

func (as *accountService) Login(loginRequestDto models.LoginRequestDto) (models.JwtTokenPairDto, error) {
	email := loginRequestDto.Email
	account, ok := as.accountRepo.Find(email)
	if !ok {
		return models.JwtTokenPairDto{}, ErrNoAccount
	}

	isAuthenticated := as.cryptoUtil.ComparePassword(loginRequestDto.Password, account.PasswordHash)
	if !isAuthenticated {
		return models.JwtTokenPairDto{}, ErrNotMatchingPasswords
	}

	accountId := account.Id

	refreshJwt, err := as.jwtUtil.GenerateRefreshToken(accountId)
	if err != nil {
		return models.JwtTokenPairDto{}, ErrFailedToGenerateRefreshJwt
	}

	accessJwt, err := as.jwtUtil.GenerateAccessToken(accountId)
	if err != nil {
		return models.JwtTokenPairDto{}, ErrFailedToGenerateAccessJwt
	}

	//	needed later for generating a new access token
	stored := as.storeRefreshJwt(refreshJwt, accountId, loginRequestDto.DeviceId)
	if !stored {
		return models.JwtTokenPairDto{}, ErrFailedToStoreRefreshJwt
	}

	tokenPair := models.JwtTokenPairDto{
		Refresh: refreshJwt,
		Access:  accessJwt,
	}

	return tokenPair, nil
}

func (as *accountService) NewAccessToken(refreshToken string) (accessToken string, err error) {
	//	a new access token should not be created for an outdated refresh token
	isRefreshTokenOk := as.jwtUtil.IsRefreshTokenOk(refreshToken)
	if !isRefreshTokenOk {
		return "", ErrFailedToCreateNewAccessToken
	}

	accountId, ok := as.jwtUtil.GetAccountId(refreshToken)
	if !ok {
		return "", ErrFailedToCreateNewAccessToken
	}

	//	important to check if the token is still stored for a given account ID in the DB
	isStored := as.isRefreshJwtStored(refreshToken, accountId)
	if !isStored {
		return "", ErrFailedToCreateNewAccessToken
	}

	newAccessToken, err := as.jwtUtil.GenerateAccessToken(accountId)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

func (as *accountService) storeRefreshJwt(refreshJwt string, accountId string, deviceId string) bool {
	hash, err := as.cryptoUtil.HashToken(refreshJwt)
	if err != nil {
		return false
	}

	ok := as.refreshTokenRepo.Create(hash, accountId, deviceId)
	return ok
}

func (as *accountService) isRefreshJwtStored(tokenString string, accountId string) bool {
	associatedTokens, ok := as.refreshTokenRepo.Find(accountId)
	if !ok {
		return false
	}

	tokensMatch := false
	for _, associatedToken := range associatedTokens {
		currentTokensMatch := as.cryptoUtil.CompareToken(tokenString, associatedToken.TokenHash)
		if currentTokensMatch {
			tokensMatch = true
			break
		}
	}

	return tokensMatch
}
