package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
)

// MedicalRepository defines the interface for interacting with the medical data store.
// It lays down the contract for all database operations related to Medical data.
type MedicalRepository interface {
	// Create creates a new Medical record in the data store.
	// Returns an error if the operation fails.
	Create(value interface{}) error

	// Update modifies an existing Medical record with the provided data.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// Find retrieves all Medical records that match the given conditions.
	// Returns an error if the operation fails.
	Find(out interface{}, conditions ...interface{}) error
}

// MedicalService is the interface that defines the methods for managing medical records in the application.
// It works with the entity layer to handle Medical data.
type MedicalService interface {
	// CreateRecordFromFile creates a new medical record using a file provided in the HTTP request context.
	// It returns the HTTP status code and an error if the operation fails.
	CreateRecordFromFile(c *gin.Context) (int, error)

	// GetAllMedicalRecords retrieves all Medical records in the application.
	// It returns a slice of Medical entities and an error if the operation fails.
	GetAllMedicalRecords() ([]*entity.Medical, error)

	// AddRatingToMedical adds a new rating to a Medical record.
	// It returns the HTTP status code and an error if the operation fails.
	AddRatingToMedical(rating *entity.MedicalRating) (int, error)
}
