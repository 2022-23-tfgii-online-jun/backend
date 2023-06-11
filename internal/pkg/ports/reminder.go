package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ReminderRepository defines the interface for interacting with the Reminder data store.
// It outlines the methods required for adding, updating, retrieving, and deleting Reminder records.
type ReminderRepository interface {
	// FindByUUID locates a Reminder in the data store by its UUID.
	// Returns the found record and an error if the operation fails.
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// CreateWithOmit inserts a new Reminder record into the data store while ignoring specific fields.
	// This can be useful when certain fields should not be affected by the operation.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// Find retrieves Reminder records from the data store that match the given conditions.
	// Returns an error if the operation fails.
	Find(model interface{}, dest interface{}, conditions ...interface{}) error

	// Update modifies an existing Reminder record in the data store with the provided data.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// Delete removes a Reminder record from the data store.
	// Returns an error if the operation fails.
	Delete(out interface{}) error
}

// ReminderService defines the methods for managing Reminder data within the application.
// It handles the business logic associated with Reminder data.
type ReminderService interface {
	// CreateReminder creates a new Reminder using the provided request data and user UUID.
	// Returns an HTTP status code and an error if the operation fails.
	CreateReminder(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateReminder) (int, error)

	// GetAllReminders retrieves all Reminder records for the given user UUID.
	// Returns a slice of Reminders and an error if the operation fails.
	GetAllReminders(c *gin.Context, userUUID uuid.UUID) ([]*entity.GetReminderResponse, error)

	// UpdateReminder updates an existing Reminder using the provided Reminder UUID and update request data.
	// Returns an HTTP status code and an error if the operation fails.
	UpdateReminder(c *gin.Context, reminderUUID uuid.UUID, updateReq *entity.RequestUpdateReminder) (int, error)

	// DeleteReminder deletes a Reminder based on the provided Reminder UUID.
	// Returns an error if the operation fails.
	DeleteReminder(c *gin.Context, reminderUUID uuid.UUID) error
}
