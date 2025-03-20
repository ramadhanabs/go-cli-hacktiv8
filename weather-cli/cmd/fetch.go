/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"go-cli-hacktiv8/weather-cli/pkg"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure the required flag is provided
		cityName, _ := cmd.Flags().GetString("city")
		if cityName == "" {
			fmt.Println("Error: --city is required")
			cmd.Help()
			return
		}

		retrieveWeather(cityName)
	},
}

func fetchWeather(city string) (pkg.WeatherResponse, error) {
	s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
	s.Suffix = " Fetching weather data..."
	s.Start()

	cities := pkg.NewCities("./cities.json")
	coordinates, err := cities.GetCoordinateByCity(city)
	if err != nil {
		return pkg.WeatherResponse{}, err
	}

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", coordinates.Lat, coordinates.Lon)

	resp, err := http.Get(url)
	if err != nil {
		return pkg.WeatherResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return pkg.WeatherResponse{}, err
	}

	var weather pkg.WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return pkg.WeatherResponse{}, err
	}

	s.Stop()

	return weather, nil
}

func retrieveWeather(city string) {
	currentDate := time.Now().Format("02012006")
	cacheFilePath := fmt.Sprintf("./cache_%s.json", currentDate)

	cacheExists := fileExists(cacheFilePath)

	var cache *pkg.CachedCity

	if cacheExists {
		cache = pkg.NewCachedCities(cacheFilePath)
		if weather, err := cache.GetWeatherDataByCity(city); err == nil {
			fmt.Println("✅ Weather data fetched from cache successfully!")
			printWeather(city, weather)
			return
		}
	}

	// fetch weather
	weather, err := fetchWeather(city)
	if err != nil {
		fmt.Printf("❌ Error fetching weather: %v\n", err)
		return
	}

	if cacheExists {
		if err := cache.AddWeather(city, weather); err != nil {
			fmt.Printf("❌ Error updating cache: %v\n", err)
		}
	} else {
		saveNewCacheFile(cacheFilePath, city, weather)
	}

	fmt.Println("✅ Weather data fetched from API successfully!")
	printWeather(city, weather)
}

func saveNewCacheFile(filename, city string, weather pkg.WeatherResponse) {
	weatherData := map[string]pkg.WeatherResponse{city: weather}
	data, err := json.MarshalIndent(weatherData, "", "  ")
	if err != nil {
		fmt.Printf("❌ Failed to encode JSON: %v\n", err)
		return
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		fmt.Printf("❌ Failed to write cache file: %v\n", err)
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func printWeather(city string, weather pkg.WeatherResponse) {
	fmt.Printf("🌍 City: %s\n", city)
	fmt.Printf("🌡️ Temperature: %.2f°C\n", weather.CurrentWeather.Temperature)
	fmt.Printf("💨 Wind Speed: %.2f km/h\n", weather.CurrentWeather.WindSpeed)
}

func init() {
	fetchCmd.Flags().StringP("city", "c", "", "City name (required, capitalized)")
	fetchCmd.MarkFlagRequired("city")
}
