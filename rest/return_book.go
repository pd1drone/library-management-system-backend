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

func (l *LMSConfig) ReturnBook(w http.ResponseWriter, r *http.Request) {

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

	req := &models.ReturnBookRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &models.ReturnBookResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.ReturnBook(l.LMSdb, req.StudentID, req.BookID)
	if err != nil {
		respondJSON(w, 200, &models.ReturnBookResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondJSON(w, 200, &models.ReturnBookResponse{
		Success: true,
		Message: "",
	})
	return
}
