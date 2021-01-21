package database

import (
	"fmt"

	"github.com/marcelovicentegc/kontrolio-api/utils"
	uuid "github.com/satori/go.uuid"
)

func GetUserByEmail(email string) *User {
	db := GetDB()

	var user User

	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil
	}

	return &user
}

func GetUserByApiKey(apiKey string) *User {
	db := GetDB()

	var user User

	result := db.Where("api_key = ?", apiKey).Take(&user)

	if result.Error != nil {
		fmt.Println("[GetUserByApiKey query]: " + result.Error.Error())
		return nil
	}

	return &user
}

func InsertUser(email string, password string) error {
	db := GetDB()

	hashedPassword, err := utils.HashPassword(password)

	if err != nil {
		return err
	}

	apiKey := uuid.NewV4().String()

	user := User{Email: email, Password: hashedPassword, ApiKey: apiKey}

	result := db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
