package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/marcelovicentegc/kontrolio-api/config"
	"github.com/marcelovicentegc/kontrolio-api/database"
	"github.com/marcelovicentegc/kontrolio-api/utils"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func parseUser(body string) (*User, error) {
	user := &User{}
	err := json.Unmarshal([]byte(body), user)

	if err != nil {
		return nil, errors.New("Sorry, something went wrong while parsing the request.")
	}

	return user, nil
}

func CreateUser(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	data, err := parseUser(req.Body)

	if err != nil {
		return apiResponse(http.StatusBadRequest, errorBody{aws.String("Sorry, something went wrong while parsing the request")})
	}

	if len(data.Password) < 8 {
		return apiResponse(http.StatusBadRequest, errorBody{aws.String("Sorry, but the password must have at least 8 characters.")})
	}

	existentUser := database.GetUser(data.Email)

	if existentUser != nil {
		return apiResponse(http.StatusBadRequest, errorBody{aws.String("Email already taken.")})
	}

	hashedPassword, err := utils.HashPassword(data.Password)

	if err != nil {
		return apiResponse(http.StatusBadRequest, errorBody{aws.String(err.Error())})
	}

	apiKey := uuid.NewV4().String()

	user := database.User{Email: data.Email, Password: hashedPassword, ApiKey: apiKey}

	result := database.GetDB().Create(&user)

	if result.Error != nil {
		return apiResponse(http.StatusBadRequest, errorBody{aws.String(result.Error.Error())})
	}

	if config.ENABLE_EMAIL_AUTH {
		utils.SendEMail(
			data.Email,
			"Account creation",
			"<h1>One more step ahead!</h1><br /><p>Click on the following URL to authenticate your account: "+"<a href="+"randomUrl"+"></a>"+"</p>",
			"One more step ahead! Click or copy and paste the following URL to your browser to authenticate your account"+"randomUrl",
			nil)
	}

	return apiResponse(http.StatusOK, result)
}
