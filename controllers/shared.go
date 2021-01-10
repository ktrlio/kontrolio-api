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

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil
}

func parseUser(body string) (*User, error) {
	user := &User{}
	err := json.Unmarshal([]byte(body), user)

	if err != nil {
		return nil, errors.New("Sorry, something went wrong while parsing the request.")
	}

	return user, nil
}

// Parses both the API key and the JWT.
func parseKey(body string) (*string, error) {
	var key *string
	err := json.Unmarshal([]byte(body), key)

	if err != nil {
		return nil, errors.New("Sorry, something went wrong while parsing the request.")
	}

	return key, nil
}
