package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/gomail.v2"
)

func sendEmailsToPassengers(body string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := passengerData.Find(ctx, bson.D{})
	if err != nil {
		Log.Error("Failed to find documents: %s", err)
	}
	defer cur.Close(ctx)

	var emails []string
	for cur.Next(ctx) {
		var result struct {
			Email string `bson:"email"`
		}
		err := cur.Decode(&result)
		if err != nil {
			Log.Error("Failed to decode document: %s", err)
		}
		emails = append(emails, result.Email)
	}
	if err := cur.Err(); err != nil {
		Log.Error("Error occurred during cursor iteration: %s", err)
	}

	if len(emails) > 0 {
		sendEmailMessage(body, emails)
	}
}

func sendEmailMessage(body string, recipients []string) {
	m := gomail.NewMessage()
	m.SetHeader("From", env.SmtpUser)
	m.SetHeader("To", recipients...)
	m.SetHeader("Subject", "Flight Notification")
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(env.SmtpHost, env.SmtpPort, env.SmtpUser, env.SmtpPass)

	if err := d.DialAndSend(m); err != nil {
		Log.Error("Failed to send email: %s", err)
	} else {
		Log.Info("Email sent!")
	}
}
