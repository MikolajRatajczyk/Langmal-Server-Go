package repositories

import "testing"

// TODO Should be changed (see other repos' unit tests) when a real DB is introduced

func TestQuizRepo_FindAll(t *testing.T) {
	sut := NewQuizRepo()

	quizzes := sut.FindAll()

	if len(quizzes) == 0 {
		t.Error("No quizzes found")
	}
}
