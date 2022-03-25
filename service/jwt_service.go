package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTServiceInterface interface {
	GenerateToken(username string) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

func NewJWTService() JWTServiceInterface {
	return &jwtService{
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

type jwtService struct {
	secret string
	issuer string
}

func (js *jwtService) GenerateToken(username string) string {
	claims := &jwtCustomClaims{
		Name: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:   js.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//	generate encoded token using the secret
	t, err := token.SignedString([]byte(js.secret))
	if err != nil {
		panic(err)
	}
	return t
}

func (js *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//	signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok == false {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(js.secret), nil
	})
}

//	Custom claims extending standard ones
type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}
