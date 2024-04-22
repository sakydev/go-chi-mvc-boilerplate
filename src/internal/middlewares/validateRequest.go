package middlewares

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go-chi-mvc-boilerplate/src/utils"
	"net/http"
)

type contextKey string

const ValidatedRequestKey contextKey = "validatedRequest"

func ValidateRequest[T any](next http.Handler, responseWriter http.ResponseWriter, request *http.Request, requestType T) {
	err := json.NewDecoder(request.Body).Decode(&requestType)
	if err != nil {
		utils.RenderBadRequestResponse(responseWriter, err.Error())

		return
	}
	defer request.Body.Close()

	validate := validator.New()
	err = validate.Struct(requestType)
	if err != nil {
		utils.RenderBadRequestResponse(responseWriter, err.Error())

		return
	}

	ctx := context.WithValue(request.Context(), ValidatedRequestKey, requestType)

	next.ServeHTTP(responseWriter, request.WithContext(ctx))
}
