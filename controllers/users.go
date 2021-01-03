package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	"github.com/marcelovicentegc/kontrolio-api/database"
	"github.com/marcelovicentegc/kontrolio-api/utils"
	uuid "github.com/satori/go.uuid"
)

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(ctx context.Context, data user) (utils.Response, error) {
	var buf bytes.Buffer

	if len(data.Password) < 8 {

		body, err := json.Marshal(map[string]interface{}{
			"message": "Sorry, but the password must have at least 8 characters.",
		})

		if err != nil {
			return utils.Response{StatusCode: 404}, err
		}

		json.HTMLEscape(&buf, body)

		return utils.Response{
			StatusCode:      400,
			IsBase64Encoded: false,
			Body:            buf.String(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			}}, errors.New("Sorry, but the password must have at least 8 characters.")
	}

	hashedPassword, err := utils.HashPassword(data.Password)

	if err != nil {
		return utils.Response{StatusCode: 500}, err
	}

	apiKey := uuid.NewV4().String()

	user := database.User{Email: data.Email, Password: hashedPassword, ApiKey: apiKey}

	result := database.GetDB().Create(&user)

	if result.Error != nil {
		return utils.Response{StatusCode: 500}, result.Error
	}

	body, err := json.Marshal(map[string]interface{}{
		"message": "Account successfully created!",
		"apiKey":  apiKey,
	})

	if err != nil {
		return utils.Response{StatusCode: 404}, err
	}

	json.HTMLEscape(&buf, body)

	resp := utils.Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}
