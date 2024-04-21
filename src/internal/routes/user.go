package routes

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
	"go-chi-mvc-boilerplate/src/internal/controllers"
	"go-chi-mvc-boilerplate/src/internal/middlewares"
)

func UserRoutes(injector *do.Injector, ctx context.Context) chi.Router {
	r := chi.NewRouter()
	controller := controllers.NewUserController(injector, ctx)

	r.Get("/", controller.ListUsers)
	r.Get("/{email}", controller.GetUser)
	r.With(middlewares.ValidateCreateUserRequest).Post("/", controller.CreateUser)

	return r
}
