package repository

import "github.com/phn00dev/go-crud/internal/models"

type AuthUserRepository interface {
	Create(user models.User) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	FindUserByUsername(username string) (*models.User, error)
}
