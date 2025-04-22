package Utils

func WeatherCodeToText(code int) string {
	switch code {
	case 0:
		return "Clear sky"
	case 1, 2, 3:
		return "Partly cloudy"
	case 45, 48:
		return "Fog"
	case 51, 53, 55:
		return "Drizzle"
	case 61, 63, 65:
		return "Rain"
	case 71, 73, 75:
		return "Snow"
	case 80, 81, 82:
		return "Showers"
	case 95:
		return "Thunderstorm"
	default:
		return "Unknown"
	}
}
