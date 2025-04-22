Weather Alert Service
=====================

This project is a simple weather alert service built with Go. It lets users check the current weather in a city and subscribe to alerts when specific weather conditions are met (like "temperature < 0"). If the condition is true, the user gets an email notification once a day.

What It Does
------------

- Get current weather for any city (via Open-Meteo API).
- Subscribe to alerts based on custom weather conditions.
- Automatically checks conditions daily at 06:00 UTC and sends notifications.
- Stores all data (subscriptions, weather conditions) in a PostgreSQL database.

Tech Stack
----------

- Go (Golang)
- Gorilla Mux for routing
- GORM for database interaction
- GoCron for scheduling
- Gomail for sending emails
- Docker & Docker Compose for easy setup
- PostgreSQL as the database

How to Use It
-------------

1. Make sure you have Docker and Docker Compose installed.
2. Clone the repo and run: `docker-compose up --build`
3. The server will start on:  `http://localhost:8080`

Available API Endpoints
-----------------------

- `GET /weather?city=Kyiv`  
→ Returns the current weather for the city.

- `POST /subscriptions`  
→ Submits a new weather alert subscription.  
 Example JSON:
 ```json
 {
   "email": "you@example.com",
   "city": "Kyiv",
   "condition": "temperature < 0"
 }
 ```

- `GET /subscriptions`  
→ Returns all existing subscriptions.

Testing
-------

Basic unit tests are included for key functionality.  
Run them with: `go test ./Tests`

Extras
------

- Daily notifications are sent using a scheduler.
- Each subscription is checked once per day.
- The app handles validation for email, city names, and condition formats.

Made for the SKELAR Warsaw Internship assignment.

Personal Notes
------

This is my first experience with Go. I really enjoyed working on this project and learning something new.
Big thanks to SKELAR for giving me the chance to try a new technology in a creative way.





