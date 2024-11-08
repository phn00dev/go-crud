package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-crud/internal/domain/post/service"

)

type postHandlerImp struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) PostHandler {
	return postHandlerImp{
		postService: postService,
	}
}

func (postHandler postHandlerImp) GetAll(ctx *gin.Context) {
	panic("post handler imp")
}

func (postHandler postHandlerImp) GetOne(ctx *gin.Context) {
	panic("post handler imp")
}

func (postHandler postHandlerImp) Create(ctx *gin.Context) {
	panic("post handler imp")
}

func (postHandler postHandlerImp) Update(ctx *gin.Context) {
	panic("post handler imp")
}

func (postHandler postHandlerImp) Delete(ctx *gin.Context) {
	panic("post handler imp")
}
