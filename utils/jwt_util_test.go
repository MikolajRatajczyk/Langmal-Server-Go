package utils

import (
	"errors"
	"testing"
)

func TestJwtUtil_GenerateRefreshTokenForValidId(t *testing.T) {
	sut := NewJWTUtil()
	validId := "abc"

	refreshToken, err := sut.GenerateRefreshToken(validId)

	if err != nil {
		t.Error("Should not fail")
	}

	if refreshToken == "" {
		t.Error("Token should not be empty")
	}
}

func TestJwtUtil_GenerateRefreshTokenForInvalidId(t *testing.T) {
	sut := NewJWTUtil()
	invalidId := ""

	refreshToken, err := sut.GenerateRefreshToken(invalidId)

	if !errors.Is(err, ErrAccountIdEmpty) {
		t.Error("Expected account ID empty error")
	}

	if refreshToken != "" {
		t.Error("Token should be empty for empty account ID")
	}
}

func TestJwtUtil_GenerateAccessTokenForValidId(t *testing.T) {
	sut := NewJWTUtil()
	validId := "abc"

	accessToken, err := sut.GenerateAccessToken(validId)

	if err != nil {
		t.Error("Should not fail")
	}

	if accessToken == "" {
		t.Error("Token should not be empty")
	}
}

func TestJwtUtil_GenerateAccessTokenForInvalidId(t *testing.T) {
	sut := NewJWTUtil()
	invalidId := ""

	accessToken, err := sut.GenerateAccessToken(invalidId)

	if !errors.Is(err, ErrAccountIdEmpty) {
		t.Error("Expected account ID empty error")
	}

	if accessToken != "" {
		t.Error("Token should be empty for empty account ID")
	}
}

func TestJwtUtil_GetAccountIdForValidToken(t *testing.T) {
	expectedId := "abc"
	sut := NewJWTUtil()
	validToken, err := sut.GenerateAccessToken(expectedId)
	if err != nil {
		t.Error("Can't generate token")
	}

	id, ok := sut.GetAccountId(validToken)

	if !ok {
		t.Error("Should not fail")
	}

	if id != expectedId {
		t.Error("IDs should match")
	}
}

func TestJwtUtil_GetAccountIdForInvalidToken(t *testing.T) {
	sut := NewJWTUtil()
	invalidToken := "foo"

	id, ok := sut.GetAccountId(invalidToken)

	if ok {
		t.Error("Should fail for invalid token")
	}

	if id != "" {
		t.Error("ID should be empty for invalid token")
	}
}

func TestJwtUtil_IsAccessTokenOkForValidToken(t *testing.T) {
	sut := NewJWTUtil()
	validToken, err := sut.GenerateAccessToken("abc")
	if err != nil {
		t.Error("Can't generate token")
	}

	ok := sut.IsAccessTokenOk(validToken)

	if !ok {
		t.Error("Expected true for valid token")
	}
}

func TestJwtUtil_IsAccessTokenOkForInvalidToken(t *testing.T) {
	sut := NewJWTUtil()
	invalidToken := "foo"

	ok := sut.IsAccessTokenOk(invalidToken)

	if ok {
		t.Error("Expected false for invalid token")
	}
}

func TestJwtUtil_IsRefreshTokenOkForValidToken(t *testing.T) {
	sut := NewJWTUtil()
	validToken, err := sut.GenerateRefreshToken("abc")
	if err != nil {
		t.Error("Can't generate token")
	}

	ok := sut.IsRefreshTokenOk(validToken)

	if !ok {
		t.Error("Expected true for valid token")
	}
}

func TestJwtUtil_IsRefreshTokenOkForInvalidToken(t *testing.T) {
	sut := NewJWTUtil()
	invalidToken := "foo"

	ok := sut.IsRefreshTokenOk(invalidToken)

	if ok {
		t.Error("Expected false for invalid token")
	}
}
