package routes

import (
	"restapi/controllers"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(rg *gin.RouterGroup) {
	{
		rg.POST("/students", controllers.Create)
		rg.GET("/students", controllers.GetAll)
		rg.GET("/students/:id", controllers.GetOne)
		rg.PUT("/students/:id", controllers.Update)
		rg.DELETE("/students/:id", controllers.Delete)
	}
}