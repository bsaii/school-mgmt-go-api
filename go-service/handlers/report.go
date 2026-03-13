package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"student-report-service/pdf"
	"student-report-service/services"
)

type ReportHandler struct {
	Client services.StudentClient
}

func (h *ReportHandler) HandleStudentReport(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "student ID is required", http.StatusBadRequest)
		return
	}

	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, "student ID must be a valid number", http.StatusBadRequest)
		return
	}

	student, err := h.Client.FetchStudent(r.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrStudentNotFound) {
			http.Error(w, "student not found", http.StatusNotFound)
			return
		}
		log.Printf("Error fetching student %s: %v", id, err)
		http.Error(w, "failed to fetch student data", http.StatusBadGateway)
		return
	}

	pdfBytes, err := pdf.GenerateStudentReport(student)
	if err != nil {
		log.Printf("Error generating PDF for student %s: %v", id, err)
		http.Error(w, "failed to generate report", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="student_report_%s.pdf"`, id))
	w.Write(pdfBytes)
}
