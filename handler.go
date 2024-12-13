package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

// GetWeather handles the HTTP request to fetch and aggregate weather data
func GetWeather(w http.ResponseWriter, r *http.Request) {
	// Validate input parameters
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	if lat == "" || lon == "" {
		http.Error(w, "Missing query parameters lat or lon", http.StatusBadRequest)
		return
	}

	// Convert lat and lon to float64
	latFloat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		http.Error(w, "Invalid latitude value", http.StatusBadRequest)
		return
	}
	lonFloat, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		http.Error(w, "Invalid longitude value", http.StatusBadRequest)
		return
	}

	// Create the result map to store aggregated weather data
	result := AggregatedWeatherData{
		OpenMeteo:  make(map[string]WeatherData),
		WeatherAPI: make(map[string]WeatherData),
	}

	// Prepare a wait group for concurrency
	var wg sync.WaitGroup

	// Fetch data concurrently
	wg.Add(2)
	go FetchWeatherFromOpenMeteo(latFloat, lonFloat, &wg, result.OpenMeteo)
	go FetchWeatherFromWeatherAPI(latFloat, lonFloat, &wg, result.WeatherAPI)

	// Wait for all goroutines to complete
	wg.Wait()

	// Return the aggregated response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}
