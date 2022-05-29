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
	questions := qs.repo.FindAll()
	questionDtos := mapQuestionsToDtos(questions)
	return questionDtos
}

func mapQuestionsToDtos(questions []entities.Question) []entities.QuestionDto {
	dtos := []entities.QuestionDto{}

	for _, question := range questions {
		dto := entities.QuestionDto{
			Title:   question.Title,
			Options: question.Options,
			Answer:  question.Answer,
		}
		dtos = append(dtos, dto)
	}

	return dtos
}
