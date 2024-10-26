package main

import (
	"os"

	"github.com/MikolajRatajczyk/Langmal-Server-Go/controllers"
	"github.com/MikolajRatajczyk/Langmal-Server-Go/middlewares"
	"github.com/MikolajRatajczyk/Langmal-Server-Go/repositories"
	"github.com/MikolajRatajczyk/Langmal-Server-Go/services"
	"github.com/MikolajRatajczyk/Langmal-Server-Go/utils"
	"github.com/gin-gonic/gin"
)

var (
	jwtUtil        = utils.NewJWTUtil(os.Getenv("LANGMAL_JWT_SECRET"))
	quizRepo       = repositories.NewQuizRepo("quizzes")
	quizService    = services.NewQuizService(quizRepo)
	quizController = controllers.QuizController{Service: quizService}

	userRepo          = repositories.NewUserRepo("users")
	blockedTokensRepo = repositories.NewBlockedTokenRepo("blocked_tokens")
	userService       = services.NewUserService(userRepo, utils.CryptoUtil{}, jwtUtil)
	userController    = controllers.UserController{
		Service:          userService,
		BlockedTokenRepo: blockedTokensRepo,
		ClaimsExtractor:  jwtUtil,
	}

	resultRepo        = repositories.NewResultRepo("results")
	resultService     = services.NewResultService(resultRepo, quizRepo)
	resultsController = controllers.ResultsController{ResultService: resultService, ClaimsExtractor: jwtUtil}

	healthController = controllers.HealthController{}

	contentMiddleware = middlewares.AuthorizeWithJWT(jwtUtil, blockedTokensRepo)
)

func main() {
	server := gin.Default()

	userRoutes := server.Group("/user")
	userRoutes.POST("/register", userController.Register)
	userRoutes.POST("/login", userController.Login)
	userRoutes.POST("/logout", userController.Logout)

	contentRoutes := server.Group("/content", contentMiddleware)
	contentRoutes.GET("/quizzes", quizController.GetQuizzes)
	contentRoutes.POST("/results", resultsController.SaveResult)
	contentRoutes.GET("/results", resultsController.GetResults)

	server.GET("/health", healthController.GetHealth)

	server.Run(":5001")
}
