package main

import (
	"log"
	"net/http"

	"awesomeProject/subscriptions"
	"awesomeProject/weather"
	"github.com/gorilla/mux"
)

func main() {
	db := weather.InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/weather", weather.GetWeatherHandler).Methods("GET")
	r.HandleFunc("/subscriptions", subscriptions.PostSubscriptionHandler(db)).Methods("POST")
	r.HandleFunc("/subscriptions", subscriptions.GetAllSubscriptionsHandler(db)).Methods("GET")

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
