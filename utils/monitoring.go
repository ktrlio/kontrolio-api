package utils

import (
	"encoding/json"
	"net/http"
)

func HealthCheck(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(string("ok"))

	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Write(response)
}
