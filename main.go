package main

import (
	"github.com/fernando-g-fraga/FinanGo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.FinanGoRoutes(r)

	r.Run(":8080")

}
