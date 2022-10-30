package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTUtilInterface interface {
	GenerateToken(id string) string
	ValidateToken(tokenString string) (*jwt.Token, error)
	GetUserId(tokenString string) (string, error)
}

func NewJWTUtil() JWTUtilInterface {
	return &jwtUtil{
		secret: getSecret(),
		issuer: "ratajczyk.dev",
	}
}

func getSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "jwt1234./"
	}
	return secret
}

type jwtUtil struct {
	secret string
	issuer string
}

func (ju *jwtUtil) GenerateToken(id string) string {
	claims := &jwtCustomClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:   ju.issuer,
			IssuedAt: time.Now().Unix(),
			Subject:  id,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//	generate encoded token using the secret
	t, err := token.SignedString([]byte(ju.secret))
	if err != nil {
		panic(err)
	}
	return t
}

func (ju *jwtUtil) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		//	signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(ju.secret), nil
	})
}

func (ju *jwtUtil) GetUserId(tokenString string) (string, error) {
	//	important: use &
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(ju.secret), nil
	})

	if err != nil {
		return "", err
	}

	//	important: use *
	customClaims, ok := token.Claims.(*jwtCustomClaims)

	if !ok {
		return "", fmt.Errorf("can't cast claims to custom claims")
	}

	return customClaims.Subject, nil
}

// Custom claims extending standard ones
// TODO: Extending is not needed (only standard properties are used)
type jwtCustomClaims struct {
	jwt.StandardClaims
}
