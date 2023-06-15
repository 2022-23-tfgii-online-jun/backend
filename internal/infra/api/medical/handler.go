package medical

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

// medicalHandler type contains an instance of MedicalService.
type medicalHandler struct {
	medicalService ports.MedicalService
}

// newHandler is a constructor function for initializing medicalHandler with the given MedicalService.
// The return is a pointer to an medicalHandler instance.
func newHandler(medicalService ports.MedicalService) *medicalHandler {
	return &medicalHandler{
		medicalService: medicalService,
	}
}

// UploadCSV handles the HTTP request for uploading and processing a CSV file.
// It calls the medicalService method to process the CSV file and save the records in the database.
// If any error occurs during this process, it will return the corresponding status code and error message.
// If the CSV file is processed and the records are saved successfully, it will return a 200 OK status.
func (m *medicalHandler) UploadCSV(c *gin.Context) {
	status, err := m.medicalService.CreateRecordFromFile(c)
	if err != nil {
		handleError(c, status, "An error occurred while processing the CSV file", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "CSV file processed and records saved successfully",
	})
}

// GetAllMedicalRecords handles the HTTP request for getting all medical records.
// It retrieves all medical records from the database.
// If any error occurs during this process, it will return a 500 Internal Server Error status.
// If the medical records are retrieved successfully, it will return a 200 OK status with the retrieved medical records.
func (m *medicalHandler) GetAllMedicalRecords(c *gin.Context) {
	medicalRecords, err := m.medicalService.GetAllMedicalRecords()
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while getting the medical records", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Medical records retrieved successfully",
		"data":    medicalRecords,
	})
}

// AddRatingToMedical handles the HTTP request for adding a rating to a medical record.
// It binds the incoming JSON payload to the req struct.
// If any error occurs during this process, it will return a 400 Bad Request status.
// If the input parameters are invalid, it will return a 400 Bad Request status with an error message.
// If the rating is added to the medical record successfully, it will return a 200 OK status.
func (m *medicalHandler) AddRatingToMedical(c *gin.Context) {
	req := &entity.MedicalRating{}

	if err := c.ShouldBindJSON(req); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	if req.MedicalID == 0 || req.ReminderID == 0 {
		handleError(c, http.StatusBadRequest, "Invalid input", fmt.Errorf("medical and reminder IDs are required"))
		return
	}

	status, err := m.medicalService.AddRatingToMedical(req)
	if err != nil {
		handleError(c, status, "An error occurred while adding the rating", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Rating added successfully",
	})
}

// handleError handles errors by sending an appropriate response to the client.
// It logs the error message and the error itself, then sends a JSON response with the status code and error message.
func handleError(c *gin.Context, status int, message string, err error) {
	log.Printf("[MedicalHandler]: %s, %v", message, err)

	c.JSON(status, gin.H{
		"code":    status,
		"message": message,
	})
}
