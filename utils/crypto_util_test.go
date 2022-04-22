package utils

import "testing"

func TestReturnsNonEmptyHash(t *testing.T) {
	sut := NewCryptoUtil()
	password := "somePassword123"

	hash, err := sut.Hash(password)

	if err != nil {
		t.Error("Received error")
	}

	if len(hash) == 0 {
		t.Error("Hash is empty")
	}
}

func TestReturnsTrueWhenMatch(t *testing.T) {
	sut := NewCryptoUtil()
	password := "somePassword123"
	passwordHash, _ := sut.Hash(password)

	result := sut.Compare(password, passwordHash)

	if result == false {
		t.Fail()
	}
}

func TestReturnsFalseWhenNotMatch(t *testing.T) {
	sut := NewCryptoUtil()
	passwordA := "somePassword123"
	passwordB := "different"
	passwordAHash, _ := sut.Hash(passwordA)

	result := sut.Compare(passwordB, passwordAHash)

	if result == true {
		t.Fail()
	}
}
