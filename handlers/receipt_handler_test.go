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

func TestReceiptWorkflow(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	t.Run("Post receipt and get points", func(t *testing.T) {
		// Setup
		handler := NewReceiptHandler()

		// First, post a receipt
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		receipt := models.Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []models.Item{
				{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
				{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
				{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
				{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
			},
			Total: "35.35",
		}

		jsonData, err := json.Marshal(receipt)
		assert.NoError(t, err)

		c.Request, _ = http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonData))
		handler.PostReceipt(c)
		assert.Equal(t, http.StatusOK, w.Code)

		// Get the receipt ID from the response
		var postResponse models.ReceiptResponse
		err = json.Unmarshal(w.Body.Bytes(), &postResponse)
		assert.NoError(t, err)
		assert.NotEmpty(t, postResponse.ID)

		// Then, get points for the receipt
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "id", Value: postResponse.ID}}

		handler.GetPoints(c)
		assert.Equal(t, http.StatusOK, w.Code)

		var pointsResponse models.PointsResponse
		err = json.Unmarshal(w.Body.Bytes(), &pointsResponse)
		assert.NoError(t, err)

		assert.Equal(t, 28, pointsResponse.Points)
	})
}
