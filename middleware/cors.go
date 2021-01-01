package middleware

import (
	"net/http"
)

func EnableCors(responseWriter *http.ResponseWriter) {
	(*responseWriter).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*responseWriter).Header().Set("Access-Control-Allow-Credentials", "true")
}
