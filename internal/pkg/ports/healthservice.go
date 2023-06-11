package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// HealthServiceRepository is an interface that acts as a contract for the data access layer,
// requiring implementations to provide methods for querying and modifying Health Service data.
type HealthServiceRepository interface {
	// FindByUUID retrieves a HealthService based on its UUID.
	// Returns the HealthService and an error if any occurred.
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// Create takes a new HealthService value and adds it to the data store.
	// Returns an error if the operation fails.
	Create(value interface{}) error

	// Update updates an existing HealthService value in the data store.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the database.
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	// Find retrieves all HealthService values that match the given conditions from the database.
	// Returns an error if the operation fails.
	Find(out interface{}, conditions ...interface{}) error

	// Delete removes an existing HealthService from the data store.
	// Returns an error if the operation fails.
	Delete(out interface{}) error
}

// HealthServiceService is an interface defining a contract for business logic operators related to Health Services.
// It works with the entity layer to manipulate Health Service data.
type HealthServiceService interface {
	// CreateHealthService takes a request to create a Health Service.
	// Returns the Health Service ID, status, and an error if any occurred.
	CreateHealthService(c *gin.Context, createReq *entity.RequestCreateHealthService) (string, int, error)

	// GetAllHealthServices retrieves all Health Services from the data store.
	// Returns a slice of HealthService entities and an error if any occurred.
	GetAllHealthServices() ([]*entity.HealthService, error)

	// AddRatingToHealthService adds a rating to a specific Health Service.
	// Returns the status and an error if any occurred.
	AddRatingToHealthService(rating *entity.HealthServiceRating) (int, error)
}
