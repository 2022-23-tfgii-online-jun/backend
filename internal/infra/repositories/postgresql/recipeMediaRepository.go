package postgresql

import (
	"errors"

	"github.com/emur-uy/backend/internal/pkg/ports"
)

type recipeMediaRepository struct {
	client *Client
}

func (r *recipeMediaRepository) Delete(value interface{}) error {
	err := r.client.db.Delete(value).Error
	if err != nil {
		return errors.New("failed to delete record: " + err.Error())
	}
	return nil
}

func (r *recipeMediaRepository) Find(model interface{}, dest interface{}, conditions ...interface{}) error {
	return r.client.db.Model(model).Find(dest, conditions...).Error
}

// NewRecipeMediaRepository creates a new instance of a PostgreSQL recipeMedia repository.
func NewRecipeMediaRepository(client *Client) ports.RecipeMediaRepository {
	return &recipeMediaRepository{client: client}
}

// Create creates a new recipeMedia in the database.
func (r *recipeMediaRepository) Create(value interface{}) error {
	return r.client.Create(value)
}
