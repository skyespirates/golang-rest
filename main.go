package main

import (
	"restapi/initializers"
	"restapi/routes"

	"github.com/gin-gonic/gin"
)

func initz() {
	initializers.LoadEnvVars()
	initializers.ConnectToDb()
}

func main() {
	

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})
	

	routes.SetupRoutes(r)

	r.Run(":3000")
}