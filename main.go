package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config))
	//router.Use(static.Serve("/", static.LocalFile("./template", true)))

	router.Static("/static", "./")

	// Define a route and its handler function
	router.GET("/listings", DirRequest)
	

	// Run the server on port 8080
	router.Run(":8080")
}
