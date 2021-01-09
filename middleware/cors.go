package middleware

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/marcelovicentegc/kontrolio-api/config"
)

func EnableCors(response events.APIGatewayProxyRequest) events.APIGatewayProxyRequest {
	response.Headers["Access-Control-Allow-Origin"] = config.CLIENT_URL
	response.Headers["Access-Control-Allow-Credentials"] = "true"
	return response
}
