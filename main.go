package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"backend_golang/config"
	"backend_golang/routes"
	"backend_golang/seeders"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	// Conectar a la base de datos (importante hacerlo antes de seeders)
	config.ConnectDatabase()

	// Ejecutar seeders UNA vez, después de la conexión a la base
	seeders.SeedAll()

	// Iniciar Gin
	router := gin.Default()

	// Configurar rutas
	routes.SetupRoutes(router)

	// Obtener puerto desde .env o usar 8080 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor en el puerto definido
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("No se pudo iniciar el servidor")
	}
}
