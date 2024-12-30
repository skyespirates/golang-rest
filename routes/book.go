package routes

import (
	"restapi/controllers"

	"github.com/gin-gonic/gin"
)


func BookRoutes(rg *gin.RouterGroup) {
	{
		rg.GET("/books", controllers.GetAllBook)
		rg.GET("/books/:id", controllers.GetSingleBook)
	}
}