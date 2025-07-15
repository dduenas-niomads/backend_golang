package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-crud/controllers"
)

func UserRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("/", controllers.GetUsers)
		users.GET("/:id", controllers.GetUserByID)
		users.POST("/", controllers.CreateUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}
}
