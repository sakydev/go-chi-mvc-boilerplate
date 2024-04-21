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

func NewUserController(injector *do.Injector, ctx context.Context) *UserController {
	return &UserController{
		userService: do.MustInvoke[*services.UserService](injector),
	}
}

func (uc *UserController) ListUsers(responseWriter http.ResponseWriter, request *http.Request) {
	users, err := uc.userService.List(request.Context())
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}

	utils.RenderResponse(responseWriter, users, http.StatusOK, nil)
}

func (uc *UserController) GetUser(responseWriter http.ResponseWriter, request *http.Request) {
	email := chi.URLParam(request, "email")
	if len(email) < 2 {
		utils.RenderNotFoundResponse(responseWriter, "user not found")

		return
	}

	users, err := uc.userService.GetByEmail(request.Context(), email)
	if err != nil {
		utils.RenderInternalServerResponse(responseWriter, fmt.Errorf("error fetching user: %w", err).Error())
	}

	utils.RenderResponse(responseWriter, users, http.StatusOK, nil)
}

func (uc *UserController) CreateUser(responseWriter http.ResponseWriter, request *http.Request) {
	requestContent := request.Context().Value("validated").(types.CreateUserRequest)

	createdUser, err := uc.userService.Create(request.Context(), requestContent)
	utils.RenderResponse(responseWriter, createdUser, http.StatusOK, err)
}
