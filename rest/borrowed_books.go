package rest

import (
	"fmt"
	"lms/database"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (l *LMSConfig) ReadAllBorrowedBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	BorrowedBooksData, err := database.ReadAllBorrowedBooks(l.LMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, BorrowedBooksData)
}

func (l *LMSConfig) ReadBorrowedBooksByStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	ID := chi.URLParam(r, "ID")
	fmt.Println(ID)

	studentID, err := strconv.Atoi(ID)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}
	fmt.Println(studentID)

	BorrowedBooksData, err := database.ReadBorrowedBooksByStudent(l.LMSdb, int64(studentID))
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, BorrowedBooksData)
}
