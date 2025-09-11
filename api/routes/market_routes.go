package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupMarketRoutes(router *gin.Engine) {
	marketGroup := router.Group("/market")
	{
		marketGroup.GET("/depth", getDepth)
		marketGroup.GET("/klines", getKlines)
		marketGroup.GET("/trades", getTrades)
		marketGroup.GET("/tickers", getTickers)
	}
}

func getDepth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Depth fetched",
	})
}

func getKlines(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Klines fetched",
	})
}

func getTrades(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Trades fetched",
	})
}

func getTickers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Tickers fetched",
	})
}
