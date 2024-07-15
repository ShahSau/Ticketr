package main

import (
	"githib.com/ShahSau/Ticketr/db"
	docs "githib.com/ShahSau/Ticketr/docs"
	"githib.com/ShahSau/Ticketr/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db.InitDB() //initializes the database

	server := gin.Default() //creats a server with some default middleware

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //serves the swagger documentation

	docs.SwaggerInfo.Title = "Ticketr API" //sets the title of the swagger documentation
	routes.RegisterRoutes(server)          //registers the routes

	server.Run(":8080") //run the server on port 8080

}
