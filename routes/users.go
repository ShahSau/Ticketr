package routes

import (
	"net/http"

	"githib.com/ShahSau/Ticketr/models"
	"githib.com/ShahSau/Ticketr/utils"
	"github.com/gin-gonic/gin"
)

// @BasePath http://localhost:8080/

// @Summary register for event
// @Schemes http
// @Description Register for an event
// @Accept json
// @Produce json
// @Param email path string true "email"
// @Param password path string true "password"
// @Success 201 { string }
// @Router /signup [post]
func signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// @Summary login
// @Schemes http
// @Description Login
// @Accept json
// @Produce json
// @Param email path string true "email"
// @Param password path string true "password"
// @Success 200 { string }
// @Router /login [post]
func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	err = user.Authenticate()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}
