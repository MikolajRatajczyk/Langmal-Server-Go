package repositories

import (
	"slices"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server-Go/models"
	"github.com/google/go-cmp/cmp"
)

const resultsDbName = "results_test"

var resultEntity = models.ResultEntity{
	Correct:   1,
	Wrong:     2,
	QuizId:    "123",
	CreatedAt: 1673122069,
	UserId:    "111",
}

func TestResultRepo_Create(t *testing.T) {
	defer removeDbFile(resultsDbName, t)
	sut := NewResultRepo(resultsDbName)

	success := sut.Create(resultEntity)

	if !success {
		t.Error("Failed to create the result")
	}
}

func TestResultRepo_FindExistingResult(t *testing.T) {
	defer removeDbFile(resultsDbName, t)
	sut := NewResultRepo(resultsDbName)
	sut.Create(resultEntity)

	foundResults := sut.Find(resultEntity.UserId)

	if !slices.Equal(foundResults, []models.ResultEntity{resultEntity}) {
		t.Error("Found results are not the same as the created one")
	}
}

func TestResultRepo_FindIfEmpty(t *testing.T) {
	defer removeDbFile(resultsDbName, t)
	sut := NewResultRepo(resultsDbName)

	foundResults := sut.Find(resultEntity.UserId)

	if !cmp.Equal(foundResults, []models.ResultEntity{}) {
		t.Error("Reported success despite no results have been created")
	}
}
