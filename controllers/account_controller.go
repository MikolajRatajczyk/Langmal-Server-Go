package controllers

import (
	"errors"
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
)

type AccountControllerInterface interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	NewAccessToken(ctx *gin.Context)
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
	var accountDto models.AccountDto
	err := ctx.BindJSON(&accountDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong account structure - failed to register an account.",
		})
		return
	}

	err = ac.accountService.Register(accountDto)
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

func (ac *accountController) Login(ctx *gin.Context) {
	var loginRequestDto models.LoginRequestDto
	err := ctx.BindJSON(&loginRequestDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong login request structure - failed to login the account.",
		})
		return
	}

	tokenPair, err := ac.accountService.Login(loginRequestDto)
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

	ctx.JSON(http.StatusOK, gin.H{
		"jwtPair": tokenPair,
	})
}

func (ac *accountController) NewAccessToken(ctx *gin.Context) {
	var request models.NewAccessTokenRequestDto
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong new access token request structure - failed to deliver the new access token.",
		})
		return
	}

	accessToken, err := ac.accountService.NewAccessToken(request.RefreshJwt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create a new access token, please login again.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
