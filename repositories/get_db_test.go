package repositories

import (
	"os"
	"testing"
)

func TestGetDb_PanicIfNameIsEmpty(t *testing.T) {
	defer func() { _ = recover() }()

	getDb("", &fakeModel{})

	t.Error("Did not panic")
}

func TestGetDb_ReturnDbIfNameIsNotEmpty(t *testing.T) {
	const dbName = "notEmptyString"
	defer func() { _ = os.Remove(dbName + ".db") }()

	db := getDb(dbName, &fakeModel{})

	if db == nil {
		t.Error("DB is nil")
	}
}

type fakeModel struct{}
