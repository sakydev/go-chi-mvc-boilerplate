package types

import "github.com/go-playground/validator/v10"

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

func (createUserRequest *CreateUserRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(createUserRequest)
	if err != nil {
		return err
	}
	
	return nil
}
