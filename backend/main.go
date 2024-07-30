package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	Init()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	r.GET("/flight-details", GetFlightDetails)
	r.POST("/update-flight-details", UpdateFlightDetails)
	r.POST("/create-flight", CreateFlight)
	r.GET("/ws", HandleConnection)

	r.Run(":8000")
}
