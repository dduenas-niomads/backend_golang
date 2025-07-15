// main.go (No necesita cambios si ya lo tienes así)
package main

import (
    "github.com/gin-gonic/gin"
    "go-gin-crud/routes"
    "go-gin-crud/config"
)

func main() {
    config.ConnectDatabase()


    SeedAllData(config.DB)

    r := gin.Default()
    // Middleware CORS
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })

    routes.SetupRoutes(r)
    r.Run(":8080")
}