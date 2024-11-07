package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-crud/internal/domain/user/dto"
	"github.com/phn00dev/go-crud/internal/domain/user/service"
	"github.com/phn00dev/go-crud/internal/utils/response"
	"github.com/phn00dev/go-crud/internal/utils/validate"

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
	// body parser
	if err := ctx.ShouldBindBodyWithJSON(&registerRequest); err != nil {
		response.Error(ctx, http.StatusBadRequest, "body parser error", err.Error())
		return
	}
	// validate
	if err := validate.ValidateStruct(registerRequest); err != nil {
		response.Error(ctx, http.StatusBadRequest, "validate error", err.Error())
		return
	}

	// register user
	registerResponse, err := authUserHandler.authUserService.RegisterUser(registerRequest)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "something wrong", err.Error())
		return
	}
	response.Success(ctx, http.StatusCreated, "user registered successfully", registerResponse)

}

func (authUserHandler authUserHandlerImp) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	// body parser
	if err := ctx.ShouldBindBodyWithJSON(&loginRequest); err != nil {
		response.Error(ctx, http.StatusBadRequest, "body parser error", err.Error())
		return
	}
	// validate
	if err := validate.ValidateStruct(loginRequest); err != nil {
		response.Error(ctx, http.StatusBadRequest, "validate error", err.Error())
		return
	}
	//login user
	loginResponse, err := authUserHandler.authUserService.LoginUser(loginRequest)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "something wrong", err.Error())
		return
	}
	response.Success(ctx, http.StatusCreated, "user login successfully", loginResponse)
}
