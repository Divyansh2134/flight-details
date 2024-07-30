package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFlightDetails(c *gin.Context) {
	cursor, err := flightsData.Find(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch flight details"})
		return
	}
	defer cursor.Close(context.TODO())

	var flightDetails []Flight
	for cursor.Next(context.TODO()) {
		var flight Flight
		if err := cursor.Decode(&flight); err != nil {
			Log.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode flight details"})
			return
		}
		flightDetails = append(flightDetails, flight)
	}

	if err := cursor.Err(); err != nil {
		Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to iterate over flight details"})
		return
	}

	c.JSON(http.StatusOK, flightDetails)
}
