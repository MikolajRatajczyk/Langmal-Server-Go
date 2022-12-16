package utils

import "testing"

func TestReturnsNonEmptyToken(t *testing.T) {
	sut := NewJWTUtil()
	accountName := "someAccount"

	token := sut.GenerateToken(accountName)

	if token == "" {
		t.Error("Token is empty")
	}
}

func TestValidatesValidToken(t *testing.T) {
	sut := NewJWTUtil()
	token := sut.GenerateToken("someAccount")

	_, err := sut.ValidateToken(token)

	if err != nil {
		t.Error("Failed to validate a valid token")
	}
}

func TestNotValidatesWrongToken(t *testing.T) {
	sut := NewJWTUtil()
	token := "Token"

	_, err := sut.ValidateToken(token)

	if err == nil {
		t.Error("No error despite wrong token")
	}
}
