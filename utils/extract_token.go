package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func ExtractToken(ctx *gin.Context) (string, error) {
	const bearerSchemaLen = len("Bearer ")
	authHeader := ctx.GetHeader("Authorization")

	if len(authHeader) <= bearerSchemaLen {
		return "", errors.New("no token")
	}

	tokenString := authHeader[bearerSchemaLen:]
	return tokenString, nil
}
