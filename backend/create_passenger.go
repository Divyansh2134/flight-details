package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePassenger(c *gin.Context) {
	var passenger Passenger

	if err := c.ShouldBindJSON(&passenger); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if passenger.Email == "" || passenger.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and ID are required"})
		return
	}

	var existingPassenger Passenger
	filter := bson.M{"id": passenger.ID}
	err := passengerData.FindOne(context.TODO(), filter).Decode(&existingPassenger)
	if err != nil && err != mongo.ErrNoDocuments {
		Log.Error("Error checking for existing passenger:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create passenger"})
		return
	}
	if err == nil {
		Log.Info("Passenger already exists")
		c.JSON(http.StatusConflict, gin.H{"error": "Passenger already exists"})
		return
	}

	result, err := passengerData.InsertOne(context.TODO(), passenger)
	if err != nil {
		Log.Error("Error inserting passenger:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create passenger"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Passenger created successfully",
		"id":      result.InsertedID,
	})
}
