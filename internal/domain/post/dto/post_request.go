package dto

type CreatePostRequest struct {
	PostTitle string `json:"post_title" validate:"required,min=3,max=100"`
	PostDesc  string `json:"post_desc" validate:"required,min=10,max=500"`
	PostImage string `json:"post_image" validate:"required,url"`
}

type UpdatePostRequest struct {
	PostTitle string `json:"post_title" validate:"required,min=3,max=100"`
	PostDesc  string `json:"post_desc" validate:"required,min=10,max=500"`
	PostImage string `json:"post_image" validate:"required,url"`
}
