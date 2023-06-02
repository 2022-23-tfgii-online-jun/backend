package medical

import (
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type medicalHandler struct {
	medicalService ports.MedicalService
}

func newHandler(medicalService ports.MedicalService) *medicalHandler {
	return &medicalHandler{
		medicalService: medicalService,
	}
}

// UploadCSV handles the HTTP request for uploading and processing a CSV file.
func (m *medicalHandler) UploadCSV(c *gin.Context) {
	// Call the service method to process the CSV file and save the records in the database
	status, err := m.medicalService.CreateRecordFromFile(c)
	if err != nil {
		handleError(c, status, "An error occurred while processing the CSV file", err)
		return
	}

	// Return a successful response
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "CSV file processed and records saved successfully",
	})
}

// GetMedicalRecord handles the HTTP request for getting a specific medical record.
func (m *medicalHandler) GetAllMedicalRecords(c *gin.Context) {

	// Get the medical record from the database.
	medicalRecord, err := m.medicalService.GetAllMedicalRecords()
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while getting the medical record", err)
		return
	}

	// Return a successful response with the retrieved medical record.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Medical record retrieved successfully",
		"data":    medicalRecord,
	})
}

// handleError handles errors by sending an appropriate response to the client.
func handleError(c *gin.Context, status int, message string, err error) {
	// Log the error message and the error itself
	log.Printf("[ArticleHandler]: %s, %v", message, err)

	// Send the JSON response with the status code and error message
	c.JSON(status, gin.H{
		"code":    status,
		"message": message,
	})
}
