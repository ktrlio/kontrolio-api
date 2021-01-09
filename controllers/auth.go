package controllers

import (
	"bytes"
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/marcelovicentegc/kontrolio-api/utils"
)

func Login(ctx context.Context, data User) (events.APIGatewayProxyResponse, error) {
	var buf bytes.Buffer

	user := utils.GetUser(data.Email)

	if user == nil {
		return events.APIGatewayProxyResponse{StatusCode: 200}, errors.New("User not found or incorrect password")
	}

	match := utils.CheckPasswordHash(data.Password, user.Password)

	if !match {
		return events.APIGatewayProxyResponse{StatusCode: 200}, errors.New("User not found or incorrect password")
	}

	// Generate and save session ID

	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
