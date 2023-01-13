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
	testRepo       repositories.TestRepoInterface = repositories.NewTestRepo()
	testService    services.TestService           = services.NewTestService(testRepo)
	testController controllers.TestController     = controllers.NewTestController(testService)

	accountRepo       repositories.AccountRepoInterface      = repositories.NewAccountRepo("accounts")
	accountService    services.AccountServiceInterface       = services.NewAccountService(accountRepo)
	accountController controllers.AccountControllerInterface = controllers.NewAccountController(accountService)

	resultRepo        repositories.ResultRepoInterface       = repositories.NewResultRepo("results")
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

	contentRoutes := server.Group("/content", middlewares.AuthorizeWithJWT())
	contentRoutes.GET("/tests", testController.GetTests)
	contentRoutes.POST("/results", resultsController.SaveResults)
	contentRoutes.GET("/results", resultsController.GetResults)

	server.Run(":5001")
}
