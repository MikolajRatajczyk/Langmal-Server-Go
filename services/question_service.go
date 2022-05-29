package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
)

type QuestionService interface {
	FindAll() []entities.QuestionDto
}

func NewQuestionService(repo repositories.QuestionRepository) QuestionService {
	return &questionServiceImpl{
		repo: repo,
	}
}

type questionServiceImpl struct {
	repo repositories.QuestionRepository
}

func (qs *questionServiceImpl) FindAll() []entities.QuestionDto {
	return qs.repo.FindAll()
}
