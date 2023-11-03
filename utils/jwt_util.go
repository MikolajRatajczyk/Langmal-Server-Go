package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var ErrAccountIdEmpty = errors.New("account ID is empty")

func NewJWTUtil() JwtUtil {
	const secretKey = "LANGMAL_JWT_SECRET"
	secret := os.Getenv(secretKey)
	if secret == "" {
		log.Println("Environment variable " + secretKey + " not set, using fallback!")
		secret = "secret_fallback_123"
	}

	return JwtUtil{
		secret: secret,
		issuer: "langmal.ratajczyk.dev",
	}
}

type JwtUtil struct {
	secret string
	issuer string
}

func (ju *JwtUtil) Generate(accountId string) (string, error) {
	if accountId == "" {
		return "", ErrAccountIdEmpty
	}

	const sixMonths = time.Hour * 24 * 30 * 6
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(sixMonths).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    ju.issuer,
		Subject:   accountId,
		Id:        uuid.New().String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString([]byte(ju.secret))
	return signedString, err
}

// Checks if a token is valid and not expired.
func (ju *JwtUtil) IsOk(tokenString string) bool {
	token, ok := ju.parse(tokenString)
	if !ok {
		return false
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return false
	}

	expired := claims.ExpiresAt < time.Now().Unix()
	return token.Valid && !expired
}

func (ju *JwtUtil) Claims(tokenString string) (*jwt.StandardClaims, bool) {
	token, ok := ju.parse(tokenString)
	if !ok {
		return nil, false
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, false
	}

	return claims, true
}

func (ju *JwtUtil) parse(tokenString string) (*jwt.Token, bool) {
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(parsedToken *jwt.Token) (any, error) {
		_, ok := parsedToken.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", parsedToken.Header["alg"])
		}
		return []byte(ju.secret), nil
	})

	return token, err == nil
}
