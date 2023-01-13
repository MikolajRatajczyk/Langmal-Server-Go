package services

import (
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
)

func TestTestService_AllfRepoIsNotEmpty(t *testing.T) {
	fakeTest := newFakeTest()
	fakeTestRepo := &FakeTestRepo{
		tests: []entities.Test{fakeTest},
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
		tests: []entities.Test{},
	}
	sut := NewTestService(fakeTestRepo)

	_, success := sut.All()

	if success {
		t.Error("Reported success despite the repo being empty")
	}
}

func newFakeTest() entities.Test {
	question := entities.Question{
		Title:   "Foo",
		Options: []string{"a", "b", "c"},
		Answer:  "a",
	}

	test := entities.Test{
		Name:      "Foo",
		Id:        "123",
		Questions: []entities.Question{question},
	}

	return test
}

type FakeTestRepo struct {
	tests []entities.Test
}

func (ftr *FakeTestRepo) FindAll() []entities.Test {
	return ftr.tests
}
