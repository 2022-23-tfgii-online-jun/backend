package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SymptomRepository defines the interface for interacting with the Symptom data store.
// It outlines the methods required for adding, updating, retrieving, and deleting Symptom records.
type SymptomRepository interface {
	// FindByUUID locates a Symptom in the data store by its UUID.
	// Returns the found record and an error if the operation fails.
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// Create inserts a new Symptom record into the data store.
	// Returns an error if the operation fails.
	Create(value interface{}) error

	// CreateWithOmit inserts a new Symptom record into the data store while ignoring specific fields.
	// This can be useful when certain fields should not be affected by the operation.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// Update modifies an existing Symptom record in the data store with the provided data.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the database.
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	// Find retrieves Symptom records from the data store that match the given conditions.
	// Returns an error if the operation fails.
	Find(out interface{}, conditions ...interface{}) error

	// Delete removes a Symptom record from the data store.
	// Returns an error if the operation fails.
	Delete(out interface{}) error
}

// SymptomService defines the methods for managing Symptom data within the application.
// It handles the business logic associated with Symptom data.
type SymptomService interface {
	// CreateSymptom creates a new Symptom using the provided request data.
	// Returns the created Symptom, an HTTP status code, and an error if the operation fails.
	CreateSymptom(c *gin.Context, createReq *entity.RequestCreateSymptom) (*entity.Symptom, int, error)

	// GetAllSymptoms retrieves all Symptom records.
	// Returns a slice of Symptoms and an error if the operation fails.
	GetAllSymptoms() ([]*entity.Symptom, error)

	// AddUserToSymptom associates a User with a Symptom based on the provided User UUID and Symptom User data.
	// Returns an HTTP status code and an error if the operation fails.
	AddUserToSymptom(userUUID uuid.UUID, symptomUser *entity.RequestCreateSymptomUser) (int, error)

	// RemoveUserFromSymptom disassociates a User from a Symptom based on the provided User UUID and Symptom User data.
	// Returns an HTTP status code and an error if the operation fails.
	RemoveUserFromSymptom(userUUID uuid.UUID, symptomUser *entity.RequestCreateSymptomUser) (int, error)

	// GetSymptomsByUser retrieves all Symptom records associated with the provided User UUID.
	// Returns a slice of Symptoms and an error if the operation fails.
	GetSymptomsByUser(userUUID uuid.UUID) ([]*entity.Symptom, error)
}
