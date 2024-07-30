package main

type EnvConfig struct {
	MongoDbUrl string
	SmtpHost   string
	SmtpPort   int
	SmtpUser   string
	SmtpPass   string
	RabbitUrl  string
}
