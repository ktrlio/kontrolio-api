package middleware

import (
	"net/http"
)

func UseMiddleware(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		EnableCors(&responseWriter)
		Authenticate(&responseWriter, request)
		f(responseWriter, request)
	}
}
