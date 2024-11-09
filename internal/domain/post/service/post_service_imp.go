package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-crud/internal/domain/post/dto"
	postRepository "github.com/phn00dev/go-crud/internal/domain/post/repository"
	"github.com/phn00dev/go-crud/internal/domain/user/repository"
	"github.com/phn00dev/go-crud/internal/models"
	filemanager "github.com/phn00dev/go-crud/internal/utils/file_manager"
	generateuniqueslug "github.com/phn00dev/go-crud/internal/utils/generateUniqueSlug"
	"github.com/phn00dev/go-crud/pkg/config"
)

type postServiceImp struct {
	postRepo postRepository.PostRepository
	userRepo repository.UserRepository
	config   *config.Config
}

func NewPostService(postRepo postRepository.PostRepository, userRepo repository.UserRepository, config *config.Config) PostService {
	return postServiceImp{
		postRepo: postRepo,
		userRepo: userRepo,
		config:   config,
	}
}

func (postService postServiceImp) GetAll(userId int) ([]models.Post, error) {
	// Get user
	user, err := postService.userRepo.GetUser(userId)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	// Get user's posts
	posts, err := postService.postRepo.GetAll(user.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %v", err)
	}
	return posts, nil
}

func (postService postServiceImp) GetOne(userId, postId int) (*models.Post, error) {
	// Get user
	user, err := postService.userRepo.GetUser(userId)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	// Get user post
	post, err := postService.postRepo.GetOne(user.Id, postId)
	if err != nil {
		return nil, fmt.Errorf("post not found: %v", err)
	}
	return post, nil
}

func (postService postServiceImp) CreatePost(ctx *gin.Context, userId int, createRequest dto.CreatePostRequest) error {
	// Get user
	user, err := postService.userRepo.GetUser(userId)
	if err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	// Upload file
	postImagePath, err := filemanager.UploadFile(ctx, "post_image", "post_images", postService.config.FolderConfig.PublicPath)
	if err != nil {
		return fmt.Errorf("file upload failed: %v", err)
	}

	// Generate unique slug for post
	postSlug, err := generateuniqueslug.GenerateUniqueSlug(createRequest.PostTitle, postService.postRepo)
	if err != nil {
		return fmt.Errorf("failed to generate unique slug: %v", err)
	}

	// Create new post
	newPost := models.Post{
		PostTitle: createRequest.PostTitle,
		PostSlug:  postSlug,
		PostDesc:  createRequest.PostDesc,
		PostImage: postImagePath,
		ViewCount: 0,
		UserId:    user.Id,
	}

	// Save post to repository
	if err := postService.postRepo.Create(newPost); err != nil {
		// Delete uploaded file if post creation fails
		if delErr := filemanager.DeleteFile(postImagePath); delErr != nil {
			return fmt.Errorf("post creation failed and could not delete file: %v", delErr)
		}
		return fmt.Errorf("failed to create post: %v", err)
	}
	return nil
}

func (postService postServiceImp) UpdatePost(ctx *gin.Context, userId, postId int, updateRequest dto.UpdatePostRequest) error {
	// Get post
	post, err := postService.postRepo.GetOne(userId, postId)
	if err != nil {
		return fmt.Errorf("post not found: %v", err)
	}

	// Check if the user is authorized to update the post
	if post.UserId != userId {
		return errors.New("user is not authorized to update this post")
	}

	// Check if file exists in the form data
	file, err := ctx.FormFile("post_image")
	if err != nil && err != http.ErrMissingFile {
		return fmt.Errorf("error getting file from request: %v", err)
	}

	// If there’s a new file, upload it and delete the old one
	if file != nil {
		// Upload new file
		postImagePath, err := filemanager.UploadFile(ctx, "post_image", "post_images", postService.config.FolderConfig.PublicPath)
		if err != nil {
			return fmt.Errorf("file upload failed: %v", err)
		}

		// Delete old file if it exists
		if post.PostImage != "" {
			if delErr := filemanager.DeleteFile(post.PostImage); delErr != nil {
				return fmt.Errorf("failed to delete old file: %v", delErr)
			}
		}

		// Update post image path
		post.PostImage = postImagePath
	}

	// Check if PostTitle has changed, then generate a new unique slug
	if post.PostTitle != updateRequest.PostTitle {
		postSlug, err := generateuniqueslug.GenerateUniqueSlug(updateRequest.PostTitle, postService.postRepo)
		if err != nil {
			return fmt.Errorf("failed to generate unique slug: %v", err)
		}
		post.PostTitle = updateRequest.PostTitle
		post.PostSlug = postSlug
	}

	// Update PostDesc if it has changed
	if post.PostDesc != updateRequest.PostDesc {
		post.PostDesc = updateRequest.PostDesc
	}

	// Update post in repository
	if err := postService.postRepo.Update(post.Id, *post); err != nil {
		return fmt.Errorf("failed to update post: %v", err)
	}

	return nil
}

func (postService postServiceImp) DeletePost(userId, postId int) error {
	// Get user
	user, err := postService.userRepo.GetUser(userId)
	if err != nil {
		return fmt.Errorf("user not found: %v", err)
	}

	// Get post
	post, err := postService.postRepo.GetOne(user.Id, postId)
	if err != nil {
		return fmt.Errorf("post not found: %v", err)
	}

	// Delete post image
	if post.PostImage != "" {
		if delErr := filemanager.DeleteFile(post.PostImage); delErr != nil {
			return fmt.Errorf("failed to delete post image: %v", delErr)
		}
	}

	// Delete post
	if err := postService.postRepo.Delete(post.UserId, post.Id); err != nil {
		return fmt.Errorf("failed to delete post: %v", err)
	}

	return nil
}
