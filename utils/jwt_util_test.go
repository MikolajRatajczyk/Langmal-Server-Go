package utils

import (
	"errors"
	"testing"
)

func TestJwtUtil_GenerateForValidId(t *testing.T) {
	sut := NewJWTUtil()
	validId := "abc"

	token, err := sut.Generate(validId)

	if err != nil {
		t.Error("Should not fail")
	}

	if token == "" {
		t.Error("Token should not be empty")
	}
}

func TestJwtUtil_GenerateForInvalidId(t *testing.T) {
	sut := NewJWTUtil()
	invalidId := ""

	token, err := sut.Generate(invalidId)

	if !errors.Is(err, ErrAccountIdEmpty) {
		t.Error("Expected account ID empty error")
	}

	if token != "" {
		t.Error("Token should be empty for empty account ID")
	}
}

func TestJwtUtil_IsOkForValidToken(t *testing.T) {
	sut := NewJWTUtil()
	validToken, err := sut.Generate("abc")
	if err != nil {
		t.Error("Can't generate token")
	}

	ok := sut.IsOk(validToken)

	if !ok {
		t.Error("Expected true for valid token")
	}
}

func TestJwtUtil_IsOkForInvalidToken(t *testing.T) {
	sut := NewJWTUtil()
	invalidToken := "foo"

	ok := sut.IsOk(invalidToken)

	if ok {
		t.Error("Expected false for invalid token")
	}
}

func TestJwtUtil_ExtractAccountIdFromValidToken(t *testing.T) {
	expectedId := "abc"
	sut := NewJWTUtil()
	validToken, err := sut.Generate(expectedId)
	if err != nil {
		t.Error("Can't generate token")
	}

	id, ok := sut.ExtractAccountId(validToken)

	if !ok {
		t.Error("Should not fail")
	}

	if id != expectedId {
		t.Error("IDs should match")
	}
}

func TestJwtUtil_ExtractAccountIdFromInvalidToken(t *testing.T) {
	sut := NewJWTUtil()
	invalidToken := "foo"

	id, ok := sut.ExtractAccountId(invalidToken)

	if ok {
		t.Error("Should fail for invalid token")
	}

	if id != "" {
		t.Error("ID should be empty for invalid token")
	}
}
