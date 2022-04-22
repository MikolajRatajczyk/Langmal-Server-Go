package controllers

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
)

type QuestionController interface {
	FindAll() []entities.Question
}

func NewQuestionController(service services.QuestionService) QuestionController {
	return &questionControllerImpl{
		service: service,
	}
}

type questionControllerImpl struct {
	service services.QuestionService
}

func (qc *questionControllerImpl) FindAll() []entities.Question {
	return qc.service.FindAll()
}
