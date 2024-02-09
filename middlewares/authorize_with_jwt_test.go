package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func TestAuthorizeWithJWT_WhenValidToken(t *testing.T) {
	jwtUtil := jwtUtilsFake{
		generateValue: "foo",
		generateError: nil,
		isOkValue:     true,
		claimsValue:   &jwt.RegisteredClaims{},
		claimsSuccess: true,
	}
	repo := blockedTokenRepoFake{
		addValue:       true,
		isBlockedValue: false,
	}
	sut := AuthorizeWithJWT(&jwtUtil, &repo)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	request := http.Request{Header: http.Header{}}
	request.Header.Set("Authorization", "Bearer foo")
	ctx.Request = &request

	sut(ctx)

	if recorder.Code != http.StatusOK {
		t.Error("Expected OK HTTP code")
	}

	responseBody := recorder.Body.String()
	if len(responseBody) != 0 {
		t.Error("Expected empty body")
	}
}

func TestAuthorizeWithJWT_WhenBlockedToken(t *testing.T) {
	jwtUtil := jwtUtilsFake{
		generateValue: "foo",
		generateError: nil,
		isOkValue:     true,
		claimsValue:   &jwt.RegisteredClaims{},
		claimsSuccess: true,
	}
	repo := blockedTokenRepoFake{
		addValue:       true,
		isBlockedValue: true,
	}
	sut := AuthorizeWithJWT(&jwtUtil, &repo)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	request := http.Request{Header: http.Header{}}
	request.Header.Set("Authorization", "Bearer foo")
	ctx.Request = &request

	sut(ctx)

	if recorder.Code != http.StatusUnauthorized {
		t.Error("Expected unauthorized HTTP code")
	}

	responseBody := recorder.Body.String()
	if len(responseBody) == 0 || responseBody == "<nil>" {
		t.Error("Expected not empty body")
	}
}

func TestAuthorizeWithJWT_WhenNoToken(t *testing.T) {
	jwtUtil := jwtUtilsFake{
		generateValue: "foo",
		generateError: nil,
		isOkValue:     true,
		claimsValue:   &jwt.RegisteredClaims{},
		claimsSuccess: true,
	}
	repo := blockedTokenRepoFake{
		addValue:       true,
		isBlockedValue: false,
	}
	sut := AuthorizeWithJWT(&jwtUtil, &repo)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	request := http.Request{Header: http.Header{}}
	request.Header.Set("Authorization", "Bearer ")
	ctx.Request = &request

	sut(ctx)

	if recorder.Code != http.StatusUnauthorized {
		t.Error("Expected unauthorized HTTP code")
	}

	responseBody := recorder.Body.String()
	if len(responseBody) == 0 || responseBody == "<nil>" {
		t.Error("Expected not empty body")
	}
}

func TestAuthorizeWithJWT_WhenNoHeader(t *testing.T) {
	jwtUtil := jwtUtilsFake{
		generateValue: "foo",
		generateError: nil,
		isOkValue:     true,
		claimsValue:   &jwt.RegisteredClaims{},
		claimsSuccess: true,
	}
	repo := blockedTokenRepoFake{
		addValue:       true,
		isBlockedValue: false,
	}
	sut := AuthorizeWithJWT(&jwtUtil, &repo)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	request := http.Request{}
	ctx.Request = &request

	sut(ctx)

	if recorder.Code != http.StatusUnauthorized {
		t.Error("Expected unauthorized HTTP code")
	}

	responseBody := recorder.Body.String()
	if len(responseBody) == 0 || responseBody == "<nil>" {
		t.Error("Expected not empty body")
	}
}

// repo
type blockedTokenRepoFake struct {
	addValue       bool
	isBlockedValue bool
}

func (btr *blockedTokenRepoFake) Add(id string) bool {
	return btr.addValue
}

func (btr *blockedTokenRepoFake) IsBlocked(id string) bool {
	return btr.isBlockedValue
}

// JWT utils
type jwtUtilsFake struct {
	generateValue string
	generateError error
	isOkValue     bool
	claimsValue   *jwt.RegisteredClaims
	claimsSuccess bool
}

func (juf *jwtUtilsFake) Generate(userId string) (string, error) {
	return juf.generateValue, juf.generateError
}

func (juf *jwtUtilsFake) IsOk(tokenString string) bool {
	return juf.isOkValue
}

func (juf *jwtUtilsFake) Claims(tokenString string) (*jwt.RegisteredClaims, bool) {
	return juf.claimsValue, juf.claimsSuccess
}
