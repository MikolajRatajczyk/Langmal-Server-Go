package controllers

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/entities"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
)

type RegisterControllerInterface interface {
	Register(ctx *gin.Context)
}

func NewRegisterController(registerService services.RegisterServiceInterface) RegisterControllerInterface {
	return &registerController{
		service: registerService,
	}
}

type registerController struct {
	service services.RegisterServiceInterface
}

func (rc *registerController) Register(ctx *gin.Context) {
	var credentialsDto entities.CredentialsDto
	err := ctx.ShouldBind(&credentialsDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong credentials structure - failed to register a user.",
		})
		return
	}

	success := rc.service.Register(credentialsDto)
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
