package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine) {
	orderGroup := router.Group("/orders")
	{
		orderGroup.GET("/", getOrders)
		orderGroup.DELETE("/", deleteAllOrders)
	}

	// Single order routes
	router.GET("/order", getOrder)
	router.POST("/order", createOrder)
	router.DELETE("/order", deleteOrder)
}

func getOrder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Orders fetched",
	})
}

func createOrder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Order executed",
	})
}

func deleteOrder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Order deleted or canceled",
	})
}

func getOrders(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Orders fetched",
	})
}

func deleteAllOrders(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Orders deleted or canceled",
	})
}
