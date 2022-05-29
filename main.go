package main

import (
	"github.com/MikolajRatajczyk/Langmal-Server/controllers"
	"github.com/MikolajRatajczyk/Langmal-Server/middlewares"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	testRepository repositories.TestRepository = repositories.NewTestRepository()
	testService    services.TestService        = services.NewTestService(testRepository)
	testController controllers.TestController  = controllers.NewTestController(testService)

	credentialsRepository repositories.CredentialsRepositoryInterface = repositories.NewCredentialsRepository()
	signInService         services.SignInServiceInterface             = services.NewSingInService(credentialsRepository)
	signInController      controllers.SignInControllerInterface       = controllers.NewSignInController(signInService)

	signUpService    services.SignUpServiceInterface       = services.NewSignUpService(credentialsRepository)
	signUpController controllers.SignUpControllerInterface = controllers.NewSignUpController(signUpService)

	resultRepo        repositories.ResultRepositoryInterface = repositories.NewResultRepository()
	resultService     services.ResultServiceInterface        = services.NewResultService(resultRepo)
	resultsController controllers.ResultsControllerInterface = controllers.NewResultsController(resultService)
)

//	TODO: HTTPS
func main() {
	server := gin.New()
	server.Use(gin.Recovery(),
		gin.Logger(),
		gindump.Dump(),
	)

	server.POST("/sign-up", signUpController.SignUp)
	server.POST("/sign-in", signInController.SignIn)

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	apiRoutes.GET("/test", testController.GetTest)
	apiRoutes.POST("/results", resultsController.SaveResults)
	apiRoutes.GET("/results", resultsController.GetResults)

	server.Run(":5001")
}
