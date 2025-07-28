package Infrastructure_test

import (
	"task-manager-go/Infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashAndComparePassword(t *testing.T) {
	service := Infrastructure.NewPasswordService()

	pw := "mypassword"
	hashed, err := service.HashPassword(pw)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)

	err = service.ComparePassword(hashed, pw)
	assert.NoError(t, err)

	err = service.ComparePassword(hashed, "wrongpw")
	assert.Error(t, err)
}
