package models

type ScanStudentQR struct {
	StudentID int64 `json:"StudentID"`
}

type ScanStudentQRResponse struct {
	Success  bool           `json:"Success"`
	Message  string         `json:"Message"`
	Response *ScanStudentQR `json:"Response"`
}
