package routes

import (
	"net/http"
	"strconv"

	"githib.com/ShahSau/Ticketr/models"
	"github.com/gin-gonic/gin"
)

// @BasePath http://localhost:8080/

// @Summary register for event
// @Schemes http
// @Description Register for an event
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 { string }
// @Router /events/{id}/register [post]
func registerForEvent(c *gin.Context) {
	userID := c.GetInt64("user_id")
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID, could not pass eventId."}) //returns an error if the id is invalid
		return
	}

	event, err := models.GetSingleEvent(eventID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"}) //returns an error if the function fails
		return
	}

	err = event.Register(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for the event"}) //returns an error if the function fails
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered for the event"}) //returns a success message if the function is successful

}

// @Summary cancel registration
// @Schemes http
// @Description Cancel registration for an event
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 { string }
// @Router /events/{id}/register [delete]
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})
}
