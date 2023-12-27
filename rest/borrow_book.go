package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"lms/database"
	"lms/models"
	"net/http"
)

func (l *LMSConfig) BorrowBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, nil)
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &models.BorrowBookRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &models.BorrowBookResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.BorrowBook(l.LMSdb, req.StudentID, req.BookID, l.overdueDate)
	if err != nil {
		respondJSON(w, 200, &models.BorrowBookResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondJSON(w, 200, &models.BorrowBookResponse{
		Success: true,
		Message: "",
	})
	return
}
