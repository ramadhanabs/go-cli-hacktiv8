package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type City struct {
	data map[string]Coordinate
}

func NewCities(filename string) *City {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var cityMap map[string]Coordinate
	if err := json.Unmarshal(data, &cityMap); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	return &City{data: cityMap}
}

func (c *City) GetCoordinateByCity(name string) (Coordinate, error) {
	city, exists := c.data[name]
	if !exists {
		return Coordinate{}, fmt.Errorf("city '%s' not found", name)
	}
	return city, nil
}

func (c *City) ListCities() []string {
	var cityNames []string
	for name := range c.data {
		cityNames = append(cityNames, name)
	}
	return cityNames
}
