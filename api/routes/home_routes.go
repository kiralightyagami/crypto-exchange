package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupHomeRoutes(router *gin.Engine) {
	router.GET("/", getHome)
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
