package rest

import (
	"encoding/csv"
	"fmt"
	"lms/database"
	"log"
	"net/http"
	"time"
)

func (l *LMSConfig) ReadOverdueBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	OverdueBooks, err := database.ReadOverdueBooks(l.LMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, OverdueBooks)
}

func (l *LMSConfig) ReadOverdueBooksCSV(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	OverdueBooks, err := database.ReadOverdueBooks(l.LMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	writer := csv.NewWriter(w)
	defer writer.Flush()

	// Write CSV header
	header := []string{"TransactionID", "StudentID", "BookID", "BorrowedDate", "ReturnedDate", "OverdueDate", "IsReturned", "Student_FirstName", "Student_MiddleName", "Student_LastName", "Student_DateOfBirth", "Student_Gender", "Student_Address", "Student_Email", "Student_PhoneNumber", "Student_RegistrationDate", "Student_LibraryCardNumber", "Student_FullName", "Book_BookID", "Book_BookShelveAddress", "Book_Title", "Book_Author", "Book_ISBN", "Book_Genre", "Book_PublicationDate", "Book_Publisher", "Book_Description", "Book_IsAvailable"}
	if err := writer.Write(header); err != nil {
		http.Error(w, "Failed to write CSV header", http.StatusInternalServerError)
		fmt.Println("Error:", err)
		return
	}

	for _, t := range OverdueBooks {
		row := []string{
			fmt.Sprintf("%d", t.TransactionID),
			fmt.Sprintf("%d", t.StudentID),
			fmt.Sprintf("%d", t.BookID),
			fmt.Sprintf("%d", t.BorrowedDate),
			fmt.Sprintf("%d", t.ReturnedDate),
			fmt.Sprintf("%d", t.OverdueDate),
			fmt.Sprintf("%t", t.IsReturned),
			t.Student.FirstName,
			t.Student.MiddleName,
			t.Student.LastName,
			t.Student.DateOfBirth,
			t.Student.Gender,
			t.Student.Address,
			t.Student.Email,
			t.Student.PhoneNumber,
			t.Student.RegistrationDate,
			t.Student.LibraryCardNumber,
			t.Student.FullName,
			fmt.Sprintf("%d", t.Book.BookID),
			t.Book.BookShelveAddress,
			t.Book.Title,
			t.Book.Author,
			t.Book.ISBN,
			t.Book.Genre,
			t.Book.PublicationDate,
			t.Book.Publisher,
			t.Book.Description,
			fmt.Sprintf("%t", t.Book.IsAvailable),
		}
		if err := writer.Write(row); err != nil {
			http.Error(w, "Failed to write CSV row", http.StatusInternalServerError)
			fmt.Println("Error:", err)
			return
		}
	}

	// Set headers for file download
	w.Header().Set("Content-Type", "text/csv")

	// Generate filename with today's date
	today := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("overdue_books_%s.csv", today)

	// Set Content-Disposition header with the generated filename
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
}
