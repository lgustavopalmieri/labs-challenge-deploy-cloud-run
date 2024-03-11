package weatherapi

import (
	"encoding/json"
	"net/http"
)

func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273
}

type Temperature struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
	Kelvin     float64 `json:"kelvin"`
}

func GetWeather(city string) (*Temperature, error) {
	url := "https://api.weatherapi.com/v1/current.json?q=" + city + "&key=360ddfd38d0d4cd3b72102808240403"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	result := &Temperature{
		Celsius:    data.Current.TemperatureC,
		Fahrenheit: CelsiusToFahrenheit(data.Current.TemperatureC),
		Kelvin:     CelsiusToKelvin(data.Current.TemperatureC),
	}

	return result, nil
}
