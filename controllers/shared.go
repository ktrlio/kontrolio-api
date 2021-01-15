package controllers

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
)

type errorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

type responseBody struct {
	Data *string `json:"data"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userResponse struct {
	Data User `json:"data"`
}

type Secret struct {
	secretString string
}

type secretResponse struct {
	Data Secret `json:"data"`
}

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil
}

func parseUser(body string) (*User, error) {
	data := &userResponse{}
	err := json.Unmarshal([]byte(body), data)

	if err != nil {
		return nil, errors.New("Sorry, something went wrong while parsing the request.")
	}

	return &data.Data, nil
}

func parseSecret(body string) (*Secret, error) {
	data := &secretResponse{}
	err := json.Unmarshal([]byte(body), data)

	if err != nil {
		return nil, errors.New("Sorry, something went wrong while parsing the request.")
	}

	return &data.Data, nil
}
