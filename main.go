package main

import (
	"fetch-rewards-takehome/handlers"

	"github.com/gin-gonic/gin"
)

// In-memory storage for receipts and their points
var receiptPoints = make(map[string]int)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	receiptHandler := handlers.NewReceiptHandler()

	r.GET("/receipts/:id/points", receiptHandler.GetPoints)
	r.POST("/receipts/process", receiptHandler.PostReceipt)
	return r
}

func main() {
	// Create a default gin router
	r := SetupRouter()

	r.Run(":8080")
}
