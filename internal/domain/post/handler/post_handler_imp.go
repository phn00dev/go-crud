package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-crud/internal/domain/post/dto"
	"github.com/phn00dev/go-crud/internal/domain/post/service"
	bindandvalidate "github.com/phn00dev/go-crud/internal/utils/bind_and_validate"
	"github.com/phn00dev/go-crud/internal/utils/response"

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
	userId, exists := ctx.Get("user_id")
	if !exists {
		response.Error(ctx, http.StatusBadRequest, "User ID is required", "User ID is required")
		return
	}
	posts, err := postHandler.postService.GetAll(userId.(int))
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "something wrong", err.Error())
		return
	}
	postResponses := dto.NewPostResponses(posts)
	// success
	response.Success(ctx, http.StatusOK, "user's all posts", postResponses)
}

func (postHandler postHandlerImp) GetOne(ctx *gin.Context) {
	userId, exists := ctx.Get("user_id")
	if !exists {
		response.Error(ctx, http.StatusBadRequest, "User ID is required", "User ID is required")
		return
	}
	postIdStr := ctx.Param("postId")
	postId, _ := strconv.Atoi(postIdStr)

	post, err := postHandler.postService.GetOne(userId.(int), postId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "something wrong", err.Error())
		return
	}
	postResponse := dto.NewPostResponse(post)
	response.Success(ctx, http.StatusOK, "post", postResponse)
}

func (postHandler postHandlerImp) Create(ctx *gin.Context) {

	userId, exists := ctx.Get("user_id")
	if !exists {
		response.Error(ctx, http.StatusBadRequest, "User ID is required", "User ID is required")
		return
	}
	var createRequest dto.CreatePostRequest
	// binding and validate
	if !bindandvalidate.BindAndValidateRequestofFormData(ctx, &createRequest) {
		return
	}
	log.Println("create_request : ", createRequest)
	// create post
	if err := postHandler.postService.CreatePost(ctx, userId.(int), createRequest); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "create post error", err.Error())
		return
	}
	response.Success(ctx, http.StatusCreated, "post created successfully", nil)
}

func (postHandler postHandlerImp) Update(ctx *gin.Context) {
	userId, exists := ctx.Get("user_id")
	if !exists {
		response.Error(ctx, http.StatusBadRequest, "User ID is required", "User ID is required")
		return
	}
	postIdStr := ctx.Param("postId")
	postId, _ := strconv.Atoi(postIdStr)

	var updateRequest dto.UpdatePostRequest

	// binding and validate
	if !bindandvalidate.BindAndValidateRequestofFormData(ctx, &updateRequest) {
		return
	}
	// update post
	if err := postHandler.postService.UpdatePost(ctx, userId.(int), postId, updateRequest); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "post update error", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "post updated successfully", nil)
}

func (postHandler postHandlerImp) Delete(ctx *gin.Context) {
	userId, exists := ctx.Get("user_id")
	if !exists {
		response.Error(ctx, http.StatusBadRequest, "User ID is required", "User ID is required")
		return
	}
	postIdStr := ctx.Param("postId")
	postId, _ := strconv.Atoi(postIdStr)
	if err := postHandler.postService.DeletePost(userId.(int), postId); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "post delete error", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "post deleted successfully", nil)
}

func (postHandler postHandlerImp) GetAllPost(ctx *gin.Context) {
	posts, err := postHandler.postService.GetAllPost()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "post find error", err.Error())
		return
	}
	postResponses := dto.NewPostResponses(posts)
	response.Success(ctx, http.StatusOK, "all posts", postResponses)
}

func (postHandler postHandlerImp) GetPostBySlug(ctx *gin.Context) {
	postSlug := ctx.Param("post_slug")
	post, err := postHandler.postService.GetPostBySlug(postSlug)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "post find error", err.Error())
		return
	}
	postResponse := dto.NewPostResponse(post)
	response.Success(ctx, http.StatusOK, "post", postResponse)
}
