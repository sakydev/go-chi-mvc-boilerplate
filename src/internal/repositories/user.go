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
	GetByEmail(ctx context.Context, database *sql.DB, email string) (types.User, error)
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

func (impl UserImpl) GetByEmail(ctx context.Context, database *sql.DB, email string) (types.User, error) {
	var user types.User
	var username, userEmail string

	err := database.QueryRowContext(ctx, `
		SELECT username, email
		FROM users WHERE email = $1
	`, email).Scan(&username, &userEmail)
	if err != nil {
		return user, err
	}

	user.Name = username
	user.Email = email

	return user, nil
}
