package controllers

import (
	"errors"
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server-Go/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server-Go/services"
	"github.com/MikolajRatajczyk/Langmal-Server-Go/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service          services.UserServiceInterface
	BlockedTokenRepo repositories.BlockedTokenRepoInterface
	ClaimsExtractor  utils.ClaimsExtractorInterface
}

func (ac *UserController) Register(ctx *gin.Context) {
	var request registerRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return
	}

	err = ac.Service.Register(request.Email, request.Password)
	if err != nil {
		var httpErrStatus int
		switch {
		case errors.Is(err, services.ErrUserAlreadyExists):
			httpErrStatus = http.StatusBadRequest
		default:
			httpErrStatus = http.StatusInternalServerError
		}

		ctx.AbortWithStatusJSON(httpErrStatus, gin.H{
			"message": "Failed to register a user: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User has been registered.",
	})
}

func (ac *UserController) Login(ctx *gin.Context) {
	var request loginRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return
	}

	token, err := ac.Service.Login(request.Email, request.Password)
	if err != nil {
		var httpErrStatus int
		switch {
		case errors.Is(err, services.ErrNoUser):
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

	response := loginResponse{Token: token}
	ctx.JSON(http.StatusOK, response)
}

func (ac *UserController) Logout(ctx *gin.Context) {
	var request logoutRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return
	}

	claims, ok := ac.ClaimsExtractor.Claims(request.Token)
	if !ok {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	success := ac.BlockedTokenRepo.Add(claims.ID)
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

type loginResponse struct {
	Token string `json:"jwt"`
}
