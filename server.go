package main

import (
	"be_go_task/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(CORSMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "root page, server oke",
		})
	})

	V1 := router.Group("/v1")
	V1.GET("/task", controllers.GetTask)
	V1.GET("/task/:id", controllers.GetTaskById)
	V1.POST("/task", controllers.NewTask)
	V1.PATCH("/task/:id", controllers.UpdateTask)
	V1.DELETE("/task/:id", controllers.DeleteTask)

	router.Run(":5000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
