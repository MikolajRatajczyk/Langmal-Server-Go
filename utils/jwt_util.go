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
var ErrTokenCreationFailed = errors.New("token creation failed")

type JWTUtilInterface interface {
	Generate(accountId string) (string, error)
	// Checks if a token is valid and not expired.
	IsOk(token string) bool
	ExtractAccountId(tokenString string) (string, bool)
	ExtractId(token string) (string, bool)
}

func NewJWTUtil() JWTUtilInterface {
	const secretKey = "LANGMAL_JWT_SECRET"
	secret := os.Getenv(secretKey)
	if secret == "" {
		log.Println("Environment variable " + secretKey + " not set, using fallback!")
		secret = "secret_fallback_123"
	}

	return &jwtUtil{
		secret: secret,
		issuer: "langmal.ratajczyk.dev",
	}
}

type jwtUtil struct {
	secret string
	issuer string
}

func (ju *jwtUtil) Generate(accountId string) (string, error) {
	if accountId == "" {
		return "", ErrAccountIdEmpty
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", ErrTokenCreationFailed
	}

	const sixMonths = time.Hour * 24 * 30 * 6
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(sixMonths).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    ju.issuer,
		Subject:   accountId,
		Id:        uuid.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString([]byte(ju.secret))
	return signedString, err
}

func (ju *jwtUtil) IsOk(token string) bool {
	//	validation
	claims, isValid := ju.getClaimsIfValid(token)
	if !isValid {
		return false
	}

	//	expiration check
	if time.Now().Unix() > claims.ExpiresAt {
		return false
	}

	return true
}

// TODO: It should be one func with ExtractId
func (ju *jwtUtil) ExtractAccountId(token string) (string, bool) {
	claims, ok := ju.getClaimsIfValid(token)
	if ok {
		return claims.Subject, true
	}

	return "", false
}

// TODO: It should be one func with ExtractAccountId
func (ju *jwtUtil) ExtractId(token string) (string, bool) {
	claims, ok := ju.getClaimsIfValid(token)
	if !ok {
		return "", false
	}

	id := claims.Id
	if len(id) == 0 {
		return "", false
	}

	return id, true
}

// Returns claims even if a token has expired.
func (ju *jwtUtil) getClaimsIfValid(tokenString string) (*jwt.StandardClaims, bool) {
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(parsedToken *jwt.Token) (any, error) {
		_, ok := parsedToken.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", parsedToken.Header["alg"])
		}
		return []byte(ju.secret), nil
	})
	if err != nil {
		return nil, false
	}

	if !token.Valid {
		return nil, false
	}

	return claims, true
}
