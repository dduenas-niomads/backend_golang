package routes

import (
    "backend_golang/controllers"
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {
        api.GET("/conceptos", controllers.GetConceptos)
        api.GET("/ciudades", controllers.GetCiudades)
        api.GET("/paises", controllers.GetPaises)
    }
}
