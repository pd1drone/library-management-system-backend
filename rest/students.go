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

func (l *LMSConfig) CreateStudents(w http.ResponseWriter, r *http.Request) {

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

	req := &models.Students{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &models.CreateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	_, err = database.CreateStudent(l.LMSdb, req.FirstName, req.MiddleName, req.LastName, req.DateOfBirth,
		req.Gender, req.Address, req.Email, req.PhoneNumber, req.LibraryCardNumber)
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

func (l *LMSConfig) ReadStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	StudentsData, err := database.ReadStudents(l.LMSdb)
	if err != nil {
		log.Print(err)
		respondJSON(w, 400, nil)
		return
	}

	respondJSON(w, 200, StudentsData)
}

func (l *LMSConfig) UpdateStudents(w http.ResponseWriter, r *http.Request) {

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

	req := &models.Students{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &models.UpdateResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.UpdateStudents(l.LMSdb, req.StudentID, req.FirstName, req.MiddleName, req.LastName, req.DateOfBirth, req.Gender, req.Address,
		req.Email, req.PhoneNumber, req.LibraryCardNumber)
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

func (l *LMSConfig) DeleteStudents(w http.ResponseWriter, r *http.Request) {

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

	req := &models.DeleteStudentRequest{}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &models.DeleteResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = database.DeleteStudents(l.LMSdb, req.StudentID)
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
