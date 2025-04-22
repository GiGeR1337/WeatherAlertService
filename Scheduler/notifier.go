package Scheduler

import (
	"awesomeProject/Models"
	"awesomeProject/Services"
	"awesomeProject/Utils"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

var notifiedToday = make(map[uint]bool)

func StartDailyNotifier(db *gorm.DB) {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At("06:00").Do(func() {
		runNotificationCheck(db)
	})

	s.StartAsync()
}

func runNotificationCheck(db *gorm.DB) {
	var subs []Models.Subscription
	if err := db.Find(&subs).Error; err != nil {
		log.Println("Error with fetching subscriptions:", err)
		return
	}

	for _, sub := range subs {
		if notifiedToday[sub.ID] {
			continue
		}

		weather, err := Services.FetchWeather(sub.City)
		if err != nil {
			log.Printf("Cannot fetch weather for %s: %v\n", sub.City, err)
			continue
		}

		if Services.EvaluateCondition(sub.Condition, *weather) {
			sendNotification(sub.Email, weather, sub.Condition)
			notifiedToday[sub.ID] = true
		}
	}

	go func() {
		now := time.Now()
		next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		time.Sleep(time.Until(next))
		notifiedToday = make(map[uint]bool)
	}()
}

func sendNotification(email string, w *Models.WeatherData, condition string) {
	subject := fmt.Sprintf("Weather Alert for %s", w.City)
	body := fmt.Sprintf("Your condition matched: %s\nTemperature: %.1f°C\nWind: %.1f m/s\n",
		condition, w.Temperature, w.WindSpeed)

	err := Utils.SendEmailGmail(email, subject, body)
	if err != nil {
		log.Printf("Failed to send email to %s: %v\n", email, err)
	} else {
		log.Printf("Email sent to %s for condition %s\n", email, condition)
	}
}
