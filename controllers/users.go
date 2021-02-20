package controllers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/marcelovicentegc/kontrolio-api/config"
	"github.com/marcelovicentegc/kontrolio-api/database"
	"github.com/marcelovicentegc/kontrolio-api/utils"
)

// CreateUser is responsible for creating users
func CreateUser(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	data, err := parseUser(req.Body)

	if err != nil {
		return apiResponse(http.StatusBadGateway, errorBody{aws.String(err.Error())})
	}

	if len(data.Password) < 8 {
		return apiResponse(http.StatusBadRequest, errorBody{aws.String("Sorry, but the password must have at least 8 characters.")})
	}

	existentUser := database.GetUserByEmail(data.Email)

	if existentUser != nil {
		return apiResponse(http.StatusBadRequest, errorBody{aws.String("Email already taken.")})
	}

	err = database.InsertUser(data.Email, data.Password)

	if config.ENABLE_EMAIL_AUTH {
		utils.SendEMail(
			data.Email,
			"Account creation",
			"<h1>One more step ahead!</h1><br /><p>Click on the following URL to authenticate your account: "+"<a href="+"randomUrl"+"></a>"+"</p>",
			"One more step ahead! Click or copy and paste the following URL to your browser to authenticate your account"+"randomUrl",
			nil)
	}

	token, err := signToken(data.Email)

	if err != nil {
		return apiResponse(http.StatusBadRequest, errorBody{aws.String(err.Error())})
	}

	secret := Secret{*token}

	if err != nil {
		return apiResponse(http.StatusBadGateway, errorBody{aws.String((err.Error()))})
	}

	return apiResponse(http.StatusCreated, secretResponseBody{secret})
}
