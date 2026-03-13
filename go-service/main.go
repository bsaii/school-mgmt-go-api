package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"student-report-service/handlers"
	"student-report-service/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	backendURL := os.Getenv("BACKEND_API_URL")
	if backendURL == "" {
		backendURL = "http://localhost:5003"
	}

	serviceKey := os.Getenv("SERVICE_API_KEY")
	if serviceKey == "" {
		log.Fatal("SERVICE_API_KEY environment variable is required")
	}

	client := services.NewHTTPStudentClient(backendURL, serviceKey)

	handler := &handlers.ReportHandler{
		Client: client,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/students/{id}/report", handler.HandleStudentReport)

	log.Printf("Student Report Service running on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
