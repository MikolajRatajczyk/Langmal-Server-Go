package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
)

type TestController interface {
	GetTests(ctx *gin.Context)
}

func NewTestController(service services.TestService) TestController {
	return &testController{
		service: service,
	}
}

type testController struct {
	service services.TestService
}

func (tc *testController) GetTests(ctx *gin.Context) {
	tests, ok := tc.service.All()

	if ok {
		ctx.JSON(http.StatusOK, tests)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "No tests found",
		})
	}
}
