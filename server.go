package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPing is the handler for the GET /info/ping route.
// This will respond by a pong JSON message if the server is alive
func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Pong": "OK"})
}

// New initialize the Unleash application based on the gin http framework
func main() {

	// Create a default gin stack
	server := gin.Default()

	// Ping route
	api := server.Group("/api")
	api.GET("/ping", GetPing)

	// Listen and serve on port defined by environment variable UNLEASH_PORT
	if err := server.Run(":2015"); err != nil {
		panic("Error while starting unleash ! Cause: " + err.Error())
	}
}
