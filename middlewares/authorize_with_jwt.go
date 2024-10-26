package middlewares

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server-Go/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server-Go/utils"
	"github.com/gin-gonic/gin"
)

func AuthorizeWithJWT(util utils.JwtUtilInterface, blockedTokensRepo repositories.BlockedTokenRepoInterface) gin.HandlerFunc {
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
				"message": "Token is invalid or expired",
			})
			return
		}

		claims, ok := util.Claims(tokenString)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		isBlocked := blockedTokensRepo.IsBlocked(claims.ID)
		if isBlocked {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token is blocked",
			})
			return
		}
	}
}
