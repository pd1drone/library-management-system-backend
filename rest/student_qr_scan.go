package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"lms/database"
	"lms/models"
	"net/http"
)

func (l *LMSConfig) ScanStudentQR(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondJSON(w, 500, &models.ScanStudentQRResponse{
			Success:  false,
			Message:  err.Error(),
			Response: nil,
		})
		return
	}

	// Restore request body after reading
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	req := &models.ScanStudentQR{}
	fmt.Println(req)

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondJSON(w, 400, &models.ScanStudentQRResponse{
			Success:  false,
			Message:  err.Error(),
			Response: nil,
		})
		return
	}

	ScanQRData, err := database.ScanStudentQR(l.LMSdb, req.StudentID)
	if err != nil {
		respondJSON(w, 200, &models.ScanStudentQRResponse{
			Success:  false,
			Message:  err.Error(),
			Response: nil,
		})
		return
	}

	respondJSON(w, 200, &models.ScanStudentQRResponse{
		Success:  true,
		Message:  "",
		Response: ScanQRData,
	})
}
