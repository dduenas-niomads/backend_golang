package controllers

import (
    "backend_golang/utils"
    "golang.org/x/crypto/bcrypt"
    "backend_golang/config"
    "backend_golang/models"
    "github.com/gin-gonic/gin"
    "net/http"
    
)

// Registro
func Register(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    var existing models.User
    if err := config.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
        c.JSON(400, gin.H{"error": "El email ya está registrado"})
        return
    }
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    input.Password = string(hashedPassword)
    if err := config.DB.Create(&input).Error; err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    token, _ := utils.GenerateJWT(input.ID)
    c.JSON(200, gin.H{"token": token, "user": gin.H{"id": input.ID, "email": input.Email}})
}

// Login
func Login(c *gin.Context) {
    var input models.User
    var user models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(401, gin.H{"error": "Usuario no encontrado"})
        return
    }
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(401, gin.H{"error": "Contraseña incorrecta"})
        return
    }
    token, _ := utils.GenerateJWT(user.ID)
    c.JSON(200, gin.H{"token": token, "user": gin.H{"id": user.ID, "email": user.Email}})
}

// CreateUser (hashea la contraseña si se envía)
func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if user.Password != "" {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hashear la contraseña"})
            return
        }
        user.Password = string(hashedPassword)
    }
    config.DB.Create(&user)
    c.JSON(http.StatusCreated, gin.H{"id": user.ID, "email": user.Email})
}
// Obtener todos los usuarios (requiere autenticación)
func GetUsers(c *gin.Context) {
    var users []models.User
    if err := config.DB.Find(&users).Error; err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    // No exponer contraseñas
    var safeUsers []gin.H
    for _, u := range users {
        safeUsers = append(safeUsers, gin.H{"id": u.ID, "email": u.Email})
    }
    c.JSON(200, safeUsers)
}

// Obtener usuario por ID (requiere autenticación)
func GetUserByID(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := config.DB.First(&user, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "Usuario no encontrado"})
        return
    }
    c.JSON(200, gin.H{"id": user.ID, "email": user.Email})
}

// UpdateUser (hashea la contraseña si se actualiza)
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := config.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
        return
    }

    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Solo actualiza los campos necesarios
    user.Email = input.Email
    if input.Password != "" {
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hashear la contraseña"})
            return
        }
        user.Password = string(hashedPassword)
    }
    config.DB.Save(&user)
    c.JSON(http.StatusOK, gin.H{"id": user.ID, "email": user.Email})
}

// DeleteUser
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := config.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
        return
    }
    if err := config.DB.Delete(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}