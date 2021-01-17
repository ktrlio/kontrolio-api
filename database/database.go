package database

import (
	"fmt"

	"github.com/marcelovicentegc/kontrolio-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)

	conn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	db = conn

	if config.IS_OFFLINE {
		Migrate()
	}
}

// Returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
