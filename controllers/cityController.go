package controllers

import (
    "net/http"
    "go-gin-crud/config"
    "go-gin-crud/models"
    "github.com/gin-gonic/gin"
)

type CityResponse struct {
    ID          uint   `json:"id"`
    Name        string `json:"name"`
    CountryID   uint   `json:"country_id"`
    CountryName string `json:"country_name"`
}

// GET /cities
func GetCities(c *gin.Context) {
    var cities []models.City
    config.DB.Find(&cities)

    var response []CityResponse
    for _, city := range cities {
        var country models.Country
        config.DB.First(&country, city.CountryID)
        response = append(response, CityResponse{
            ID:          city.ID,
            Name:        city.Name,
            CountryID:   city.CountryID,
            CountryName: country.Name,
        })
    }
    c.JSON(http.StatusOK, gin.H{"data": response})
}

// POST /cities
func CreateCity(c *gin.Context) {
    var city models.City
    if err := c.ShouldBindJSON(&city); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&city)
    var country models.Country
    config.DB.First(&country, city.CountryID)
    response := CityResponse{
        ID:          city.ID,
        Name:        city.Name,
        CountryID:   city.CountryID,
        CountryName: country.Name,
    }
    c.JSON(http.StatusCreated, gin.H{"data": response})
}

// PUT /cities/:id
func UpdateCity(c *gin.Context) {
    var city models.City
    id := c.Param("id")
    if err := config.DB.First(&city, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
        return
    }
    if err := c.ShouldBindJSON(&city); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Save(&city)
    var country models.Country
    config.DB.First(&country, city.CountryID)
    response := CityResponse{
        ID:          city.ID,
        Name:        city.Name,
        CountryID:   city.CountryID,
        CountryName: country.Name,
    }
    c.JSON(http.StatusOK, gin.H{"data": response})
}

// DELETE /cities/:id
func DeleteCity(c *gin.Context) {
    var city models.City
    id := c.Param("id")
    if err := config.DB.First(&city, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
        return
    }
    config.DB.Delete(&city)
    c.JSON(http.StatusOK, gin.H{"data": true})
}