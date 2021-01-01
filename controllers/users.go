package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/marcelovicentegc/kontrolio-api/utils"
	uuid "github.com/satori/go.uuid"
)

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(ctx context.Context, user user) (utils.Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})

	apiKey := uuid.NewV4().String()

	if err != nil {
		return utils.Response{StatusCode: 404}, err
	}

	json.HTMLEscape(&buf, body)

	fmt.Println(user.Email)
	fmt.Println(apiKey)

	resp := utils.Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}
