package database

import (
	"lms/models"

	"github.com/jmoiron/sqlx"
)

func ReadAvailableBooks(db sqlx.Ext) ([]*models.Books, error) {

	booksArr := make([]*models.Books, 0)
	var BookID int64
	var BookShelveAddress string
	var Title string
	var Author string
	var ISBN string
	var Genre string
	var PublicationDate string
	var Publisher string
	var Description string
	var IsAvailable bool

	rows, err := db.Queryx(`SELECT BookID,
	BookShelveAddress,
	Title,
	Author,
	ISBN,
	Genre,
	PublicationDate,
	Publisher,
	Description,
	IsAvailable FROM Books WHERE IsAvailable=true`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&BookID, &BookShelveAddress, &Title, &Author, &ISBN, &Genre,
			&PublicationDate, &Publisher, &Description, &IsAvailable)
		if err != nil {
			return nil, err
		}
		booksArr = append(booksArr, &models.Books{
			BookID:            BookID,
			BookShelveAddress: BookShelveAddress,
			Title:             Title,
			Author:            Author,
			ISBN:              ISBN,
			Genre:             Genre,
			PublicationDate:   PublicationDate,
			Publisher:         Publisher,
			Description:       Description,
			IsAvailable:       IsAvailable,
		})
	}
	return booksArr, nil
}
