package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var ErrAccountIdEmpty = errors.New("account ID is empty")

type JWTUtilInterface interface {
	GenerateRefreshToken(accountId string) (string, error)
	GenerateAccessToken(accountId string) (string, error)

	// Checks if a token is valid and not expired.
	IsRefreshTokenOk(tokenString string) bool
	// Checks if a token is valid and not expired.
	IsAccessTokenOk(tokenString string) bool

	GetAccountId(tokenString string) (string, bool)
}

func NewJWTUtil() JWTUtilInterface {
	return &jwtUtil{
		refreshSecret: getEnv("REFRESH_JWT_SECRET", "refresh_secret_fallback"),
		accessSecret:  getEnv("ACCESS_JWT_SECRET", "access_secret_fallback"),
		issuer:        "langmal.ratajczyk.dev",
	}
}

func getEnv(key string, fallback string) string {
	env := os.Getenv(key)
	if env == "" {
		log.Println("Environment variable " + key + " not set, using fallback!")
		env = fallback
	}
	return env
}

type jwtUtil struct {
	refreshSecret string
	accessSecret  string
	issuer        string
}

func (ju *jwtUtil) GenerateRefreshToken(accountId string) (string, error) {
	const sixMonths = time.Hour * 24 * 30 * 6
	return ju.generateToken(accountId, sixMonths, ju.refreshSecret)
}

func (ju *jwtUtil) GenerateAccessToken(accountId string) (string, error) {
	const fifteenMinutes = time.Minute * 15
	return ju.generateToken(accountId, fifteenMinutes, ju.accessSecret)
}

func (ju *jwtUtil) IsRefreshTokenOk(tokenString string) bool {
	return ju.isTokenOk(tokenString, ju.refreshSecret)
}

func (ju *jwtUtil) IsAccessTokenOk(tokenString string) bool {
	return ju.isTokenOk(tokenString, ju.accessSecret)
}

func (ju *jwtUtil) GetAccountId(tokenString string) (string, bool) {
	claims, ok := ju.getClaimsIfValid(tokenString, ju.refreshSecret)
	if ok {
		return claims.Subject, true
	}

	claims, ok = ju.getClaimsIfValid(tokenString, ju.accessSecret)
	if ok {
		return claims.Subject, true
	}

	return "", false
}

func (ju *jwtUtil) generateToken(accountId string, lifespan time.Duration, secret string) (string, error) {
	if accountId == "" {
		return "", ErrAccountIdEmpty
	}

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(lifespan).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    ju.issuer,
		Subject:   accountId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString([]byte(secret))
	return signedString, err
}

func (ju *jwtUtil) isTokenOk(tokenString string, secret string) bool {
	//	validation
	claims, isValid := ju.getClaimsIfValid(tokenString, secret)
	if !isValid {
		return false
	}

	//	expiration check
	if time.Now().Unix() > claims.ExpiresAt {
		return false
	}

	return true
}

// Returns claims even if a token has expired.
func (*jwtUtil) getClaimsIfValid(tokenString string, secret string) (*jwt.StandardClaims, bool) {
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(parsedToken *jwt.Token) (any, error) {
		_, ok := parsedToken.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", parsedToken.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, false
	}

	if !token.Valid {
		return nil, false
	}

	return claims, true
}
