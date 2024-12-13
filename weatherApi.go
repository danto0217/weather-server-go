package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// FetchWeatherFromWeatherAPI fetches the weather data for a specific day from WeatherAPI
func FetchWeatherFromWeatherAPI(lat, lon float64, apiKey string, wg *sync.WaitGroup, result map[string]WeatherData) {
	defer wg.Done()

	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%f,%f&days=5", apiKey, lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching data from WeatherAPI: %v", err)
		return
	}
	defer resp.Body.Close()

	var weatherAPIResponse WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherAPIResponse); err != nil {
		log.Printf("Error decoding WeatherAPI response: %v", err)
		return
	}

	for _, day := range weatherAPIResponse.Forecast.Forecastday {
		result[day.Date] = WeatherData{TemperatureMax: day.Day.MaxtempC}
	}

}
