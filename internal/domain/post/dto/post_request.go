package dto

type CreatePostRequest struct {
	PostTitle string `form:"post_title" binding:"required"`
	PostDesc  string `form:"post_desc" binding:"required" `
	PostImage string `form:"post_image"  binding:"-"`
}

type UpdatePostRequest struct {
	PostTitle string `form:"post_title" binding:"required" validate:"required,min=3,max=100"`
	PostDesc  string `form:"post_desc" binding:"required" validate:"required,min=10,max=500"`
	PostImage string `form:"post_image"`
}
