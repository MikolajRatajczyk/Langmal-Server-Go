package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
)

type SignInControllerInterface interface {
	SignIn(ctx *gin.Context)
}

func NewSignInController(signInService services.SignInServiceInterface) SignInControllerInterface {
	return &signInController{
		signInService: signInService,
	}
}

type signInController struct {
	signInService services.SignInServiceInterface
}

func (sic *signInController) SignIn(ctx *gin.Context) {
	var credentialsDto entities.CredentialsDto
	err := ctx.ShouldBind(&credentialsDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong credentials structure.",
		})
		return
	}

	token, err := sic.signInService.SignIn(credentialsDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
