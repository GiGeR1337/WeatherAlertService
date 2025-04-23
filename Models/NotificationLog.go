package Models

import "gorm.io/gorm"

type NotificationLog struct {
	gorm.Model
	SubscriptionID uint
	Email          string
	City           string
	Condition      string
	MatchedAt      string
}
