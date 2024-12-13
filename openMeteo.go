package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// FetchWeatherFromOpenMeteo fetches the weather data for a specific day from Open-Meteo API
func FetchWeatherFromOpenMeteo(lat, lon float64, wg *sync.WaitGroup, result map[string]WeatherData) {
	defer wg.Done()

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&forecast_days=5&daily=temperature_2m_max", lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching data from Open-Meteo: %v", err)
		return
	}
	defer resp.Body.Close()

	var openMeteoResponse OpenMeteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&openMeteoResponse); err != nil {
		log.Printf("Error decoding Open-Meteo response: %v", err)
		return
	}

	for i, date := range openMeteoResponse.Daily.Time {
		result[date] = WeatherData{
			TemperatureMax: openMeteoResponse.Daily.Temperature2MMax[i],
		}
	}
}
