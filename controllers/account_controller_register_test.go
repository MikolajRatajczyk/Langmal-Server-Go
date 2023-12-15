package controllers

import (
	"net/http/httptest"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/MikolajRatajczyk/Langmal-Server/testutils"
	"github.com/gin-gonic/gin"
)

func TestAccountController_RegisterWhenRequestOk(t *testing.T) {
	service := serviceFake{registerError: nil, loginError: nil}
	testAccountController_Register(&service, credentials, 200, t)
}

func TestAccountController_RegisterWhenAccountAlreadyExists(t *testing.T) {
	service := serviceFake{registerError: services.ErrAccountAlreadyExists, loginError: nil}
	testAccountController_Register(&service, credentials, 400, t)
}

func TestAccountController_RegisterWhenRequestBodyEmpty(t *testing.T) {
	service := serviceFake{registerError: nil, loginError: nil}
	testAccountController_Register(&service, gin.H{}, 400, t)
}

func testAccountController_Register(
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

	sut.Register(ctx)

	if recorder.Code != expectedCode {
		t.Errorf("Expected %d status code, got %d", expectedCode, recorder.Code)
	}
}
