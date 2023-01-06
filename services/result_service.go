package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
)

type ResultServiceInterface interface {
	Save(result entities.ResultDto, token string) bool
	Find(token string) ([]entities.ResultDto, bool)
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

func (rs *resultService) Save(resultDto entities.ResultDto, token string) bool {
	accountId, err := rs.jwtUtil.GetAccountId(token)
	if err != nil {
		return false
	}

	result := mapResultDtoToResult(resultDto, accountId)

	success := rs.repo.Create(result)
	return success
}

func (rs *resultService) Find(token string) ([]entities.ResultDto, bool) {
	accountId, err := rs.jwtUtil.GetAccountId(token)
	if err != nil {
		return []entities.ResultDto{}, false
	}

	results, success := rs.repo.Find(accountId)
	return mapResultsToDtos(results), success
}

func mapResultDtoToResult(resultDto entities.ResultDto, accountId string) entities.Result {
	result := entities.Result{
		Correct:   resultDto.Correct,
		Wrong:     resultDto.Wrong,
		TestId:    resultDto.TestId,
		AccountId: accountId,
		CreatedAt: resultDto.CreatedAt,
	}
	return result
}

func mapResultsToDtos(results []entities.Result) []entities.ResultDto {
	resultDtos := []entities.ResultDto{}

	for _, result := range results {
		resultDto := entities.ResultDto{
			Correct:   result.Correct,
			Wrong:     result.Wrong,
			TestId:    result.TestId,
			CreatedAt: result.CreatedAt,
		}
		resultDtos = append(resultDtos, resultDto)
	}

	return resultDtos
}
