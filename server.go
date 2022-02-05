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
		//	TODO: add id
		//	TODO: create model, controller... and architecure
		//	TODO: https
		apiRoutes.GET("/question", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"title":   "Some question from the server",
				"options": []string{"Answer A", "Answer B", "Answer C"},
				"answer":  "Answer B",
			})
		})
	}

	server.Run(":5001")
}
