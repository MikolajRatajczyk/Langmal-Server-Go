package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/gin-gonic/gin"
)

type ResultsController struct {
	ResultService services.ResultServiceInterface
	JwtUtil       utils.JWTUtilInterface
}

func (rc *ResultsController) SaveResult(ctx *gin.Context) {
	var request saveRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return
	}

	tokenString, err := utils.ExtractToken(ctx.Request.Header)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	accountId, ok := rc.extractAccountId(tokenString)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to extract account ID from the token.",
		})
	}

	result := models.ResultEntity{
		Correct:   request.Correct,
		Wrong:     request.Wrong,
		QuizId:    request.QuizId,
		CreatedAt: request.CreatedAt,
		AccountId: accountId,
	}
	saved := rc.ResultService.Save(result, accountId)
	if saved {
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Result saved.",
		})
	} else {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save the result.",
		})
	}
}

func (rc *ResultsController) GetResults(ctx *gin.Context) {
	tokenString, err := utils.ExtractToken(ctx.Request.Header)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	accountId, ok := rc.extractAccountId(tokenString)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	resultDtos := rc.ResultService.Find(accountId)
	ctx.JSON(http.StatusOK, resultDtos)
}

func (rc *ResultsController) extractAccountId(tokenString string) (string, bool) {
	claims, ok := rc.JwtUtil.Claims(tokenString)
	return claims.Subject, ok
}

type saveRequest struct {
	Correct   int    `json:"correct" binding:"number"`
	Wrong     int    `json:"wrong" binding:"number"`
	QuizId    string `json:"quiz_id" binding:"required"`
	CreatedAt int64  `json:"created_at" binding:"required"`
}
