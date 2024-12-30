package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	routerGroup := r.Group("/api")
	StudentRoutes(routerGroup)
	BookRoutes(routerGroup)
}