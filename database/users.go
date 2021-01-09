package database

import "fmt"

func GetUser(email string) *User {
	db := GetDB()

	var user User

	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	return &user
}
