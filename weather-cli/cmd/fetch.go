/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"go-cli-hacktiv8/weather-cli/pkg"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

type WeatherResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		WindSpeed   float64 `json:"windspeed"`
	} `json:"current_weather"`
}

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

		fetchWeather(cityName)
	},
}

func fetchWeather(city string) {
	s := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
	s.Suffix = " Fetching weather data..."
	s.Start()

	cities := pkg.NewCities("./cities.json")
	coordinates, err := cities.GetCoordinateByCity(city)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", coordinates.Lat, coordinates.Lon)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	s.Stop()

	fmt.Println("✅ Weather data fetched successfully!")
	fmt.Printf("City: %s \n", city)
	fmt.Printf("Temperature: %f \n", weather.CurrentWeather.Temperature)
	fmt.Printf("Wind Speed: %f \n", weather.CurrentWeather.WindSpeed)
}

func init() {
	fetchCmd.Flags().StringP("city", "c", "", "City name (required, capitalized)")
	fetchCmd.MarkFlagRequired("city")
}
