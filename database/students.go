package database

import (
	"lms/models"
	"time"

	"github.com/jmoiron/sqlx"
)

func CreateStudent(db sqlx.Ext, FirstName string, MiddleName string, LastName string, DateOfBirth string, Gender string,
	Address string, Email string, PhoneNumber string, LibraryCardNumber string) (int64, error) {

	currentTime := time.Now()
	// Format the time as "YYYY-MM-DD 03:04 PM"
	formattedTime := currentTime.Format("2006-01-02 03:04 PM")

	query, err := db.Exec(`INSERT INTO Students (
		FirstName,
		MiddleName,
		LastName,
		DateOfBirth,
		Gender,
		Address,
		Email,
		PhoneNumber,
		RegistrationDate,
		LibraryCardNumber
	)
	Values(?,?,?,?,?,?,?,?,?,?)`,
		FirstName,
		MiddleName,
		LastName,
		DateOfBirth,
		Gender,
		Address,
		Email,
		PhoneNumber,
		formattedTime,
		LibraryCardNumber,
	)

	if err != nil {
		return 0, err
	}

	StudentID, err := query.LastInsertId()
	if err != nil {
		return 0, err
	}
	return StudentID, nil
}

func ReadStudents(db sqlx.Ext) ([]*models.Students, error) {

	studentsArr := make([]*models.Students, 0)
	var StudentID int64
	var FirstName string
	var MiddleName string
	var LastName string
	var DateOfBirth string
	var Gender string
	var Address string
	var Email string
	var PhoneNumber string
	var RegistrationDate string
	var LibraryCardNumber string

	rows, err := db.Queryx(`SELECT StudentID,
	FirstName,
	MiddleName,
	LastName,
	DateOfBirth,
	Gender,
	Address,
	Email,
	PhoneNumber,
	RegistrationDate,
	LibraryCardNumber FROM Students`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&StudentID, &FirstName, &MiddleName, &LastName, &DateOfBirth, &Gender,
			&Address, &Email, &PhoneNumber, &RegistrationDate, &LibraryCardNumber)
		if err != nil {
			return nil, err
		}
		studentsArr = append(studentsArr, &models.Students{
			StudentID:         StudentID,
			FirstName:         FirstName,
			MiddleName:        MiddleName,
			LastName:          LastName,
			DateOfBirth:       DateOfBirth,
			Gender:            Gender,
			Address:           Address,
			Email:             Email,
			PhoneNumber:       PhoneNumber,
			RegistrationDate:  RegistrationDate,
			LibraryCardNumber: LibraryCardNumber,
		})
	}
	return studentsArr, nil
}

func UpdateStudents(db sqlx.Ext, StudentID int64, FirstName string, MiddleName string, LastName string, DateOfBirth string, Gender string,
	Address string, Email string, PhoneNumber string, LibraryCardNumber string) error {

	_, err := db.Exec(`UPDATE Students SET 
		FirstName = ?,
		MiddleName = ?,
		LastName = ?,
		DateOfBirth = ?,
		Gender = ?,
		Address = ?,
		Email = ?,
		PhoneNumber = ?,
		LibraryCardNumber = ? WHERE StudentID= ?`,
		FirstName,
		MiddleName,
		LastName,
		DateOfBirth,
		Gender,
		Address,
		Email,
		PhoneNumber,
		LibraryCardNumber,
		StudentID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteStudents(db sqlx.Ext, StudentID int64) error {

	_, err := db.Exec(`DELETE FROM Students WHERE StudentID = ? `, StudentID)

	if err != nil {
		return err
	}

	return nil
}

func ReadStudentsByID(db sqlx.Ext, stdID int64) (*models.Students, error) {

	students := &models.Students{}
	var StudentID int64
	var FirstName string
	var MiddleName string
	var LastName string
	var DateOfBirth string
	var Gender string
	var Address string
	var Email string
	var PhoneNumber string
	var RegistrationDate string
	var LibraryCardNumber string

	rows, err := db.Queryx(`SELECT StudentID,
	FirstName,
	MiddleName,
	LastName,
	DateOfBirth,
	Gender,
	Address,
	Email,
	PhoneNumber,
	RegistrationDate,
	LibraryCardNumber FROM Students WHERE StudentID = ?`, stdID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&StudentID, &FirstName, &MiddleName, &LastName, &DateOfBirth, &Gender,
			&Address, &Email, &PhoneNumber, &RegistrationDate, &LibraryCardNumber)
		if err != nil {
			return nil, err
		}
		students = &models.Students{
			StudentID:         StudentID,
			FirstName:         FirstName,
			MiddleName:        MiddleName,
			LastName:          LastName,
			DateOfBirth:       DateOfBirth,
			Gender:            Gender,
			Address:           Address,
			Email:             Email,
			PhoneNumber:       PhoneNumber,
			RegistrationDate:  RegistrationDate,
			LibraryCardNumber: LibraryCardNumber,
			FullName:          FirstName + " " + MiddleName + " " + LastName,
		}
	}
	return students, nil
}
