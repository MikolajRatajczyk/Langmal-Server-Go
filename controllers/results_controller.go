package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/gin-gonic/gin"
)

type ResultsControllerInterface interface {
	SaveResults(ctx *gin.Context)
	GetResults(ctx *gin.Context)
}

func NewResultsController(resultService services.ResultServiceInterface) ResultsControllerInterface {
	return &resultsController{
		resultService: resultService,
	}
}

type resultsController struct {
	resultService services.ResultServiceInterface
}

func (rc *resultsController) SaveResults(ctx *gin.Context) {
	var resultDto entities.ResultDto
	err := ctx.ShouldBind(&resultDto)
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

	saved := rc.resultService.Save(resultDto, tokenString)
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

func (rc *resultsController) GetResults(ctx *gin.Context) {
	tokenString, err := utils.ExtractToken(ctx.Request.Header)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	resultDtos, success := rc.resultService.Find(tokenString)

	if !success {
		ctx.JSON(http.StatusOK, []entities.ResultDto{})
	}

	ctx.JSON(http.StatusOK, resultDtos)
}
