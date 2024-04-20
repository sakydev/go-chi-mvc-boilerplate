package utils

import (
	"encoding/json"
	"net/http"
)

const defaultContentType = "application/json"

func RenderResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", defaultContentType)
	w.WriteHeader(status)

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func RenderNotFoundResponse(w http.ResponseWriter) {
	RenderResponse(w, nil, http.StatusNotFound)
}

func RenderBadRequestResponse(w http.ResponseWriter) {
	RenderResponse(w, nil, http.StatusBadRequest)
}
