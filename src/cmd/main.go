package main

import (
	"context"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/samber/do"
	"go-chi-mvc-boilerplate/src/internal"
	"go-chi-mvc-boilerplate/src/internal/config"
	"go-chi-mvc-boilerplate/src/internal/routes"
	"net/http"
)

func main() {
	ctx := context.Background()
	injector := do.DefaultInjector
	database, err := config.GetDatabase()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	setup(injector, ctx, database)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/users", routes.UserRoutes(injector, ctx, database))

	http.ListenAndServe(":3000", r)
}

func setup(injector *do.Injector, ctx context.Context, database *sql.DB) {
	internal.WireDependencies(injector)
}
