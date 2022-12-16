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
	loginService          services.LoginServiceInterface              = services.NewLoginService(credentialsRepository)
	registerService       services.RegisterServiceInterface           = services.NewRegisterService(credentialsRepository)
	accountController     controllers.AccountControllerInterface      = controllers.NewAccountController(registerService, loginService)

	resultRepo        repositories.ResultRepositoryInterface = repositories.NewResultRepository()
	resultService     services.ResultServiceInterface        = services.NewResultService(resultRepo)
	resultsController controllers.ResultsControllerInterface = controllers.NewResultsController(resultService)
)

// TODO: HTTPS
func main() {
	server := gin.New()
	server.Use(gin.Recovery(),
		gin.Logger(),
		gindump.Dump(),
	)

	accountRoutes := server.Group("/account")
	accountRoutes.POST("/register", accountController.Register)
	accountRoutes.POST("/login", accountController.Login)

	contentRoutes := server.Group("/content", middlewares.AuthorizeJWT())
	contentRoutes.GET("/tests", testController.GetTests)
	contentRoutes.POST("/results", resultsController.SaveResults)
	contentRoutes.GET("/results", resultsController.GetResults)

	server.Run(":5001")
}
