package Controllers

import (
	"awesomeProject/Database"
	"awesomeProject/Models"
	"encoding/json"
	"net/http"
)

func GetLogsHandler(w http.ResponseWriter, r *http.Request) {
	var logs []Models.NotificationLog
	if err := Database.DB.Find(&logs).Error; err != nil {
		http.Error(w, "Failed to fetch logs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}
