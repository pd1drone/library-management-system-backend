package rest

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (wt *LMSConfig) check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func Routes() {
	// Check if the log file exists
	if _, err := os.Stat("/root/library-management-system-backend/http_logs.log"); os.IsNotExist(err) {
		// Create a log file
		logFile, err := os.Create("/root/library-management-system-backend/http_logs.log")
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()
	}

	// Open the log file
	logFile, err := os.OpenFile("/root/library-management-system-backend/http_logs.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// // Create a new logger using the log file
	logger := middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  log.New(logFile, "", log.LstdFlags), // Use the log file as the output
		NoColor: true,
	})

	log.Print("Starting Library Management System Backend Service.....")
	r := chi.NewRouter()

	newLMS, err := New()
	if err != nil {
		log.Fatal(err)
	}
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(logger)
	r.Use(func(next http.Handler) http.Handler {
		return logPayloads(logFile, next)
	})

	// Check route if service is running
	r.Get("/check", newLMS.check)

	// Login route
	r.Post("/login", newLMS.Login)

	// Students route
	r.Get("/students", newLMS.ReadStudents)
	r.Delete("/students", newLMS.DeleteStudents)
	r.Put("/students", newLMS.UpdateStudents)
	r.Post("/students", newLMS.CreateStudents)

	// Books route
	r.Get("/books", newLMS.ReadBooks)
	r.Delete("/books", newLMS.DeleteBooks)
	r.Put("/books", newLMS.UpdateBooks)
	r.Post("/books", newLMS.CreateBooks)

	// Get Available Books
	r.Get("/available_books", newLMS.ReadAvailableBooks)

	// Borrow a book
	r.Post("/borrow", newLMS.BorrowBook)

	// Return a book
	r.Post("/return", newLMS.ReturnBook)

	// Read all borrowed books
	r.Get("/all_borrowed_books", newLMS.ReadAllBorrowedBooks)

	// Read borrowed books by student
	r.Get("/borrowed_books/{ID}", newLMS.ReadBorrowedBooksByStudent)

	// Read overdue books
	r.Get("/overdue_books", newLMS.ReadOverdueBooks)

	// Scan StudentID in RFID
	r.Post("/scan", newLMS.ScanStudentQR)

	r.Get("/overdue_books_csv", newLMS.ReadOverdueBooksCSV)
	r.Get("/borrowed_books_csv", newLMS.ReadAllBorrowedBooksCSV)

	log.Fatal(http.ListenAndServe("0.0.0.0:8090", r))
}

func logPayloads(logFile *os.File, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Record start time of request processing
		startTime := time.Now()

		// Call the next middleware/handler in the chain and capture response
		responseRecorder := httptest.NewRecorder()
		next.ServeHTTP(responseRecorder, r)

		// Record end time of request processing
		endTime := time.Now()

		// Read request payload
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}

		// Restore request body after reading
		r.Body = io.NopCloser(bytes.NewBuffer(requestBody))

		// Log the request and response payloads
		log.SetOutput(logFile)
		log.Printf("Request: %s %s\n", r.Method, r.URL.Path)

		if len(requestBody) > 0 {
			log.Printf("Request Payload: %s\n", string(requestBody))
		} else {
			log.Println("Request Payload: [Empty]")
		}

		// Read response payload
		responseBody := responseRecorder.Body.Bytes()

		if len(responseBody) > 0 {
			log.Printf("Request Response: %d %s\nPayload: %s\n", responseRecorder.Code, http.StatusText(responseRecorder.Code), string(responseBody))
		} else {
			log.Printf("Request Response: %d %s\nPayload: [Empty]\n", responseRecorder.Code, http.StatusText(responseRecorder.Code))
		}
		log.Printf("Duration: %v\n", endTime.Sub(startTime))

		// Write response back to original response writer
		for k, v := range responseRecorder.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(responseRecorder.Code)
		w.Write(responseRecorder.Body.Bytes())

		log.Printf("\n")
	})
}
