package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
)

type ResultServiceInterface interface {
	Save(result models.ResultEntity, userId string) bool
	Find(userId string) []models.ResultDto
}

func NewResultService(resultRepo repositories.ResultRepoInterface, quizRepo repositories.QuizRepoInterface) ResultServiceInterface {
	return &resultService{
		resultRepo: resultRepo,
		quizRepo:   quizRepo,
	}
}

type resultService struct {
	resultRepo repositories.ResultRepoInterface
	quizRepo   repositories.QuizRepoInterface
}

func (rs *resultService) Save(result models.ResultEntity, userId string) bool {
	success := rs.resultRepo.Create(result)
	return success
}

func (rs *resultService) Find(userId string) []models.ResultDto {
	results := rs.resultRepo.Find(userId)
	return rs.addQuizTitleToResults(results)
}

func (rs *resultService) addQuizTitleToResults(results []models.ResultEntity) []models.ResultDto {
	resultDtos := []models.ResultDto{}

	for _, result := range results {
		quiz, _ := rs.quizRepo.Find(result.QuizId)
		resultDto := models.ResultDto{
			Correct:   result.Correct,
			Wrong:     result.Wrong,
			QuizId:    result.QuizId,
			CreatedAt: result.CreatedAt,
			QuizTitle: quiz.Title,
		}
		resultDtos = append(resultDtos, resultDto)
	}

	return resultDtos
}
