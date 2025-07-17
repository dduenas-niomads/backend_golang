package routes

import (
	"backend_golang/controllers"
	"backend_golang/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Conceptos
		api.GET("/conceptos", controllers.GetConceptos)

		// Countries
		countries := api.Group("/countries")
		{
			countries.GET("/", controllers.GetCountries)
			countries.POST("/", controllers.CreateCountry)
			countries.PUT("/:id", controllers.UpdateCountry)
			countries.DELETE("/:id", controllers.DeleteCountry)
		}

		// Cities
		cities := api.Group("/cities")
		{
			cities.GET("/", controllers.GetCities)
			cities.POST("/", controllers.CreateCity)
			cities.PUT("/:id", controllers.UpdateCity)
			cities.DELETE("/:id", controllers.DeleteCity)
		}

		// Users
		users := api.Group("/users")
		{
			users.POST("/register", controllers.Register)
			users.POST("/login", controllers.Login)

			users.GET("/", middleware.JWTAuth(), controllers.GetUsers)
			users.GET("/:id", middleware.JWTAuth(), controllers.GetUserByID)
			users.POST("/", middleware.JWTAuth(), controllers.CreateUser)
			users.PUT("/:id", middleware.JWTAuth(), controllers.UpdateUser)
			users.DELETE("/:id", middleware.JWTAuth(), controllers.DeleteUser)
		}
	}
}