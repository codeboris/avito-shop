package middleware

import (
	"github.com/codeboris/avito-shop/pkg/jwtutil"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем заголовок Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}

		// Проверяем, что токен начинается с "Bearer "
		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			return
		}

		// Извлекаем сам токен
		tokenString := strings.TrimPrefix(authHeader, prefix)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "empty token"})
			return
		}

		// Валидируем токен
		userID, err := jwtutil.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid or expired token"})
			return
		}

		// Передаём userID в контекст
		c.Set("userID", userID)
		c.Next()
	}
}
