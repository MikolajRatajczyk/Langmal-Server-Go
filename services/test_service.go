package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
)

type TestService interface {
	//	TODO: Implement finding using param
	Find() ([]entities.TestDto, bool)
}

func NewTestService(repo repositories.TestRepository) TestService {
	return &testService{
		repo: repo,
	}
}

type testService struct {
	repo repositories.TestRepository
}

func (qs *testService) Find() ([]entities.TestDto, bool) {
	tests := qs.repo.FindAll()

	if !(len(tests) > 0) {
		return []entities.TestDto{}, false
	}

	testDto := mapTestsToDtos(tests)
	return testDto, true
}

func mapTestsToDtos(tests []entities.Test) []entities.TestDto {
	dtos := []entities.TestDto{}
	for _, test := range tests {
		dto := entities.TestDto{
			Name:      test.Name,
			Id:        test.Id,
			Questions: mapQuestionsToDtos(test.Questions),
		}
		dtos = append(dtos, dto)
	}

	return dtos
}

func mapQuestionsToDtos(questions []entities.Question) []entities.QuestionDto {
	dtos := []entities.QuestionDto{}

	for _, question := range questions {
		dto := entities.QuestionDto(question)
		dtos = append(dtos, dto)
	}

	return dtos
}
