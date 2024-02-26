package models

type Students struct {
	StudentID         int64  `json:"StudentID"`
	FirstName         string `json:"FirstName"`
	MiddleName        string `json:"MiddleName"`
	LastName          string `json:"LastName"`
	DateOfBirth       string `json:"DateOfBirth"`
	Gender            string `json:"Gender"`
	Address           string `json:"Address"`
	Email             string `json:"Email"`
	PhoneNumber       string `json:"PhoneNumber"`
	RegistrationDate  string `json:"RegistrationDate"`
	LibraryCardNumber string `json:"LibraryCardNumber"`
	FullName          string `json:"FullName"`
}
