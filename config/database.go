package config

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "go-gin-crud/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    database, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("No se pudo conectar a la base de datos:", err)
    }
    DB = database
    DB.AutoMigrate(&models.User{}, &models.Country{}, &models.City{})
}