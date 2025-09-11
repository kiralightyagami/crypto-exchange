package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", createUser)
	}
}

func createUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User created",
	})
}
