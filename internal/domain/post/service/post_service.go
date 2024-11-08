package service

import (
	"github.com/phn00dev/go-crud/internal/domain/post/dto"
	"github.com/phn00dev/go-crud/internal/models"
)

type PostService interface {
	GetAll(userId int) ([]models.Post, error)
	GetOne(userId, postId int) (*models.Post, error)
	CreatePost(userId int, createRequest dto.CreatePostRequest) error
	UpdatePost(userId, postId int, updateRequest dto.UpdatePostRequest) error
	DeletePost(userId, postId int) error
}
