package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
)

type TestController interface {
	GetTest(ctx *gin.Context)
}

func NewTestController(service services.TestService) TestController {
	return &testController{
		service: service,
	}
}

type testController struct {
	service services.TestService
}

func (tc *testController) GetTest(ctx *gin.Context) {
	test, ok := tc.service.Find()

	if ok {
		ctx.JSON(http.StatusOK, test)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "No test found",
		})
	}
}
