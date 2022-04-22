package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
)

type QuestionController interface {
	Questions(ctx *gin.Context)
}

func NewQuestionController(service services.QuestionService) QuestionController {
	return &questionControllerImpl{
		service: service,
	}
}

type questionControllerImpl struct {
	service services.QuestionService
}

func (qc *questionControllerImpl) Questions(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, qc.service.FindAll())
}
