package postgresql

import (
	"github.com/emur-uy/backend/internal/pkg/ports"
)

type mediaRepository struct {
	client *Client
}

// NewMediaRepository creates a new instance of a PostgreSQL media repository.
func NewMediaRepository(client *Client) ports.MediaRepository {
	return &mediaRepository{client: client}
}

// CreateWithOmit creates a new media in the database.
func (r *mediaRepository) CreateWithOmit(omit string, value interface{}) error {
	return r.client.CreateWithOmit(omit, value)
}
