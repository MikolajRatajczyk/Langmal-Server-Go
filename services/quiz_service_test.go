package services

import (
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
)

func TestQuizService_AllfRepoIsNotEmpty(t *testing.T) {
	fakeQuiz := newFakeQuiz()
	fakeQuizRepo := &FakeQuizRepo{
		quizzes: []models.QuizEntity{fakeQuiz},
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
		quizzes: []models.QuizEntity{},
	}
	sut := NewQuizService(fakeQuizRepo)

	_, success := sut.All()

	if success {
		t.Error("Reported success despite the repo being empty")
	}
}

func newFakeQuiz() models.QuizEntity {
	question := models.QuestionEntity{
		Title:   "Foo",
		Options: []string{"a", "b", "c"},
		Answer:  "a",
	}

	quiz := models.QuizEntity{
		Title:     "Foo",
		Id:        "123",
		Questions: []models.QuestionEntity{question},
	}

	return quiz
}

type FakeQuizRepo struct {
	quizzes []models.QuizEntity
}

func (fqr *FakeQuizRepo) FindAll() []models.QuizEntity {
	return fqr.quizzes
}

func (fqr *FakeQuizRepo) Find(id string) (models.QuizEntity, bool) {
	for _, quiz := range fqr.quizzes {
		if quiz.Id == id {
			return quiz, true
		}
	}
	return models.QuizEntity{}, false
}
