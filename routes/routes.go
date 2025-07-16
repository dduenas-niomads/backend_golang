package routes

import (
    "github.com/gin-gonic/gin"
    "backend_golang/controllers"
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
            users.GET("/", controllers.GetUsers)
            users.GET("/:id", controllers.GetUserByID)
            users.POST("/", controllers.CreateUser)
            users.PUT("/:id", controllers.UpdateUser)
            users.DELETE("/:id", controllers.DeleteUser)
        }
    }
}