package utils

import (
	"net/http"
	"testing"
)

const authHeaderKey = "Authorization"

func TestExtractToken_ForValidHeader(t *testing.T) {
	const validToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
		"eyJpYXQiOjE2NzM2MTcyODIsImlzcyI6InJhdGFqY3p5ay5kZXYiLCJzdWIiOiJhNGM0ODI3My02NzMyLTQ5MDYtODk5NS1mZWU5MTg0OGI5YmEifQ." +
		"zcZd6aOqTGIoESE7Bg3CEtlyFc53uJMSuwTRLnJoaZs"

	headerFake := http.Header{authHeaderKey: []string{"Bearer " + validToken}}

	extractedToken, err := ExtractToken(headerFake)

	if err != nil {
		t.Error("Error should be nil for a valid token")
	}

	if extractedToken != validToken {
		t.Error("Tokens don't match")
	}
}

func TestExtractToken_ForHeaderWithNoToken(t *testing.T) {
	headerFake := http.Header{authHeaderKey: []string{}}
	_, err := ExtractToken(headerFake)

	if err == nil {
		t.Error("Error should be returned for a header with no token")
	}
}

func TestExtractToken_ForNoHeaders(t *testing.T) {
	headerFake := http.Header{}
	_, err := ExtractToken(headerFake)

	if err == nil {
		t.Error("Error should be returned for an empty header")
	}
}
