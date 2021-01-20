package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

type errorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

type responseBody struct {
	Data *string `json:"data"`
}

type userResponseBody struct {
	Data User `json:"data"`
}

type secretResponseBody struct {
	Data Secret `json:"data"`
}

type recordResponseBody struct {
	Data Record `json:"data"`
}

type recordRequestBody struct {
	Data PartialRecord `json:"data"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Secret struct {
	SecretString string `json:"secretString"`
}

type PartialRecord struct {
	Time   string `json:"time"`
	ApiKey string `json:"apiKey"`
}

type Record struct {
	Time       string `json:"time"`
	RecordType string `json:"recordType"`
}

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil
}

func parseUser(body string) (*User, error) {
	data := &userResponseBody{}
	err := json.Unmarshal([]byte(body), data)

	if err != nil {
		fmt.Println("Could not parse user object: " + err.Error())
		return nil, errors.New("Sorry, something went wrong while parsing the request.")
	}

	return &data.Data, nil
}

func parseSecret(body string) (*string, error) {
	data := &responseBody{}
	err := json.Unmarshal([]byte(body), data)

	if err != nil {
		fmt.Println("Could not parse secret object: " + err.Error())
		return nil, errors.New("Sorry, something went wrong while parsing the request.")
	}

	quotedSecret := *data.Data

	firstChar := string((quotedSecret)[0])
	lastChar := string((quotedSecret)[len(quotedSecret)-1])

	if firstChar != `"` && lastChar != firstChar {
		quotedSecret = strconv.Quote(quotedSecret)
	}

	return &quotedSecret, nil
}
