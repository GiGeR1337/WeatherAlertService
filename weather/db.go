package weather

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	connStr := "host=db user=postgres password=postgres dbname=weather sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS subscriptions (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL,
		city TEXT NOT NULL,
		condition TEXT NOT NULL
	);`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Failed to run migration:", err)
	}

	log.Println("Database connected and migrated")
	return db
}
