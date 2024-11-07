package dto

import "github.com/phn00dev/go-crud/internal/models"

type UserAuthResponse struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

func NewUserAuthResponse(user *models.User, accessToken string) *UserAuthResponse {
	return &UserAuthResponse{
		Id:          user.Id,
		Username:    user.Username,
		AccessToken: accessToken,
	}
}
