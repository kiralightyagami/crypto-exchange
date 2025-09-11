package main

import (
	"crypto-exchange-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Setup all route groups
	routes.SetupHomeRoutes(router)
	routes.SetupUserRoutes(router)
	routes.SetupMarketRoutes(router)
	routes.SetupOrderRoutes(router)

	router.Run()
}
