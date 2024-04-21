package repositories

import (
	"context"
	"github.com/samber/do"
	"go-chi-mvc-boilerplate/src/internal/database"
	"go-chi-mvc-boilerplate/src/internal/types"
)

func InjectUserRepository(i *do.Injector) (UserRepository, error) {
	return UserImpl{
		db: do.MustInvoke[database.Database](i),
	}, nil
}

type UserImpl struct {
	db database.Database
}

type UserRepository interface {
	List(ctx context.Context) ([]types.User, error)
	GetByEmail(ctx context.Context, email string) (types.User, error)
}

func (impl UserImpl) List(ctx context.Context) ([]types.User, error) {
	var users []types.User

	err := impl.db.Select(ctx, &users, "SELECT username, email FROM users")

	return users, err
}

func (impl UserImpl) GetByEmail(ctx context.Context, email string) (types.User, error) {
	var user types.User

	err := impl.db.Get(ctx, &user, "SELECT username, email FROM users WHERE email = $1", email)
	if err != nil {
		return types.User{}, err
	}

	return user, nil
}
