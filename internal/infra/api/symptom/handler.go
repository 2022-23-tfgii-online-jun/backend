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

// handleError handles errors by sending an appropriate response to the client.
func handleError(c *gin.Context, status int, message string, err error) {
	log.Printf("[SymptomHandler]: %s, %v", message, err)
	c.JSON(status, gin.H{
		"code":    status,
		"message": message,
	})
}
