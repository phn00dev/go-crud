package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-crud/internal/domain/user/dto"
	"github.com/phn00dev/go-crud/internal/domain/user/service"
	bindandvalidate "github.com/phn00dev/go-crud/internal/utils/bind_and_validate"
	"github.com/phn00dev/go-crud/internal/utils/response"
)

type authUserHandlerImp struct {
	authUserService service.AuthUserService
}

func NewAuthUserHandler(authUserService service.AuthUserService) AuthUserHandler {
	return authUserHandlerImp{
		authUserService: authUserService,
	}
}

func (authUserHandler authUserHandlerImp) Register(ctx *gin.Context) {
	var registerRequest dto.RegisterRequest
	// bind and validate
	if !bindandvalidate.BindAndValidateRequest(ctx, &registerRequest) {
		return
	}
	// register user
	registerResponse, err := authUserHandler.authUserService.RegisterUser(registerRequest)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "something wrong", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "user registered successfully", registerResponse)
}

func (authUserHandler authUserHandlerImp) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	// bind and validate
	if !bindandvalidate.BindAndValidateRequest(ctx, &loginRequest) {
		return
	}
	// login user
	loginResponse, err := authUserHandler.authUserService.LoginUser(loginRequest)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "something wrong", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "user login successfully", loginResponse)
}
