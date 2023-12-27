package rest

import (
	"lms/database"
	"log"
	"net/http"
)

func (l *LMSConfig) ReadAvailableBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	BooksData, err := database.ReadAvailableBooks(l.LMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, BooksData)
}
