package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
)

type QuizController struct {
	Service services.QuizService
}

func (qc *QuizController) GetQuizzes(ctx *gin.Context) {
	quizzes, ok := qc.Service.All()

	if ok {
		ctx.JSON(http.StatusOK, quizzes)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "No quizzes found",
		})
	}
}
