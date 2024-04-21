package routes

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
	"go-chi-mvc-boilerplate/src/internal/controllers"
)

func UserRoutes(i *do.Injector, ctx context.Context) chi.Router {
	r := chi.NewRouter()
	controller := controllers.NewUserController(i, ctx)

	r.Get("/", controller.ListUsers)
	r.Get("/{email}", controller.GetUser)

	return r
}
