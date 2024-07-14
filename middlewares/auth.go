package middlewares

import (
	"net/http"

	"githib.com/ShahSau/Ticketr/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization") //gets the Authorization header from the request

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access"}) //returns an error if the token is empty
		return
	}

	userId, err := utils.ValidateToken(token) //calls the ValidateToken function from models/event.go

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access"}) //returns an error if the function fails
		return
	}

	c.Set("userId", userId) //sets the userId in the context

	c.Next() //calls the next middleware
}
