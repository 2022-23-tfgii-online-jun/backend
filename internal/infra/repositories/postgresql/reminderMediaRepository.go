package postgresql

import (
	"errors"
	"github.com/emur-uy/backend/internal/pkg/ports"
)

type reminderMediaRepository struct {
	client *Client
}

func (r *reminderMediaRepository) Delete(value interface{}) error {
	err := r.client.db.Delete(value).Error
	if err != nil {
		return errors.New("failed to delete record: " + err.Error())
	}
	return nil
}

func (r *reminderMediaRepository) Find(model interface{}, dest interface{}, conditions ...interface{}) error {
	return r.client.db.Model(model).Find(dest, conditions...).Error
}

// NewReminderMediaRepository creates a new instance of a PostgreSQL reminderMedia repository.
func NewReminderMediaRepository(client *Client) ports.ReminderMediaRepository {
	return &reminderMediaRepository{client: client}
}

// Create creates a new reminderMedia in the database.
func (r *reminderMediaRepository) Create(value interface{}) error {
	return r.client.Create(value)
}
