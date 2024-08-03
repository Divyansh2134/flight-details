package main

func startEmailConsumer() {
	msgs, err := rabbitMQChannel.Consume(
		emailQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		Log.Error("Failed to register a RabbitMQ consumer:", "err", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			Log.Error("Received a message:", "err", string(d.Body))
			sendEmailsToPassengers(string(d.Body))
		}
	}()

	Log.Info(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
