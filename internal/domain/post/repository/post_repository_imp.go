package repository

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-crud/internal/models"
)

type postRepositoryImp struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return postRepositoryImp{
		db: db,
	}
}

func (postRepo postRepositoryImp) GetAll(userId int) ([]models.Post, error) {
	panic("post repo imp")
}

func (postRepo postRepositoryImp) GetOne(userId, postId int) (*models.Post, error) {
	panic("post repo imp")
}

func (postRepo postRepositoryImp) Create(post models.Post) error {
	panic("post repo imp")
}

func (postRepo postRepositoryImp) Update(postId int, post models.Post) error {
	panic("post repo imp")
}

func (postRepo postRepositoryImp) Delete(userId, postId int) error {
	panic("post repo imp")
}
