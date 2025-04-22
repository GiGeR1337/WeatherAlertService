package main

import (
	"awesomeProject/Controllers"
	"awesomeProject/Database"
	"awesomeProject/Scheduler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := Database.InitDB()

	Scheduler.StartDailyNotifier(db)

	r := mux.NewRouter()
	r.HandleFunc("/weather", Controllers.GetWeatherHandler).Methods("GET")
	r.HandleFunc("/subscriptions", Controllers.PostSubscriptionHandler(db)).Methods("POST")
	r.HandleFunc("/subscriptions", Controllers.GetAllSubscriptionsHandler(db)).Methods("GET")

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
