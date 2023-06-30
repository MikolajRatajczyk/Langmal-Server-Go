package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
)

type QuizController interface {
	GetQuizzes(ctx *gin.Context)
}

func NewQuizController(service services.QuizService) QuizController {
	return &quizController{
		service: service,
	}
}

type quizController struct {
	service services.QuizService
}

func (qc *quizController) GetQuizzes(ctx *gin.Context) {
	quizzes, ok := qc.service.All()

	if ok {
		ctx.JSON(http.StatusOK, quizzes)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "No quizzes found",
		})
	}
}
