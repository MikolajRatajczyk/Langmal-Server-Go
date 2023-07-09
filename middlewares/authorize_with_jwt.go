package middlewares

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/gin-gonic/gin"
)

func AuthorizeWithJWT(util utils.JWTUtilInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := utils.ExtractToken(ctx.Request.Header)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		ok := util.IsOk(tokenString)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Wrong token",
			})
			return
		}
	}
}
