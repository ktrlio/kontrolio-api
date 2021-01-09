package controllers

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/marcelovicentegc/kontrolio-api/config"
	"github.com/marcelovicentegc/kontrolio-api/database"
	"github.com/marcelovicentegc/kontrolio-api/utils"
	uuid "github.com/satori/go.uuid"
)

func CreateUser(ctx context.Context, data User) (*events.APIGatewayProxyResponse, error) {
	if data == (User{}) {
		return utils.ApiResponse(http.StatusBadRequest, utils.ErrorBody{aws.String("Sorry, but you must provide a email and a password.")})
	}

	if len(data.Password) < 8 {
		return utils.ApiResponse(http.StatusBadRequest, utils.ErrorBody{aws.String("Sorry, but the password must have at least 8 characters.")})
	}

	existentUser := utils.GetUser(data.Email)

	if existentUser != nil {
		return utils.ApiResponse(http.StatusBadRequest, utils.ErrorBody{aws.String("Email already taken.")})
	}

	hashedPassword, err := utils.HashPassword(data.Password)

	if err != nil {
		return utils.ApiResponse(http.StatusBadRequest, utils.ErrorBody{aws.String(err.Error())})
	}

	apiKey := uuid.NewV4().String()

	user := database.User{Email: data.Email, Password: hashedPassword, ApiKey: apiKey}

	result := database.GetDB().Create(&user)

	if result.Error != nil {
		return utils.ApiResponse(http.StatusBadRequest, utils.ErrorBody{aws.String(result.Error.Error())})
	}

	if config.ENABLE_EMAIL_AUTH {
		utils.SendEMail(
			data.Email,
			"Account creation",
			"<h1>One more step ahead!</h1><br /><p>Click on the following URL to authenticate your account: "+"<a href="+"randomUrl"+"></a>"+"</p>",
			"One more step ahead! Click or copy and paste the following URL to your browser to authenticate your account"+"randomUrl",
			nil)
	}

	return utils.ApiResponse(http.StatusOK, result)
}
