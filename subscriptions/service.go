package subscriptions

import (
	"database/sql"
	"errors"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[\w._%+\-]+@[\w.\-]+\.[a-zA-Z]{2,}$`)

func SaveSubscription(db *sql.DB, sub Subscription) error {
	if !emailRegex.MatchString(sub.Email) {
		return errors.New("invalid email")
	}
	if strings.TrimSpace(sub.City) == "" || strings.TrimSpace(sub.Condition) == "" {
		return errors.New("city and condition required")
	}

	_, err := db.Exec("INSERT INTO subscriptions (email, city, condition) VALUES ($1, $2, $3)",
		sub.Email, sub.City, sub.Condition)
	return err
}

func GetAllSubscriptions(db *sql.DB) ([]Subscription, error) {
	rows, err := db.Query("SELECT email, city, condition FROM subscriptions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []Subscription

	for rows.Next() {
		var sub Subscription
		if err := rows.Scan(&sub.Email, &sub.City, &sub.Condition); err != nil {
			return nil, err
		}
		subs = append(subs, sub)
	}

	return subs, nil
}
