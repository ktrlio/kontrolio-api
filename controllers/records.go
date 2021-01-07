package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/marcelovicentegc/kontrolio-api/utils"
)

func CreateRecord(ctx context.Context, record Record) (utils.Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})

	if err != nil {
		return utils.Response{StatusCode: 404}, err
	}

	json.HTMLEscape(&buf, body)

	fmt.Println(record.ApiKey)

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
