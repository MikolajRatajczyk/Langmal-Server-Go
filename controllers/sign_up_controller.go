package controllers

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
)

type SignUpControllerInterface interface {
	SignUp(ctx *gin.Context) bool
}

func NewSignUpController(signUpService services.SignUpServiceInterface) SignUpControllerInterface {
	return &signUpController{
		service: signUpService,
	}
}

type signUpController struct {
	service services.SignUpServiceInterface
}

func (suc *signUpController) SignUp(ctx *gin.Context) bool {
	var credentials entities.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		log.Println("Wrong credentials structure - can't sign-up")
		return false
	}
	success := suc.service.SignUp(credentials)
	return success
}
