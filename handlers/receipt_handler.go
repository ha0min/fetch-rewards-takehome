package handlers

import (
	"fetch-rewards-takehome/models"
	"fetch-rewards-takehome/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReceiptHandler struct {
	receipts   map[string]int
	calculator *services.PointsCalculator
}

func NewReceiptHandler() *ReceiptHandler {
	return &ReceiptHandler{
		calculator: services.NewPointsCalculator(),
		receipts:   make(map[string]int),
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
	// create new id
	id := uuid.New().String()

	if err != nil {
		fmt.Println("[ERROR] ID: ", id, ", Error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt"})
		return
	}

	// calculate points
	points := h.calculator.CalculatePoints(&receipt)
	h.receipts[id] = points

	fmt.Println("[INFO] ID: ", id, ", Points: ", points)

	c.JSON(http.StatusOK, models.ReceiptResponse{ID: id})
}
