package models

type BorrowedBooks struct {
	TransactionID int64    `json:"TransactionID"`
	StudentID     int64    `json:"StudentID"`
	BookID        int64    `json:"BookID"`
	BorrowedDate  int64    `json:"BorrowedDate"`
	ReturnedDate  int64    `json:"ReturnedDate"`
	OverdueDate   int64    `json:"OverdueDate"`
	IsReturned    bool     `json:"IsReturned"`
	Student       Students `json:"Student"`
	Book          Books    `json:"Book"`
}

type BorrowBookRequest struct {
	StudentID int64 `json:"StudentID"`
	BookID    int64 `json:"BookID"`
}

type BorrowBookResponse struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

type ReturnBookRequest struct {
	StudentID int64 `json:"StudentID"`
	BookID    int64 `json:"BookID"`
}

type ReturnBookResponse struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

type ReadBorrowBooksByStudentRequest struct {
	StudentID int64 `json:"StudentID"`
}
