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
	}

	server.GET("/events/:id", getSingleEvent)
	server.POST(("/signup"), signup)
	server.POST(("/login"), login)
}
