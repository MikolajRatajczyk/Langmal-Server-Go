package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
)

type ResultServiceInterface interface {
	Save(result entities.ResultDto, token string) bool
	Find(token string) []entities.ResultDto
}

func NewResultService(repo repositories.ResultRepositoryInterface) ResultServiceInterface {
	return &resultService{
		jwtUtil: utils.NewJWTUtil(),
		repo:    repo,
	}
}

type resultService struct {
	jwtUtil utils.JWTUtilInterface
	repo    repositories.ResultRepositoryInterface
}

func (rs *resultService) Save(resultDto entities.ResultDto, token string) bool {
	userId, err := rs.jwtUtil.GetUserId(token)
	if err != nil {
		return false
	}

	result := mapResultDtoToResult(resultDto, userId)

	success := rs.repo.Create(result)
	return success
}

func (rs *resultService) Find(token string) []entities.ResultDto {
	userId, err := rs.jwtUtil.GetUserId(token)
	if err != nil {
		return []entities.ResultDto{}
	}

	results := rs.repo.Find(userId)
	return mapResultsToDtos(results)
}

func mapResultDtoToResult(resultDto entities.ResultDto, userId string) entities.Result {
	result := entities.Result{
		Correct:   resultDto.Correct,
		Wrong:     resultDto.Wrong,
		TestId:    resultDto.TestId,
		UserId:    userId,
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
