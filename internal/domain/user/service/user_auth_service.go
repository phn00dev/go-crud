package service

import (
	"github.com/phn00dev/go-crud/internal/domain/user/dto"
)

type AuthUserService interface {
	RegisterUser(registerRequest dto.RegisterRequest) (*dto.UserAuthResponse, error)
	LoginUser(loginRequest dto.LoginRequest) (*dto.UserAuthResponse, error)
}
