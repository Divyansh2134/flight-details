package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateFlight(c *gin.Context) {
	var flight Flight

	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingFlight Flight
	filter := bson.M{"flight": flight.Flight}
	err := flightsData.FindOne(context.TODO(), filter).Decode(&existingFlight)
	if err != nil && err != mongo.ErrNoDocuments {
		Log.Error("Error checking for existing flight:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create flight"})
		return
	}
	if err == nil {
		Log.Info("Flight already exists")
		c.JSON(http.StatusConflict, gin.H{"error": "Flight already exists"})
		return
	}

	result, err := flightsData.InsertOne(context.TODO(), flight)
	if err != nil {
		Log.Error("Error inserting flight:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create flight"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Flight created successfully",
		"id":      result.InsertedID,
	})
}
