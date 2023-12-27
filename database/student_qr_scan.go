package database

import (
	"fmt"
	"lms/models"

	"github.com/jmoiron/sqlx"
)

func ScanStudentQR(db sqlx.Ext, StudentID int64) (*models.ScanStudentQR, error) {

	counter := 0
	var id int64

	rows, err := db.Queryx(`SELECT StudentID FROM Students
	WHERE StudentID=?`,
		StudentID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		counter++
	}

	if counter == 0 {
		return nil, fmt.Errorf("Student does not exists")
	}

	return &models.ScanStudentQR{
		StudentID: id,
	}, nil
}
