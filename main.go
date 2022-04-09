package main

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/controller"
	"github.com/MikolajRatajczyk/Langmal-Server/middlewares"
	"github.com/MikolajRatajczyk/Langmal-Server/repository"
	"github.com/MikolajRatajczyk/Langmal-Server/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	questionRepository repository.QuestionRepository = repository.NewQuestionRepository()
	questionService    service.QuestionService       = service.NewQuestionService(questionRepository)
	questionController controller.QuestionController = controller.NewQuestionController(questionService)

	loginService    service.LoginServiceInterface       = service.NewLoginService()
	jwtService      service.JWTServiceInterface         = service.NewJWTService()
	loginController controller.LoginControllerInterface = controller.NewLoginController(loginService, jwtService)

	credentialsRepository repository.CredentialsRepositoryInterface = repository.NewCredentialsRepository()
	cryptoService         service.CryptoServiceInterface            = service.NewCryptoService()
	signUpService         service.SignUpServiceInterface            = service.NewSignUpService(credentialsRepository, cryptoService)
	signUpController      controller.SignUpControllerInterface      = controller.NewSignUpController(signUpService)
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(),
		gin.Logger(),
		gindump.Dump(),
	)

	//	Login endpoint: authentication + token creation
	//	TODO: update (sign-up-->sign-in-->JWT flow)
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	server.POST("/sign-up", func(ctx *gin.Context) {
		success := signUpController.SignUp(ctx)
		if success {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "User has been created.",
			})
		} else {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "Failed to create a user.",
			})
		}
	})

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		//	TODO: https
		apiRoutes.GET("/questions", func(c *gin.Context) {
			c.JSON(200, questionController.FindAll())
		})
	}

	server.Run(":5001")
}
