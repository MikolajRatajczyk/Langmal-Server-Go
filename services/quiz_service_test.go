package services

import (
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
)

func TestQuizService_AllfRepoIsNotEmpty(t *testing.T) {
	fakeQuiz := newFakeQuiz()
	fakeQuizRepo := &FakeQuizRepo{
		quizzes: []models.Quiz{fakeQuiz},
	}
	sut := NewQuizService(fakeQuizRepo)

	foundQuizzes, success := sut.All()

	if !success {
		t.Error("Reported failure despite the repo being not empty")
	}

	if len(foundQuizzes) == 0 {
		t.Error("No found quizzes despite the repo being not empty")
	}
}

func TestQuizService_AllIfRepoIsEmpty(t *testing.T) {
	fakeQuizRepo := &FakeQuizRepo{
		quizzes: []models.Quiz{},
	}
	sut := NewQuizService(fakeQuizRepo)

	_, success := sut.All()

	if success {
		t.Error("Reported success despite the repo being empty")
	}
}

func newFakeQuiz() models.Quiz {
	question := models.Question{
		Title:   "Foo",
		Options: []string{"a", "b", "c"},
		Answer:  "a",
	}

	quiz := models.Quiz{
		Name:      "Foo",
		Id:        "123",
		Questions: []models.Question{question},
	}

	return quiz
}

type FakeQuizRepo struct {
	quizzes []models.Quiz
}

func (fqr *FakeQuizRepo) FindAll() []models.Quiz {
	return fqr.quizzes
}
