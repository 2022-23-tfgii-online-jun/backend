package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CategoryRepository is the interface that defines the methods for accessing the category data store.
type CategoryRepository interface {
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// CreateWithOmit creates a new category record while omitting specific fields.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// UpdateCategory updates an existing category record with the provided category data.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the database
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	Find(out interface{}, conditions ...interface{}) error

	Delete(out interface{}) error
}

// CategoryService is the interface that defines the methods for managing categories in the application.
type CategoryService interface {
	CreateCategory(c *gin.Context, createReq *entity.Category) (int, error)
	UpdateCategory(categoryUUID uuid.UUID, updateReq *entity.Category) (int, error)
	DeleteCategory(c *gin.Context, categoryUUID uuid.UUID) (int, error)
	GetAllCategories() ([]*entity.Category, error)
}
