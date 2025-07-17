package config

import (
	// "os"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "backend_golang/models"
)

var DB *gorm.DB

func ConnectDatabase() {

	dbURL := "data.db"

    database, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
    if err != nil {
        log.Fatal("No se pudo conectar a la base de datos:", err)
    }
    DB = database
    DB.AutoMigrate(&models.User{}, &models.Country{}, &models.City{})
}