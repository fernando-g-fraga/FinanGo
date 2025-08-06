package routes

import (
	"github.com/fernando-g-fraga/FinanGo/controller"
	"github.com/gin-gonic/gin"
)

func FinanGoRoutes(r *gin.Engine) {
	routes := r.Group("/api")

	routes.GET("/test", controller.FinanGoController)
}
