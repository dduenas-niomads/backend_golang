package controllers

import (
	"net/http"
	"strconv"

	"backend_golang/database"
	"backend_golang/models"

	"github.com/gin-gonic/gin"
)

// GET /concepts
func GetConcepts(c *gin.Context) {
	var concepts []models.Concept
	if err := database.DB.Find(&concepts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, concepts)
}

// GET /concepts/create
func CreateMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "create"})
}

// POST /concepts
func CreateConcept(c *gin.Context) {
	var concept models.Concept
	if err := c.ShouldBindJSON(&concept); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&concept).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, concept)
}

// GET /concepts/:id
func GetConceptByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var concept models.Concept
	if err := database.DB.First(&concept, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Concepto no encontrado"})
		return
	}

	c.JSON(http.StatusOK, concept)
}

// GET /concepts/:id/edit
func EditMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "edit"})
}

// PUT /concepts/:id
func UpdateConcept(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var concept models.Concept
	if err := database.DB.First(&concept, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Concepto no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&concept); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&concept).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, concept)
}

// DELETE /concepts/:id
func DeleteConcept(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var concept models.Concept
	if err := database.DB.First(&concept, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Concepto no encontrado"})
		return
	}

	if err := database.DB.Delete(&concept).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
