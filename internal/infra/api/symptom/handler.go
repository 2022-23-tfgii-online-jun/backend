package symptom

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// symptomHandler type contains an instance of SymptomService
type symptomHandler struct {
	symptomService ports.SymptomService
}

// newHandler is a constructor function for initializing symptomHandler with the given SymptomService.
// The return is a pointer to a symptomHandler instance.
func newHandler(symptomService ports.SymptomService) *symptomHandler {
	return &symptomHandler{
		symptomService: symptomService,
	}
}

// CreateSymptom handles the HTTP request for creating a symptom.
// It parses the incoming JSON payload, binds it to the reqCreate struct,
// and calls the symptom service to create the symptom.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the symptom is created successfully, it returns a 200 OK status with the created symptom.
func (h *symptomHandler) CreateSymptom(c *gin.Context) {
	reqCreate := &entity.RequestCreateSymptom{}

	if err := c.ShouldBindJSON(reqCreate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	createdSymptom, statusCode, err := h.symptomService.CreateSymptom(c, reqCreate)
	if err != nil {
		handleError(c, statusCode, "An error occurred while creating the symptom", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"message": "Symptom created successfully",
		"data": gin.H{
			"name": createdSymptom,
		},
	})
}

// GetAllSymptoms handles the HTTP request for getting all symptoms.
// It retrieves all symptoms from the service.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the symptoms are fetched successfully, it returns a 200 OK status with the retrieved symptoms.
func (h *symptomHandler) GetAllSymptoms(c *gin.Context) {
	symptoms, err := h.symptomService.GetAllSymptoms()
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while getting the symptoms", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Symptoms retrieved successfully",
		"data":    symptoms,
	})
}

// AddUserToSymptom handles the HTTP request for adding a user to a symptom.
// It parses the user UUID from the context and the symptom UUID from the request JSON payload,
// and calls the symptom service to add the user to the symptom.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the user is added to the symptom successfully, it returns a 200 OK status.
func (h *symptomHandler) AddUserToSymptom(c *gin.Context) {
	req := &entity.RequestCreateSymptomUser{}

	userUUID, err := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid user UUID", err)
		return
	}

	if err := c.ShouldBindJSON(req); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	if req.SymptomUUID == uuid.Nil {
		handleError(c, http.StatusBadRequest, "Invalid input", fmt.Errorf("symptom UUID is required"))
		return
	}

	status, err := h.symptomService.AddUserToSymptom(userUUID, req)
	if err != nil {
		handleError(c, status, "An error occurred while adding the user to the symptom", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "User added to symptom successfully",
	})
}

// RemoveUserFromSymptom handles the HTTP request for removing a user from a symptom.
// It parses the user UUID from the context and the symptom UUID from the request JSON payload,
// and calls the symptom service to remove the user from the symptom.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the user is removed from the symptom successfully, it returns a 200 OK status.
func (h *symptomHandler) RemoveUserFromSymptom(c *gin.Context) {
	req := &entity.RequestCreateSymptomUser{}

	userUUID, err := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid user UUID", err)
		return
	}

	if err := c.ShouldBindJSON(req); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	if req.SymptomUUID == uuid.Nil {
		handleError(c, http.StatusBadRequest, "Invalid input", fmt.Errorf("symptom UUID is required"))
		return
	}

	status, err := h.symptomService.RemoveUserFromSymptom(userUUID, req)
	if err != nil {
		handleError(c, status, "An error occurred while removing the user from the symptom", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "User removed from symptom successfully",
	})
}

// GetSymptomsByUser handles the HTTP request for getting all symptoms related to a user.
// It parses the user UUID from the context and calls the symptom service to get all symptoms related to the user.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the symptoms are fetched successfully, it returns a 200 OK status with the retrieved symptoms.
func (h *symptomHandler) GetSymptomsByUser(c *gin.Context) {
	userUUID, err := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid user UUID", err)
		return
	}

	// Call the service method to get all symptoms related to the user
	symptoms, err := h.symptomService.GetSymptomsByUser(userUUID)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while getting symptoms", err)
		return
	}

	// Return the symptoms in the HTTP response
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Symptoms retrieved successfully",
		"data":    symptoms,
	})
}

// handleError is a generic error handler that logs the error and responds.
func handleError(c *gin.Context, status int, message string, err error) {
	log.Printf("[SymptomHandler]: %s, %v", message, err)
	c.JSON(status, gin.H{
		"code":    status,
		"message": message,
	})
}
