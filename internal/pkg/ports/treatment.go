package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TreatmentRepository is the interface that defines the methods for accessing the treatment data store.
// This allows you to perform operations like find, create, update, and delete on treatment records.
type TreatmentRepository interface {
	// FindByUUID retrieves a treatment from the data store using its UUID.
	// Returns the found treatment and an error if the operation fails.
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// FindItemByIDs retrieves a treatment from the data store using two ID fields.
	// Useful when the data store's structure has composite keys or relations.
	FindItemByIDs(firstID, secondID int, tableName string, column1Name string, column2Name string, dest interface{}) error

	// Create inserts a new treatment record into the data store.
	// Returns an error if the operation fails.
	Create(value interface{}) error

	// CreateWithOmit inserts a new treatment record into the data store while ignoring specific fields.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// Update modifies an existing treatment record in the data store.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the data store.
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	// Find retrieves treatment records from the data store that match the given conditions.
	// Returns an error if the operation fails.
	Find(out interface{}, conditions ...interface{}) error

	// Delete removes a treatment record from the data store.
	// Returns an error if the operation fails.
	Delete(out interface{}) error
}

// TreatmentService is the interface that defines the methods for managing treatments in the application.
// This includes operations like creating, updating, deleting, and retrieving treatments.
type TreatmentService interface {
	// CreateTreatment creates a new treatment record in the application using the provided request data.
	// Returns the created treatment, a status code, and an error if the operation fails.
	CreateTreatment(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateTreatment) (*entity.Treatment, int, error)

	// UpdateTreatment updates an existing treatment record in the application.
	// Returns a status code and an error if the operation fails.
	UpdateTreatment(treatmentUUID uuid.UUID, updateReq *entity.RequestUpdateTreatment) (int, error)

	// DeleteTreatment removes an existing treatment record from the application.
	// Returns a status code and an error if the operation fails.
	DeleteTreatment(treatmentUUID uuid.UUID) (int, error)

	// GetAllTreatments retrieves all treatment records for a specific user from the application.
	// Returns a list of treatments and an error if the operation fails.
	GetAllTreatments(userUUID uuid.UUID) ([]*entity.Treatment, error)
}
