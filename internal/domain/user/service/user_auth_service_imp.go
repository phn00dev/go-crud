package service

import (
	"errors"

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

func (authUserService authUserServiceImp) RegisterUser(registerRequest dto.RegisterRequest) (*dto.UserAuthResponse, error) {
	existingUser, err := authUserService.authUserRepo.GetUserByUsername(registerRequest.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	newUser := models.User{
		Username:     registerRequest.Username,
		PasswordHash: passwordhash.GeneratePassword(registerRequest.Password),
	}
	// create user
	user, err := authUserService.authUserRepo.Create(newUser)
	if err != nil {
		return nil, err
	}
	// generate token
	accessToken, err := jwttoken.GenerateJwtToken(user.Id, user.Username)
	if err != nil {
		return nil, err
	}
	// register response
	registerResponse := dto.NewUserAuthResponse(user, accessToken)
	return registerResponse, nil
}
func (authUserService authUserServiceImp) LoginUser(loginRequest dto.LoginRequest) (*dto.UserAuthResponse, error) {
	user, err := authUserService.authUserRepo.GetUserByUsername(loginRequest.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("username or password wrong")
	}

	if err := passwordhash.CheckPasswordHash(loginRequest.Password, user.PasswordHash); err != nil {
		return nil, errors.New("username or password wrong")
	}

	accessToken, err := jwttoken.GenerateJwtToken(user.Id, user.Username)
	if err != nil {
		return nil, err
	}

	loginResponse := dto.NewUserAuthResponse(user, accessToken)
	return loginResponse, nil
}
