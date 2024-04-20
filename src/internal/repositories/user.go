package repositories

import (
	"context"
	"database/sql"
	"github.com/samber/do"
	"go-chi-mvc-boilerplate/src/internal/types"
)

func InjectUserRepository(i *do.Injector) (UserRepository, error) {
	return UserImpl{}, nil
}

type UserImpl struct{}

type UserRepository interface {
	List(ctx context.Context, database *sql.DB) ([]types.User, error)
}

func (impl UserImpl) List(ctx context.Context, database *sql.DB) ([]types.User, error) {
	var users []types.User

	rows, err := database.QueryContext(ctx, `
		SELECT username, email
		FROM users
	`)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var username, email string
		err = rows.Scan(&username, &email)
		if err != nil {
			return users, err
		}

		user := types.User{Name: username, Email: email}
		users = append(users, user)
	}

	return users, nil
}
