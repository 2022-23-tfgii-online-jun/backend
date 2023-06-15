package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RecipeRepository defines the interface for interacting with the Recipe data store.
// It outlines the methods needed for adding, updating, retrieving, and deleting Recipe records.
type RecipeRepository interface {
	// FindByUUID locates a Recipe in the data store by its UUID.
	// Returns the found record and an error if the operation fails.
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// FindItemByIDs locates a Recipe by the provided first and second ID in the specified table.
	// Returns an error if the operation fails.
	FindItemByIDs(firstID, secondID int, tableName string, column1Name string, column2Name string, dest interface{}) error

	// Create inserts a new Recipe record into the data store.
	// Returns an error if the operation fails.
	Create(value interface{}) error

	// CreateWithOmit creates a new Recipe record in the data store while ignoring specific fields.
	// This can be useful when certain fields should not be affected by the operation.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// Update modifies an existing Recipe record in the data store with the provided data.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first Recipe record from the data store that matches the provided conditions.
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	// Find retrieves Recipe records from the data store that match the given conditions.
	// Returns an error if the operation fails.
	Find(out interface{}, conditions ...interface{}) error

	// Delete removes a Recipe record from the data store.
	// Returns an error if the operation fails.
	Delete(out interface{}) error
}

// RecipeService defines the methods for managing Recipe data within the application.
// It handles the business logic associated with Recipe data.
type RecipeService interface {
	// CreateRecipe creates a new Recipe using the provided request data and user UUID.
	// Returns an HTTP status code and an error if the operation fails.
	CreateRecipe(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateRecipe) (*entity.Recipe, error)

	// UpdateRecipe updates an existing Recipe using the provided Recipe UUID and update request data.
	// Returns an HTTP status code and an error if the operation fails.
	UpdateRecipe(c *gin.Context, recipeUUID uuid.UUID, updateReq *entity.RequestUpdateRecipe) (int, error)

	// DeleteRecipe deletes a Recipe based on the provided Recipe UUID.
	// Returns an HTTP status code and an error if the operation fails.
	DeleteRecipe(c *gin.Context, recipeUUID uuid.UUID) (int, error)

	// GetAllRecipes retrieves all Recipe records.
	// Returns a slice of Recipes and an error if the operation fails.
	GetAllRecipes() ([]*entity.RecipeWithMediaURLs, error)

	// VoteRecipe enables users to vote for a Recipe using the provided user UUID, recipe UUID and vote value.
	// Returns an HTTP status code and an error if the operation fails.
	VoteRecipe(c *gin.Context, userUUID, recipeUUID uuid.UUID, vote int) (int, error)
}
