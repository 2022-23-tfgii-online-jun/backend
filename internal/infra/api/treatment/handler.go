package treatment

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// treatmentHandler type contains an instance of TreatmentService
type treatmentHandler struct {
	treatmentService ports.TreatmentService
}

// newHandler is a constructor function for initializing treatmentHandler with the given TreatmentService.
// The return is a pointer to a treatmentHandler instance.
func newHandler(treatmentService ports.TreatmentService) *treatmentHandler {
	return &treatmentHandler{
		treatmentService: treatmentService,
	}
}

// CreateTreatment handles the HTTP request for creating a treatment.
// It parses the incoming JSON payload, binds it to the reqCreate struct,
// and calls the treatment service to create the treatment.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the treatment is created successfully, it returns a 200 OK status with the created treatment.
func (t *treatmentHandler) CreateTreatment(c *gin.Context) {
	reqCreate := &entity.RequestCreateTreatment{}

	userUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))

	if err := c.ShouldBind(reqCreate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	treatment, statusCode, err := t.treatmentService.CreateTreatment(c, userUUID, reqCreate)
	if err != nil {
		handleError(c, statusCode, "An error occurred while creating the treatment", err)
		return
	}

	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": "Treatment created successfully",
		"data":    treatment,
	})
}

// GetAllTreatments handles the HTTP request for getting all treatments of a user.
// It retrieves all treatments of the user from the service.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the treatments are fetched successfully, it returns a 200 OK status with the retrieved treatments.
func (t *treatmentHandler) GetAllTreatments(c *gin.Context) {
	// Get user UUID from context
	userUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))

	// Get all treatments of the user from the database.
	treatments, err := t.treatmentService.GetAllTreatments(userUUID)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while retrieving treatments", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Treatments retrieved successfully",
		"data":    treatments,
	})
}

// handleError is a generic error handler that logs the error and responds.
func handleError(c *gin.Context, status int, msg string, err error) {
	log.Println(err)
	c.JSON(status, gin.H{
		"code":    status,
		"message": msg,
		"error":   err.Error(),
	})
}

// DeleteTreatment handles the HTTP request for deleting a treatment.
// It parses the treatment UUID from the path parameter and calls the treatment service to delete the treatment.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the treatment is deleted successfully, it returns a 200 OK status.
func (t *treatmentHandler) DeleteTreatment(c *gin.Context) {
	// Parse the treatment UUID from the path parameter.
	treatmentUUID, _ := uuid.Parse(c.Param("uuid"))

	// Delete the treatment record.
	statusCode, err := t.treatmentService.DeleteTreatment(treatmentUUID)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while deleting the treatment", err)
		return
	}

	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": "Treatment deleted successfully",
	})
}

// UpdateTreatment handles the HTTP request for updating a treatment.
// It parses the treatment UUID from the path parameter and binds the incoming JSON payload to the updateReq struct.
// Then it calls the treatment service to update the treatment.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the treatment is updated successfully, it returns a 200 OK status.
func (t *treatmentHandler) UpdateTreatment(c *gin.Context) {
	// Parse the treatment UUID from the path parameter.
	treatmentUUID, _ := uuid.Parse(c.Param("uuid"))

	// Bind the incoming JSON to a struct.
	updateReq := &entity.RequestUpdateTreatment{}
	if err := c.ShouldBind(updateReq); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Update the treatment record.
	statusCode, err := t.treatmentService.UpdateTreatment(treatmentUUID, updateReq)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while updating the treatment", err)
		return
	}

	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": "Treatment updated successfully",
	})
}
