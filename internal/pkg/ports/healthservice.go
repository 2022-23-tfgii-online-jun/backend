package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// HealthServiceRepository is the interface that defines the methods for accessing the health service data store.
type HealthServiceRepository interface {
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// CreateWithOmit creates a new health service record while omitting specific fields.
	// Returns an error if the operation fails.
	Create(value interface{}) error

	// UpdateHealthService updates an existing health service record with the provided health service data.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the database
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	Find(out interface{}, conditions ...interface{}) error

	Delete(out interface{}) error
}

// HealthServiceService is the interface that defines the methods for managing health services in the application.
type HealthServiceService interface {
	CreateHealthService(c *gin.Context, createReq *entity.RequestCreateHealthService) (string, int, error)
	GetAllHealthServices() ([]*entity.HealthService, error)
	AddRatingToHealthService(rating *entity.HealthServiceRating) (int, error)
}
