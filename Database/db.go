package Database

import (
	"awesomeProject/Models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "host=db user=postgres password=postgres dbname=weather port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&Models.Subscription{})
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	DB = db
	log.Println("GORM connected and migrated")
	return db
}
