package Services

import (
	"awesomeProject/Models"
	"errors"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[\w._%+\-]+@[\w.\-]+\.[a-zA-Z]{2,}$`)

func SaveSubscription(db *gorm.DB, sub Models.Subscription) error {
	if !emailRegex.MatchString(sub.Email) {
		return errors.New("invalid email")
	}
	if strings.TrimSpace(sub.City) == "" || strings.TrimSpace(sub.Condition) == "" {
		return errors.New("city and condition required")
	}

	result := db.Create(&sub)
	return result.Error
}

func GetAllSubscriptions(db *gorm.DB) ([]Models.Subscription, error) {
	var subs []Models.Subscription

	result := db.Find(&subs)
	if result.Error != nil {
		return nil, result.Error
	}

	return subs, nil
}
