package repositories

import (
	"os"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/google/go-cmp/cmp"
)

const usersDbName = "users_test"

var user = models.UserEntity{
	Id:           "123",
	Email:        "foo@foo.com",
	PasswordHash: []byte{},
}

func TestUserRepo_Create(t *testing.T) {
	defer removeDbFile(usersDbName, t)
	sut := NewUserRepo(usersDbName)

	success := sut.Create(user)

	if !success {
		t.Error("Failed to create a user")
	}
}

func TestUserRepo_FindExistingUser(t *testing.T) {
	defer removeDbFile(usersDbName, t)
	sut := NewUserRepo(usersDbName)
	success := sut.Create(user)
	if !success {
		t.Fatal("Can't create a user and continue the test")
	}

	foundUser, success := sut.Find(user.Email)

	if !success {
		t.Error("Reported failure despite a user has been created")
	}

	if !cmp.Equal(foundUser, user) {
		t.Error("Found user is not the same as the created one")
	}
}

func TestUserRepo_FindNonExistingUser(t *testing.T) {
	defer removeDbFile(usersDbName, t)
	sut := NewUserRepo(usersDbName)

	_, success := sut.Find(user.Email)

	if success {
		t.Error("Reported success despite no users have been created")
	}
}

func removeDbFile(name string, t *testing.T) {
	filename := name + ".db"
	err := os.Remove(filename)
	if err != nil {
		t.Fatal("Can't remove temporary DB file named " + filename)
	}
}
