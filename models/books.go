package models

type Books struct {
	BookID            int64  `json:"BookID"`
	BookShelveAddress string `json:"BookShelveAddress"`
	Title             string `json:"Title"`
	Author            string `json:"Author"`
	ISBN              string `json:"ISBN"`
	Genre             string `json:"Genre"`
	PublicationDate   string `json:"PublicationDate"`
	Publisher         string `json:"Publisher"`
	Description       string `json:"Description"`
	IsAvailable       bool   `json:"IsAvailable"`
}
