package services

import (
	"context"
	"go-chi-mvc-boilerplate/src/internal/repositories"
	"go-chi-mvc-boilerplate/src/internal/types"
	"go-chi-mvc-boilerplate/src/views"

	"github.com/samber/do"
)

func InjectUserService(injector *do.Injector) (*UserService, error) {
	return &UserService{
		userRepository: do.MustInvoke[repositories.UserRepository](injector),
	}, nil
}

type UserService struct {
	userRepository repositories.UserRepository
}

func (s UserService) List(ctx context.Context) ([]types.UserResponse, error) {
	users, err := s.userRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	return views.NewUserListResponse(users), nil
}

func (s UserService) GetByEmail(ctx context.Context, email string) (types.UserResponse, error) {
	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return types.UserResponse{}, err
	}

	return views.NewUserResponse(user), nil
}

func (s UserService) Create(ctx context.Context, requestContent types.CreateUserRequest) (types.UserResponse, error) {
	created, err := s.userRepository.Create(ctx, requestContent)
	if err != nil {
		return types.UserResponse{}, err
	}

	user, err := s.userRepository.GetById(ctx, created)
	if err != nil {
		return types.UserResponse{}, err
	}

	return views.NewUserResponse(user), nil
}

func (s UserService) UpdateUsername(ctx context.Context, requestContent types.UpdateUsernameRequest) (types.UserResponse, error) {
	err := s.userRepository.UpdateUsername(ctx, requestContent)
	if err != nil {
		return types.UserResponse{}, err
	}

	user, err := s.userRepository.GetByEmail(ctx, requestContent.Email)
	if err != nil {
		return types.UserResponse{}, err
	}

	return views.NewUserResponse(user), nil
}
