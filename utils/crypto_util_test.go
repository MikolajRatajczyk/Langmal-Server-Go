package utils

import "testing"

func TestCryptoUtil_HashPasswordReturnNotEmptyHash(t *testing.T) {
	sut := CryptoUtil{}
	const password = "somePassword123"

	hash, err := sut.HashPassword(password)

	if err != nil {
		t.Error("Received error")
	}

	if len(hash) == 0 {
		t.Error("Hash is empty")
	}
}

func TestCryptoUtil_ComparePasswordReturnTrueWhenMatch(t *testing.T) {
	sut := CryptoUtil{}
	const password = "somePassword123"
	passwordHash, _ := sut.HashPassword(password)

	result := sut.ComparePassword(password, passwordHash)

	if result == false {
		t.Fail()
	}
}

func TestCryptoUtil_ComparePasswordReturnFalseWhenNotMatch(t *testing.T) {
	sut := CryptoUtil{}
	const passwordA = "somePassword123"
	const passwordB = "different"
	passwordAHash, _ := sut.HashPassword(passwordA)

	result := sut.ComparePassword(passwordB, passwordAHash)

	if result == true {
		t.Fail()
	}
}
