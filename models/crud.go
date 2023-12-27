package models

type CreateResponse struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

type UpdateResponse struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

type DeleteResponse struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

type DeleteStudentRequest struct {
	StudentID int64 `json:"StudentID"`
}

type DeleteBookRequest struct {
	BookID int64 `json:"BookID"`
}
