package services

import (
	"context"
	"database/sql"
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

func (s UserService) List(ctx context.Context, database *sql.DB) ([]types.UserResponse, error) {
	users, err := s.userRepository.List(ctx, database)
	if err != nil {
		return nil, err
	}

	return views.NewUserListResponse(users), nil
}
