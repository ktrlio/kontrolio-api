package controllers

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Record struct {
	Time       string `json:"time"`
	RecordType string `json:"recordType"`
	ApiKey     string `json:"apiKey"`
}
