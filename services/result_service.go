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
	username, err := rs.jwtUtil.GetUsername(token)
	if err != nil {
		return false
	}

	result := entities.Result{
		Correct:   resultDto.Correct,
		Total:     resultDto.Total,
		TestId:    resultDto.TestId,
		Username:  username,
		CreatedAt: resultDto.CreatedAt,
	}

	success := rs.repo.Create(result)
	return success
}

func (rs *resultService) Find(token string) []entities.ResultDto {
	username, err := rs.jwtUtil.GetUsername(token)
	if err != nil {
		return []entities.ResultDto{}
	}

	results := rs.repo.Find(username)
	return mapResultsToDtos(results)
}

func mapResultsToDtos(results []entities.Result) []entities.ResultDto {
	resultDtos := []entities.ResultDto{}

	for _, result := range results {
		resultDto := entities.ResultDto{
			Correct:   result.Correct,
			Total:     result.Total,
			TestId:    result.TestId,
			CreatedAt: result.CreatedAt,
		}
		resultDtos = append(resultDtos, resultDto)
	}

	return resultDtos
}
