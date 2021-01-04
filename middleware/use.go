package middleware

import (
	"github.com/marcelovicentegc/kontrolio-api/utils"
)

func UseMiddleware(f func(request utils.AuthRequest)) func(utils.AuthRequest) {
	return func(request utils.AuthRequest) {
		Authenticate(request)
		f(request)
	}
}
