# Weather HTTP Server with goLang
## 1. How to run
go run main.go handler.go models.go openMeteo.go weatherApi.go

## 2. How to test
curl --location 'http://localhost:8080/weather?lon=13.41&lat=52.52'

## 3. Test Result
![image](https://github.com/user-attachments/assets/3b479aa1-ffd6-4808-a428-cfd91766b986)


## 4. Explanation
### Concurrency:
Used goroutines and wiat group to fetch weather data for each day in parallel.

Each day’s forecast is fetched concurrently for both APIs (Open-Meteo and WeatherAPI).

The results are sent through a channel and collected in the main function.
### Extensibility:
WeatherProvider interface! 

The interface defines the method FetchForecast(lat, lon string) (map[string]interface{}, error) which is implemented by both the Open-Meteo and WeatherAPI services. 

This abstraction ensures that the server doesn't need to worry about the details of each API provider.
    
### Data Aggregation:
The aggregated data is collected from both APIs and returned in a single JSON response, organized by API and day.
    
### Error Handling:
Basic error handling is implemented for API calls, and if any error occurs, a message is logged, and the result for that particular API call is skipped.
    
### API Integration:
Open-Meteo does not require an API key, and the request is straightforward.

WeatherAPI requires an API key, which is passed as a query parameter.
    
### HTTP Endpoint:

The server exposes the /weather endpoint, which accepts latitude (lat), longitude (lon), and apiKey as query parameters.

