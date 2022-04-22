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

	credentialsRepository repository.CredentialsRepositoryInterface = repository.NewCredentialsRepository()
	signInService         service.SignInServiceInterface            = service.NewSingInService(credentialsRepository)
	signInController      controller.SignInControllerInterface      = controller.NewSignInController(signInService)

	signUpService    service.SignUpServiceInterface       = service.NewSignUpService(credentialsRepository)
	signUpController controller.SignUpControllerInterface = controller.NewSignUpController(signUpService)
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(),
		gin.Logger(),
		gindump.Dump(),
	)

	//	Sign-in endpoint: authentication + token creation
	server.POST("/sign-in", func(ctx *gin.Context) {
		token := signInController.SignIn(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "User not authenticated.",
			})
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
