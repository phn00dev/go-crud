package service

import (
	"github.com/phn00dev/go-crud/internal/domain/user/dto"
	"github.com/phn00dev/go-crud/internal/models"

)

type UserService interface {
	GetUser(userId int) (*models.User, error)
	Update(userId int, updateRequest dto.UpdateUserRequest) error
	UpdatePassword(userId int, request dto.UpdateUserPasswordRequest) error
	Delete(userId int) error
}
