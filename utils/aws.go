package utils

import "github.com/aws/aws-lambda-go/events"

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse
type AuthResponse events.APIGatewayCustomAuthorizerResponse

type Request events.APIGatewayProxyRequest
type AuthRequest events.APIGatewayCustomAuthorizerRequest
