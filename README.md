Weather Alert Service
=====================

This project is built in Go for the SKELAR Warsaw Internship.
It lets users get current weather for any city, subscribe to weather condition alerts,
and receive daily email notifications when their conditions are met.

Features
--------

- Check current weather using /weather
- Subscribe to alerts using /subscriptions
- Automatically checks weather daily at 06:00 UTC
- Sends email notifications when conditions are met
- Stores all subscriptions and notification history in the database
- Dockerized setup with PostgreSQL
- Unit tests for main logic

Tech Stack
----------

- Go (Golang)
- Gorilla Mux (Routing)
- GORM (PostgreSQL ORM)
- GoCron (Task scheduler)
- Gomail (Email sender)
- Docker & Docker Compose
- Open-Meteo API (Weather data)

API Endpoints
-------------

1. `GET /weather?city=Kyiv`
   - Returns current weather for the specified city.

2. `POST /subscriptions`
   - Creates a new alert subscription.
   - JSON example:
```
     {
       "email": "user@example.com",
       "city": "Kyiv",
       "condition": "temperature < 0"
     }
```

3. `GET /subscriptions`
   - Returns a list of all active subscriptions.

Notification Logic
------------------

- Each day at 06:00 UTC, the service checks all saved subscriptions.
- If a user's condition is met, an email is sent (only once per day).
- Every notification is saved to the database (internally).

Running the Project
-------------------

1. Make sure Docker and Docker Compose are installed.
2. Run the app with: `docker-compose up --build`
3. The server will be available at: `http://localhost:8080`

Important Setup
---------------

Before running the project, open the file `Utils/Email.go` and replace these variables with your Gmail account details:
```
   var FromEmail = "YOUR_MAIL"
   var AppPassword = "YOUR_PASSWORD"
```

Note: You'll need to generate an "App Password" from your Gmail account security settings.

Testing
-------

To run unit tests: `go test ./Tests`

Personal Notes
------

This is my first experience with Go. I really enjoyed working on this project and learning something new.
Big thanks to SKELAR for giving me the chance to try a new technology in a creative way.

