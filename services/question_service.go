package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
)

type QuestionService interface {
	FindAll() []entities.Question
}

func NewQuestionService(repo repositories.QuestionRepository) QuestionService {
	return &questionServiceImpl{
		repo: repo,
	}
}

type questionServiceImpl struct {
	repo repositories.QuestionRepository
}

func (qs *questionServiceImpl) FindAll() []entities.Question {
	return qs.repo.FindAll()
}
