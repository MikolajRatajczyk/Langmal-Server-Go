package middlewares

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/gin-gonic/gin"
)

func AuthorizeWithJWT(util utils.JwtUtil, blockedTokensRepo repositories.BlockedTokensRepoInterface) gin.HandlerFunc {
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

		isBlocked := blockedTokensRepo.IsBlocked(claims.Id)
		if isBlocked {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token is blocked",
			})
			return
		}
	}
}
