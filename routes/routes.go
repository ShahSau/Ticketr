package routes

import (
	"githib.com/ShahSau/Ticketr/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.AuthMiddleware)
	{
		authenticated.POST("/events", createEvent)
		authenticated.PUT("/events/:id", updateEvent)
		authenticated.DELETE("/events/:id", deleteEvent)
		authenticated.POST("/events/:id/register", registerForEvent)
		authenticated.DELETE("/events/:id/register", cancelRegistration)
	}

	server.GET("/events/:id", getSingleEvent)
	server.POST(("/signup"), signup)
	server.POST(("/login"), login)

}
