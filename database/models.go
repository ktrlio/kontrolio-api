package database

import (
	"time"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Time       time.Time
	RecordType string
	UserID     uint
	User       User
}

type User struct {
	gorm.Model
	ApiKey   string
	Email    string
	Password string
}
