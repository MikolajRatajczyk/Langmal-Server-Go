package services

import (
	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
)

type ResultServiceInterface interface {
	Save(result models.ResultWriteDto, userId string) bool
	Find(userId string) []models.ResultReadDto
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

func (rs *resultService) Save(result models.ResultWriteDto, userId string) bool {
	resultEntity := models.ResultEntity{
		Correct:   result.Correct,
		Wrong:     result.Wrong,
		QuizId:    result.QuizId,
		CreatedAt: result.CreatedAt,
		UserId:    userId,
	}

	success := rs.resultRepo.Create(resultEntity)
	return success
}

func (rs *resultService) Find(userId string) []models.ResultReadDto {
	results := rs.resultRepo.Find(userId)
	return rs.addQuizTitleToResults(results)
}

func (rs *resultService) addQuizTitleToResults(results []models.ResultEntity) []models.ResultReadDto {
	resultDtos := []models.ResultReadDto{}
	for _, result := range results {
		quiz, _ := rs.quizRepo.Find(result.QuizId)
		resultDto := models.ResultReadDto{
			ResultWriteDto: models.ResultWriteDto{
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
