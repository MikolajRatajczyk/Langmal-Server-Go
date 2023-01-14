package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
)

type ResultServiceInterface interface {
	Save(result models.ResultDto, token string) bool
	Find(token string) ([]models.ResultDto, bool)
}

func NewResultService(repo repositories.ResultRepoInterface) ResultServiceInterface {
	return &resultService{
		jwtUtil: utils.NewJWTUtil(),
		repo:    repo,
	}
}

type resultService struct {
	jwtUtil utils.JWTUtilInterface
	repo    repositories.ResultRepoInterface
}

func (rs *resultService) Save(resultDto models.ResultDto, token string) bool {
	accountId, err := rs.jwtUtil.GetAccountId(token)
	if err != nil {
		return false
	}

	result := mapResultDtoToResult(resultDto, accountId)

	success := rs.repo.Create(result)
	return success
}

func (rs *resultService) Find(token string) ([]models.ResultDto, bool) {
	accountId, err := rs.jwtUtil.GetAccountId(token)
	if err != nil {
		return []models.ResultDto{}, false
	}

	results := rs.repo.Find(accountId)
	return mapResultsToDtos(results), true
}

func mapResultDtoToResult(resultDto models.ResultDto, accountId string) models.Result {
	result := models.Result{
		Correct:   resultDto.Correct,
		Wrong:     resultDto.Wrong,
		TestId:    resultDto.TestId,
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
			TestId:    result.TestId,
			CreatedAt: result.CreatedAt,
		}
		resultDtos = append(resultDtos, resultDto)
	}

	return resultDtos
}
