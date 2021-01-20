package database

import (
	"fmt"
	"time"

	"github.com/marcelovicentegc/kontrolio-api/utils"
)

type recordTypeRegistry struct {
	In  string
	Out string
}

func newRecordTypeRegistry() *recordTypeRegistry {
	return &recordTypeRegistry{
		In:  "IN",
		Out: "OUT",
	}
}

var RecordTypeRegistry = newRecordTypeRegistry()

func GetLastRecord(userId uint) *Record {
	db := GetDB()

	var record Record

	result := db.Where("user_id = ?", userId).Last(&record)

	if result.Error != nil {
		fmt.Println("[GetLastRecord query]: " + result.Error.Error())
		return nil
	}

	return &record
}

func InsertRecord(userId uint, clientTime string, recordType string) (*Record, error) {
	db := GetDB()

	parsedTime, err := time.Parse(utils.RecordTimeFormat, clientTime)

	if err != nil {
		return nil, err
	}

	record := Record{UserID: userId, Time: parsedTime, RecordType: recordType}

	result := db.Create(&record)

	if result.Error != nil {
		return nil, result.Error
	}

	return &record, nil

}
