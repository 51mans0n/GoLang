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
		// Trim the "Bearer" prefix to read the token correctly
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		userID, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Retrieving a user from a database by ID
		user, err := userRepo.FindByID(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

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

		// Checking if a user is an administrator
		userObj := user.(*models.User)
		if !userObj.IsAdmin() {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
