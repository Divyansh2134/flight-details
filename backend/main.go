package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	Init()

	r := gin.Default()
	r.GET("/flight-details", GetFlightDetails)
	r.POST("/update-flight-details", UpdateFlightDetails)

	r.Run(":8000")
}
