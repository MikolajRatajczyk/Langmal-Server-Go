package repositories

import (
	"os"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/google/go-cmp/cmp"
)

const resultsDbName = "results_test"

var result = entities.Result{
	Correct:   1,
	Wrong:     2,
	TestId:    "123",
	CreatedAt: 1673122069,
	AccountId: "111",
}

func TestResultRepo_Create(t *testing.T) {
	defer removeResultsDbFile()
	sut := NewResultRepo(resultsDbName)

	success := sut.Create(result)

	if !success {
		t.Error("Failed to create the result")
	}
}

func TestResultRepo_FindExistingResult(t *testing.T) {
	defer removeResultsDbFile()
	sut := NewResultRepo(resultsDbName)
	sut.Create(result)

	foundResults, success := sut.Find(result.AccountId)

	if !success {
		t.Error("Reported failure despite a result has been created")
	}

	if !cmp.Equal(foundResults, []entities.Result{result}) {
		t.Error("Found results are not the same as the created one")
	}
}

func TestResultRepo_FindNonExistingResult(t *testing.T) {
	defer removeResultsDbFile()
	sut := NewResultRepo(resultsDbName)

	_, success := sut.Find(result.AccountId)

	if success {
		t.Error("Reported success despite no results have been created")
	}
}

func removeResultsDbFile() {
	os.Remove(resultsDbName + ".db")
}
