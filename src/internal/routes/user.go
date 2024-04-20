package routes

import (
	"context"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
	"go-chi-mvc-boilerplate/src/internal/controllers"
)

func UserRoutes(i *do.Injector, ctx context.Context, database *sql.DB) chi.Router {
	r := chi.NewRouter()
	controller := controllers.NewUserController(i, ctx, database)

	r.Get("/", controller.ListUsers)

	return r
}
