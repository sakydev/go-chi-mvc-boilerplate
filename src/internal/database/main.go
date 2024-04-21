package database

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/samber/do"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func InjectDatabaseService(injector *do.Injector) (Database, error) {
	accessor := &DatabaseImpl{
		Connection: nil,
	}
	err := accessor.Connect()
	return accessor, err
}

type Database interface {
	Connect() error
	GetConnection() *pgx.Conn
	Get(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, dst interface{}, query string, args ...interface{}) error
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) pgx.Row
	Exec(ctx context.Context, query string, arguments ...any) (pgconn.CommandTag, error)
}

type DatabaseImpl struct {
	Connection *pgx.Conn
}

func (d *DatabaseImpl) Connect() error {
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		return fmt.Errorf("error converting DATABASE_PORT to integer: %w", err)
	}

	databaseOptions := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s port=%d",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_SSL_MODE"),
		port,
	)
	conn, err := pgx.Connect(context.Background(), databaseOptions)
	if err != nil {
		return fmt.Errorf("error connecting to the database: %w", err)
	}
	d.Connection = conn
	return nil
}

func (d *DatabaseImpl) GetConnection() *pgx.Conn {
	return d.Connection
}

func (d *DatabaseImpl) Get(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	err := pgxscan.Get(ctx, d.GetConnection(), dst, query, args...)

	return err
}

func (d *DatabaseImpl) Select(ctx context.Context, dst interface{}, query string, args ...interface{}) error {
	err := pgxscan.Select(ctx, d.GetConnection(), dst, query, args...)

	return err
}

func (d *DatabaseImpl) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return d.Connection.Query(ctx, query, args...)
}

func (d *DatabaseImpl) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	return d.Connection.QueryRow(ctx, query, args...)
}

func (d *DatabaseImpl) Exec(ctx context.Context, query string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return d.Connection.Exec(ctx, query, arguments...)
}
