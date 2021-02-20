package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

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
		return nil, errors.New("Sorry, something went wrong while parsing the request")
	}

	return &data.Data, nil
}

func ensureSecretIsUnderQuotes(secret string) string {
	firstChar := string((secret)[0])
	lastChar := string((secret)[len(secret)-1])

	if firstChar != `"` && lastChar != firstChar {
		secret = strconv.Quote(secret)
	}

	return secret
}

func parseSecret(body string) (*string, error) {
	data := &responseBody{}
	err := json.Unmarshal([]byte(body), data)

	if err != nil {
		fmt.Println("Could not parse secret object: " + err.Error())
		return nil, errors.New("Sorry, something went wrong while parsing the request")
	}

	quotedSecret := ensureSecretIsUnderQuotes(*data.Data)

	return &quotedSecret, nil
}

func parseRecord(body string) (*PartialRecord, error) {
	data := &recordRequestBody{}
	err := json.Unmarshal([]byte(body), data)

	if err != nil {
		fmt.Println("Could not parse record object: " + err.Error())
		return nil, errors.New("Sorry, something went wrong while parsing the request")
	}

	return &data.Data, nil
}

func parseRecordsRequest(body string) (*RecordsRequestBody, error) {
	data := &recordsRequestBody{}
	err := json.Unmarshal([]byte(body), data)

	if err != nil {
		fmt.Println("Could not parse records filter object: " + err.Error())
		return nil, errors.New("Sorry, something went wrong while parsing the request")
	}

	data.Data.Auth.SecretString = ensureSecretIsUnderQuotes(data.Data.Auth.SecretString)

	return &data.Data, nil
}
