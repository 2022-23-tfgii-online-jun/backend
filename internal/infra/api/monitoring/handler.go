package monitoring

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// monitoringHandler type contains an instance of MonitoringService
type monitoringHandler struct {
	monitoringService ports.MonitoringService
}

// newHandler is a constructor function for initializing monitoringHandler with the given MonitoringService.
// The return is a pointer to an monitoringHandler instance.
func newHandler(monitoringService ports.MonitoringService) *monitoringHandler {
	return &monitoringHandler{
		monitoringService: monitoringService,
	}
}

// CreateMonitoring handles the HTTP request for creating a monitoring.
// It binds the incoming JSON payload to the reqCreate struct and calls the monitoring service to create the monitoring.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the monitoring is created successfully, it returns a 200 OK status with the created monitoring.
func (h *monitoringHandler) CreateMonitoring(c *gin.Context) {
	reqCreate := &entity.RequestCreateMonitoring{}

	// Get user UUID from JWT token
	userUUID, err := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid user UUID", err)
		return
	}

	// Bind request JSON to struct
	if err := c.ShouldBindJSON(reqCreate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Call the monitoring service to create the monitoring
	createdMonitoring, statusCode, err := h.monitoringService.CreateMonitoring(c, userUUID, reqCreate)
	if err != nil {
		handleError(c, statusCode, "An error occurred while creating the monitoring", err)
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"message": "Monitoring created successfully",
		"data": gin.H{
			"monitoring": createdMonitoring,
		},
	})
}

// GetAllMonitorings handles the HTTP request for getting all monitorings.
// It retrieves all monitorings from the database.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the monitorings are retrieved successfully, it returns a 200 OK status with the retrieved monitorings.
func (h *monitoringHandler) GetAllMonitorings(c *gin.Context) {
	// Get user UUID from JWT token
	userUUID, err := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid user UUID", err)
		return
	}

	// Get all monitorings for the user from the database
	monitorings, statusCode, err := h.monitoringService.GetAllMonitorings(c, userUUID)
	if err != nil {
		handleError(c, statusCode, "An error occurred while getting the monitorings", err)
		return
	}

	// Return the response with the retrieved monitorings
	c.JSON(http.StatusOK, gin.H{
		"code":    statusCode,
		"message": "Monitorings retrieved successfully",
		"data":    monitorings,
	})
}

// handleError handles errors by sending an appropriate response to the client.
// It takes the gin.Context, status code, error message, and error as parameters.
func handleError(c *gin.Context, status int, message string, err error) {
	log.Printf("[MonitoringHandler]: %s, %v", message, err)
	c.JSON(status, gin.H{
		"code":    status,
		"message": err.Error(),
	})
}
