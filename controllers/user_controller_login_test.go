package controllers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server-Go/services"
	"github.com/MikolajRatajczyk/Langmal-Server-Go/testutils"
	"github.com/gin-gonic/gin"
)

var credentials = gin.H{"email": "foo@foo.com", "password": "123"}

func TestUserController_LoginWhenRequestOk(t *testing.T) {
	service := userServiceFake{registerError: nil, loginError: nil}
	testUserController_Login(&service, credentials, 200, true, t)
}

func TestUserController_LoginWhenNoUser(t *testing.T) {
	service := userServiceFake{registerError: nil, loginError: services.ErrNoUser}
	testUserController_Login(&service, credentials, 401, false, t)
}

func TestUserController_LoginWhenRequestBodyEmpty(t *testing.T) {
	service := userServiceFake{registerError: nil, loginError: nil}
	testUserController_Login(&service, gin.H{}, 400, false, t)
}

func testUserController_Login(
	service services.UserServiceInterface,
	requestBody gin.H,
	expectedCode int,
	responseWithJwtFieldExpected bool,
	t *testing.T,
) {
	request := testutils.CreatePostJsonRequest(requestBody)
	if request == nil {
		t.Fatal("Can't create a request and continue the test")
	}
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = request

	sut := UserController{
		Service:          service,
		BlockedTokenRepo: nil,
		ClaimsExtractor:  nil,
	}

	sut.Login(ctx)

	if recorder.Code != expectedCode {
		t.Errorf("Expected %d status code, got %d", expectedCode, recorder.Code)
	}

	if responseWithJwtFieldExpected && !hasFilledJwtField(recorder.Body) {
		t.Errorf("Expected filled jwt field in the response body")
	}
}

func hasFilledJwtField(body *bytes.Buffer) bool {
	var parsedJson map[string]string
	err := json.Unmarshal(body.Bytes(), &parsedJson)
	if err != nil {
		return false
	}

	jwt := parsedJson["jwt"]
	return len(jwt) != 0
}

type userServiceFake struct {
	registerError error
	loginError    error
}

func (sf *userServiceFake) Register(email string, password string) error {
	return sf.registerError
}

func (sf *userServiceFake) Login(email string, password string) (token string, err error) {
	if sf.loginError != nil {
		return "", sf.loginError
	} else {
		return "foo", nil
	}
}
