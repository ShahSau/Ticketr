package main

import (
	"net/http"

	"githib.com/ShahSau/Ticketr/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() //creats a server with some default middleware

	server.GET("/events", getEvents)    //creates a route that listens to GET requests on /events
	server.POST("/events", createEvent) //creates a route that listens to POST requests on /events

	server.Run(":8080") //run the server on port 8080

}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents() //calls the GetEvents function from models/event.go

	c.JSON(http.StatusOK, events) //returns the events as a JSON response
}

func createEvent(c *gin.Context) {
	var event models.Event //creates a new event

	if err := c.ShouldBindJSON(&event); err != nil { //binds the incoming JSON to the event struct
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //returns an error if the binding fails
		return
	}

	event.SaveEvent() //calls the SaveEvent function from models/event.go

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	}) //returns the event as a JSON response
}
