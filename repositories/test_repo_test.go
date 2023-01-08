package repositories

import "testing"

// TODO Should be changed (see other repos' unit tests) when a real DB is introduced

func TestTestRepo_FindAll(t *testing.T) {
	sut := NewTestRepo()

	tests := sut.FindAll()

	if len(tests) == 0 {
		t.Error("No tests found")
	}
}
