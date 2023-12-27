package database

import (
	"lms/models"

	"github.com/jmoiron/sqlx"
)

func CreateBooks(db sqlx.Ext, BookShelveAddress string, Title string, Author string, ISBN string, Genre string, PublicationDate string,
	Publisher string, Description string) (int64, error) {

	query, err := db.Exec(`INSERT INTO Books (
		BookShelveAddress,
		Title,
		Author,
		ISBN,
		Genre,
		PublicationDate,
		Publisher,
		Description,
		IsAvailable
	)
	Values(?,?,?,?,?,?,?,?,?)`,
		BookShelveAddress,
		Title,
		Author,
		ISBN,
		Genre,
		PublicationDate,
		Publisher,
		Description,
		true,
	)

	if err != nil {
		return 0, err
	}

	BookID, err := query.LastInsertId()
	if err != nil {
		return 0, err
	}
	return BookID, nil
}

func ReadBooks(db sqlx.Ext) ([]*models.Books, error) {

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
	IsAvailable FROM Books`)

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

func UpdateBooks(db sqlx.Ext, BookID int64, BookShelveAddress string, Title string, Author string, ISBN string, Genre string, PublicationDate string,
	Publisher string, Description string) error {

	_, err := db.Exec(`UPDATE Books SET 
		BookShelveAddress = ?,
		Title = ?,
		Author = ?,
		ISBN = ?,
		Genre = ?,
		PublicationDate = ?,
		Publisher = ?,
		Description = ? WHERE BookID= ?`,
		BookShelveAddress,
		Title,
		Author,
		ISBN,
		Genre,
		PublicationDate,
		Publisher,
		Description,
		BookID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteBooks(db sqlx.Ext, BookID int64) error {

	_, err := db.Exec(`DELETE FROM Books WHERE BookID = ? `, BookID)

	if err != nil {
		return err
	}

	return nil
}

func ReadBooksByID(db sqlx.Ext, bkID int64) (*models.Books, error) {

	books := &models.Books{}
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
	IsAvailable FROM Books WHERE BookID = ?`, bkID)

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
		books = &models.Books{
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
		}
	}
	return books, nil
}
