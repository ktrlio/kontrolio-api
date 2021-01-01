package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/marcelovicentegc/kontrolio-api/utils"
)

func Authenticate(responseWriter *http.ResponseWriter, request *http.Request) (http.ResponseWriter, http.Request) {
	(*responseWriter).Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

	token := strings.SplitN(request.Header.Get("Authorization"), " ", 2)

	if len(token) != 2 {
		http.Error((*responseWriter), "Not authorized", 401)
		return (*responseWriter), *request
	}

	pair := strings.SplitN(token[1], ":", 2)
	if len(pair) != 2 {
		http.Error((*responseWriter), "Not authorized", 401)
		return (*responseWriter), *request
	}

	authKey := os.Getenv("AUTH_TOKEN_KEY")
	authValue := os.Getenv("AUTH_TOKEN_VALUE")

	if !utils.CheckPasswordHash(authKey, pair[0]) || !utils.CheckPasswordHash(authValue, pair[1]) {
		http.Error((*responseWriter), "Not authorized", 401)
		return (*responseWriter), *request
	}

	return (*responseWriter), *request
}
