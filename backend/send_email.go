package main

import "github.com/streadway/amqp"

func sendEmail(message string) {
	err := rabbitMQChannel.Publish(
		"",              // exchange
		emailQueue.Name, // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		Log.Error("Failed to publish a message: %s", err)
	}
}
