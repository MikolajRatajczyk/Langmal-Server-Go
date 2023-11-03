package main

import (
	"github.com/MikolajRatajczyk/Langmal-Server/controllers"
	"github.com/MikolajRatajczyk/Langmal-Server/middlewares"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/gin-gonic/gin"
)

var (
	jwtUtil = utils.NewJWTUtil()

	quizRepo       = repositories.NewQuizRepo("quizzes")
	quizService    = services.NewQuizService(quizRepo)
	quizController = controllers.QuizController{Service: quizService}

	accountRepo       = repositories.NewAccountRepo("accounts")
	blockedTokensRepo = repositories.NewBlockedTokenRepo("blocked_tokens")
	accountService    = services.NewAccountService(accountRepo, utils.CryptoUtil{}, jwtUtil)
	accountController = controllers.AccountController{
		Service:          accountService,
		BlockedTokenRepo: blockedTokensRepo,
		JwtUtil:          jwtUtil,
	}

	resultRepo        = repositories.NewResultRepo("results")
	resultService     = services.NewResultService(resultRepo, quizRepo)
	resultsController = controllers.ResultsController{ResultService: resultService, JwtUtil: jwtUtil}
)

func main() {
	server := gin.Default()

	accountRoutes := server.Group("/account")
	accountRoutes.POST("/register", accountController.Register)
	accountRoutes.POST("/login", accountController.Login)
	accountRoutes.POST("/logout", accountController.Logout)

	contentRoutes := server.Group("/content", middlewares.AuthorizeWithJWT(jwtUtil, blockedTokensRepo))
	contentRoutes.GET("/quizzes", quizController.GetQuizzes)
	contentRoutes.POST("/results", resultsController.SaveResult)
	contentRoutes.GET("/results", resultsController.GetResults)

	server.Run(":5001")
}
