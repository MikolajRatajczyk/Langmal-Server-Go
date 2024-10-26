package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server-Go/services"
	"github.com/gin-gonic/gin"
)

type QuizController struct {
	Service services.QuizServiceInterface
}

func (qc *QuizController) GetQuizzes(ctx *gin.Context) {
	quizzes, ok := qc.Service.All()

	if ok {
		ctx.JSON(http.StatusOK, quizzes)
	} else {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "No quizzes found",
		})
	}
}
