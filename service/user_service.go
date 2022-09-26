package service

import (
	"github.com/mehmetcanhoroz/digital-marketplace/repository"
)

type UserService interface {
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return userService{
		repository: repository,
	}
}
