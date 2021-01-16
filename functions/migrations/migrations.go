package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/marcelovicentegc/kontrolio-api/database"
)

func main() {
	lambda.Start(database.Migrate)
}
