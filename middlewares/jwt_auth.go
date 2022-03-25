package middlewares

import (
	"log"
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//	Validates the token from the http request, returning 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const bearerSchemaLen = len("Bearer ")
		authHeader := ctx.GetHeader("Authorization")

		if len(authHeader) < bearerSchemaLen {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
			return
		}

		tokenString := authHeader[bearerSchemaLen:]
		token, err := service.NewJWTService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("claims.name: ", claims["name"])
			log.Println("claims.iss: ", claims["iss"])
			log.Println("claims.iat: ", claims["iat"])
		} else {
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		}
	}
}
