package Services

import (
	"awesomeProject/Models"
	"awesomeProject/Utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func FetchWeather(city string) (*Models.WeatherData, error) {
	lat, lon, resolvedName, err := getCoordinates(city)
	if err != nil {
		return nil, fmt.Errorf("City not found")
	}

	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%.4f&longitude=%.4f&current_weather=true",
		lat, lon,
	)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Weather API error")
	}
	defer resp.Body.Close()

	var apiResp Models.OpenMeteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("Failed to parse weather data")
	}

	return &Models.WeatherData{
		City:        resolvedName,
		Temperature: apiResp.CurrentWeather.Temperature,
		WindSpeed:   apiResp.CurrentWeather.Windspeed,
		Condition:   Utils.WeatherCodeToText(apiResp.CurrentWeather.Weathercode),
	}, nil
}

func getCoordinates(city string) (float64, float64, string, error) {
	geoURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1", strings.TrimSpace(city))

	resp, err := http.Get(geoURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return 0, 0, "", fmt.Errorf("Geocoding error")
	}
	defer resp.Body.Close()

	var geoResp Models.GeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoResp); err != nil {
		return 0, 0, "", err
	}

	if len(geoResp.Results) == 0 {
		return 0, 0, "", fmt.Errorf("City not found")
	}

	requested := strings.ToLower(city)
	returned := strings.ToLower(geoResp.Results[0].Name)

	if requested != returned {
		return 0, 0, "", fmt.Errorf("No exact match for city")
	}

	res := geoResp.Results[0]
	return res.Latitude, res.Longitude, res.Name, nil
}
