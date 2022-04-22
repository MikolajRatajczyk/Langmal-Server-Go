package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTUtilInterface interface {
	GenerateToken(username string) string
	ValidateToken(tokenString string) (*jwt.Token, error)
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

func (ju *jwtUtil) GenerateToken(username string) string {
	claims := &jwtCustomClaims{
		Name: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:   ju.issuer,
			IssuedAt: time.Now().Unix(),
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
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok == false {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(ju.secret), nil
	})
}

//	Custom claims extending standard ones
type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}
