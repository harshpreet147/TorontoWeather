# Toronto Weather

## Docker Image

```
docker pull jasbirnetwork/torontoweather:v2
```

```
docker run -p 7575:7575 jasbirnetwork/torontoweather:v2
```

## Access weather app

Open [http://localhost:7575/weather](http://localhost:7575/weather) with your browser to see the result.


## You should get reponse like below:

```
{
    "city": "Toronto",
    "temperature": "12.90",
    "description": "overcast clouds",
    "humidity": 75,
    "wind_speed": 8.23,
    "weather_icon": "04d"
}

```


## Explaination

```
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
```

* The function weatherHandler() has the Request where we call the API `http.Get(url)` mentioned above.
* After we get the response from the GET request, we read body data using `io.ReadAll(response.Body)` method.

```
var data map[string]interface{}
err = json.Unmarshal(body, &data)
if err != nil {
	http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusInternalServerError)
	return
}
```

* Unmarshal parses the JSON-encoded data and stores the result in the `data`
* If Unmarshal return any error than it will return `http.Error` as mentioned above code snipped.

## Preparing response

```
func kelvinToCelsius(kelvin float64) string {
	celsius := kelvin - 273.15
	return fmt.Sprintf("%.2f", celsius)
}
```

```
temperatureCelsius := kelvinToCelsius(mainData["temp"].(float64))

weatherData := WeatherData{
	City:        city,
	Temperature: temperatureCelsius,
	Description: weatherDescription["description"].(string),
	Humidity:    mainData["humidity"].(float64),
	WindSpeed:   data["wind"].(map[string]interface{})["speed"].(float64),
	WeatherIcon: weatherDescription["icon"].(string),
}
```
* Convert the temperature from Kelvin to Celsius `kelvinToCelsius(mainData["temp"].(float64))`
* Create a WeatherData struct with the extracted information

```
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(weatherData)
```
* Convert the WeatherData struct to JSON and write it to the response

## Run `http` server and define api endpoint `/weather`

```
func main() {
	http.HandleFunc("/weather", weatherHandler)

	port := 7575
	fmt.Printf("Server is running on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
```
* Application will expose on port number `:7575`
* Open [http://localhost:7575/weather](http://localhost:7575/weather) with your browser to see the result.

## Test app 

* Define testcase inside `main_test.go` file.

```
req, err := http.NewRequest("GET", "/weather", nil)
if err != nil {
	t.Fatal(err)
}
```
* Create a mock HTTP request
* If `http.NewRequest` throw any error it raise fatal error.
```
rr := httptest.NewRecorder()

handler := http.HandlerFunc(weatherHandler)

handler.ServeHTTP(rr, req)
```
* Create a response recorder to record the response
* Create an HTTP handler from the weatherHandler function
* Serve the HTTP request and record the response

### Test cases #1

```
if status := rr.Code; status != http.StatusOK {
	t.Errorf("Handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
}
```
* Check the status code 200.
* If got any other status code fail the test case and return error info.

### Test cases #2

```
expectedContentType := "application/json"
if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
	t.Errorf("Handler returned unexpected content type: got %v want %v",
		contentType, expectedContentType)
}
```
* Check the content type header is `application/json` or not.
* If got content type other than `application/json` test case will fail and return error info.

### Test cases #3

```
var weatherData WeatherData
if err := json.NewDecoder(rr.Body).Decode(&weatherData); err != nil {
	t.Errorf("Error decoding JSON response: %v", err)
}
```
* Parse the JSON response.
* While decode JOSN response according to struct `weatherData` if we got error test case will fail and return error info. 

### Test cases #4

```
expectedCity := "Toronto"
if weatherData.City != expectedCity {
	t.Errorf("Unexpected city: got %v want %v", weatherData.City, expectedCity)
}
```
* Example assertion check if the city is as expected.
* If got `Unexpected city` test case will fail and return error info. 

### Run test cases

```
go test
```

* If all test will pass than you will see below response

```
PASS
ok      github.com/harshpreet147/torontoTime    0.005s
```
