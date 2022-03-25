package service

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entity"
	"github.com/MikolajRatajczyk/Langmal-Server/repository"
)

type QuestionService interface {
	FindAll() []entity.Question
}

func NewQuestionService(repo repository.QuestionRepository) QuestionService {
	return &questionServiceImpl{
		repo: repo,
	}
}

type questionServiceImpl struct {
	repo repository.QuestionRepository
}

func (qs *questionServiceImpl) FindAll() []entity.Question {
	return qs.repo.FindAll()
}
