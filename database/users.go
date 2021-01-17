package database

import "fmt"

func GetUser(email string) *User {
	db := GetDB()

	var user User

	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		fmt.Println(fmt.Errorf("Something went wrong while getting a user by email: %w", result.Error))
		return nil
	}

	return &user
}
