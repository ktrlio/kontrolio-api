package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DB_USER           string
	DB_PASSWORD       string
	DB_NAME           string
	DB_HOST           string
	DB_TYPE           string
	CLIENT_URL        string
	SENDER_EMAIL      string
	ENABLE_EMAIL_AUTH bool
)

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
	CLIENT_URL = os.Getenv("CLIENT_URL")
	SENDER_EMAIL = os.Getenv("SENDER_EMAIL")

	if shouldEnableEmailAuth, err := strconv.Atoi(os.Getenv("ENABLE_EMAIL_AUTH")); err == nil {
		ENABLE_EMAIL_AUTH = shouldEnableEmailAuth != 0
	}

}
