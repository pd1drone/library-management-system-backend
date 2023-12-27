package database

import (
	"github.com/jmoiron/sqlx"
)

func ReturnBook(db sqlx.Ext, StudentID int64, BookID int64) error {

	_, err := db.Exec(`DELETE FROM BorrowedBooks WHERE BookID = ? AND StudentID = ? `, BookID, StudentID)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	_, err = db.Exec(`UPDATE Books SET 
	IsAvailable = ? WHERE BookID= ?`,
		true,
		BookID,
	)

	if err != nil {
		return err
	}

	return nil
}
