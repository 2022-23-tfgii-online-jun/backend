package recipe

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	ErrTypeAssertionFailed = errors.New("type assertion failed")
	ErrCreatingRecipe      = errors.New("error creating recipe")
	ErrUpdatingRecipe      = errors.New("error updating recipe")
	ErrDeletingRecipe      = errors.New("error deleting recipe")
)

type service struct {
	repo ports.RecipeRepository
}

// NewService returns a new instance of the recipe service with the given recipe repository.
func NewService(recipeRepo ports.RecipeRepository) ports.RecipeService {
	return &service{
		repo: recipeRepo,
	}
}

// CreateRecipe is the service for creating a recipe and saving it in the database.
func (s *service) CreateRecipe(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateRecipe) (*entity.Recipe, error) {
	user := &entity.User{}

	// Find user by UUID
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		return nil, err
	}

	// Perform type assertion to convert foundUser to *entity.User
	user, ok := foundUser.(*entity.User)
	if !ok {
		return nil, ErrTypeAssertionFailed
	}

	// Create a new recipe
	recipe := &entity.Recipe{
		Name:        createReq.Name,
		Ingredients: createReq.Ingredients,
		Elaboration: createReq.Elaboration,
		Time:        createReq.Time,
		UserID:      user.ID,
		Category:    createReq.Category,
	}

	// Save the recipe to the database
	err = s.repo.CreateWithOmit("uuid", recipe)
	if err != nil {
		return nil, ErrCreatingRecipe
	}

	return recipe, nil
}

// GetAllRecipes returns all recipes stored in the database.
func (s *service) GetAllRecipes() ([]*entity.Recipe, error) {
	// Get all recipes from the database
	var recipes []*entity.Recipe
	if err := s.repo.Find(&recipes); err != nil {
		return nil, err
	}

	return recipes, nil
}

// UpdateRecipe is the service for updating a recipe in the database.
func (s *service) UpdateRecipe(recipeUUID uuid.UUID, updateReq *entity.RequestUpdateRecipe) (int, error) {
	// Find the existing recipe by UUID
	recipe := &entity.Recipe{}
	foundRecipe, err := s.repo.FindByUUID(recipeUUID, recipe)
	if err != nil {
		// Return error if the recipe is not found
		return http.StatusNotFound, err
	}

	// Perform type assertion to convert foundRecipe to *entity.Recipe
	recipe, ok := foundRecipe.(*entity.Recipe)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	// Update the recipe fields with the new data from the update request
	recipe.Name = updateReq.Name

	// Update the recipe in the database
	err = s.repo.Update(recipe)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error updating recipe: %s", err)
	}

	// Return the HTTP OK status code if the update is successful
	return http.StatusOK, nil
}

// DeleteRecipe deletes a recipe from the database by its UUID.
func (s *service) DeleteRecipe(c *gin.Context, recipeUUID uuid.UUID) (int, error) {
	// Retrieve the recipe from the repository by its UUID.
	recipe := &entity.Recipe{}
	foundRecipe, err := s.repo.FindByUUID(recipeUUID, recipe)
	if err != nil {
		// Return an error response if the recipe is not found.
		return http.StatusNotFound, err
	}

	// Perform type assertion to convert foundRecipe to *entity.Recipe.
	recipe, ok := foundRecipe.(*entity.Recipe)
	if !ok {
		return http.StatusInternalServerError, ErrTypeAssertionFailed
	}

	// Delete the recipe from the repository.
	err = s.repo.Delete(recipe)
	if err != nil {
		// Return an error response if there was an issue deleting the recipe.
		return http.StatusInternalServerError, ErrDeletingRecipe
	}

	// Return a success response.
	return http.StatusOK, nil
}

// VoteRecipe is the service for voting a recipe in the database.
func (s *service) VoteRecipe(c *gin.Context, userUUID uuid.UUID, recipeUUID uuid.UUID, vote int) (int, error) {
	if vote < 1 || vote > 5 {
		return http.StatusBadRequest, errors.New("invalid vote value, must be between 1 and 5")
	}

	// Get user by UUID
	user := &entity.User{}
	_, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		return http.StatusNotFound, fmt.Errorf("error finding user: %s", err)
	}

	// Get recipe by UUID
	recipe := &entity.Recipe{}
	_, err = s.repo.FindByUUID(recipeUUID, recipe)
	if err != nil {
		return http.StatusNotFound, fmt.Errorf("error finding recipe: %s", err)
	}

	type RatingRecipe struct {
		ID       int
		UserID   int
		RecipeID int
		Level    int
	}

	// Check if the user has already voted for this recipe
	existingVote := &RatingRecipe{}
	err = s.repo.FindItemByIDs(user.ID, recipe.ID, "rating_recipes", "user_id", "recipe_id", existingVote)
	if err == nil {
		// Update existing vote value
		existingVote.Level = vote
		return http.StatusOK, s.repo.Update(existingVote)
	}

	// Create a new vote record
	recipeVote := &RatingRecipe{
		UserID:   user.ID,
		RecipeID: recipe.ID,
		Level:    vote,
	}
	return http.StatusOK, s.repo.Create(recipeVote)
}
