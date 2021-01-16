package database

import "gorm.io/gorm"

// Returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
