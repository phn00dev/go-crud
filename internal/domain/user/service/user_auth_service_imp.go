package service

import (
	"errors"
	"fmt"

	"github.com/phn00dev/go-crud/internal/domain/user/dto"
	"github.com/phn00dev/go-crud/internal/domain/user/repository"
	"github.com/phn00dev/go-crud/internal/models"
	passwordhash "github.com/phn00dev/go-crud/internal/utils/password_hash"
	jwttoken "github.com/phn00dev/go-crud/pkg/jwt_token"

)

type authUserServiceImp struct {
	authUserRepo repository.AuthUserRepository
}

func NewAuthUserService(authUserRepo repository.AuthUserRepository) AuthUserService {
	return authUserServiceImp{
		authUserRepo: authUserRepo,
	}
}

// RegisterUser handles user registration
func (authUserService authUserServiceImp) RegisterUser(registerRequest dto.RegisterRequest) (*dto.UserAuthResponse, error) {
	// Check if user already exists
	existingUser, err := authUserService.authUserRepo.FindUserByUsername(registerRequest.Username)

	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username or email already exists")
	}

	// Create new user
	newUser := models.User{
		Username:     registerRequest.Username,
		PasswordHash: passwordhash.GeneratePassword(registerRequest.Password),
	}

	// Save user to the repository
	user, err := authUserService.authUserRepo.Create(newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	// Generate JWT token for user
	accessToken, err := jwttoken.GenerateJwtToken(user.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to generate JWT token: %v", err)
	}

	// Return successful response
	registerResponse := dto.NewUserAuthResponse(user, accessToken)
	return registerResponse, nil
}

// LoginUser handles user login
func (authUserService authUserServiceImp) LoginUser(loginRequest dto.LoginRequest) (*dto.UserAuthResponse, error) {
	// Find user by username
	user, err := authUserService.authUserRepo.GetUserByUsername(loginRequest.Username)
	if err != nil {
		return nil, fmt.Errorf("username or password is incorrect")
	}
	if user == nil {
		return nil, errors.New("username or password is incorrect")
	}

	// Verify password
	if err := passwordhash.CheckPasswordHash(loginRequest.Password, user.PasswordHash); err != nil {
		return nil, errors.New("username or password is incorrect")
	}

	// Generate JWT token
	accessToken, err := jwttoken.GenerateJwtToken(user.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to generate JWT token: %v", err)
	}

	// Return successful response
	loginResponse := dto.NewUserAuthResponse(user, accessToken)
	return loginResponse, nil
}
