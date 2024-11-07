package repository

import "github.com/phn00dev/go-crud/internal/models"

type UserRepository interface {
	// user profile data
	GetUser(userId int, username string) (*models.User, error)
	Update(userId int, user models.User) error
	UpdatePassword(userId int, password string) error
	Delete(userId int, username string) error
}
