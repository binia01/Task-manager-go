package Usecases

import (
	"fmt"

	"task-manager-go/Domain"
	"task-manager-go/Infrastructure"
	"task-manager-go/Repositories"
)

type UserUsecase interface {
	Login(user Domain.User) (string, error)
	Register(user Domain.User) (*Domain.User, error)
	UpdateRole(username string) error
}

type userUsecase struct {
	userRepo        Repositories.UserRepository
	passwordService Infrastructure.PasswordService
	jwtService      Infrastructure.JWTService
}

func NewUserUsecase(userRepo Repositories.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo:        userRepo,
		passwordService: Infrastructure.NewPasswordService(),
		jwtService:      Infrastructure.NewJWTService(string(Infrastructure.JwtSecret)),
	}
}

func (uu *userUsecase) Login(user Domain.User) (string, error) {
	existingUser, err := uu.userRepo.FindByUsername(user.Username)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	err = uu.passwordService.ComparePassword(existingUser.Password, user.Password)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	token, err := uu.jwtService.GenerateToken(*existingUser)
	if err != nil {
		return "", fmt.Errorf("error generating token")
	}

	return token, nil
}

func (uu *userUsecase) Register(user Domain.User) (*Domain.User, error) {
	hashedPassword, err := uu.passwordService.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password")
	}
	user.Password = hashedPassword

	return uu.userRepo.CreateUser(user)
}

func (uu *userUsecase) UpdateRole(username string) error {
	return uu.userRepo.UpdateUserRole(username, Domain.AdminRole)
}
