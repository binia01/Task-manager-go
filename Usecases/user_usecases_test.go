package Usecases_test

import (
	"task-manager-go/Domain"
	"task-manager-go/Usecases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct{ mock.Mock }
type MockPasswordService struct{ mock.Mock }
type MockJWTService struct{ mock.Mock }

func (m *MockUserRepo) FindByUsername(username string) (*Domain.User, error) {
	args := m.Called(username)
	return args.Get(0).(*Domain.User), args.Error(1)
}
func (m *MockUserRepo) CreateUser(user Domain.User) (*Domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(*Domain.User), args.Error(1)
}
func (m *MockUserRepo) UpdateUserRole(username string, role Domain.Role) error {
	args := m.Called(username, role)
	return args.Error(0)
}
func (m *MockPasswordService) HashPassword(pw string) (string, error) {
	args := m.Called(pw)
	return args.String(0), args.Error(1)
}
func (m *MockPasswordService) ComparePassword(hashed, pw string) error {
	args := m.Called(hashed, pw)
	return args.Error(0)
}
func (m *MockJWTService) GenerateToken(user Domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func TestLogin_Success(t *testing.T) {
	user := Domain.User{Username: "test", Password: "plain", Role: Domain.UserRole}
	stored := user
	stored.Password = "hashed"

	mockUserRepo := new(MockUserRepo)
	mockPw := new(MockPasswordService)
	mockJwt := new(MockJWTService)
	uc := Usecases.NewUserUsecase(mockUserRepo, mockPw, mockJwt)

	mockUserRepo.On("FindByUsername", "test").Return(&stored, nil)
	mockPw.On("ComparePassword", "hashed", "plain").Return(nil)
	mockJwt.On("GenerateToken", stored).Return("token123", nil)

	token, err := uc.Login(user)

	assert.NoError(t, err)
	assert.Equal(t, "token123", token)
}
