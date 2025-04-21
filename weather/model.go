package weather

type WeatherData struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	WindSpeed   float64 `json:"wind_speed"`
	Condition   string  `json:"condition"`
}

type OpenMeteoResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		Windspeed   float64 `json:"windspeed"`
		Weathercode int     `json:"weathercode"`
	} `json:"current_weather"`
}

type GeoResponse struct {
	Results []struct {
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"results"`
}
