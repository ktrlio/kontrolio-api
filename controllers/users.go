package controllers

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user user) (string, error) {
	fmt.Println(user.Email)
	apiKey := uuid.NewV4().String()
	return apiKey, nil
}
