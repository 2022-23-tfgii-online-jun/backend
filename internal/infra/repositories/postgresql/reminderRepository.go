package postgresql

import (
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/google/uuid"
)

// ReminderRepository is the repository that handles all database interactions related to reminders
type ReminderRepository struct {
	client *Client
}

// NewReminderRepository creates a new instance of ReminderRepository
func NewReminderRepository(client *Client) ports.ReminderRepository {
	return &ReminderRepository{
		client: client,
	}
}

// FindByUUID retrieves a record from the database based on the provided UUID
func (r *ReminderRepository) FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error) {
	return r.client.FindByUUID(uuid, out)
}

// CreateWithOmit stores a new record in the database and omits the specified columns
func (r *ReminderRepository) CreateWithOmit(omitColumns string, value interface{}) error {
	return r.client.CreateWithOmit(omitColumns, value)
}
