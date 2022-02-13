package main

import (
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(),
		gin.Logger(),
		gindump.Dump(),
	)

	apiRoutes := server.Group("/api")
	{
		//	TODO: create model, controller... and architecure
		//	TODO: https
		apiRoutes.GET("/questions", func(c *gin.Context) {
			c.JSON(200, createQuestions())
		})
	}

	server.Run(":5001")
}

//	TODO: remove
func createQuestions() []gin.H {
	question1 := gin.H{
		"title":   "First question from the server",
		"options": []string{"Answer A", "Answer B", "Answer C"},
		"answer":  "Answer A",
	}
	question2 := gin.H{
		"title":   "Second question from the server",
		"options": []string{"Answer A", "Answer B", "Answer C"},
		"answer":  "Answer B",
	}
	question3 := gin.H{
		"title":   "Third question from the server",
		"options": []string{"Answer A", "Answer B", "Answer C"},
		"answer":  "Answer C",
	}
	return []gin.H{question1, question2, question3}
}
