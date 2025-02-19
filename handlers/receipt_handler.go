package handlers

import (
	"fetch-rewards-takehome/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h *ReceiptHandler) PostReceipt(c *gin.Context) {
	var receipt models.Receipt
	err := c.ShouldBindBodyWithJSON(&receipt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt"})
		return
	}

	// create new id
	id := uuid.New().String()

	// TODO calculate points

	h.receipts[id] = 100

	c.JSON(http.StatusOK, models.ReceiptResponse{ID: id})
}
