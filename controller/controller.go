package controller

import "github.com/gin-gonic/gin"

func FinanGoController(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Message": "All working now..Keep Going!",
	})
}
