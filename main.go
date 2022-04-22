package main

import (
	"net/http"

	"github.com/MikolajRatajczyk/Langmal-Server/controllers"
	"github.com/MikolajRatajczyk/Langmal-Server/middlewares"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	questionRepository repositories.QuestionRepository = repositories.NewQuestionRepository()
	questionService    services.QuestionService        = services.NewQuestionService(questionRepository)
	questionController controllers.QuestionController  = controllers.NewQuestionController(questionService)

	credentialsRepository repositories.CredentialsRepositoryInterface = repositories.NewCredentialsRepository()
	signInService         services.SignInServiceInterface             = services.NewSingInService(credentialsRepository)
	signInController      controllers.SignInControllerInterface       = controllers.NewSignInController(signInService)

	signUpService    services.SignUpServiceInterface       = services.NewSignUpService(credentialsRepository)
	signUpController controllers.SignUpControllerInterface = controllers.NewSignUpController(signUpService)
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(),
		gin.Logger(),
		gindump.Dump(),
	)

	server.POST("/sign-in", signInController.SignIn)

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
