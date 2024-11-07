package service

import (
	"github.com/phn00dev/go-crud/internal/domain/user/dto"
	"github.com/phn00dev/go-crud/internal/models"
)

type UserService interface {
	GetUser(userId int, username string) (*models.User, error)
	Update(updateRequest dto.UpdateUserRequest) error
	UpdatePassword(request dto.UpdateUserPasswordRequest) error
	Delete(userId int, username string) error
}
