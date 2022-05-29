package middlewares

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/gin-gonic/gin"
)

//	Validates the token from the http request, returning 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := GetTokenString(ctx)
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "No token",
			})
			return
		}

		_, err := utils.NewJWTUtil().ValidateToken(tokenString)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Wrong token: " + err.Error(),
			})
			return
		}
	}
}

func GetTokenString(ctx *gin.Context) string {
	const bearerSchemaLen = len("Bearer ")
	authHeader := ctx.GetHeader("Authorization")

	if len(authHeader) < bearerSchemaLen {
		return ""
	}

	tokenString := authHeader[bearerSchemaLen:]
	return tokenString
}
