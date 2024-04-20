package controllers

import (
	"context"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
	"go-chi-mvc-boilerplate/src/utils"
	"net/http"

	"go-chi-mvc-boilerplate/src/internal/services"
)

type UserController struct {
	ctx         context.Context
	database    *sql.DB
	userService *services.UserService
}

func NewUserController(i *do.Injector, ctx context.Context, database *sql.DB) *UserController {
	return &UserController{
		ctx:         ctx,
		database:    database,
		userService: do.MustInvoke[*services.UserService](i),
	}
}

func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.userService.List(uc.ctx, uc.database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	utils.RenderResponse(w, users, http.StatusOK)
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	if len(email) < 2 {
		utils.RenderNotFoundResponse(w)

		return
	}

	users, err := uc.userService.GetByEmail(uc.ctx, uc.database, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	utils.RenderResponse(w, users, http.StatusOK)
}
