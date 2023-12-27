package database

import (
	"lms/models"

	"github.com/jmoiron/sqlx"
)

func ReadAllBorrowedBooks(db sqlx.Ext) ([]*models.BorrowedBooks, error) {

	borrowedBooksArr := make([]*models.BorrowedBooks, 0)
	var TransactionID int64
	var StudentID int64
	var BookID int64
	var BorrowedDate int64
	var ReturnedDate int64
	var OverdueDate int64
	var IsReturned bool

	rows, err := db.Queryx(`SELECT
	TransactionID,
	StudentID,
	BookID,
	BorrowedDate,
	ReturnedDate,
	OverdueDate,
	IsReturned FROM BorrowedBooks`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&TransactionID, &StudentID, &BookID, &BorrowedDate, &ReturnedDate, &OverdueDate, &IsReturned)
		if err != nil {
			return nil, err
		}

		book, err := ReadBooksByID(db, BookID)
		if err != nil {
			return nil, err
		}

		student, err := ReadStudentsByID(db, StudentID)
		if err != nil {
			return nil, err
		}

		borrowedBooksArr = append(borrowedBooksArr, &models.BorrowedBooks{
			TransactionID: TransactionID,
			StudentID:     StudentID,
			BookID:        BookID,
			BorrowedDate:  BorrowedDate,
			ReturnedDate:  ReturnedDate,
			OverdueDate:   OverdueDate,
			IsReturned:    IsReturned,
			Book:          *book,
			Student:       *student,
		})
	}
	return borrowedBooksArr, nil
}

func ReadBorrowedBooksByStudent(db sqlx.Ext, studentID int64) ([]*models.BorrowedBooks, error) {

	borrowedBooksArr := make([]*models.BorrowedBooks, 0)
	var TransactionID int64
	var StudentID int64
	var BookID int64
	var BorrowedDate int64
	var ReturnedDate int64
	var OverdueDate int64
	var IsReturned bool

	rows, err := db.Queryx(`SELECT
	TransactionID,
	StudentID,
	BookID,
	BorrowedDate,
	ReturnedDate,
	OverdueDate,
	IsReturned FROM BorrowedBooks WHERE StudentID = ?`, studentID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&TransactionID, &StudentID, &BookID, &BorrowedDate, &ReturnedDate, &OverdueDate, &IsReturned)
		if err != nil {
			return nil, err
		}

		book, err := ReadBooksByID(db, BookID)
		if err != nil {
			return nil, err
		}

		student, err := ReadStudentsByID(db, BookID)
		if err != nil {
			return nil, err
		}

		borrowedBooksArr = append(borrowedBooksArr, &models.BorrowedBooks{
			TransactionID: TransactionID,
			StudentID:     StudentID,
			BookID:        BookID,
			BorrowedDate:  BorrowedDate,
			ReturnedDate:  ReturnedDate,
			OverdueDate:   OverdueDate,
			IsReturned:    IsReturned,
			Book:          *book,
			Student:       *student,
		})
	}
	return borrowedBooksArr, nil
}
