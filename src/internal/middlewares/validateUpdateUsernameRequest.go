package middlewares

import (
	"go-chi-mvc-boilerplate/src/internal/types"
	"net/http"
)

func ValidateUpdateUsernameRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		ValidateRequest(next, responseWriter, request, types.UpdateUsernameRequest{})
	})
}
