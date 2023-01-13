package utils

import (
	"errors"
	"net/http"
)

type HeaderGetter interface {
	GetHeader(key string) string
}

func ExtractToken(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")

	const bearerSchemaLen = len("Bearer ")
	if len(authHeader) <= bearerSchemaLen {
		return "", errors.New("no token")
	}

	tokenString := authHeader[bearerSchemaLen:]
	return tokenString, nil
}
