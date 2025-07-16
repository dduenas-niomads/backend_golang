package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"backend_golang/database"
	"backend_golang/routes"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	// Conectar a la base de datos
	database.ConnectDB()

	// Iniciar Gin
	router := gin.Default()

	// Configurar rutas
	routes.RegisterRoutes(router)

	// Puerto desde .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("No se pudo iniciar el servidor")
	}
}
