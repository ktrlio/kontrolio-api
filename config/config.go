package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB_USER string
var DB_PASSWORD string
var DB_NAME string
var DB_HOST string
var DB_TYPE string

func init() {
	ENV := os.Getenv("ENV")

	if ENV != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_TYPE = "postgres"

}
