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

func NewAccountController(accountService services.AccountServiceInterface) AccountControllerInterface {
	return &accountController{
		accountService: accountService,
	}
}

type accountController struct {
	accountService services.AccountServiceInterface
}

func (ac *accountController) Register(ctx *gin.Context) {
	var accountDto entities.AccountDto
	err := ctx.ShouldBind(&accountDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong account structure - failed to register an account.",
		})
		return
	}

	success := ac.accountService.Register(accountDto)
	if !success {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register an account.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Account has been registered.",
	})
}

func (ac *accountController) Login(ctx *gin.Context) {
	var accountDto entities.AccountDto
	err := ctx.ShouldBind(&accountDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong account structure - failed to login an account.",
		})
		return
	}

	token, err := ac.accountService.Login(accountDto)
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
