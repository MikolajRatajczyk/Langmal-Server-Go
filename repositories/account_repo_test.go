package repositories

import (
	"os"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/google/go-cmp/cmp"
)

const accountsDbName = "accounts_test"

var account = models.AccountEntity{
	Id:           "123",
	Email:        "foo@foo.com",
	PasswordHash: []byte{},
}

func TestAccountRepo_Create(t *testing.T) {
	defer removeAccountsDbFile()
	sut := NewAccountRepo(accountsDbName)

	success := sut.Create(account)

	if !success {
		t.Error("Failed to create an account")
	}
}

func TestAccountRepo_FindExistingAccount(t *testing.T) {
	defer removeAccountsDbFile()
	sut := NewAccountRepo(accountsDbName)
	success := sut.Create(account)
	if !success {
		t.Error("Can't create an account and continue the test")
	}

	foundAccount, success := sut.Find(account.Email)

	if !success {
		t.Error("Reported failure despite an account has been created")
	}

	if !cmp.Equal(foundAccount, account) {
		t.Error("Found account is not the same as the created one")
	}
}

func TestAccountRepo_FindNonExistingAccount(t *testing.T) {
	defer removeAccountsDbFile()
	sut := NewAccountRepo(accountsDbName)

	_, success := sut.Find(account.Email)

	if success {
		t.Error("Reported success despite no accounts have been created")
	}
}

func removeAccountsDbFile() {
	os.Remove(accountsDbName + ".db")
}
