package database

import (
	"time"

	"gorm.io/gorm"

	uuid "github.com/satori/go.uuid"
)

type Record struct {
	gorm.Model
	Uuid       uuid.UUID
	Time       time.Time
	RecordType string
}

type User struct {
	gorm.Model
	Uuid     uuid.UUID
	ApiKey   string
	Email    string
	Password string
}
