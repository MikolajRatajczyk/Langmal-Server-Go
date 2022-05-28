package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
)

type SignUpControllerInterface interface {
	SignUp(ctx *gin.Context)
}

func NewSignUpController(signUpService services.SignUpServiceInterface) SignUpControllerInterface {
	return &signUpController{
		service: signUpService,
	}
}

type signUpController struct {
	service services.SignUpServiceInterface
}

func (suc *signUpController) SignUp(ctx *gin.Context) {
	var credentialsDto entities.CredentialsDto
	err := ctx.ShouldBind(&credentialsDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong credentials structure - failed to create a user.",
		})
		return
	}

	success := suc.service.SignUp(credentialsDto)
	if success == false {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create a user.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User has been created.",
	})
}
