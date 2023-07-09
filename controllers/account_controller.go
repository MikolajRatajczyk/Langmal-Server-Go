package controllers

import (
	"errors"
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	Service services.AccountServiceInterface
}

func (ac *AccountController) Register(ctx *gin.Context) {
	var accountDto models.AccountDto
	err := ctx.BindJSON(&accountDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong account structure - failed to register an account.",
		})
		return
	}

	err = ac.Service.Register(accountDto)
	if err != nil {
		var httpErrStatus int
		switch {
		case errors.Is(err, services.ErrAccountAlreadyExists):
			httpErrStatus = http.StatusBadRequest
		default:
			httpErrStatus = http.StatusInternalServerError
		}

		ctx.JSON(httpErrStatus, gin.H{
			"message": "Failed to register an account: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Account has been registered.",
	})
}

func (ac *AccountController) Login(ctx *gin.Context) {
	var loginRequestDto models.LoginRequestDto
	err := ctx.BindJSON(&loginRequestDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong login request structure - failed to login the account.",
		})
		return
	}

	token, err := ac.Service.Login(loginRequestDto)
	if err != nil {
		var httpErrStatus int
		switch {
		case errors.Is(err, services.ErrNoAccount):
			httpErrStatus = http.StatusUnauthorized
		case errors.Is(err, services.ErrNotMatchingPasswords):
			httpErrStatus = http.StatusForbidden
		default:
			httpErrStatus = http.StatusInternalServerError
		}

		ctx.JSON(httpErrStatus, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, token)
}
