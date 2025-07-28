package Infrastructure_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"task-manager-go/Infrastructure"
)

func generateToken(secret string, role string) string {
	claims := jwt.MapClaims{
		"user_id":  "1",
		"username": "testuser",
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
	secret := "testsecret"
	token := generateToken(secret, "Admin")
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(Infrastructure.AuthMiddleware([]byte(secret)))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(Infrastructure.AuthMiddleware([]byte("secret")))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}
