package utils

import (
	"fmt"

	"github.com/marcelovicentegc/kontrolio-api/database"
)

func GetUser(email string) *database.User {
	db := database.GetDB()

	var user database.User

	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	return &user
}
