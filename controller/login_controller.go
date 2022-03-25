package controller

import (
	"log"

	"github.com/MikolajRatajczyk/Langmal-Server/entity"
	"github.com/MikolajRatajczyk/Langmal-Server/service"
	"github.com/gin-gonic/gin"
)

type LoginControllerInterface interface {
	Login(ctx *gin.Context) string
}

func NewLoginController(loginService service.LoginServiceInterface,
	jwtService service.JWTServiceInterface) LoginControllerInterface {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

type loginController struct {
	loginService service.LoginServiceInterface
	jwtService   service.JWTServiceInterface
}

func (lc *loginController) Login(ctx *gin.Context) string {
	var credentials entity.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		log.Println("Wrong credentials structure")
		return ""
	}

	username := credentials.Username
	password := credentials.Password
	isAuthenticated := lc.loginService.Login(username, password)
	if isAuthenticated {
		return lc.jwtService.GenerateToken(username)
	} else {
		log.Println("User not authenticated")
		return ""
	}
}
