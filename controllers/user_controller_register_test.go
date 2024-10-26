package controllers

import (
	"net/http/httptest"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server-Go/services"
	"github.com/MikolajRatajczyk/Langmal-Server-Go/testutils"
	"github.com/gin-gonic/gin"
)

func TestUserController_RegisterWhenRequestOk(t *testing.T) {
	service := userServiceFake{registerError: nil, loginError: nil}
	testUserController_Register(&service, credentials, 200, t)
}

func TestUserController_RegisterWhenUserAlreadyExists(t *testing.T) {
	service := userServiceFake{registerError: services.ErrUserAlreadyExists, loginError: nil}
	testUserController_Register(&service, credentials, 400, t)
}

func TestUserController_RegisterWhenRequestBodyEmpty(t *testing.T) {
	service := userServiceFake{registerError: nil, loginError: nil}
	testUserController_Register(&service, gin.H{}, 400, t)
}

func testUserController_Register(
	service services.UserServiceInterface,
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

	sut := UserController{
		Service:          service,
		BlockedTokenRepo: nil,
		ClaimsExtractor:  nil,
	}

	sut.Register(ctx)

	if recorder.Code != expectedCode {
		t.Errorf("Expected %d status code, got %d", expectedCode, recorder.Code)
	}
}
