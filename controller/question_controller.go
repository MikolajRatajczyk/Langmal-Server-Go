package controller

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entity"
	"github.com/MikolajRatajczyk/Langmal-Server/service"
)

type QuestionController interface {
	FindAll() []entity.Question
}

func NewQuestionController(service service.QuestionService) QuestionController {
	return &questionControllerImpl{
		service: service,
	}
}

type questionControllerImpl struct {
	service service.QuestionService
}

func (qc *questionControllerImpl) FindAll() []entity.Question {
	return qc.service.FindAll()
}
