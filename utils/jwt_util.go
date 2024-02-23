package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var ErrUserIdEmpty = errors.New("user ID is empty")

type ClaimsExtractorInterface interface {
	Claims(tokenString string) (*jwt.RegisteredClaims, bool)
}

type JwtUtilInterface interface {
	Generate(userId string) (string, error)
	// Checks if a token is valid and not expired.
	IsOk(tokenString string) bool
	ClaimsExtractorInterface
}

func NewJWTUtil(secret string) JwtUtilInterface {
	if secret == "" {
		const escape = "\033["
		const setRed = escape + "31m"
		const setDefault = escape + "0m"
		const message = "JWT secret not set, using fallback!"
		log.Println(setRed + message + setDefault)

		secret = "secret_fallback_123"
	}

	return &jwtUtil{
		secret:        secret,
		issuer:        "langmal.ratajczyk.dev",
		signingMethod: jwt.SigningMethodHS256,
	}
}

type jwtUtil struct {
	secret        string
	issuer        string
	signingMethod *jwt.SigningMethodHMAC
}

func (ju *jwtUtil) Generate(userId string) (string, error) {
	if userId == "" {
		return "", ErrUserIdEmpty
	}

	now := time.Now()
	const sixMonths = time.Hour * 24 * 30 * 6
	claims := jwt.RegisteredClaims{
		Issuer:    ju.issuer,
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(now.Add(sixMonths)),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        uuid.New().String(),
	}

	token := jwt.NewWithClaims(ju.signingMethod, claims)

	signedString, err := token.SignedString([]byte(ju.secret))
	return signedString, err
}

func (ju *jwtUtil) IsOk(tokenString string) bool {
	token, ok := ju.parse(tokenString)
	if !ok {
		return false
	}

	expirationDate, err := token.Claims.GetExpirationTime()
	if err != nil {
		return false
	}

	expired := time.Now().After(expirationDate.Time)
	return token.Valid && !expired
}

func (ju *jwtUtil) Claims(tokenString string) (*jwt.RegisteredClaims, bool) {
	token, ok := ju.parse(tokenString)
	if !ok {
		return nil, false
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	return claims, ok
}

func (ju *jwtUtil) parse(tokenString string) (*jwt.Token, bool) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(_ *jwt.Token) (any, error) { return []byte(ju.secret), nil },
		jwt.WithValidMethods([]string{ju.signingMethod.Name}))

	return token, err == nil
}
