package Infrastructure

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtSecret = []byte("wN7z@JrV3uK!dXp2qT$eGh8yF9cLb6mZ")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const authHeaderKey = "Authorization"

		authHeader := c.GetHeader(authHeaderKey)
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenStr := authParts[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return JwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid JWT claims"})
			c.Abort()
			return
		}

		// Save claims in context for later use
		c.Set("claims", claims)

		c.Next()
	}
}

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, exists := c.Get("claims")
		if !exists {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, ok := val.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid JWT claims"})
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok || role != requiredRole {
			c.JSON(403, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
