package utils

import "testing"

func TestCryptoUtil_HashReturnNotEmptyHash(t *testing.T) {
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

func TestCryptoUtil_CompareReturnTrueWhenMatch(t *testing.T) {
	sut := NewCryptoUtil()
	password := "somePassword123"
	passwordHash, _ := sut.Hash(password)

	result := sut.Compare(password, passwordHash)

	if result == false {
		t.Fail()
	}
}

func TestCryptoUtil_CompareReturnFalseWhenNotMatch(t *testing.T) {
	sut := NewCryptoUtil()
	passwordA := "somePassword123"
	passwordB := "different"
	passwordAHash, _ := sut.Hash(passwordA)

	result := sut.Compare(passwordB, passwordAHash)

	if result == true {
		t.Fail()
	}
}
