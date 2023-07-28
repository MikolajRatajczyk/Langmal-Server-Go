package services

import (
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
)

var result = models.Result{
	Correct:   1,
	Wrong:     2,
	QuizId:    "123",
	CreatedAt: 1673122069,
	AccountId: "111",
}

var fakeQuizRepo = FakeQuizRepo{
	quizzes: []models.Quiz{newFakeQuiz()},
}

func TestResultService_FindIfRepoIsEmpty(t *testing.T) {
	fakeResultRepo := FakeResultRepo{
		resultToFind: nil,
	}
	sut := NewResultService(&fakeResultRepo, &fakeQuizRepo)

	results := sut.Find("123")

	if len(results) != 0 {
		t.Error("There should be no results for an empty repo")
	}
}

func TestResultService_FindIfRepoIsNotEmpty(t *testing.T) {
	fakeResultRepo := FakeResultRepo{
		resultToFind: &result,
	}
	sut := NewResultService(&fakeResultRepo, &fakeQuizRepo)

	foundResults := sut.Find("123")

	if len(foundResults) == 0 {
		t.Error("Found results should not be empty for not empty repo")
	}
}

func TestResultService_SaveIfRepoFails(t *testing.T) {
	fakeResultRepo := FakeResultRepo{
		isCreateAlwaysSuccess: false,
	}
	sut := NewResultService(&fakeResultRepo, &fakeQuizRepo)

	success := sut.Save(models.ResultDtoSave{}, "123")

	if success {
		t.Error("Should fail if repo fails")
	}
}

func TestResultService_SaveIfRepoSucceeds(t *testing.T) {
	fakeResultRepo := FakeResultRepo{
		isCreateAlwaysSuccess: true,
	}
	sut := NewResultService(&fakeResultRepo, &fakeQuizRepo)

	success := sut.Save(models.ResultDtoSave{}, "123")

	if !success {
		t.Error("Should not fail if repo succeeds")
	}
}

type FakeResultRepo struct {
	isCreateAlwaysSuccess bool
	resultToFind          *models.Result
}

func (frr *FakeResultRepo) Create(result models.Result) bool {
	return frr.isCreateAlwaysSuccess
}

func (frr *FakeResultRepo) Find(accountId string) []models.Result {
	if frr.resultToFind != nil {
		return []models.Result{*frr.resultToFind}
	} else {
		return []models.Result{}
	}
}
