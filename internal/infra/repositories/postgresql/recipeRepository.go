package postgresql

import (
	"errors"

	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/google/uuid"
)

// RecipeRepository is the repository that handles all database interactions related to recipes
type RecipeRepository struct {
	client *Client
}

// NewRecipeRepository creates a new instance of RecipeRepository
func NewRecipeRepository(client *Client) ports.RecipeRepository {
	return &RecipeRepository{
		client: client,
	}
}

// FindByUUID retrieves a record from the database based on the provided UUID
func (r *RecipeRepository) FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error) {
	return r.client.FindByUUID(uuid, out)
}

// Create stores a new record in the database.
// This function creates a new record in the database using the given value and returns an error if the operation fails.
func (r *RecipeRepository) Create(value interface{}) error {
	if value == nil {
		return errors.New("input value cannot be nil")
	}
	err := r.client.db.Create(value).Error
	if err != nil {
		return errors.New("failed to create record: " + err.Error())
	}
	return nil
}

// FindItemByIDs retrieves a record from the database based on the provided IDs
func (r *RecipeRepository) FindItemByIDs(firstID, secondID int, tableName, column1Name, column2Name string, dest interface{}) error {
	return r.client.FindItemByIDs(firstID, secondID, tableName, column1Name, column2Name, dest)
}

// First returns the first record that matches the given conditions.
func (r *RecipeRepository) First(dest interface{}, conditions ...interface{}) error {
	return r.client.First(dest, conditions...)
}

// CreateWithOmit stores a new record in the database and omits the specified columns
func (r *RecipeRepository) CreateWithOmit(omitColumns string, value interface{}) error {
	return r.client.CreateWithOmit(omitColumns, value)
}

// Find return records that match given conditions.
func (r *RecipeRepository) Find(dest interface{}, conditions ...interface{}) error {
	return r.client.db.Model(dest).Find(dest, conditions...).Error
}
func (r *RecipeRepository) Update(value interface{}) error {
	if value == nil {
		return errors.New("input value cannot be nil")
	}
	err := r.client.db.Save(value).Error
	if err != nil {
		return errors.New("failed to update record: " + err.Error())
	}
	return nil
}

func (r *RecipeRepository) Delete(out interface{}) error {
	err := r.client.db.Delete(out).Error
	if err != nil {
		return errors.New("failed to delete record: " + err.Error())
	}
	return nil
}
