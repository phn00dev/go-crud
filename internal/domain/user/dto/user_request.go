package dto

type UpdateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
}

type UpdateUserPasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required"`
	Password        string `json:"password" validate:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
