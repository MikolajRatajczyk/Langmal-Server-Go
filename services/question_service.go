package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
)

type TestService interface {
	//	TODO: Implement finding using param (now it returns only first test)
	Find() (entities.TestDto, bool)
}

func NewTestService(repo repositories.TestRepository) TestService {
	return &testService{
		repo: repo,
	}
}

type testService struct {
	repo repositories.TestRepository
}

func (qs *testService) Find() (entities.TestDto, bool) {
	tests := qs.repo.FindAll()

	if len(tests) > 0 == false {
		return entities.TestDto{}, false
	}

	testDto := mapTestToDto(tests[0])
	return testDto, true
}

func mapTestToDto(test entities.Test) entities.TestDto {
	return entities.TestDto{
		Name:      test.Name,
		Id:        test.Id,
		Questions: mapQuestionsToDtos(test.Questions),
	}
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
