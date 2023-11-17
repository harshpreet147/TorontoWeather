package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWeatherHandler(t *testing.T) {
	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/weather", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Create an HTTP handler from the weatherHandler function
	handler := http.HandlerFunc(weatherHandler)

	// Serve the HTTP request and record the response
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the content type header
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Handler returned unexpected content type: got %v want %v",
			contentType, expectedContentType)
	}

	// Parse the JSON response
	var weatherData WeatherData
	if err := json.NewDecoder(rr.Body).Decode(&weatherData); err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	// Perform additional assertions on the weatherData struct if needed
	// ...

	// Example assertion (check if the city is as expected)
	expectedCity := "Toronto"
	if weatherData.City != expectedCity {
		t.Errorf("Unexpected city: got %v want %v", weatherData.City, expectedCity)
	}
}
