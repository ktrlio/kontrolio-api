package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/marcelovicentegc/kontrolio-api/config"
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
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)

	conn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db = conn

	db.Debug().AutoMigrate(&Record{})
}

// Returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
