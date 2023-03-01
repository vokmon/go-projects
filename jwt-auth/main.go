package main

import (
	routes "jwt-auth/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	v1 := router.Group("v1")
	{
		v1.GET("/api-1", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": "Access granted for api-1"})
		})

		v1.GET("/api-2", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": "Access granted for api-2"})
		})
	}

	router.Run(":" + port)
}
