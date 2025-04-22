package Controllers

import (
	weather2 "awesomeProject/Services"
	"encoding/json"
	"net/http"
)

func GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "Missing city parameter", http.StatusBadRequest)
		return
	}

	weather, err := weather2.FetchWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
