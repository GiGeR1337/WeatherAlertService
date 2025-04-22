package Services

import (
	"awesomeProject/Models"
	"strconv"
	"strings"
)

func EvaluateCondition(condition string, weather Models.WeatherData) bool {
	parts := strings.Fields(condition)
	if len(parts) != 3 {
		return false
	}

	field, op, valStr := parts[0], parts[1], parts[2]
	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return false
	}

	switch field {
	case "temperature":
		switch op {
		case "<":
			return weather.Temperature < val
		case ">":
			return weather.Temperature > val
		case "==":
			return weather.Temperature == val
		}
	case "wind":
		switch op {
		case "<":
			return weather.WindSpeed < val
		case ">":
			return weather.WindSpeed > val
		case "==":
			return weather.WindSpeed == val
		}
	}

	return false
}
