package main

import (
	"restapi/controllers"
	"restapi/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDb()
}

func main() {
	

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})
	
	// students routes
	r.POST("/students", controllers.Create)
	r.GET("/students", controllers.GetAll)
	r.GET("/students/:id", controllers.GetOne)
	r.PUT("/students/:id", controllers.Update)
	r.DELETE("/students/:id", controllers.Delete)


	r.Run()
}