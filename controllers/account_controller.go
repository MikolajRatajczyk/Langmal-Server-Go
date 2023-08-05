package controllers

import (
	"errors"
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/models"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	Service          services.AccountServiceInterface
	BlockedTokenRepo repositories.BlockedTokensRepoInterface
	JwtUtil          utils.JWTUtilInterface
}

func (ac *AccountController) Register(ctx *gin.Context) {
	var credentialsDto models.CredentialsDto
	err := ctx.BindJSON(&credentialsDto)
	if err != nil {
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

func (ac *AccountController) Logout(ctx *gin.Context) {
	tokenString, err := utils.ExtractToken(ctx.Request.Header)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	tokenId, ok := ac.JwtUtil.ExtractId(tokenString)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Token doesn't contain an ID",
		})
		return
	}

	success := ac.BlockedTokenRepo.Add(tokenId)
	if !success {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Already logged-out",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged-out (token has been blocked)",
	})
}
