package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
)

type TestService interface {
	All() ([]models.TestDto, bool)
}

func NewTestService(repo repositories.TestRepoInterface) TestService {
	return &testService{
		repo: repo,
	}
}

type testService struct {
	repo repositories.TestRepoInterface
}

func (qs *testService) All() ([]models.TestDto, bool) {
	tests := qs.repo.FindAll()

	if !(len(tests) > 0) {
		return []models.TestDto{}, false
	}

	testDto := mapTestsToDtos(tests)
	return testDto, true
}

func mapTestsToDtos(tests []models.Test) []models.TestDto {
	dtos := []models.TestDto{}
	for _, test := range tests {
		dto := models.TestDto{
			Name:      test.Name,
			Id:        test.Id,
			Questions: mapQuestionsToDtos(test.Questions),
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
