package Controllers

import (
	"awesomeProject/Models"
	"awesomeProject/Services"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

func PostSubscriptionHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sub Models.Subscription
		if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if err := Services.SaveSubscription(db, sub); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Subscription created",
		})
	}
}

func GetAllSubscriptionsHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subs, err := Services.GetAllSubscriptions(db)
		if err != nil {
			http.Error(w, "Failed to fetch subscriptions", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(subs)
	}
}
