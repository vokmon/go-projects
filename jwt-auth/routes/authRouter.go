package routes

import (
	"jwt-auth/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	// apply middleware for a grop
	//rg1 := r.Group("/group1", middleware1())
	// router.Group("/test").Use(middleWare())

	// apply middleware for a specific route
	// r.GET("/validate", middlewares.RequireAuth, controllers.Validate)

	auth := incomingRoutes.Group("/auth")
	{
		auth.POST("/signup", controllers.Signup())
		auth.POST("/login", controllers.Login())
	}
}
