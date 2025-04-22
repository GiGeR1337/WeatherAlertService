package Models

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	Email     string `json:"email"`
	City      string `json:"city"`
	Condition string `json:"condition"`
}
