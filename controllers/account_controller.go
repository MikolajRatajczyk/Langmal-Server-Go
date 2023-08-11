package controllers

import (
	"errors"
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	Service          services.AccountServiceInterface
	BlockedTokenRepo repositories.BlockedTokensRepoInterface
	JwtUtil          utils.JwtUtil
}

func (ac *AccountController) Register(ctx *gin.Context) {
	var request registerRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return
	}

	err = ac.Service.Register(request.Email, request.Password)
	if err != nil {
		var httpErrStatus int
		switch {
		case errors.Is(err, services.ErrAccountAlreadyExists):
			httpErrStatus = http.StatusBadRequest
		default:
			httpErrStatus = http.StatusInternalServerError
		}

		ctx.AbortWithStatusJSON(httpErrStatus, gin.H{
			"message": "Failed to register an account: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Account has been registered.",
	})
}

func (ac *AccountController) Login(ctx *gin.Context) {
	var request loginRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return
	}

	token, err := ac.Service.Login(request.Email, request.Password)
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

		ctx.AbortWithStatusJSON(httpErrStatus, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (ac *AccountController) Logout(ctx *gin.Context) {
	var request logoutRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return
	}

	claims, ok := ac.JwtUtil.Claims(request.Token)
	if !ok {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	success := ac.BlockedTokenRepo.Add(claims.Id)
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

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type registerRequest loginRequest

type logoutRequest struct {
	Token string `json:"token" binding:"required"`
}
