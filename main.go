package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/static"

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

	router.Static("/static", "./")
	
	router.NoRoute(func(c *gin.Context) {
		c.File("./template/index.html")
	})
	// Define a route and its handler function
	router.GET("/listings", DirRequest)
	

	// Run the server on port 8080
	router.Run(":8080")
}
