package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

func BorrowBook(db sqlx.Ext, StudentID int64, BookID int64, overdue int64) error {

	timeNow := time.Now().Unix()
	overdueDate := timeNow + overdue

	_, err := db.Exec(`INSERT INTO BorrowedBooks (
		StudentID,
		BookID,
		BorrowedDate,
		ReturnedDate,
		OverdueDate,
		IsReturned
	)
	Values(?,?,?,?,?,?)`,
		StudentID,
		BookID,
		timeNow,
		0,
		overdueDate,
		false,
	)

	if err != nil {
		return err
	}

	_, err = db.Exec(`UPDATE Books SET 
	IsAvailable = ? WHERE BookID= ?`,
		false,
		BookID,
	)

	if err != nil {
		return err
	}

	return nil
}
