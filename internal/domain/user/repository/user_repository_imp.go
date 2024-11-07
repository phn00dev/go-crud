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

func (userRepo userRepositoryImp) GetUser(userId int) (*models.User, error) {
	var user models.User
	if err := userRepo.db.Where("id =?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo userRepositoryImp) Update(userId int, user models.User) error {
	return userRepo.db.Where("id=?", userId).Updates(&user).Error
}

func (userRepo userRepositoryImp) UpdatePassword(userId int, password string) error {
	return userRepo.db.Model(&models.User{}).Where("id=?", userId).Update("password_hash", password).Error
}

func (userRepo userRepositoryImp) Delete(userId int) error {
	return userRepo.db.Where("id=?", userId).Error
}
