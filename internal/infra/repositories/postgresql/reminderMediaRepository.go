package postgresql

import (
	"github.com/emur-uy/backend/internal/pkg/ports"
)

type reminderMediaRepository struct {
	client *Client
}

// NewReminderMediaRepository creates a new instance of a PostgreSQL reminderMedia repository.
func NewReminderMediaRepository(client *Client) ports.ReminderMediaRepository {
	return &reminderMediaRepository{client: client}
}

// Create creates a new reminderMedia in the database.
func (r *reminderMediaRepository) Create(value interface{}) error {
	return r.client.Create(value)
}
