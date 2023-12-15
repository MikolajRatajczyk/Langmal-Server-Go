package controllers

import (
	"net/http/httptest"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/MikolajRatajczyk/Langmal-Server/testutils"
	"github.com/gin-gonic/gin"
)

var credentials = gin.H{"email": "foo@foo.com", "password": "123"}

func TestAccountController_LoginWhenRequestOk(t *testing.T) {
	service := serviceFake{registerError: nil, loginError: nil}
	testAccountController_Login(&service, credentials, 200, t)
}

func TestAccountController_LoginWhenNoAccount(t *testing.T) {
	service := serviceFake{registerError: nil, loginError: services.ErrNoAccount}
	testAccountController_Login(&service, credentials, 401, t)
}

func TestAccountController_LoginWhenRequestBodyEmpty(t *testing.T) {
	service := serviceFake{registerError: nil, loginError: nil}
	testAccountController_Login(&service, gin.H{}, 400, t)
}

func testAccountController_Login(
	service services.AccountServiceInterface,
	requestBody gin.H,
	expectedCode int,
	t *testing.T,
) {
	request := testutils.CreatePostJsonRequest(requestBody)
	if request == nil {
		t.Fatal("Can't create a request and continue the test")
	}
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = request

	sut := AccountController{
		Service:          service,
		BlockedTokenRepo: nil,
		JwtUtil:          nil,
	}

	sut.Login(ctx)

	if recorder.Code != expectedCode {
		t.Errorf("Expected %d status code, got %d", expectedCode, recorder.Code)
	}
}

type serviceFake struct {
	registerError error
	loginError    error
}

func (sf *serviceFake) Register(email string, password string) error {
	return sf.registerError
}

func (sf *serviceFake) Login(email string, password string) (token string, err error) {
	if sf.loginError != nil {
		return "", sf.loginError
	} else {
		return "foo", nil
	}
}
