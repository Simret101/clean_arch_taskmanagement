package middleware

import (
	"net/http"
	"strings"
	"task/usecase"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	UserUsecase *usecase.LoginUsecase
}

func NewAuthMiddleware(userUsecase *usecase.LoginUsecase) *AuthMiddleware {
	return &AuthMiddleware{UserUsecase: userUsecase}
}

func (m *AuthMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			c.Abort()
			return
		}

		claims, err := m.UserUsecase.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
