package controllers

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/samber/do"
	"go-chi-mvc-boilerplate/src/internal/types"
	"go-chi-mvc-boilerplate/src/utils"
	"net/http"

	"go-chi-mvc-boilerplate/src/internal/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(i *do.Injector, ctx context.Context) *UserController {
	return &UserController{
		userService: do.MustInvoke[*services.UserService](i),
	}
}

func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.userService.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	utils.RenderResponse(w, users, http.StatusOK, nil)
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	if len(email) < 2 {
		utils.RenderNotFoundResponse(w, "user not found")

		return
	}

	users, err := uc.userService.GetByEmail(r.Context(), email)
	if err != nil {
		utils.RenderInternalServerResponse(w, fmt.Sprintf("error fetching user: %w", err))
	}

	utils.RenderResponse(w, users, http.StatusOK, nil)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestContent := r.Context().Value("validated").(types.CreateUserRequest)

	createdUser, err := uc.userService.Create(r.Context(), requestContent)
	utils.RenderResponse(w, createdUser, http.StatusOK, err)
}
