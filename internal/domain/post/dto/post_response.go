package dto

import (
	"github.com/phn00dev/go-crud/internal/domain/user/dto"
	"github.com/phn00dev/go-crud/internal/models"

)

type PostResponse struct {
	Id        int              `json:"id"`
	PostTitle string           `json:"post_title"`
	PostSlug  string           `json:"post_slug"`
	PostDesc  string           `json:"post_desc"`
	PostImage string           `json:"post_image"`
	ViewCount int              `json:"view_count"`
	CreatedAt string           `json:"created_at"`
	UserId    int              `json:"user_id"`
	User      dto.UserResponse `json:"writer"`
}

func NewPostResponses(posts []models.Post) []PostResponse {
	var postResponses []PostResponse
	for _, post := range posts {
		postResponse := NewPostResponse(&post)
		postResponses = append(postResponses, postResponse)
	}
	return postResponses
}

func NewPostResponse(post *models.Post) PostResponse {
	return PostResponse{
		Id:        post.Id,
		PostTitle: post.PostTitle,
		PostSlug:  post.PostSlug,
		PostDesc:  post.PostDesc,
		PostImage: post.PostImage,
		ViewCount: post.ViewCount,
		CreatedAt: post.CreatedAt.Format("02-01-2006 15:04:05"),
		UserId:    post.UserId,
		User: dto.UserResponse{
			Id:        post.User.Id,
			Username:  post.User.Username,
			CreatedAt: post.User.CreatedAt.Format("02-01-2006 15:04:05"),
		},
	}
}
