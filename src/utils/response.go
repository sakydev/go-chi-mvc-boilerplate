package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

const defaultContentType = "application/json"

type Response struct {
	Error   string      `json:"error,omitempty"`
	Content interface{} `json:"content"`
}

func RenderResponse(responseWriter http.ResponseWriter, content interface{}, status int, error error) {
	responseWriter.Header().Set("Content-Type", defaultContentType)
	responseWriter.WriteHeader(status)

	errorMessage := ""
	if error != nil {
		errorMessage = error.Error()
	}

	jsonData, err := json.Marshal(Response{
		Error:   errorMessage,
		Content: content,
	})
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Write(jsonData)
}

func RenderNotFoundResponse(responseWriter http.ResponseWriter, message string) {
	RenderResponse(responseWriter, nil, http.StatusNotFound, errors.New(message))
}

func RenderBadRequestResponse(responseWriter http.ResponseWriter, message string) {
	RenderResponse(responseWriter, nil, http.StatusBadRequest, errors.New(message))
}

func RenderInternalServerResponse(responseWriter http.ResponseWriter, message string) {
	RenderResponse(responseWriter, nil, http.StatusInternalServerError, errors.New(message))
}
