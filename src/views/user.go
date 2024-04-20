package views

import (
	"go-chi-mvc-boilerplate/src/internal/types"
)

func NewUserResponse(user types.User) types.UserResponse {
	return types.UserResponse{
		Name:  user.Name,
		Email: user.Email,
	}
}

func NewUserListResponse(users []types.User) []types.UserResponse {
	var response []types.UserResponse

	for _, user := range users {
		currentUserResponse := NewUserResponse(user)
		response = append(response, currentUserResponse)
	}

	return response
}
