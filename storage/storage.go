package storage

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"backend_golang/models"
)

var DB *gorm.DB

func Connect() {
	// Asegurarse de que el directorio de la base de datos exista
	err := os.MkdirAll("database", os.ModePerm)
	if err != nil {
		log.Fatal("No se pudo crear el directorio de base de datos:", err)
	}

	// Conectarse a SQLite y crear/abrir el archivo
	db, err := gorm.Open(sqlite.Open("database/data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	// Migrar el modelo Country (crear tabla si no existe)
	err = db.AutoMigrate(&models.Country{})
	if err != nil {
		log.Fatal("Error en AutoMigrate:", err)
	}

	log.Println("Conexión a la base de datos establecida y modelo migrado.")
	DB = db
}
