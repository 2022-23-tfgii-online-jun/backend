package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MonitoringRepository defines the interface for interacting with the Monitoring data store.
// It lays down the contract for all database operations related to Monitoring data.
type MonitoringRepository interface {
	// FindByUUID finds a Monitoring by its UUID in the data store.
	// Returns the found record and an error if the operation fails.
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// Create adds a new Monitoring record to the data store.
	// Returns an error if the operation fails.
	Create(value interface{}) error

	// FindItemByIDs finds a record based on provided IDs and other parameters.
	// It's a flexible method to retrieve a record based on multiple conditions.
	// Returns an error if the operation fails.
	FindItemByIDs(firstID, secondID int, tableName string, column1Name string, column2Name string, dest interface{}) error

	// Find retrieves records that match the given conditions from the data store.
	// Returns an error if the operation fails.
	Find(out interface{}, conditions ...interface{}) error
}

// MonitoringService defines the interface for managing Monitorings in the application.
// It works with the entity layer to handle Monitoring data.
type MonitoringService interface {
	// CreateMonitoring creates a new Monitoring using the provided data and context.
	// Returns the newly created Monitoring, the HTTP status code, and an error if the operation fails.
	CreateMonitoring(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateMonitoring) (*entity.Monitoring, int, error)

	// GetAllMonitorings retrieves all Monitoring records for a given user UUID.
	// Returns the retrieved Monitorings, the HTTP status code, and an error if the operation fails.
	GetAllMonitorings(c *gin.Context, userUUID uuid.UUID) ([]*entity.Monitoring, int, error)
}
