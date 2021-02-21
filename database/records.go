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

// RecordTypeRegistry holds the record types: IN, or OUT
var RecordTypeRegistry = newRecordTypeRegistry()

// GetLastRecord gets the last entry. This is particularly useful to check whether
// the next entry should be of type IN, or OUT
func GetLastRecord(userID uint) *Record {
	db := GetDB()

	var record Record

	result := db.Where("user_id = ?", userID).Last(&record)

	if result.Error != nil {
		fmt.Println("[GetLastRecord query]: " + result.Error.Error())
		return nil
	}

	return &record
}

// InsertRecord inserts a new record on the database
func InsertRecord(userID uint, clientTime string, recordType string) (*Record, error) {
	db := GetDB()

	parsedTime, err := time.Parse(utils.RecordTimeFormat, clientTime)

	if err != nil {
		return nil, err
	}

	record := Record{UserID: userID, Time: parsedTime, RecordType: recordType}

	result := db.Create(&record)

	if result.Error != nil {
		fmt.Println("[InsertRecord query]: " + result.Error.Error())
		return nil, result.Error
	}

	return &record, nil

}

// QueryRecords query records considering the pagination settings
func QueryRecords(userID uint, limit uint, offset uint, startDate *time.Time, endDate *time.Time) (*[]Record, uint) {
	db := GetDB()

	var records []Record
    var count int64

	countResult := db.Where("user_id = ?", userID).Find(&records).Count(&count)
	result := db.Where("user_id = ?", userID).Limit(int(limit)).Offset(int(offset)).Order("id DESC").Find(&records)

	if countResult.Error != nil {
		fmt.Println("[GetRecords query [1]]: " + result.Error.Error())
		return nil, 0
	}

	if result.Error != nil {
		fmt.Println("[GetRecords query [2]]: " + result.Error.Error())
		return nil, 0
	}

	return &records, uint(count)
}

// QueryAllRecords simply gets every record ever saved
func QueryAllRecords(userID uint) *[]Record {
	db := GetDB()

	var records []Record

	result := db.Where("user_id = ?", userID).Order("id DESC").Find(&records)

	if result.Error != nil {
		fmt.Println("[GetAllRecords query]: " + result.Error.Error())
		return nil
	}

	return &records
}
