package handler

import (
	"github.com/gin-gonic/gin"
)

type AuthUserHandler interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}
