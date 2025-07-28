package Usecases

import (
	"fmt"

	"task-manager-go/Domain"
)

type UserUsecase struct {
	userRepo        Domain.IUserRepository
	passwordService Domain.IPasswordService
	jwtService      Domain.IJWTService
}

func NewUserUsecase(userRepo Domain.IUserRepository, passwordService Domain.IPasswordService, jwtService Domain.IJWTService) *UserUsecase {
	return &UserUsecase{
		userRepo:        userRepo,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

func (uu *UserUsecase) Login(user Domain.User) (string, error) {
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

func (uu *UserUsecase) Register(user Domain.User) (*Domain.User, error) {
	hashedPassword, err := uu.passwordService.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("error hashing password")
	}
	user.Password = hashedPassword

	return uu.userRepo.CreateUser(user)
}

func (uu *UserUsecase) UpdateRole(username string) error {
	return uu.userRepo.UpdateUserRole(username, Domain.AdminRole)
}
