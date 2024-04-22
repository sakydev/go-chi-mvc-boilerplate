package repositories

import (
	"context"
	"github.com/samber/do"
	"go-chi-mvc-boilerplate/src/internal/database"
	"go-chi-mvc-boilerplate/src/internal/types"
)

func InjectUserRepository(injector *do.Injector) (UserRepository, error) {
	return UserImpl{
		db: do.MustInvoke[database.Database](injector),
	}, nil
}

type UserImpl struct {
	db database.Database
}

type UserRepository interface {
	List(ctx context.Context) ([]types.User, error)
	GetById(ctx context.Context, userId int64) (types.User, error)
	GetByEmail(ctx context.Context, email string) (types.User, error)
	Create(ctx context.Context, requestContent types.CreateUserRequest) (int64, error)
}

func (repo UserImpl) List(ctx context.Context) ([]types.User, error) {
	var users []types.User
	err := repo.db.Select(ctx, &users, "SELECT username, email FROM users")

	return users, err
}

func (repo UserImpl) GetById(ctx context.Context, userId int64) (types.User, error) {
	var user types.User
	err := repo.db.Get(ctx, &user, "SELECT username, email FROM users WHERE id = $1", userId)

	return user, err
}

func (repo UserImpl) GetByEmail(ctx context.Context, email string) (types.User, error) {
	var user types.User
	err := repo.db.Get(ctx, &user, "SELECT username, email FROM users WHERE email = $1", email)

	return user, err
}

func (repo UserImpl) Create(ctx context.Context, requestContent types.CreateUserRequest) (int64, error) {
	var userId int64
	err := repo.db.QueryRow(ctx, `
		INSERT INTO users (username, email) VALUES ($1, $2)
		RETURNING id
	`, requestContent.Username, requestContent.Email).Scan(&userId)
	if err != nil {
		return userId, err
	}

	return userId, nil
}
