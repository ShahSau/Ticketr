package routes

import (
	"net/http"
	"strconv"

	"githib.com/ShahSau/Ticketr/models"
	"github.com/gin-gonic/gin"
)

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

func getSingleEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64) //gets the id from the URL

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID, could not pass eventId."}) //returns an error if the id is invalid
		return
	}

	event, err := models.GetSingleEvent(id) //calls the GetSingleEvent function from models/event.go

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"}) //returns an error if the function fails
		return
	}

	c.JSON(http.StatusOK, event) //returns the event as a JSON response
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64) //gets the id from the URL

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID, could not pass eventId."})
		return
	}

	_, err = models.GetSingleEvent(id) //calls the GetSingleEvent function from models/event.go

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	var updatedEvent models.Event //creates a new event

	err = c.ShouldBindJSON(&updatedEvent) //binds the incoming JSON to the event struct

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not pass the request data."})
		return
	}

	updatedEvent.ID = id //sets the id of the updated event

	err = updatedEvent.UpdateEvent() //calls the UpdateEvent function from models/event.go

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent}) //returns the updated event as a JSON response

}
