package middlewares

import (
	"context"
	"encoding/json"
	"go-chi-mvc-boilerplate/src/internal/types"
	"go-chi-mvc-boilerplate/src/utils"
	"net/http"
)

func ValidateCreateUserRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		var requestBody types.CreateUserRequest

		err := json.NewDecoder(request.Body).Decode(&requestBody)
		if err != nil {
			utils.RenderBadRequestResponse(responseWriter, err.Error())

			return
		}
		defer request.Body.Close()

		err = requestBody.Validate()
		if err != nil {
			utils.RenderBadRequestResponse(responseWriter, err.Error())

			return
		}

		ctx := context.WithValue(request.Context(), "validated", requestBody)

		next.ServeHTTP(responseWriter, request.WithContext(ctx))
	})
}
