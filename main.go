package main

import (
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
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(),
		gin.Logger(),
		gindump.Dump(),
	)

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		//	TODO: https
		apiRoutes.GET("/questions", func(c *gin.Context) {
			c.JSON(200, questionController.FindAll())
		})
	}

	server.Run(":5001")
}
