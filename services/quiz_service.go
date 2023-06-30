package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
)

type QuizService interface {
	All() ([]models.QuizDto, bool)
}

func NewQuizService(repo repositories.QuizRepoInterface) QuizService {
	return &quizService{
		repo: repo,
	}
}

type quizService struct {
	repo repositories.QuizRepoInterface
}

func (qs *quizService) All() ([]models.QuizDto, bool) {
	quizzes := qs.repo.FindAll()

	if !(len(quizzes) > 0) {
		return []models.QuizDto{}, false
	}

	quizDto := mapQuizzesToDtos(quizzes)
	return quizDto, true
}

func mapQuizzesToDtos(quizzes []models.Quiz) []models.QuizDto {
	dtos := []models.QuizDto{}
	for _, quiz := range quizzes {
		dto := models.QuizDto{
			Name:      quiz.Name,
			Id:        quiz.Id,
			Questions: mapQuestionsToDtos(quiz.Questions),
		}
		dtos = append(dtos, dto)
	}

	return dtos
}

func mapQuestionsToDtos(questions []models.Question) []models.QuestionDto {
	dtos := []models.QuestionDto{}

	for _, question := range questions {
		dto := models.QuestionDto(question)
		dtos = append(dtos, dto)
	}

	return dtos
}
