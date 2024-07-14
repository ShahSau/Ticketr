package main

import (
	"githib.com/ShahSau/Ticketr/db"
	"githib.com/ShahSau/Ticketr/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() //initializes the database

	server := gin.Default() //creats a server with some default middleware

	routes.RegisterRoutes(server) //registers the routes

	server.Run(":8080") //run the server on port 8080

}
