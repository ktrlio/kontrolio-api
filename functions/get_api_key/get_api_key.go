package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/marcelovicentegc/kontrolio-api/controllers"
)

func main() {
	lambda.Start(controllers.GetApiKey)
}
