package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
)

type LoginControllerInterface interface {
	Login(ctx *gin.Context)
}

func NewLoginController(loginService services.LoginServiceInterface) LoginControllerInterface {
	return &loginController{
		loginService: loginService,
	}
}

type loginController struct {
	loginService services.LoginServiceInterface
}

func (lc *loginController) Login(ctx *gin.Context) {
	var credentialsDto entities.CredentialsDto
	err := ctx.ShouldBind(&credentialsDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong credentials structure.",
		})
		return
	}

	token, err := lc.loginService.Login(credentialsDto)
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
