package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/marcelovicentegc/kontrolio-api/database"
	"github.com/marcelovicentegc/kontrolio-api/utils"
)

func CreateRecord(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	parsedRecord, err := parseRecord(req.Body)

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

func GetRecords(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	requestBody, err := parseRecordsRequest(req.Body)

	if err != nil {
		return apiResponse(http.StatusBadGateway, errorBody{aws.String(err.Error())})
	}

	formattedSecret, err := json.Marshal(responseBody{&requestBody.Auth.SecretString})

	if err != nil {
		return apiResponse(http.StatusBadGateway, errorBody{aws.String(err.Error())})
	}

	email, err := isLoggedIn(string(formattedSecret))

	if err != nil {
		return apiResponse(http.StatusBadGateway, errorBody{aws.String(err.Error())})
	}

	user := database.GetUserByEmail(*email)

	if user == nil {
		return apiResponse(http.StatusBadGateway, errorBody{aws.String("User not found.")})
	}

	records, count := database.QueryRecords(
		user.ID,
		requestBody.Filter.Pagination.Limit,
		requestBody.Filter.Pagination.Offset,
		requestBody.Filter.DateRange.StartDate,
		requestBody.Filter.DateRange.EndDate,
	)

	var response RecordsResponseBody

	response.Count = count
	response.CurrentPage = uint(math.Ceil(float64(requestBody.Filter.Pagination.Offset) / float64(requestBody.Filter.Pagination.Limit)))
	response.TotalPages = uint(math.Ceil(float64(count) / float64(requestBody.Filter.Pagination.Limit)))

	for _, record := range *records {
		formattedRecord := Record{record.Time.Format(utils.RecordTimeFormat), record.RecordType}
		response.Results = append(response.Results, formattedRecord)
	}

	return apiResponse(http.StatusOK, recordsResponseBody{response})
}

// Gets all records. This function is only triggered by the CLI client,
// in order to sync offline and online records.
func GetAllRecords(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	apiKey, err := parseSecret(req.Body)

	if err != nil {
		fmt.Println("[getAllRecords] failed with: ", err.Error())
		return apiResponse(http.StatusBadRequest, errorBody{aws.String(err.Error())})
	}

	user := database.GetUserByApiKey(*apiKey)

	if user == nil {
		return apiResponse(http.StatusBadGateway, errorBody{aws.String("User not found.")})
	}

	records := database.QueryAllRecords(user.ID)

	var response AllRecordsResponseBody

	for _, record := range *records {
		formattedRecord := Record{record.Time.Format(utils.RecordTimeFormat), record.RecordType}
		response.Results = append(response.Results, formattedRecord)
	}

	return apiResponse(http.StatusOK, allRecordsResponseBody{response})
}
