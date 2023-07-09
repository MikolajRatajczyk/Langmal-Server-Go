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
	var credentialsDto models.CredentialsDto
	err := ctx.BindJSON(&credentialsDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong credentials structure - failed to register the account.",
		})
		return
	}

	err = ac.Service.Register(credentialsDto)
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
	var credentialsDto models.CredentialsDto
	err := ctx.BindJSON(&credentialsDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong credentials structure - failed to login the account.",
		})
		return
	}

	token, err := ac.Service.Login(credentialsDto)
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
