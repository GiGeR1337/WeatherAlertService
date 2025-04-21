package weather

import (
	"encoding/json"
	"net/http"
)

func GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "Missing city parameter", http.StatusBadRequest)
		return
	}

	weather, err := FetchWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
