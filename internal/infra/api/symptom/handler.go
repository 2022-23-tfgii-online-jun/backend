package symptom

import (
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type symptomHandler struct {
	symptomService ports.SymptomService
}

func newHandler(symptomService ports.SymptomService) *symptomHandler {
	return &symptomHandler{
		symptomService: symptomService,
	}
}

// CreateSymptom handles the HTTP request for creating a symptom.
func (h *symptomHandler) CreateSymptom(c *gin.Context) {
	reqCreate := &entity.RequestCreateSymptom{}

	// Bind incoming JSON payload to the reqCreate struct.
	if err := c.ShouldBindJSON(reqCreate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Create the symptom and store it in the database.
	createdSymptom, statusCode, err := h.symptomService.CreateSymptom(c, reqCreate)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while creating the symptom", err)
		return
	}

	// Return a successful response with the name of the created symptom.
	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"message": "Symptom created successfully",
		"data": gin.H{
			"name": createdSymptom,
		},
	})
}

// GetAllSymptoms handles the HTTP request for getting all symptoms.
func (h *symptomHandler) GetAllSymptoms(c *gin.Context) {

	// Get all symptoms from the database.
	symptoms, err := h.symptomService.GetAllSymptoms()
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while getting the symptoms", err)
		return
	}

	// Return a successful response with the retrieved symptoms.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Symptoms retrieved successfully",
		"data":    symptoms,
	})
}

// handleError handles errors by sending an appropriate response to the client.
func handleError(c *gin.Context, status int, message string, err error) {
	// Log the error message and the error itself
	log.Printf("[SymptomHandler]: %s, %v", message, err)

	// Send the JSON response with the status code and error message
	c.JSON(status, gin.H{
		"code":    status,
		"message": message,
	})
}
