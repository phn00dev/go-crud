package service

import (
	"errors"

	"github.com/phn00dev/go-crud/internal/domain/user/dto"
	"github.com/phn00dev/go-crud/internal/domain/user/repository"
	"github.com/phn00dev/go-crud/internal/models"
	passwordhash "github.com/phn00dev/go-crud/internal/utils/password_hash"
)

type userServiceImp struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userServiceImp{
		userRepo: repo,
	}
}

func (userService userServiceImp) GetUser(userId int) (*models.User, error) {

	// get user
	user, err := userService.userRepo.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userService userServiceImp) Update(userId int, updateRequest dto.UpdateUserRequest) error {
	if userId == 0 {
		return errors.New("something wrong")
	}
	user, err := userService.userRepo.GetUser(userId)
	if err != nil {
		return err
	}
	user.Username = updateRequest.Username
	return userService.userRepo.Update(user.Id, *user)
}

func (userService userServiceImp) UpdatePassword(userId int, request dto.UpdateUserPasswordRequest) error {
	if userId == 0 {
		return errors.New("something wrong")
	}
	user, err := userService.userRepo.GetUser(userId)
	if err != nil {
		return err
	}
	if err := passwordhash.CheckPasswordHash(request.OldPassword, user.PasswordHash); err != nil {
		return errors.New("old password wrong")
	}
	hashPassword := passwordhash.GeneratePassword(request.Password)
	// update password
	return userService.userRepo.UpdatePassword(user.Id, hashPassword)
}

func (userService userServiceImp) Delete(userId int) error {
	// get user
	user, err := userService.userRepo.GetUser(userId)
	if err != nil {
		return err
	}
	return userService.userRepo.Delete(user.Id)
}
