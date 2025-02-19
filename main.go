package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Create a default gin router
	r := gin.Default()

	// Define a simple route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Run the server on port 8080
	r.Run(":8080")
}