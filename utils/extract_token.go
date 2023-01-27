package utils

import (
	"errors"
	"net/http"
)

var ErrNoToken = errors.New("no token")

type HeaderGetter interface {
	GetHeader(key string) string
}

func ExtractToken(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")

	const bearerSchemaLen = len("Bearer ")
	if len(authHeader) <= bearerSchemaLen {
		return "", ErrNoToken
	}

	tokenString := authHeader[bearerSchemaLen:]
	return tokenString, nil
}
