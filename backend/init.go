package main

import (
	"flag"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	Log         *log.Logger
	postgresDB  *gorm.DB
	mongoClient *mongo.Client
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

	var err error
	mongoClient, err = connectMongo()
	if err != nil {
		Log.Error("Could not connect to MongoDB: %v", err)
	}
	// database := mongoClient.Database("flights-data")
	// passengersCollection := database.Collection("passengers")
	// flightsCollection := database.Collection("flights")
}
