package repository

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-crud/internal/models"
)

type userRepositoryImp struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepositoryImp{
		db: db,
	}
}

func (userRepo userRepositoryImp) GetUser(userId int, username string) (*models.User, error) {
	panic("user repo imp")
}

func (userRepo userRepositoryImp) Update(user models.User) error {
	panic("user repo imp")
}

func (userRepo userRepositoryImp) UpdatePassword(password string) error {
	panic("user repo imp")
}

func (userRepo userRepositoryImp) Delete(userId int, username string) error {
	panic("user repo imp")
}
