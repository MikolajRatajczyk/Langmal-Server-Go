package main

import (
	"github.com/MikolajRatajczyk/Langmal-Server/controllers"
	"github.com/MikolajRatajczyk/Langmal-Server/middlewares"
	"github.com/MikolajRatajczyk/Langmal-Server/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server/services"
	"github.com/MikolajRatajczyk/Langmal-Server/utils"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	jwtUtil utils.JWTUtilInterface = utils.NewJWTUtil()

	quizRepo       repositories.QuizRepoInterface = repositories.NewQuizRepo()
	quizService    services.QuizService           = services.NewQuizService(quizRepo)
	quizController controllers.QuizController     = controllers.QuizController{Service: quizService}

	accountRepo      repositories.AccountRepoInterface      = repositories.NewAccountRepo("accounts")
	refreshTokenRepo repositories.RefreshTokenRepoInterface = repositories.NewRefreshTokenRepo("refresh_tokens")
	accountService   services.AccountServiceInterface       = services.NewAccountService(
		accountRepo,
		refreshTokenRepo,
		utils.NewCryptoUtil(),
		jwtUtil)
	accountController controllers.AccountController = controllers.AccountController{Service: accountService}

	resultRepo        repositories.ResultRepoInterface = repositories.NewResultRepo("results")
	resultService     services.ResultServiceInterface  = services.NewResultService(resultRepo)
	resultsController controllers.ResultsController    = controllers.ResultsController{
		ResultService: resultService,
		JwtUtil:       jwtUtil,
	}
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
	accountRoutes.POST("/new-access-token", accountController.NewAccessToken)

	contentRoutes := server.Group("/content", middlewares.AuthorizeWithJWT())
	contentRoutes.GET("/quizzes", quizController.GetQuizzes)
	contentRoutes.POST("/results", resultsController.SaveResults)
	contentRoutes.GET("/results", resultsController.GetResults)

	server.Run(":5001")
}
