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
	var posts []models.Post
	if err := postRepo.db.Where("user_id=?", userId).Order("id desc").Preload("User").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (postRepo postRepositoryImp) GetOne(userId, postId int) (*models.Post, error) {
	var post models.Post
	if err := postRepo.db.Where("id=? AND user_id=?", postId, userId).Preload("User").First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (postRepo postRepositoryImp) Create(post models.Post) error {
	return postRepo.db.Create(&post).Error
}

func (postRepo postRepositoryImp) Update(postId int, post models.Post) error {
	return postRepo.db.Model(&models.Post{}).Where("id=?", postId).Updates(&post).Error
}

func (postRepo postRepositoryImp) Delete(userId, postId int) error {
	return postRepo.db.Where("id =? AND user_id=?", postId, userId).Delete(&models.Post{}).Error
}

// postRepositoryImp structynyň içinde ýazylýar
func (postRepo postRepositoryImp) SlugExists(slug string) bool {
	var count int64
	// Slug bilen deň gelýän ýazgylar sanalýar
	postRepo.db.Model(&models.Post{}).Where("post_slug = ?", slug).Count(&count)
	return count > 0
}

func (postRepo postRepositoryImp) GetAllPost() ([]models.Post, error) {
	var posts []models.Post
	if err := postRepo.db.Order("id desc").Preload("User").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (postRepo postRepositoryImp) GetPostBySlug(postSlug string) (*models.Post, error) {
	var post models.Post
	if err := postRepo.db.Where("post_slug = ?", postSlug).Preload("User").First(&post).Error; err != nil {
		return nil, err
	}
	post.ViewCount += 1
	if err := postRepo.db.Model(&post).UpdateColumn("view_count", post.ViewCount).Error; err != nil {
		return nil, err
	}

	return &post, nil
}
