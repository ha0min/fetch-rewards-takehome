package handlers

import (
	"fetch-rewards-takehome/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReceiptHandler struct {
	receipts map[string]int
}

func NewReceiptHandler() *ReceiptHandler {
	return &ReceiptHandler{
		receipts: make(map[string]int),
	}
}

func (h *ReceiptHandler) GetPoints(c *gin.Context) {
	id := c.Param("id")

	points, exists := h.receipts[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	c.JSON(http.StatusOK, models.PointsResponse{Points: points})
}
