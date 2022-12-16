package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
)

type AccountControllerInterface interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

func NewAccountController(registerService services.RegisterServiceInterface,
	loginService services.LoginServiceInterface) AccountControllerInterface {
	return &accountController{
		registerService: registerService,
		loginService:    loginService,
	}
}

type accountController struct {
	registerService services.RegisterServiceInterface
	loginService    services.LoginServiceInterface
}

func (ac *accountController) Register(ctx *gin.Context) {
	var credentialsDto entities.CredentialsDto
	err := ctx.ShouldBind(&credentialsDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong credentials structure - failed to register a user.",
		})
		return
	}

	success := ac.registerService.Register(credentialsDto)
	if !success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register a user.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User has been registered.",
	})
}

func (ac *accountController) Login(ctx *gin.Context) {
	var credentialsDto entities.CredentialsDto
	err := ctx.ShouldBind(&credentialsDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong credentials structure - failed to login a user.",
		})
		return
	}

	token, err := ac.loginService.Login(credentialsDto)
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
