package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/thegalactiks/giteway/api"
)

func TestRespondError(t *testing.T) {
	// Create a test context
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/error", func(c *gin.Context) {
		api.RespondError(c, http.StatusInternalServerError, "Internal Server Error")
	})

	// Perform a GET request to the test endpoint
	req, _ := http.NewRequest("GET", "/error", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/problem+json", w.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"error": "Internal Server Error"}`, w.Body.String())
}

func TestRespondJSON(t *testing.T) {
	// Create a test context
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/json", func(c *gin.Context) {
		api.RespondJSON(c, http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	// Perform a GET request to the test endpoint
	req, _ := http.NewRequest("GET", "/json", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"message": "Hello, World!"}`, w.Body.String())
}
