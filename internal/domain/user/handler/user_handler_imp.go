package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-crud/internal/domain/user/dto"
	"github.com/phn00dev/go-crud/internal/domain/user/service"
	bindandvalidate "github.com/phn00dev/go-crud/internal/utils/bind_and_validate"
	"github.com/phn00dev/go-crud/internal/utils/response"

)

type userHandlerImp struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return userHandlerImp{
		userService: service,
	}
}

func (userHandler userHandlerImp) GetUser(ctx *gin.Context) {
	userId, exists := ctx.Get("id")
	if !exists {
		response.Error(ctx, http.StatusBadRequest, "User ID is required", "User ID is required")
		return
	}
	username, exists := ctx.Get("username")
	if !exists {
		response.Error(ctx, http.StatusBadRequest, "Username is required", "User ID is required")
		return
	}
	user, err := userHandler.userService.GetUser(userId.(int), username.(string))
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Error fetching user", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "User data", user)
}

func (userHandler userHandlerImp) Update(ctx *gin.Context) {
	userId, exists := ctx.Get("id")
	if !exists {
		response.Error(ctx, http.StatusBadRequest, "User ID is required", "User ID is required")
		return
	}

	var updateRequest dto.UpdateUserRequest
	if !bindandvalidate.BindAndValidateRequest(ctx, &updateRequest) {
		return
	}
	// update user data
	if err := userHandler.userService.Update(userId.(int), updateRequest); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "update error", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "update successfully", nil)
}

func (userHandler userHandlerImp) UpdatePassword(ctx *gin.Context) {
	userId, exists := ctx.Get("id")
	if !exists {
		response.Error(ctx, http.StatusBadRequest, "User ID is required", "User ID is required")
		return
	}

	var updatePasswordRequest dto.UpdateUserPasswordRequest
	if !bindandvalidate.BindAndValidateRequest(ctx, &updatePasswordRequest) {
		return
	}
	// update user data
	if err := userHandler.userService.UpdatePassword(userId.(int), updatePasswordRequest); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "update password error", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "update  password successfully", nil)
}

func (userHandler userHandlerImp) Delete(ctx *gin.Context) {
	userId, exists := ctx.Get("id")
	if !exists {
		response.Error(ctx, http.StatusBadRequest, "User ID is required", "User ID is required")
		return
	}
	username, exists := ctx.Get("username")
	if !exists {
		response.Error(ctx, http.StatusBadRequest, "Username is required", "User ID is required")
		return
	}

	if err := userHandler.userService.Delete(userId.(int), username.(string)); err != nil {
		response.Error(ctx, http.StatusBadRequest, "delete error", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "delete successfully", nil)
}
