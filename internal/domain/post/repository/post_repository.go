package repository

import "github.com/phn00dev/go-crud/internal/models"

type PostRepository interface {
	GetAll(userId int) ([]models.Post, error)
	GetOne(userId, postId int) (*models.Post, error)
	Create(post models.Post) error
	Update(postId int, post models.Post) error
	Delete(userId, postId int) error
	//
	SlugExists(slug string) bool

	// public Post
	GetAllPost() ([]models.Post, error)
	GetPostBySlug(postSlug string) (*models.Post, error)
}
