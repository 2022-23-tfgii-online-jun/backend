package postgresql

import (
	"errors"
	"github.com/emur-uy/backend/internal/pkg/ports"
)

type mediaRepository struct {
	client *Client
}

func (r *mediaRepository) Delete(value interface{}) error {
	err := r.client.db.Delete(value).Error
	if err != nil {
		return errors.New("failed to delete record: " + err.Error())
	}
	return nil
}

func (r *mediaRepository) Find(model interface{}, dest interface{}, conditions ...interface{}) error {
	return r.client.db.Model(model).Find(dest, conditions...).Error
}

// NewMediaRepository creates a new instance of a PostgreSQL media repository.
func NewMediaRepository(client *Client) ports.MediaRepository {
	return &mediaRepository{client: client}
}

// CreateWithOmit creates a new media in the database.
func (r *mediaRepository) CreateWithOmit(omit string, value interface{}) error {
	return r.client.CreateWithOmit(omit, value)
}
