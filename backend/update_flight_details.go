package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Flight struct {
	Flight  string `json:"flight"`
	Status  string `json:"status"`
	Gate    string `json:"gate"`
	Remarks string `json:"remarks"`
}

func UpdateFlightDetails(c *gin.Context) {
	var flight Flight
	if err := c.BindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := mongoClient.Database("yourdb").Collection("flights")
	_, err := collection.InsertOne(context.TODO(), flight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flight created successfully"})
}
