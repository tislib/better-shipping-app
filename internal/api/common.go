package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

// handleHttpError is a helper function to handle http errors
func handleHttpError(writer http.ResponseWriter, statusCode int, message string) {
	writer.WriteHeader(statusCode)

	// write the error message
	var apiError = Error{
		Message: message,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	if err := json.NewEncoder(writer).Encode(apiError); err != nil {
		log.Error(err)
	}
}

// respondJsonBody is a helper function to respond with a json body
func respondJsonBody(writer http.ResponseWriter, body interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)

	if err := json.NewEncoder(writer).Encode(body); err != nil {
		log.Error(err)
	}
}
