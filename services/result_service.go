package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
)

type ResultServiceInterface interface {
	Save(result models.ResultDtoSave, accountId string) bool
	Find(accountId string) []models.ResultDtoRead
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

func (rs *resultService) Save(resultDto models.ResultDtoSave, accountId string) bool {
	result := mapResultDtoToResult(resultDto, accountId)

	success := rs.resultRepo.Create(result)
	return success
}

func (rs *resultService) Find(accountId string) []models.ResultDtoRead {
	results := rs.resultRepo.Find(accountId)
	return rs.addQuizTitleToResults(results)
}

func mapResultDtoToResult(resultDto models.ResultDtoSave, accountId string) models.Result {
	result := models.Result{
		Correct:   resultDto.Correct,
		Wrong:     resultDto.Wrong,
		QuizId:    resultDto.QuizId,
		AccountId: accountId,
		CreatedAt: resultDto.CreatedAt,
	}
	return result
}

func (rs *resultService) addQuizTitleToResults(results []models.Result) []models.ResultDtoRead {
	resultDtos := []models.ResultDtoRead{}

	for _, result := range results {
		quiz, _ := rs.quizRepo.Find(result.QuizId)
		resultDto := models.ResultDtoRead{
			ResultDtoSave: models.ResultDtoSave{
				Correct:   result.Correct,
				Wrong:     result.Wrong,
				QuizId:    result.QuizId,
				CreatedAt: result.CreatedAt,
			},
			QuizTitle: quiz.Title,
		}
		resultDtos = append(resultDtos, resultDto)
	}

	return resultDtos
}
