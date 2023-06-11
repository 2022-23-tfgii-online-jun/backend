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

type treatmentHandler struct {
	treatmentService ports.TreatmentService
}

func newHandler(treatmentService ports.TreatmentService) *treatmentHandler {
	return &treatmentHandler{
		treatmentService: treatmentService,
	}
}

// CreateTreatment handler for creating a treatment
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

// GetAllTreatments handler for getting all treatments of a user
func (t *treatmentHandler) GetAllTreatments(c *gin.Context) {
	// Get user uuid from context
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

func handleError(c *gin.Context, status int, msg string, err error) {
	log.Println(err)
	c.JSON(status, gin.H{
		"code":    status,
		"message": msg,
		"error":   err.Error(),
	})
}

// DeleteTreatment handler for deleting a treatment
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

// UpdateTreatment handler for updating a treatment
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
