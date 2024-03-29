package healthservice

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

// healthServiceHandler type contains an instance of HealthServiceService.
type healthServiceHandler struct {
	healthService ports.HealthServiceService
}

// newHandler is a constructor function for initializing healthServiceHandler with the given HealthServiceService.
// The return is a pointer to an healthServiceHandler instance.
func newHandler(healthService ports.HealthServiceService) *healthServiceHandler {
	return &healthServiceHandler{
		healthService: healthService,
	}
}

// CreateHealthService handles the HTTP request for creating a health service.
// It binds the incoming JSON payload to the reqCreate struct.
// If any error occurs during this process, it will return a 400 Bad Request status.
// If the health service is created successfully, it will return a 200 OK status with the name of the created health service.
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
// It retrieves all health services from the database.
// If any error occurs during this process, it will return a 500 Internal Server Error status.
// If the health services are retrieved successfully, it will return a 200 OK status with the retrieved health services.
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

// AddRatingToHealthService handles the HTTP request for adding a rating to a health service.
// It binds the incoming JSON payload to the req struct and validates the input parameters.
// If any error occurs during this process, it will return the corresponding status code and error message.
// If the rating is added successfully, it will return a 200 OK status.
func (h *healthServiceHandler) AddRatingToHealthService(c *gin.Context) {
	req := &entity.HealthServiceRating{}

	// Bind incoming JSON payload to the req struct.
	if err := c.ShouldBindJSON(req); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Validate input parameters.
	if req.HealthServiceID == 0 || req.Rating == 0 {
		handleError(c, http.StatusBadRequest, "Invalid input", fmt.Errorf("health service ID and rating are required"))
		return
	}

	// Call the service method to add the rating to the health service.
	status, err := h.healthService.AddRatingToHealthService(req)
	if err != nil {
		handleError(c, status, "An error occurred while adding the rating", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Rating added successfully",
	})
}

// handleError handles errors by sending an appropriate response to the client.
func handleError(c *gin.Context, status int, message string, err error) {
	// Log the error message and the error itself.
	log.Printf("[HealthServiceHandler]: %s, %v", message, err)

	// Send the JSON response with the status code and error message.
	c.JSON(status, gin.H{
		"code":    status,
		"message": message,
	})
}
