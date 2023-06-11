package medicalrecord

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type medicalRecordHandler struct {
	medicalRecordService ports.MedicalRecordService
}

func newHandler(medicalRecordService ports.MedicalRecordService) *medicalRecordHandler {
	return &medicalRecordHandler{
		medicalRecordService: medicalRecordService,
	}
}

// CreateMedicalRecord handles the HTTP request for creating a medical record.
func (h *medicalRecordHandler) CreateMedicalRecord(c *gin.Context) {
	reqCreate := &entity.MedicalRecord{}

	// Bind incoming JSON payload to the reqCreate struct.
	if err := c.ShouldBindJSON(reqCreate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Get user UUID from JWT token
	userUUID, err := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid user UUID", err)
		return
	}

	// Create the medical record and store it in the database.
	createdMedicalRecord, statusCode, err := h.medicalRecordService.CreateMedicalRecord(c, userUUID, reqCreate)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while creating the medical record", err)
		return
	}

	// Return a successful response with the created medical record.
	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"message": "Medical record created successfully",
		"data": gin.H{
			"medical_record": createdMedicalRecord,
		},
	})
}

// GetMedicalRecord handles the HTTP request for getting a medical record.
func (h *medicalRecordHandler) GetMedicalRecord(c *gin.Context) {
	// Get user UUID from JWT token
	userUUID, err := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid user UUID", err)
		return
	}

	// Get the medical record for the user from the database
	medicalRecord, statusCode, err := h.medicalRecordService.GetMedicalRecord(c, userUUID)
	if err != nil {
		handleError(c, statusCode, "An error occurred while getting the medical record", err)
		return
	}

	// Return a successful response with the retrieved medical record
	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"message": "Medical record retrieved successfully",
		"data":    medicalRecord,
	})
}

// UpdateMedicalRecord handles the HTTP request for updating a medical record.
func (h *medicalRecordHandler) UpdateMedicalRecord(c *gin.Context) {
	// Get medical record ID from path parameter
	medicalRecordUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.Param("uuid")))

	reqUpdate := &entity.MedicalRecord{}

	// Bind incoming JSON payload to the reqUpdate struct.
	if err := c.ShouldBindJSON(reqUpdate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Get user UUID from JWT token
	userUUID, err := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid user UUID", err)
		return
	}

	// Update the medical record in the database
	updatedMedicalRecord, statusCode, err := h.medicalRecordService.UpdateMedicalRecord(c, userUUID, medicalRecordUUID, reqUpdate)
	if err != nil {
		handleError(c, statusCode, "An error occurred while updating the medical record", err)
		return
	}

	// Return a successful response with the updated medical record
	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"message": "Medical record updated successfully",
		"data": gin.H{
			"medical_record": updatedMedicalRecord,
		},
	})
}

// handleError handles errors by sending an appropriate response to the client.
func handleError(c *gin.Context, status int, message string, err error) {
	// Log the error message and the error itself
	log.Printf("[MedicalRecordHandler]: %s, %v", message, err)

	// Send the JSON response with the status code and error message
	c.JSON(status, gin.H{
		"code":    status,
		"message": err.Error(),
	})
}
