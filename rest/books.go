package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"lms/database"
	"lms/models"
	"log"
	"net/http"
)

func (l *LMSConfig) CreateBooks(w http.ResponseWriter, r *http.Request) {

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

	req := &models.Books{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &models.CreateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	_, err = database.CreateBooks(l.LMSdb, req.BookShelveAddress, req.Title, req.Author, req.ISBN, req.Genre,
		req.PublicationDate, req.Publisher, req.Description)
	if err != nil {
		respondJSON(w, 200, &models.CreateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondJSON(w, 200, &models.CreateResponse{
		Success: true,
		Message: "",
	})
	return
}

func (l *LMSConfig) ReadBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	BooksData, err := database.ReadBooks(l.LMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, BooksData)
}

func (l *LMSConfig) UpdateBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, &models.UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &models.Books{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &models.UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.UpdateBooks(l.LMSdb, req.BookID, req.BookShelveAddress, req.Title, req.Author, req.ISBN,
		req.Genre, req.PublicationDate, req.Publisher, req.Description)
	if err != nil {
		respondJSON(w, 200, &models.UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondJSON(w, 200, &models.UpdateResponse{
		Success: true,
		Message: "",
	})
}

func (l *LMSConfig) DeleteBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, &models.DeleteResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &models.DeleteBookRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &models.DeleteResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.DeleteBooks(l.LMSdb, req.BookID)
	if err != nil {
		respondJSON(w, 200, &models.DeleteResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondJSON(w, 200, &models.DeleteResponse{
		Success: true,
		Message: "",
	})
	return
}
