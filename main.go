package main

import (
	"go-gin-crud/config"
	"go-gin-crud/models"
	"go-gin-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Conexión y migración
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})

	
	// Rutas
	routes.UserRoutes(r)

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}
