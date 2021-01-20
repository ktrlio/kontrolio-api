package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/marcelovicentegc/kontrolio-api/database"
	"github.com/marcelovicentegc/kontrolio-api/utils"
)

func parseRecord(body string) (*PartialRecord, error) {
	data := &recordRequestBody{}
	err := json.Unmarshal([]byte(body), data)

	if err != nil {
		fmt.Println("Could not parse record object: " + err.Error())
		return nil, errors.New("Sorry, something went wrong while parsing the request.")
	}

	return &data.Data, nil
}

func CreateRecord(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	parsedRecord, err := parseRecord(req.Body)

	fmt.Println(parsedRecord)

	if err != nil {
		return apiResponse(http.StatusBadGateway, errorBody{aws.String(err.Error())})
	}

	user := database.GetUserByApiKey(parsedRecord.ApiKey)

	if user == nil {
		return apiResponse(http.StatusBadGateway, errorBody{aws.String("User not found.")})
	}

	record := database.GetLastRecord(user.ID)

	recordType := func() string {
		if record == nil || record.RecordType == database.RecordTypeRegistry.Out {
			return database.RecordTypeRegistry.In
		} else {
			return database.RecordTypeRegistry.Out
		}
	}()

	newRecord, err := database.InsertRecord(user.ID, parsedRecord.Time, recordType)

	if err != nil {
		fmt.Println("Failed to insert record with: " + err.Error())
		return apiResponse(http.StatusBadRequest, errorBody{aws.String("Sorry, we couldn't save your record this time. Try again soon.")})
	}

	recordResponse := Record{newRecord.Time.Format(utils.RecordTimeFormat), newRecord.RecordType}

	return apiResponse(http.StatusOK, recordResponseBody{recordResponse})
}
