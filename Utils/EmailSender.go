package Utils

import (
	"gopkg.in/gomail.v2"
	"log"
)

var FromEmail = "YOUR_MAIL"
var AppPassword = "YOUR_PASSWORD" //Gmail App Password

func SendEmailGmail(to string, subject string, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", FromEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, FromEmail, AppPassword)

	err := d.DialAndSend(m)
	if err != nil {
		log.Println("Email send error:", err)
	}
	return err
}
