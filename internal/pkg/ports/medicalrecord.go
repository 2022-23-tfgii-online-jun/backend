package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// MedicalRecordRepository defines the interface for interacting with the MedicalRecord data store.
// It lays down the contract for all database operations related to MedicalRecord data.
type MedicalRecordRepository interface {
	// FindByUUID finds a MedicalRecord by its UUID in the data store.
	// Returns the found record and an error if the operation fails.
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// CreateWithOmit creates a new MedicalRecord in the data store while omitting specific fields.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// Update modifies an existing MedicalRecord in the data store.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the data store.
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error
}

// MedicalRecordService defines the interface for managing MedicalRecords in the application.
// It works with the entity layer to handle MedicalRecord data.
type MedicalRecordService interface {
	// CreateMedicalRecord creates a new MedicalRecord using the provided data and context.
	// Returns the newly created MedicalRecord, the HTTP status code, and an error if the operation fails.
	CreateMedicalRecord(c *gin.Context, userUUID uuid.UUID, createReq *entity.MedicalRecord) (*entity.MedicalRecord, int, error)

	// GetMedicalRecord retrieves a MedicalRecord for a given user UUID.
	// Returns the retrieved MedicalRecord, the HTTP status code, and an error if the operation fails.
	GetMedicalRecord(c *gin.Context, userUUID uuid.UUID) (*entity.MedicalRecord, int, error)

	// UpdateMedicalRecord updates an existing MedicalRecord for a given user and MedicalRecord UUIDs using the provided data and context.
	// Returns the updated MedicalRecord, the HTTP status code, and an error if the operation fails.
	UpdateMedicalRecord(c *gin.Context, userUUID uuid.UUID, medicalRecordUUID uuid.UUID, updateReq *entity.MedicalRecord) (*entity.MedicalRecord, int, error)
}
