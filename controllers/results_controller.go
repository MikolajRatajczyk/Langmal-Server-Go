package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/middlewares"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong result structure.",
		})
	}

	tokenString := middlewares.GetTokenString(ctx)
	if tokenString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "No token.",
		})
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
	tokenString := middlewares.GetTokenString(ctx)
	resultDtos := rc.resultService.Find(tokenString)
	ctx.JSON(http.StatusOK, resultDtos)
}
