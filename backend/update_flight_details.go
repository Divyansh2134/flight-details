package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateFlightDetails(c *gin.Context) {
	var flight Flight
	if err := c.BindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingFlight Flight
	filter := bson.M{"flight": flight.Flight}
	err := flightsData.FindOne(context.TODO(), filter).Decode(&existingFlight)
	if err == mongo.ErrNoDocuments {
		Log.Error("Flight Not exist")
		c.JSON(http.StatusNotFound, gin.H{"message": "Flight not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	changes := bson.M{}
	var message string

	if existingFlight.Status != flight.Status {
		changes["status"] = flight.Status
		message += fmt.Sprintf("For flight %s, status changed to %s. ", flight.Flight, flight.Status)
	}
	if existingFlight.Gate != flight.Gate {
		changes["gate"] = flight.Gate
		message += fmt.Sprintf("For flight %s, gate changed to %s. ", flight.Flight, flight.Gate)
	}
	if existingFlight.Remarks != flight.Remarks {
		changes["remarks"] = flight.Remarks
		message += fmt.Sprintf("For flight %s, remarks updated to: %s. ", flight.Flight, flight.Remarks)
	}

	if len(changes) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No update necessary; details are unchanged"})
		return
	}

	update := bson.M{
		"$set": changes,
	}

	result, err := flightsData.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	NotifyClients(message)
	c.JSON(http.StatusOK, gin.H{"message": "Flight details updated successfully", "matchedCount": result.MatchedCount, "modifiedCount": result.ModifiedCount})
}
