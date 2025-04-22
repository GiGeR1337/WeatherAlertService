package Tests

import (
	"awesomeProject/Models"
	"awesomeProject/Services"
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestEvaluateCondition(t *testing.T) {
	weather := Models.WeatherData{
		Temperature: -2.0,
		WindSpeed:   3.5,
	}

	tests := []struct {
		condition string
		expected  bool
	}{
		{"temperature < 0", true},
		{"temperature > 0", false},
		{"wind > 3", true},
		{"wind < 3", false},
		{"temperature == -2", true},
		{"temperature == 0", false},
		{"invalid", false},
		{"wind < wrong", false},
	}

	for _, tt := range tests {
		result := Services.EvaluateCondition(tt.condition, weather)
		if result != tt.expected {
			t.Errorf("EvaluateCondition(%q) = %v; want %v", tt.condition, result, tt.expected)
		}
	}
}

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	_ = db.AutoMigrate(&Models.Subscription{})
	return db
}

func TestSaveSubscription(t *testing.T) {
	db := setupTestDB()

	valid := Models.Subscription{
		Email:     "test@example.com",
		City:      "Kyiv",
		Condition: "temperature < 0",
	}

	invalidEmail := Models.Subscription{
		Email:     "invalid-email",
		City:      "Kyiv",
		Condition: "temperature < 0",
	}

	missingFields := Models.Subscription{
		Email:     "test@example.com",
		City:      "",
		Condition: "",
	}

	tests := []struct {
		name     string
		input    Models.Subscription
		expected error
	}{
		{"valid input", valid, nil},
		{"invalid email", invalidEmail, errors.New("invalid email")},
		{"missing fields", missingFields, errors.New("city and condition required")},
	}

	for _, tt := range tests {
		err := Services.SaveSubscription(db, tt.input)
		if (err != nil) != (tt.expected != nil) {
			t.Errorf("Test '%s' failed: expected error: %v, got: %v", tt.name, tt.expected != nil, err)
		}
	}
}

func TestGetAllSubscriptions(t *testing.T) {
	db := setupTestDB()

	_ = Services.SaveSubscription(db, Models.Subscription{
		Email:     "test1@example.com",
		City:      "Lviv",
		Condition: "temperature < 0",
	})
	_ = Services.SaveSubscription(db, Models.Subscription{
		Email:     "test2@example.com",
		City:      "Kyiv",
		Condition: "wind > 5",
	})

	subs, err := Services.GetAllSubscriptions(db)
	if err != nil {
		t.Fatal("GetAllSubscriptions failed:", err)
	}

	if len(subs) != 2 {
		t.Errorf("Expected 2 subscriptions, got %d", len(subs))
	}
}
