package main

import (
	"fmt"
	"net/http"
	"os"
)

type WeatherData struct {
	City        string  `json:"city"`
	Temperature string  `json:"temperature"`
	Description string  `json:"description"`
	Humidity    float64 `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
	WeatherIcon string  `json:"weather_icon"`
}

func kelvinToCelsius(kelvin float64) string {
	celsius := kelvin - 273.15
	return fmt.Sprintf("%.2f", celsius)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, openWeatherMapAPIKey)

	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching weather", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusInternalServerError)
		return
	}

	// Log the JSON response for debugging
	fmt.Println("JSON Response:", string(body))

	// Check if the "main" key exists in the map
	mainData, ok := data["main"].(map[string]interface{})
	if !ok {
		http.Error(w, "Error extracting main weather information", http.StatusInternalServerError)
		return
	}

	weatherArray, ok := data["weather"].([]interface{})
	if !ok || len(weatherArray) == 0 {
		http.Error(w, "Error extracting weather information", http.StatusInternalServerError)
		return
	}

	weatherDescription, ok := weatherArray[0].(map[string]interface{})
	if !ok {
		http.Error(w, "Error extracting weather description", http.StatusInternalServerError)
		return
	}

	// Convert the temperature from Kelvin to Celsius
	temperatureCelsius := kelvinToCelsius(mainData["temp"].(float64))

	// Create a WeatherData struct with the extracted information
	weatherData := WeatherData{
		City:        city,
		Temperature: temperatureCelsius,
		Description: weatherDescription["description"].(string),
		Humidity:    mainData["humidity"].(float64),
		WindSpeed:   data["wind"].(map[string]interface{})["speed"].(float64),
		WeatherIcon: weatherDescription["icon"].(string),
	}

	// Convert the WeatherData struct to JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)
}

func main() {
	http.HandleFunc("/weather", weatherHandler)

	port := 7575
	fmt.Printf("Server is running on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
