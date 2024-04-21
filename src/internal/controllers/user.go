package controllers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
	"go-chi-mvc-boilerplate/src/utils"
	"net/http"

	"go-chi-mvc-boilerplate/src/internal/services"
)

type UserController struct {
	ctx         context.Context
	userService *services.UserService
}

func NewUserController(i *do.Injector, ctx context.Context) *UserController {
	return &UserController{
		ctx:         ctx,
		userService: do.MustInvoke[*services.UserService](i),
	}
}

func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.userService.List(uc.ctx)
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

	users, err := uc.userService.GetByEmail(uc.ctx, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	utils.RenderResponse(w, users, http.StatusOK)
}
