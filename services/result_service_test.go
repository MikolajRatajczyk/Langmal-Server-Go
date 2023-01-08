package services

import (
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
)

const invalidToken = "foo"

var validToken = utils.NewJWTUtil().GenerateToken("123")
var result = entities.Result{
	Correct:   1,
	Wrong:     2,
	TestId:    "123",
	CreatedAt: 1673122069,
	AccountId: "111",
}

func TestResultService_FindIfTokenIsInvalid(t *testing.T) {
	sut := NewResultService(&FakeResultRepo{})

	_, success := sut.Find(invalidToken)

	if success {
		t.Error("Should fail for invalid token")
	}
}

func TestResultService_FindIfTokenIsValid(t *testing.T) {
	sut := NewResultService(&FakeResultRepo{
		isCreateAlwaysSuccess: false,
		resultToFind:          &result,
	})

	foundResults, success := sut.Find(validToken)

	if !success {
		t.Error("Should not fail for valid token")
	}

	if len(foundResults) == 0 {
		t.Error("Found results should not be empty for valid token and not empty repo")
	}
}

func TestResultService_SaveIfTokenIsInvalid(t *testing.T) {
	sut := NewResultService(&FakeResultRepo{})

	success := sut.Save(entities.ResultDto{}, invalidToken)

	if success {
		t.Error("Should fail for invalid token")
	}
}

func TestResultService_SaveIfTokenIsValid(t *testing.T) {
	sut := NewResultService(&FakeResultRepo{
		isCreateAlwaysSuccess: true,
	})

	success := sut.Save(entities.ResultDto{}, validToken)

	if !success {
		t.Error("Should not fail for valid token")
	}
}

type FakeResultRepo struct {
	isCreateAlwaysSuccess bool
	resultToFind          *entities.Result
}

func (frr *FakeResultRepo) Create(result entities.Result) bool {
	return frr.isCreateAlwaysSuccess
}

func (frr *FakeResultRepo) Find(accountId string) ([]entities.Result, bool) {
	if frr.resultToFind != nil {
		return []entities.Result{*frr.resultToFind}, true
	} else {
		return []entities.Result{}, false
	}
}
