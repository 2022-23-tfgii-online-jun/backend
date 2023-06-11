package maps

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type mapHandler struct {
	mapService ports.MapService
}

func newHandler(mapService ports.MapService) *mapHandler {
	return &mapHandler{
		mapService: mapService,
	}
}

// CreateMap handler for creating a map
func (m *mapHandler) CreateMap(c *gin.Context) {
	reqCreate := &entity.RequestCreateUpdateMap{}

	// Bind incoming JSON payload to the reqCreate struct.
	if err := c.ShouldBindJSON(reqCreate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Create the map and store it in the database.
	createdMap, statusCode, err := m.mapService.CreateMap(c, reqCreate)
	if err != nil {
		handleError(c, statusCode, "An error occurred while creating the map", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Map created successfully",
		"data":    createdMap,
	})
}

// GetAllMaps handles the HTTP request for getting all maps.
func (m *mapHandler) GetAllMaps(c *gin.Context) {
	// Get all maps from the database.
	maps, err := m.mapService.GetAllMaps()
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while getting the maps", err)
		return
	}

	// Return a successful response with the retrieved maps.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Maps retrieved successfully",
		"data":    maps,
	})
}

// UpdateMap handler for updating a map
func (m *mapHandler) UpdateMap(c *gin.Context) {
	// Parse the map UUID from the URL parameter.
	mapUUID, err := uuid.Parse(fmt.Sprintf("%v", c.Param("uuid")))

	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid UUID format", err)
		return
	}

	// Bind the incoming JSON payload to an UpdateMap struct.
	reqUpdate := &entity.RequestCreateUpdateMap{}
	if err := c.ShouldBind(reqUpdate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Update the map in the database.
	updatedMap, statusCode, err := m.mapService.UpdateMap(c, mapUUID, reqUpdate)
	if err != nil {
		handleError(c, statusCode, "An error occurred while updating the map", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Map updated successfully",
		"data":    updatedMap,
	})
}

// DeleteMap handler for deleting a map
func (m *mapHandler) DeleteMap(c *gin.Context) {
	// Parse the map UUID from the URL parameter.
	mapUUID, err := uuid.Parse(fmt.Sprintf("%v", c.Param("uuid")))

	// Delete the map from the database.
	statusCode, err := m.mapService.DeleteMap(c, mapUUID)
	if err != nil {
		handleError(c, statusCode, "An error occurred while deleting the map", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Map deleted successfully",
	})
}

// handleError is a generic error handler that logs the error and responds
func handleError(c *gin.Context, statusCode int, message string, err error) {
	// Log the error message and the error itself
	log.Printf("[MapHandler]: %s, %v", message, err)

	// Send the JSON response with the status code and error message
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
		"data":    nil,
	})
}
