package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"backend_golang/config"
)

var DB *gorm.DB

func ConnectDB() {
	dbURL := config.GetEnv("DATABASE_URL", "./data.db")

	database, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a la base de datos: ", err)
	}

	DB = database
}

