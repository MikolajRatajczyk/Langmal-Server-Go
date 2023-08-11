package repositories

import (
	"os"
	"testing"
)

const blockedTokensDbName = "blocked_tokens_test"

var blockedTokenId = "123"

func TestBlockedTokensRepo_Add(t *testing.T) {
	defer removeBlockedTokensDbFile()
	sut := NewBlockedTokenRepo(blockedTokensDbName)

	ok := sut.Add(blockedTokenId)

	if !ok {
		t.Error("Adding a token should succeed")
	}
}

func TestBlockedTokensRepo_IsBlockedForAdded(t *testing.T) {
	defer removeBlockedTokensDbFile()
	sut := NewBlockedTokenRepo(blockedTokensDbName)
	sut.Add(blockedTokenId)

	isBlocked := sut.IsBlocked(blockedTokenId)

	if !isBlocked {
		t.Error("Previously added token should be reported as blocked")
	}
}

func TestBlockedTokensRepo_IsBlockedForNotAdded(t *testing.T) {
	defer removeBlockedTokensDbFile()
	sut := NewBlockedTokenRepo(blockedTokensDbName)

	isBlocked := sut.IsBlocked("foo")

	if isBlocked {
		t.Error("Not added token should be reported as NOT blocked")
	}
}

func removeBlockedTokensDbFile() {
	os.Remove(blockedTokensDbName + ".db")
}