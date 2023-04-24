package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RecipeRepository is the interface that defines the methods for accessing the recipe data store.
type RecipeRepository interface {
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)
	FindItemByIDs(firstID, secondID int, tableName string, column1Name string, column2Name string, dest interface{}) error

	Create(value interface{}) error

	// CreateWithOmit creates a new recipe record while omitting specific fields.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// UpdateRecipe updates an existing recipe record with the provided recipe data.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the database
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	Find(out interface{}, conditions ...interface{}) error

	Delete(out interface{}) error
}

// RecipeService is the interface that defines the methods for managing recipes in the application.
type RecipeService interface {
	CreateRecipe(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateRecipe) (int, error)
	UpdateRecipe(recipeUUID uuid.UUID, updateReq *entity.RequestUpdateRecipe) (int, error)
	DeleteRecipe(c *gin.Context, recipeUUID uuid.UUID) (int, error)
	GetAllRecipes() ([]*entity.Recipe, error)
	VoteRecipe(c *gin.Context, userUUID, recipeUUID uuid.UUID, vote int) (int, error)
}
