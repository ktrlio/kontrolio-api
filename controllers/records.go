package controllers

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Record struct {
	Time       string `json:"time"`
	RecordType string `json:"recordType"`
	ApiKey     string `json:"apiKey"`
}

func CreateRecord(ctx context.Context, record Record) (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusOK, "Function executed successfully")
}
