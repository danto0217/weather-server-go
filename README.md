# Weather HTTP Server with goLang
## 1. How to run
go run main.go handler.go models.go openMeteo.go weatherApi.go

## 2. How to test
curl --location 'http://localhost:8080/weather?lon=13.41&lat=52.52&apiKey=c6740c40812a44719b215848241312'

## 3. Test Result
{
    "Open-Meteo": {
        "2024-12-13": {
            "temperature_max": 2.6
        },
        "2024-12-14": {
            "temperature_max": 2.4
        },
        "2024-12-15": {
            "temperature_max": 6.6
        },
        "2024-12-16": {
            "temperature_max": 10.6
        },
        "2024-12-17": {
            "temperature_max": 9.2
        }
    },
    "WeatherAPI": {
        "2024-12-13": {
            "temperature_max": 3.2
        },
        "2024-12-14": {
            "temperature_max": 2.8
        },
        "2024-12-15": {
            "temperature_max": 5
        },
        "2024-12-16": {
            "temperature_max": 9.9
        },
        "2024-12-17": {
            "temperature_max": 9
        }
    }
}

## 4. Explanation
### Concurrency:
    Used goroutines and wiat group to fetch weather data for each day in parallel.
    Each day’s forecast is fetched concurrently for both APIs (Open-Meteo and WeatherAPI).
    The results are sent through a channel and collected in the main function.
    
### Data Aggregation:
    The aggregated data is collected from both APIs and returned in a single JSON response, organized by API and day.
    
### Error Handling:
    Basic error handling is implemented for API calls, and if any error occurs, a message is logged, and the result for that particular API call is skipped.
    
### API Integration:
    Open-Meteo does not require an API key, and the request is straightforward.
    WeatherAPI requires an API key, which is passed as a query parameter.
    
### HTTP Endpoint:
    The server exposes the /weather endpoint, which accepts latitude (lat), longitude (lon), and apiKey as query parameters.

