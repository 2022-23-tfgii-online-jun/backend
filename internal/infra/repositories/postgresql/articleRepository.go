package postgresql

import (
	"errors"

	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/google/uuid"
)

// ArticleRepository is the repository that handles all database interactions related to articles.
type ArticleRepository struct {
	client *Client
}

// NewArticleRepository creates a new instance of ArticleRepository.
func NewArticleRepository(client *Client) ports.ArticleRepository {
	return &ArticleRepository{
		client: client,
	}
}

// FindByUUID retrieves a record from the database based on the provided UUID.
func (r *ArticleRepository) FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error) {
	return r.client.FindByUUID(uuid, out)
}

// Create stores a new record in the database.
// This function creates a new record in the database using the given value and returns an error if the operation fails.
func (r *ArticleRepository) Create(value interface{}) error {
	if value == nil {
		return errors.New("input value cannot be nil")
	}
	err := r.client.db.Create(value).Error
	if err != nil {
		return errors.New("failed to create record: " + err.Error())
	}
	return nil
}

// FindItemByIDs retrieves a record from the database based on the provided IDs.
func (r *ArticleRepository) FindItemByIDs(firstID, secondID int, tableName, column1Name, column2Name string, dest interface{}) error {
	return r.client.FindItemByIDs(firstID, secondID, tableName, column1Name, column2Name, dest)
}

// First returns the first record that matches the given conditions.
func (r *ArticleRepository) First(dest interface{}, conditions ...interface{}) error {
	return r.client.First(dest, conditions...)
}

// CreateWithOmit stores a new record in the database and omits the specified columns.
func (r *ArticleRepository) CreateWithOmit(omitColumns string, value interface{}) error {
	return r.client.CreateWithOmit(omitColumns, value)
}

// Find return records that match given conditions.
func (r *ArticleRepository) Find(dest interface{}, conditions ...interface{}) error {
	return r.client.db.Model(dest).Find(dest, conditions...).Error
}

// Update updates a record in the database.
func (r *ArticleRepository) Update(value interface{}) error {
	if value == nil {
		return errors.New("input value cannot be nil")
	}
	err := r.client.db.Save(value).Error
	if err != nil {
		return errors.New("failed to update record: " + err.Error())
	}
	return nil
}

// Delete deletes a record from the database.
func (r *ArticleRepository) Delete(out interface{}) error {
	err := r.client.db.Delete(out).Error
	if err != nil {
		return errors.New("failed to delete record: " + err.Error())
	}
	return nil
}
