package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-crud/internal/domain/user/service"

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
	panic("user handler imp")
}

func (userHandler userHandlerImp) Update(ctx *gin.Context) {
	panic("user handler imp")
}

func (userHandler userHandlerImp) UpdatePassword(ctx *gin.Context) {
	panic("user handler imp")
}

func (userHandler userHandlerImp) Delete(ctx *gin.Context) {
	panic("user handler imp")
}
