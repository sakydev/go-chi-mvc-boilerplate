package internal

import (
	"github.com/samber/do"
	"go-chi-mvc-boilerplate/src/internal/database"
	"go-chi-mvc-boilerplate/src/internal/repositories"
	"go-chi-mvc-boilerplate/src/internal/services"
)

func WireDependencies(i *do.Injector) {
	do.Provide(i, database.InjectDatabaseService)

	repositories.Wire(i)
	services.Wire(i)
}
