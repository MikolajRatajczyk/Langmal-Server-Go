package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/gin-gonic/gin"
)

type ResultsController struct {
	ResultService   services.ResultServiceInterface
	ClaimsExtractor utils.ClaimsExtractorInterface
}

func (rc *ResultsController) SaveResult(ctx *gin.Context) {
	var result models.ResultWriteDto
	err := ctx.BindJSON(&result)
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

	userId, ok := rc.extractUserId(tokenString)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to extract user ID from the token.",
		})
	}

	saved := rc.ResultService.Save(result, userId)
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

	userId, ok := rc.extractUserId(tokenString)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	resultDtos := rc.ResultService.Find(userId)
	ctx.JSON(http.StatusOK, resultDtos)
}

func (rc *ResultsController) extractUserId(tokenString string) (string, bool) {
	claims, ok := rc.ClaimsExtractor.Claims(tokenString)
	if ok {
		return claims.Subject, true
	} else {
		return "", false
	}
}
