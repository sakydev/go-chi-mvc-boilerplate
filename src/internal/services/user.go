package services

import (
	"context"
	"go-chi-mvc-boilerplate/src/internal/repositories"
	"go-chi-mvc-boilerplate/src/internal/types"
	"go-chi-mvc-boilerplate/src/views"

	"github.com/samber/do"
)

func InjectUserService(i *do.Injector) (*UserService, error) {
	return &UserService{
		userRepository: do.MustInvoke[repositories.UserRepository](i),
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
