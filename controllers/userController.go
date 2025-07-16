package controllers


import (
	"backend_golang/config"
	"backend_golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv" 
)

func GetUsers(c *gin.Context) {
	var users []models.User
	search := c.Query("search")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "5")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	offset := (pageInt - 1) * limitInt

	query := config.DB

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	result := query.Limit(limitInt).Offset(offset).Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Total count
	var total int64
	config.DB.Model(&models.User{}).Count(&total)

	c.JSON(http.StatusOK, gin.H{
		"data":  users,
		"total": total,
		"page":  pageInt,
		"limit": limitInt,
	})
}


func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.Delete(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}
