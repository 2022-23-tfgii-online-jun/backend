package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
)

// MedicalRepository is the interface that defines the methods for accessing the medical data store.
type MedicalRepository interface {
	Create(value interface{}) error

	// Update updates an existing medical record with the provided data.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	Find(out interface{}, conditions ...interface{}) error
}

// MedicalService is the interface that defines the methods for managing medical records in the application.
type MedicalService interface {
	CreateRecordFromFile(c *gin.Context) (int, error)
	GetAllMedicalRecords() ([]*entity.Medical, error)
	AddRatingToMedical(rating *entity.MedicalRating) (int, error)
}
