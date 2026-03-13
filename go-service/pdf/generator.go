package pdf

import (
	"bytes"
	"fmt"
	"time"

	"github.com/go-pdf/fpdf"

	"student-report-service/models"
)

func deref(s *string) string {
	if s == nil {
		return "N/A"
	}
	return *s
}

func derefInt(i *int) string {
	if i == nil {
		return "N/A"
	}
	return fmt.Sprintf("%d", *i)
}

func GenerateStudentReport(student *models.Student) ([]byte, error) {
	doc := fpdf.New("P", "mm", "A4", "")
	doc.SetAutoPageBreak(true, 20)
	doc.AddPage()

	// Header
	doc.SetFont("Arial", "B", 20)
	doc.CellFormat(190, 12, "Student Report", "", 1, "C", false, 0, "")
	doc.SetFont("Arial", "", 10)
	doc.SetTextColor(128, 128, 128)
	doc.CellFormat(190, 6, fmt.Sprintf("Generated on %s", time.Now().Format("January 02, 2006")), "", 1, "C", false, 0, "")
	doc.SetTextColor(0, 0, 0)
	doc.Ln(8)

	// Divider
	doc.SetDrawColor(200, 200, 200)
	doc.Line(10, doc.GetY(), 200, doc.GetY())
	doc.Ln(6)

	// Personal Information
	sectionHeader(doc, "Personal Information")
	addRow(doc, "Name", student.Name)
	addRow(doc, "Email", student.Email)
	addRow(doc, "Phone", deref(student.Phone))
	addRow(doc, "Gender", deref(student.Gender))
	addRow(doc, "Date of Birth", deref(student.DOB))
	doc.Ln(4)

	// Academic Information
	sectionHeader(doc, "Academic Information")
	addRow(doc, "Class", deref(student.Class))
	addRow(doc, "Section", deref(student.Section))
	addRow(doc, "Roll Number", derefInt(student.Roll))
	addRow(doc, "Admission Date", deref(student.AdmissionDate))
	doc.Ln(4)

	// Guardian Information
	sectionHeader(doc, "Guardian Information")
	addRow(doc, "Father's Name", deref(student.FatherName))
	addRow(doc, "Father's Phone", deref(student.FatherPhone))
	addRow(doc, "Mother's Name", deref(student.MotherName))
	addRow(doc, "Mother's Phone", deref(student.MotherPhone))
	addRow(doc, "Guardian's Name", deref(student.GuardianName))
	addRow(doc, "Guardian's Phone", deref(student.GuardianPhone))
	addRow(doc, "Relation", deref(student.RelationOfGuardian))
	doc.Ln(4)

	// Address
	sectionHeader(doc, "Address")
	addRow(doc, "Current Address", deref(student.CurrentAddress))
	addRow(doc, "Permanent Address", deref(student.PermanentAddress))
	doc.Ln(4)

	// Administrative
	sectionHeader(doc, "Administrative")
	systemAccess := "Inactive"
	if student.SystemAccess {
		systemAccess = "Active"
	}
	addRow(doc, "System Access", systemAccess)
	addRow(doc, "Registered By", deref(student.ReporterName))

	var buf bytes.Buffer
	if err := doc.Output(&buf); err != nil {
		return nil, fmt.Errorf("failed to generate PDF: %w", err)
	}

	return buf.Bytes(), nil
}

func sectionHeader(doc *fpdf.Fpdf, title string) {
	doc.SetFont("Arial", "B", 13)
	doc.SetTextColor(40, 80, 160)
	doc.CellFormat(190, 8, title, "", 1, "L", false, 0, "")
	doc.SetTextColor(0, 0, 0)
	doc.SetDrawColor(200, 200, 200)
	doc.Line(10, doc.GetY(), 200, doc.GetY())
	doc.Ln(3)
}

func addRow(doc *fpdf.Fpdf, label, value string) {
	doc.SetFont("Arial", "B", 10)
	doc.CellFormat(50, 7, label, "", 0, "L", false, 0, "")
	doc.SetFont("Arial", "", 10)
	doc.CellFormat(140, 7, value, "", 1, "L", false, 0, "")
}
