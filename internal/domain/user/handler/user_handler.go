package handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUser(ctx *gin.Context)
	Update(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
