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

func RenderResponse(w http.ResponseWriter, content interface{}, status int, error error) {
	w.Header().Set("Content-Type", defaultContentType)
	w.WriteHeader(status)

	errorMessage := ""
	if error != nil {
		errorMessage = error.Error()
	}

	jsonData, err := json.Marshal(Response{
		Error:   errorMessage,
		Content: content,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func RenderNotFoundResponse(w http.ResponseWriter, message string) {
	RenderResponse(w, nil, http.StatusNotFound, errors.New(message))
}

func RenderBadRequestResponse(w http.ResponseWriter, message string) {
	RenderResponse(w, nil, http.StatusBadRequest, errors.New(message))
}

func RenderInternalServerResponse(w http.ResponseWriter, message string) {
	RenderResponse(w, nil, http.StatusInternalServerError, errors.New(message))
}
