package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type WeatherResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		WindSpeed   float64 `json:"windspeed"`
	} `json:"current_weather"`
}

type CachedCity struct {
	data map[string]WeatherResponse
}

var filename string

func NewCachedCities(_filename string) *CachedCity {
	data, err := os.ReadFile(_filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var cityMap map[string]WeatherResponse
	if err := json.Unmarshal(data, &cityMap); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	filename = _filename

	return &CachedCity{data: cityMap}
}

func (c *CachedCity) GetWeatherDataByCity(name string) (WeatherResponse, error) {
	weather, exists := c.data[name]
	if !exists {
		return WeatherResponse{}, fmt.Errorf("city '%s' not found", name)
	}
	return weather, nil
}

func (c *CachedCity) AddWeather(city string, weather WeatherResponse) error {
	c.data[city] = weather

	updatedData, err := json.MarshalIndent(c.data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	if err := os.WriteFile(filename, updatedData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
