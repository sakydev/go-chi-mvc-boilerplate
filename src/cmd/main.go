package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/samber/do"
	"go-chi-mvc-boilerplate/src/internal"
	"go-chi-mvc-boilerplate/src/internal/config"
	"go-chi-mvc-boilerplate/src/internal/routes"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	injector := do.DefaultInjector

	database, err := config.GetDatabase()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	setup(injector, ctx, database)
	startServer(injector, ctx, database)
}

func setup(injector *do.Injector, ctx context.Context, database *sql.DB) {
	internal.WireDependencies(injector)
}

func startServer(injector *do.Injector, ctx context.Context, database *sql.DB) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/users", routes.UserRoutes(injector, ctx, database))

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)

		return
	}
}
