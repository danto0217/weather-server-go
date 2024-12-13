package main

// WeatherData represents the weather data for a single day
type WeatherData struct {
	TemperatureMax float64 `json:"temperature_max"`
}

// WeatherAPIResponse represents the response structure from WeatherAPI
type WeatherAPIResponse struct {
	Forecast struct {
		Forecastday []struct {
			Date string `json:"date"`
			Day  struct {
				MaxtempC float64 `json:"maxtemp_c"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

// OpenMeteoResponse represents the response structure from Open-Meteo
type OpenMeteoResponse struct {
	Daily struct {
		Time             []string  `json:"time"`
		Temperature2MMax []float64 `json:"temperature_2m_max"`
	} `json:"daily"`
}

// AggregatedWeatherData represents the aggregated weather data from both APIs
type AggregatedWeatherData struct {
	OpenMeteo  map[string]WeatherData `json:"Open-Meteo"`
	WeatherAPI map[string]WeatherData `json:"WeatherAPI"`
}
