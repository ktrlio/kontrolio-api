package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

var db *gorm.DB

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

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")

	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", DB_HOST, DB_USER, DB_NAME, DB_PASSWORD)

	conn, e := gorm.Open("postgres", connectionString)
	if e != nil {
		fmt.Println(err)
	}

	db = conn

	db.Debug().AutoMigrate(&Record{})
}

// Returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
