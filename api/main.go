package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.POST("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "User created",
		})
	})

	router.GET("/depth", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Depth fetched",
		})
	})

	router.GET("/klines", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Klines fetched",
		})
	})

	router.GET("/market", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Market fetched",
		})
	})

	router.GET("/tickers", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Tickers fetched",
		})
	})

	router.GET("/order", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Orders fetched",
		})
	})

	router.POST("/order", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Order executed",
		})
	})

	router.DELETE("/order", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Order deleted or canceled",
		})
	})

	router.GET("/orders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Orders fetched",
		})
	})

	router.DELETE("/orders", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Order deleted or canceled",
		})
	})

	router.Run()

}
