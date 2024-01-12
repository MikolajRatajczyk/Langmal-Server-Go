package controllers

import (
	"net/http/httptest"
	"testing"

	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/testutils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var token = gin.H{"token": "foo"}

func TestUserController_LogoutWhenRequestOk(t *testing.T) {
	testUserController_Logout(
		&repoFake{addSuccessful: true},
		token,
		200,
		t,
	)
}

func TestUserController_LogoutWhenEmptyRequest(t *testing.T) {
	testUserController_Logout(
		&repoFake{addSuccessful: true},
		gin.H{},
		400,
		t,
	)
}

func TestUserController_LogoutWhenAlreadyLoggedOut(t *testing.T) {
	testUserController_Logout(
		&repoFake{addSuccessful: false},
		token,
		400,
		t,
	)
}

func testUserController_Logout(
	repo repositories.BlockedTokenRepoInterface,
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
		Service:          nil,
		BlockedTokenRepo: repo,
		JwtUtil:          &util{},
	}

	sut.Logout(ctx)

	if recorder.Code != expectedCode {
		t.Errorf("Expected %d status code, got %d", expectedCode, recorder.Code)
	}
}

// fakes

type repoFake struct {
	addSuccessful bool
}

func (rf *repoFake) Add(id string) bool {
	return rf.addSuccessful
}

func (*repoFake) IsBlocked(id string) bool {
	return false
}

type util struct{}

func (*util) Generate(userId string) (string, error) {
	return "foo", nil
}

func (*util) IsOk(tokenString string) bool {
	return true
}

func (*util) Claims(tokenString string) (*jwt.StandardClaims, bool) {
	return &jwt.StandardClaims{}, true
}
