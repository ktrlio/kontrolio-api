package controllers

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/marcelovicentegc/kontrolio-api/utils"
)

type Record struct {
	Time       string `json:"time"`
	RecordType string `json:"recordType"`
	ApiKey     string `json:"apiKey"`
}

func CreateRecord(ctx context.Context, record Record) (*events.APIGatewayProxyResponse, error) {
	return utils.ApiResponse(http.StatusOK, "Function executed successfully")
}
