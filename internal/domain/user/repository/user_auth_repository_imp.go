package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/phn00dev/go-crud/internal/models"
)

type authUserRepositoryImp struct {
	db *gorm.DB
}

func NewAuthUserRepository(db *gorm.DB) AuthUserRepository {
	return authUserRepositoryImp{
		db: db,
	}
}

func (authUserRepo authUserRepositoryImp) Create(user models.User) (*models.User, error) {
	if err := authUserRepo.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (authUserRepo authUserRepositoryImp) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := authUserRepo.db.Where("username =?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (authUserRepo authUserRepositoryImp) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := authUserRepo.db.Select("users.username").Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
