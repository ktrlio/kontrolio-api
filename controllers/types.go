package controllers

import "time"

type errorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

type responseBody struct {
	Data *string `json:"data"`
}

type userResponseBody struct {
	Data User `json:"data"`
}

type secretResponseBody struct {
	Data Secret `json:"data"`
}

type recordResponseBody struct {
	Data Record `json:"data"`
}

type recordsResponseBody struct {
	Data RecordsResponseBody `json:"data"`
}

type allRecordsResponseBody struct {
	Data AllRecordsResponseBody `json:"data"`
}

type recordRequestBody struct {
	Data PartialRecord `json:"data"`
}

type recordsRequestBody struct {
	Data RecordsRequestBody `json:"data"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Secret struct {
	SecretString string `json:"secretString"`
}

type PartialRecord struct {
	Time   string `json:"time"`
	ApiKey string `json:"apiKey"`
}

type Record struct {
	Time       string `json:"time"`
	RecordType string `json:"recordType"`
}

type RecordsRequestBody struct {
	Auth   Secret        `json:"auth"`
	Filter RecordsFilter `json:"filter"`
}

type RecordsResponseBody struct {
	Count       uint     `json:"count"`
	CurrentPage uint     `json:"currentPage"`
	TotalPages  uint     `json:"totalPages"`
	Results     []Record `json:"results"`
}

type AllRecordsResponseBody struct {
	Results []Record `json:"results"`
}

type RecordsFilter struct {
	DateRange  DateRange  `json:"dateRange"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Offset uint `json:"offset"`
	Limit  uint `json:"limit"`
}

type DateRange struct {
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
}
