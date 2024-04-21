package middlewares

import (
	"context"
	"encoding/json"
	"go-chi-mvc-boilerplate/src/internal/types"
	"go-chi-mvc-boilerplate/src/utils"
	"net/http"
)

func ValidateCreateUserRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestBody types.CreateUserRequest

		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			utils.RenderBadRequestResponse(w, err.Error())

			return
		}
		defer r.Body.Close()

		err = requestBody.Validate()
		if err != nil {
			utils.RenderBadRequestResponse(w, err.Error())

			return
		}

		ctx := context.WithValue(r.Context(), "validated", requestBody)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
