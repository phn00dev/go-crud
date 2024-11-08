package service

import (
	"github.com/phn00dev/go-crud/internal/domain/post/dto"
	postRepository "github.com/phn00dev/go-crud/internal/domain/post/repository"
	"github.com/phn00dev/go-crud/internal/domain/user/repository"
	"github.com/phn00dev/go-crud/internal/models"
	"github.com/phn00dev/go-crud/pkg/config"

)

type postServiceImp struct {
	postRepo postRepository.PostRepository
	userRepo repository.UserRepository
	config   *config.Config
}

func NewPostService(postRepo postRepository.PostRepository, userRepo repository.UserRepository,
	config *config.Config) PostService {
	return postServiceImp{
		postRepo: postRepo,
		userRepo: userRepo,
		config:   config,
	}
}

func (postService postServiceImp) GetAll(userId int) ([]models.Post, error) {
	panic("post service imp")
}

func (postService postServiceImp) GetOne(userId, postId int) (*models.Post, error) {
	panic("post service imp")
}

func (postService postServiceImp) CreatePost(userId int, createRequest dto.CreatePostRequest) error {
	panic("post service imp")
}

func (postService postServiceImp) UpdatePost(userId, postId int, updateRequest dto.UpdatePostRequest) error {
	panic("post service imp")
}

func (postService postServiceImp) DeletePost(userId, postId int) error {
	panic("post service imp")
}
