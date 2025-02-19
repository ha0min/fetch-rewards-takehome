package handlers

import (
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
