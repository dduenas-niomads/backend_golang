package middleware

import (
    "strings"
    "backend_golang/utils"
    "github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.AbortWithStatusJSON(401, gin.H{"error": "Token requerido"})
            return
        }
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := utils.ValidateJWT(tokenString)
        if err != nil {
            c.AbortWithStatusJSON(401, gin.H{"error": "Token inválido"})
            return
        }
        c.Set("user_id", claims["user_id"])
        c.Next()
    }
}
