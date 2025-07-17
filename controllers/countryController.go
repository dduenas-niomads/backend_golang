package controllers

import (
    "net/http"
    "backend_golang/config"
    "backend_golang/models"
    "github.com/gin-gonic/gin"
)

func GetCountries(c *gin.Context) {
    var countries []models.Country
    config.DB.Find(&countries)
    c.JSON(http.StatusOK, gin.H{"data": countries})
}

func CreateCountry(c *gin.Context) {
    var country models.Country
    if err := c.ShouldBindJSON(&country); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&country)
    c.JSON(http.StatusCreated, gin.H{"data": country})
}

func UpdateCountry(c *gin.Context) {
    var country models.Country
    id := c.Param("id")
    if err := config.DB.First(&country, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Country not found"})
        return
    }
    if err := c.ShouldBindJSON(&country); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Save(&country)
    c.JSON(http.StatusOK, gin.H{"data": country})
}

func DeleteCountry(c *gin.Context) {
    var country models.Country
    id := c.Param("id")
    if err := config.DB.First(&country, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Country not found"})
        return
    }
    config.DB.Delete(&country)
    c.JSON(http.StatusOK, gin.H{"data": true})
}