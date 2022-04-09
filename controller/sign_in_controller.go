package controller

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entity"
	"github.com/MikolajRatajczyk/Langmal-Server/service"
	"github.com/gin-gonic/gin"
)

type SignInControllerInterface interface {
	SignIn(ctx *gin.Context) string
}

func NewSignInController(signInService service.SignInServiceInterface) SignInControllerInterface {
	return &signInController{
		signInService: signInService,
	}
}

type signInController struct {
	signInService service.SignInServiceInterface
}

func (sic *signInController) SignIn(ctx *gin.Context) string {
	var credentials entity.Credentials
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
