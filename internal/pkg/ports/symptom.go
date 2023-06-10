package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SymptomRepository is the interface that defines the methods for accessing the symptom data store.
type SymptomRepository interface {
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// Create creates a new symptom record.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// UpdateSymptom updates an existing symptom record with the provided symptom data.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the database
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	Find(out interface{}, conditions ...interface{}) error

	Delete(out interface{}) error
}

// SymptomService is the interface that defines the methods for managing symptoms in the application.
type SymptomService interface {
	CreateSymptom(c *gin.Context, createReq *entity.RequestCreateSymptom) (*entity.Symptom, int, error)
	GetAllSymptoms() ([]*entity.Symptom, error)
}
