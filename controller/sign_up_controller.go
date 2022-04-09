package controller

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entity"
	"github.com/MikolajRatajczyk/Langmal-Server/service"
	"github.com/gin-gonic/gin"
)

type SignUpControllerInterface interface {
	SignUp(ctx *gin.Context) bool
}

func NewSignUpController(signUpService service.SignUpServiceInterface) SignUpControllerInterface {
	return &signUpController{
		service: signUpService,
	}
}

type signUpController struct {
	service service.SignUpServiceInterface
}

func (suc *signUpController) SignUp(ctx *gin.Context) bool {
	var credentials entity.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		log.Println("Wrong credentials structure - can't sign-up")
		return false
	}
	success := suc.service.SignUp(credentials)
	return success
}
