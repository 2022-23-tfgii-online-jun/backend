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

type monitoringHandler struct {
	monitoringService ports.MonitoringService
}

func newHandler(monitoringService ports.MonitoringService) *monitoringHandler {
	return &monitoringHandler{
		monitoringService: monitoringService,
	}
}

// CreateMonitoring handles the HTTP request for creating a monitoring.
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

// handleError handles errors by sending an appropriate response to the client.
func handleError(c *gin.Context, status int, message string, err error) {
	log.Printf("[MonitoringHandler]: %s, %v", message, err)
	c.JSON(status, gin.H{
		"code":    status,
		"message": err.Error(),
	})
}
