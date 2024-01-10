// weather.go
package main

import "fmt"

type WeatherAPI interface {
	Fetch(city string) (string, error)
}

type WeatherService struct {
	api WeatherAPI
}

func NewWeatherService(api WeatherAPI) *WeatherService {
	return &WeatherService{api: api}
}

func (ws *WeatherService) GetWeather(city string) (string, error) {
	return ws.api.Fetch(city)
}

func main() {
	fmt.Println("Weather Service")
}
