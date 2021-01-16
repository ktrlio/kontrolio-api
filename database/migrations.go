package database

import (
	"fmt"

	"github.com/marcelovicentegc/kontrolio-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() {
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)

	conn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	db = conn

	db.Debug().AutoMigrate(&Record{}, &User{})
}
