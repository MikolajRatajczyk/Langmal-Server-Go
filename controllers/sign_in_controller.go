package controllers

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
)

type SignInControllerInterface interface {
	SignIn(ctx *gin.Context) string
}

func NewSignInController(signInService services.SignInServiceInterface) SignInControllerInterface {
	return &signInController{
		signInService: signInService,
	}
}

type signInController struct {
	signInService services.SignInServiceInterface
}

func (sic *signInController) SignIn(ctx *gin.Context) string {
	var credentials entities.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		log.Println("Wrong credentials structure")
		return ""
	}

	token, err := sic.signInService.SignIn(credentials)
	if err != nil {
		log.Println("User not authenticated")
		return ""
	} else {
		return token
	}
}
