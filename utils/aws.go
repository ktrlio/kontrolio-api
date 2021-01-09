package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func ApiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil
}
