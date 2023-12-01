// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrentTimeEndpoint(t *testing.T) {
	// Set up a test router
	router := gin.Default()
	router.GET("/current-time", getCurrentTime)

	// Create a mock HTTP request to the /current-time endpoint
	req, err := http.NewRequest("GET", "/current-time", nil)
	assert.NoError(t, err)

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllTimesEndpoint(t *testing.T) {
	// Set up a test router
	router := gin.Default()
	router.GET("/all-times", getAllTimes)

	// Create a mock HTTP request to the /all-times endpoint
	req, err := http.NewRequest("GET", "/all-times", nil)
	assert.NoError(t, err)

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)
}
