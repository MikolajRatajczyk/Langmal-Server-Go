package services

import (
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
)

func TestTestService_AllfRepoIsNotEmpty(t *testing.T) {
	fakeTest := newFakeTest()
	fakeTestRepo := &FakeTestRepo{
		tests: []models.Test{fakeTest},
	}
	sut := NewTestService(fakeTestRepo)

	foundTests, success := sut.All()

	if !success {
		t.Error("Reported failure despite the repo being not empty")
	}

	if len(foundTests) == 0 {
		t.Error("No found tests despite the repo being not empty")
	}
}

func TestTestService_AllIfRepoIsEmpty(t *testing.T) {
	fakeTestRepo := &FakeTestRepo{
		tests: []models.Test{},
	}
	sut := NewTestService(fakeTestRepo)

	_, success := sut.All()

	if success {
		t.Error("Reported success despite the repo being empty")
	}
}

func newFakeTest() models.Test {
	question := models.Question{
		Title:   "Foo",
		Options: []string{"a", "b", "c"},
		Answer:  "a",
	}

	test := models.Test{
		Name:      "Foo",
		Id:        "123",
		Questions: []models.Question{question},
	}

	return test
}

type FakeTestRepo struct {
	tests []models.Test
}

func (ftr *FakeTestRepo) FindAll() []models.Test {
	return ftr.tests
}
