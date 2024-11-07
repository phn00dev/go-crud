package service

import (
	"github.com/phn00dev/go-crud/internal/domain/user/dto"
	"github.com/phn00dev/go-crud/internal/domain/user/repository"
	"github.com/phn00dev/go-crud/internal/models"
)

type userServiceImp struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userServiceImp{
		userRepo: repo,
	}
}

func (userService userServiceImp) GetUser(userId int, username string) (*models.User, error) {
	panic("user service imp")
}

func (userService userServiceImp) Update(updateRequest dto.UpdateUserRequest) error {
	panic("user service imp")
}

func (userService userServiceImp) UpdatePassword(request dto.UpdateUserPasswordRequest) error {
	panic("user service imp")
}

func (userService userServiceImp) Delete(userId int, username string) error {
	panic("user service imp")
}
