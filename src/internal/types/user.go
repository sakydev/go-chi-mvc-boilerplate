package types

type User struct {
	Username string `query:"username"`
	Email    string `query:"email"`
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type UpdateUsernameRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}
