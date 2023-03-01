package routes

import (
	"jwt-auth/controllers"
	"jwt-auth/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	users := incomingRoutes.Group("/users", middlewares.Authenticate())
	{
		users.GET("/", controllers.GetUsers())
		users.GET("/:user_id", controllers.GetUser())
	}
}
