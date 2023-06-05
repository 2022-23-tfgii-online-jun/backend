package healthservice

import (
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type healthServiceHandler struct {
	healthService ports.HealthServiceService
}

func newHandler(healthService ports.HealthServiceService) *healthServiceHandler {
	return &healthServiceHandler{
		healthService: healthService,
	}
}

// CreateHealthService handles the HTTP request for creating a health service.
func (h *healthServiceHandler) CreateHealthService(c *gin.Context) {
	reqCreate := &entity.RequestCreateHealthService{}

	// Bind incoming JSON payload to the reqCreate struct.
	if err := c.ShouldBindJSON(reqCreate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Create the health service and store it in the database.
	createdHealthService, statusCode, err := h.healthService.CreateHealthService(c, reqCreate)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while creating the health service", err)
		return
	}

	// Return a successful response with the name of the created health service.
	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"message": "Health service created successfully",
		"data": gin.H{
			"name": createdHealthService,
		},
	})
}

// GetAllHealthServices handles the HTTP request for getting all health services.
func (h *healthServiceHandler) GetAllHealthServices(c *gin.Context) {

	// Get all health services from the database.
	healthServices, err := h.healthService.GetAllHealthServices()
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while getting the health services", err)
		return
	}

	// Return a successful response with the retrieved health services.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Health services retrieved successfully",
		"data":    healthServices,
	})
}

// handleError handles errors by sending an appropriate response to the client.
func handleError(c *gin.Context, status int, message string, err error) {
	// Log the error message and the error itself
	log.Printf("[HealthServiceHandler]: %s, %v", message, err)

	// Send the JSON response with the status code and error message
	c.JSON(status, gin.H{
		"code":    status,
		"message": message,
	})
}
