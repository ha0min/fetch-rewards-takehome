package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"fetch-rewards-takehome/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPoints(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	t.Run("Receipt not found", func(t *testing.T) {
		// Setup
		handler := NewReceiptHandler()
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "id", Value: "non-existent-id"}}

		// Execute
		handler.GetPoints(c)

		// Assert
		assert.Equal(t, http.StatusNotFound, w.Code)

		var response gin.H
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Receipt not found", response["error"])
	})

	t.Run("Receipt found with preset", func(t *testing.T) {
		// Preset the points for receipt
		handler := NewReceiptHandler()
		handler.receipts["test-id"] = 100

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "id", Value: "test-id"}}

		// Execute
		handler.GetPoints(c)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response models.PointsResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 100, response.Points)
	})
}

func TestPostReceipt(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
	t.Run("Valid receipt", func(t *testing.T) {
		// Setup
		handler := NewReceiptHandler()
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		receipt := models.Receipt{
			Retailer:     "Test Retailer",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []models.Item{
				{ShortDescription: "Item 1", Price: "10.00"},
			},
			Total: "10.00",
		}
		// Create json body
		jsonData, err := json.Marshal(receipt)
		assert.NoError(t, err)

		// Create request
		c.Request, _ = http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonData))
		handler.PostReceipt(c)
		assert.Equal(t, http.StatusOK, w.Code)

		var response models.ReceiptResponse
		err = json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.ID)
	})
}
