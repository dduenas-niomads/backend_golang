package routes

import (
    "github.com/gin-gonic/gin"
    "go-gin-crud/controllers"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/countries/", controllers.GetCountries)
    r.POST("/countries/", controllers.CreateCountry)
    r.PUT("/countries/:id", controllers.UpdateCountry)
    r.DELETE("/countries/:id", controllers.DeleteCountry)

    // Ciudades
    r.GET("/cities/", controllers.GetCities)
    r.POST("/cities/", controllers.CreateCity)
    r.PUT("/cities/:id", controllers.UpdateCity)
    r.DELETE("/cities/:id", controllers.DeleteCity)

	// Usuarios
    r.GET("/users/", controllers.GetUsers)
    r.GET("/users/:id", controllers.GetUserByID)
    r.POST("/users/", controllers.CreateUser)
    r.PUT("/users/:id", controllers.UpdateUser)
    r.DELETE("/users/:id", controllers.DeleteUser)

	
}