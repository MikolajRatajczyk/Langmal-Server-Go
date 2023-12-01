package repositories

import (
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/google/go-cmp/cmp"
)

const quizzesDbName = "quizzes_test"

var (
	quizId = "123"

	quiz = models.QuizEntity{
		Title:     "Quiz title",
		Id:        quizId,
		Questions: []models.QuestionEntity{question},
	}

	question = models.QuestionEntity{
		Title:        "Question title",
		Id:           "456",
		Options:      []string{"option1", "option2", "option3"},
		Answer:       0,
		QuizEntityId: quizId,
	}
)

func TestQuizRepo_Create(t *testing.T) {
	defer removeDbFile(quizzesDbName, t)
	sut := NewQuizRepo(quizzesDbName)

	success := sut.Create(quiz)

	if !success {
		t.Error("Failed to create a quiz")
	}
}

func TestQuizRepo_FindDefaultQuizzes(t *testing.T) {
	defer removeDbFile(quizzesDbName, t)
	sut := NewQuizRepo(quizzesDbName)

	all := sut.FindAll()

	if len(all) != 2 {
		t.Error("Two default quizzes expected")
	}
}

func TestQuizRepo_FindExistingQuiz(t *testing.T) {
	defer removeDbFile(quizzesDbName, t)
	sut := NewQuizRepo(quizzesDbName)
	sut.Create(quiz)

	foundQuiz, success := sut.Find(quiz.Id)

	if !success {
		t.Error("Reported failure even though the quiz was created")
	}

	if !cmp.Equal(foundQuiz, quiz) {
		t.Error("Found quiz is not the same as the created one")
	}
}

func TestQuizRepo_FindNonExistingQuiz(t *testing.T) {
	defer removeDbFile(quizzesDbName, t)
	sut := NewQuizRepo(quizzesDbName)

	_, success := sut.Find(quiz.Id)

	if success {
		t.Error("Reported success even though no quizzes were created")
	}
}
