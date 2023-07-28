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

func (rc *ResultsController) SaveResults(ctx *gin.Context) {
	var resultDto models.ResultDtoSave
	err := ctx.BindJSON(&resultDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Wrong result structure.",
		})
		return
	}

	tokenString, err := utils.ExtractToken(ctx.Request.Header)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	accountId, ok := rc.JwtUtil.ExtractAccountId(tokenString)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to extract account ID from the token.",
		})
	}

	saved := rc.ResultService.Save(resultDto, accountId)
	if saved {
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Result saved.",
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
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

	accountId, ok := rc.JwtUtil.ExtractAccountId(tokenString)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to extract account ID from the token.",
		})
	}

	resultDtos := rc.ResultService.Find(accountId)
	ctx.JSON(http.StatusOK, resultDtos)
}
