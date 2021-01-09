package middleware

import "github.com/aws/aws-lambda-go/events"

func UseMiddleware(f func(request events.APIGatewayCustomAuthorizerRequest)) func(events.APIGatewayCustomAuthorizerRequest) {
	return func(request events.APIGatewayCustomAuthorizerRequest) {
		Authenticate(request)
		f(request)
	}
}
