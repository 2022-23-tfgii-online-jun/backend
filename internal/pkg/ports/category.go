package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CategoryRepository is an interface that acts as a contract for the data access layer,
// requiring implementations to provide methods for querying and modifying category data.
type CategoryRepository interface {
	// FindByUUID retrieves a Category based on its UUID.
	// Returns the category and an error if any occurred.
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// CreateWithOmit creates a new category record while omitting specific fields.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// Update updates an existing Category value in the data store.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the database.
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	// Find retrieves all Category values that match the given conditions from the database.
	// Returns an error if the operation fails.
	Find(out interface{}, conditions ...interface{}) error

	// Delete removes an existing Category from the data store.
	// Returns an error if the operation fails.
	Delete(out interface{}) error
}

// CategoryService is an interface defining a contract for business logic operators related to Categories.
// It works with the entity layer to manipulate Category data.
type CategoryService interface {
	// CreateCategory takes a request to create a Category.
	// Returns the status and an error if any occurred.
	CreateCategory(c *gin.Context, createReq *entity.Category) (int, error)

	// UpdateCategory updates an existing Category using the provided update request.
	// Returns the status and an error if any occurred.
	UpdateCategory(categoryUUID uuid.UUID, updateReq *entity.Category) (int, error)

	// DeleteCategory removes an existing Category using the provided category UUID.
	// Returns the status and an error if any occurred.
	DeleteCategory(c *gin.Context, categoryUUID uuid.UUID) (int, error)

	// GetAllCategories retrieves all categories from the data store.
	// Returns a slice of Category entities and an error if any occurred.
	GetAllCategories() ([]*entity.Category, error)
}
