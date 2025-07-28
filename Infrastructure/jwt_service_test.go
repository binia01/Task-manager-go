package Infrastructure_test

import (
	"task-manager-go/Domain"
	"task-manager-go/Infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	service := Infrastructure.NewJWTService("secret")
	user := Domain.User{ID: [12]byte{}, Username: "testuser", Role: Domain.UserRole}

	token, err := service.GenerateToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
