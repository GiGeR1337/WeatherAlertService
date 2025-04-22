package Tests

import (
	"awesomeProject/Controllers"
	"awesomeProject/Models"
	"awesomeProject/Services"
	"bytes"
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	_ = db.AutoMigrate(&Models.Subscription{})
	return db
}

func TestPostSubscriptionHandler_Valid(t *testing.T) {
	db := SetupTestDB()
	handler := Controllers.PostSubscriptionHandler(db)

	sub := Models.Subscription{
		Email:     "test@example.com",
		City:      "Kyiv",
		Condition: "temperature < 0",
	}
	body, _ := json.Marshal(sub)

	req := httptest.NewRequest("POST", "/subscriptions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", rr.Code)
	}
}

func TestPostSubscriptionHandler_InvalidJSON(t *testing.T) {
	db := SetupTestDB()
	handler := Controllers.PostSubscriptionHandler(db)

	req := httptest.NewRequest("POST", "/subscriptions", bytes.NewBufferString("invalid json"))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 for invalid JSON, got %d", rr.Code)
	}
}

func TestPostSubscriptionHandler_InvalidEmail(t *testing.T) {
	db := SetupTestDB()
	handler := Controllers.PostSubscriptionHandler(db)

	sub := Models.Subscription{
		Email:     "bad-email",
		City:      "Kyiv",
		Condition: "temperature < 0",
	}
	body, _ := json.Marshal(sub)

	req := httptest.NewRequest("POST", "/subscriptions", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected 400 for invalid email, got %d", rr.Code)
	}
}

func TestGetAllSubscriptionsHandler(t *testing.T) {
	db := SetupTestDB()
	_ = Services.SaveSubscription(db, Models.Subscription{
		Email:     "user@example.com",
		City:      "Lviv",
		Condition: "wind > 3",
	})

	handler := Controllers.GetAllSubscriptionsHandler(db)
	req := httptest.NewRequest("GET", "/subscriptions", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}

	var subs []Models.Subscription
	if err := json.Unmarshal(rr.Body.Bytes(), &subs); err != nil {
		t.Fatal("Failed to parse JSON:", err)
	}

	if len(subs) != 1 || subs[0].Email != "user@example.com" {
		t.Errorf("Unexpected response body: %+v", subs)
	}
}
