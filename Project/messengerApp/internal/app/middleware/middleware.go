package middleware

import (
	"messengerApp/internal/app/models"
	"messengerApp/internal/app/repository"
	"messengerApp/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware(userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		// Обрезать префикс "Bearer " для корректного чтения токена
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		userID, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Извлечение пользователя из базы данных по ID
		user, err := userRepo.FindByID(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Сохраняем пользователя в контексте для дальнейшего использования
		c.Set("user", user)
		c.Next()
	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Проверка, является ли пользователь администратором
		userObj := user.(*models.User)
		if !userObj.IsAdmin() {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
