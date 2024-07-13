package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() //creats a server with some default middleware

	server.GET("/events", getEvents) //creates a route that listens to GET requests on /events

	server.Run(":8080") //run the server on port 8080

}

func getEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
