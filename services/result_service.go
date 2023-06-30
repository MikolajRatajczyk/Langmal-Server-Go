package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
)

type ResultServiceInterface interface {
	Save(result models.ResultDto, accountId string) bool
	Find(accountId string) []models.ResultDto
}

func NewResultService(repo repositories.ResultRepoInterface) ResultServiceInterface {
	return &resultService{
		repo: repo,
	}
}

type resultService struct {
	repo repositories.ResultRepoInterface
}

func (rs *resultService) Save(resultDto models.ResultDto, accountId string) bool {
	result := mapResultDtoToResult(resultDto, accountId)

	success := rs.repo.Create(result)
	return success
}

func (rs *resultService) Find(accountId string) []models.ResultDto {
	results := rs.repo.Find(accountId)
	return mapResultsToDtos(results)
}

func mapResultDtoToResult(resultDto models.ResultDto, accountId string) models.Result {
	result := models.Result{
		Correct:   resultDto.Correct,
		Wrong:     resultDto.Wrong,
		QuizId:    resultDto.QuizId,
		AccountId: accountId,
		CreatedAt: resultDto.CreatedAt,
	}
	return result
}

func mapResultsToDtos(results []models.Result) []models.ResultDto {
	resultDtos := []models.ResultDto{}

	for _, result := range results {
		resultDto := models.ResultDto{
			Correct:   result.Correct,
			Wrong:     result.Wrong,
			QuizId:    result.QuizId,
			CreatedAt: result.CreatedAt,
		}
		resultDtos = append(resultDtos, resultDto)
	}

	return resultDtos
}
