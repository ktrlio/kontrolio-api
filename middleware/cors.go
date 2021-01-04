package middleware

import (
	"github.com/marcelovicentegc/kontrolio-api/config"
	"github.com/marcelovicentegc/kontrolio-api/utils"
)

func EnableCors(response utils.Response) utils.Response {
	response.Headers["Access-Control-Allow-Origin"] = config.CLIENT_URL
	response.Headers["Access-Control-Allow-Credentials"] = "true"
	return response
}
