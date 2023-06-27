package postgresql

import (
	"errors"

	"github.com/emur-uy/backend/internal/pkg/ports"
)

type articleMediaRepository struct {
	client *Client
}

func (r *articleMediaRepository) Delete(value interface{}) error {
	err := r.client.db.Delete(value).Error
	if err != nil {
		return errors.New("failed to delete record: " + err.Error())
	}
	return nil
}

func (r *articleMediaRepository) Find(model interface{}, dest interface{}, conditions ...interface{}) error {
	return r.client.db.Model(model).Find(dest, conditions...).Error
}

// NewArticleMediaRepository creates a new instance of a PostgreSQL articleMedia repository.
func NewArticleMediaRepository(client *Client) ports.ArticleMediaRepository {
	return &articleMediaRepository{client: client}
}

// Create creates a new articleMedia in the database.
func (r *articleMediaRepository) Create(value interface{}) error {
	return r.client.Create(value)
}
