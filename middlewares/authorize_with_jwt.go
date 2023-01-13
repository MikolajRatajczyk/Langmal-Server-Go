package middlewares

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/gin-gonic/gin"
)

func AuthorizeWithJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := utils.ExtractToken(ctx.Request.Header)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		_, err = utils.NewJWTUtil().ValidateToken(tokenString)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Wrong token: " + err.Error(),
			})
			return
		}
	}
}
