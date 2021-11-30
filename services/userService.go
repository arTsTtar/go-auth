package services

import (
	"go-auth/entity"
	"go-auth/repository"
)

type UserService interface {
	GetAllUsers() (entity.Users, error)
}

type userService struct {
	authService    AuthService
	userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository, as AuthService) UserService {
	return userService{
		userRepository: ur,
		authService:    as,
	}
}

func (u userService) GetAllUsers() (entity.Users, error) {
	return u.userRepository.FindAllUsers()
}
