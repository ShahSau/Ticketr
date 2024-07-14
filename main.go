package main

import (
	"net/http"

	"githib.com/ShahSau/Ticketr/db"
	"githib.com/ShahSau/Ticketr/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() //initializes the database

	server := gin.Default() //creats a server with some default middleware

	server.GET("/events", getEvents)    //creates a route that listens to GET requests on /events
	server.POST("/events", createEvent) //creates a route that listens to POST requests on /events

	server.Run(":8080") //run the server on port 8080

}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents() //calls the GetEvents function from models/event.go

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //returns an error if the function fails
		return
	}

	c.JSON(http.StatusOK, events) //returns the events as a JSON response
}

func createEvent(c *gin.Context) {
	var event models.Event //creates a new event

	if err := c.ShouldBindJSON(&event); err != nil { //binds the incoming JSON to the event struct
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not pass the request data."}) //returns an error if the binding fails
		return
	}

	err := event.SaveEvent() //calls the SaveEvent function from models/event.go

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnot create event. Please try again later."}) //returns an error if the function fails
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	}) //returns the event as a JSON response
}
