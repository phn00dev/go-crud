package dto

import "github.com/phn00dev/go-crud/internal/models"

type UserResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
}

func NewUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Format("02-01-2006 15:04:05"),
	}
}
