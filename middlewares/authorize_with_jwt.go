package middlewares

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/gin-gonic/gin"
)

func AuthorizeWithJWT(util utils.JWTUtilInterface, blockedTokensRepo repositories.BlockedTokensRepoInterface) gin.HandlerFunc {
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

		tokenId, ok := util.ExtractId(tokenString)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token doesn't contain an ID",
			})
			return
		}

		isBlocked := blockedTokensRepo.IsBlocked(tokenId)
		if isBlocked {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token is blocked",
			})
			return
		}
	}
}
