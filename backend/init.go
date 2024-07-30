package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	Log             *log.Logger
	postgresDB      *gorm.DB
	env             = EnvConfig{}
	mongoClient     *mongo.Client
	flightsData     *mongo.Collection
	passengerData   *mongo.Collection
	rabbitMQConn    *amqp.Connection
	rabbitMQChannel *amqp.Channel
	emailQueue      amqp.Queue
)

func Init() {
	verbose := flag.Bool("v", false, "Enable verbose logging (debug level)")
	verboseLong := flag.Bool("verbose", false, "Enable verbose logging (debug level)")
	errorStyle := lipgloss.NewStyle().
		SetString("ERROR").
		Padding(0, 0, 0, 0).
		Bold(true).
		Foreground(lipgloss.AdaptiveColor{
			Light: "203",
			Dark:  "204",
		})
	debugStyle := lipgloss.NewStyle().
		SetString("DEBUG").
		Padding(0, 0, 0, 0).
		Bold(true).
		Foreground(lipgloss.Color("63"))
	fatalStyle := lipgloss.NewStyle().
		SetString("FATAL").
		Padding(0, 0, 0, 0).
		Bold(true).
		Foreground(lipgloss.AdaptiveColor{
			Light: "133",
			Dark:  "134",
		})
	defaultStyle := log.DefaultStyles()
	defaultStyle.Levels[log.ErrorLevel] = errorStyle
	defaultStyle.Levels[log.FatalLevel] = fatalStyle
	defaultStyle.Levels[log.DebugLevel] = debugStyle
	Log = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      "2006-01-02 15:04:05",
		Prefix:          "Logs",
	})
	Log.SetStyles(defaultStyle)
	if *verbose || *verboseLong {
		Log.SetLevel(log.DebugLevel)
	}

	godotenv.Load()
	env.MongoDbUrl = os.Getenv("MONGO_DB_URL")
	env.SmtpHost = os.Getenv("SMTP_HOST")
	env.SmtpPort, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	env.SmtpUser = os.Getenv("SMTP_USER")
	env.SmtpPass = os.Getenv("SMTP_PASS")
	env.RabbitUrl = os.Getenv("RABBIT_URL")

	var err error
	mongoClient, err = connectMongo()
	if err != nil {
		Log.Error("Could not connect to MongoDB: %v", err)
	}
	database := mongoClient.Database("Database")
	flightsData = database.Collection("flights")
	passengerData = database.Collection("passengers")

	rabbitMQConn, err = amqp.Dial(env.RabbitUrl)
	if err != nil {
		Log.Error("Failed to connect to RabbitMQ: %s", err)
	}
	rabbitMQChannel, err = rabbitMQConn.Channel()
	if err != nil {
		Log.Error("Failed to open a RabbitMQ channel: %s", err)
	}
	emailQueue, err = rabbitMQChannel.QueueDeclare(
		"email_queue", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		Log.Error("Failed to declare a RabbitMQ queue: %s", err)
	}
}
