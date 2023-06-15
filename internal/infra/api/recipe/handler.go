package recipe

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// recipeHandler type contains an instance of RecipeService
type recipeHandler struct {
	recipeService ports.RecipeService
}

// newHandler is a constructor function for initializing recipeHandler with the given RecipeService.
// The return is a pointer to an recipeHandler instance.
func newHandler(recipeService ports.RecipeService) *recipeHandler {
	return &recipeHandler{
		recipeService: recipeService,
	}
}

// CreateRecipe handles the HTTP request for creating a recipe.
// It binds the incoming form-data payload to the reqCreate struct and calls the recipe service to create the recipe.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the recipe is created successfully, it returns a 200 OK status with the created recipe.
func (r *recipeHandler) CreateRecipe(c *gin.Context) {
	reqCreate := &entity.RequestCreateRecipe{}

	// Get user UUID from context
	userUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))

	// Bind incoming form-data payload to the reqCreate struct.
	if err := c.ShouldBind(reqCreate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Create the recipe and store it in the database.
	createdRecipe, err := r.recipeService.CreateRecipe(c, userUUID, reqCreate)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while creating the recipe", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Recipe created successfully",
		"data":    createdRecipe,
	})
}

// GetAllRecipes handles the HTTP request for getting all recipes.
// It retrieves all recipes from the database.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the recipes are retrieved successfully, it returns a 200 OK status with the retrieved recipes.
func (r *recipeHandler) GetAllRecipes(c *gin.Context) {
	// Get all recipes from the database.
	recipes, err := r.recipeService.GetAllRecipes()
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while getting the recipes", err)
		return
	}

	// Return a successful response with the retrieved recipes.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Recipes retrieved successfully",
		"data":    recipes,
	})
}

// UpdateRecipe handles the HTTP request for updating a recipe.
// It parses the recipe UUID from the URL parameter, binds the incoming JSON payload to an UpdateRecipe struct,
// and calls the recipe service to update the recipe in the database.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the recipe is updated successfully, it returns a 200 OK status with the updated recipe.
func (r *recipeHandler) UpdateRecipe(c *gin.Context) {
	// Parse the recipe UUID from the URL parameter.
	recipeUUID, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid UUID format", err)
		return
	}

	// Bind the incoming JSON payload to an UpdateRecipe struct.
	reqUpdate := &entity.RequestUpdateRecipe{}
	if err := c.ShouldBind(reqUpdate); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Update the recipe in the database.
	updatedRecipe, err := r.recipeService.UpdateRecipe(c, recipeUUID, reqUpdate)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "An error occurred while updating the recipe", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Recipe updated successfully",
		"data":    updatedRecipe,
	})
}

// DeleteRecipe handles the HTTP request for deleting a recipe.
// It gets the recipe UUID from the path parameter, calls the recipe service to delete the recipe from the database.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the recipe is deleted successfully, it returns a 200 OK status.
func (r *recipeHandler) DeleteRecipe(c *gin.Context) {
	// Get recipe UUID from path parameter.
	recipeUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.Param("uuid")))

	// Delete the recipe from the database.
	statusCode, err := r.recipeService.DeleteRecipe(c, recipeUUID)
	if err != nil {
		handleError(c, statusCode, "An error occurred while deleting the recipe", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Recipe deleted successfully",
	})
}

// handleError is a generic error handler that logs the error and responds.
func handleError(c *gin.Context, statusCode int, message string, err error) {
	// Log the error message and the error itself.
	log.Printf("[RecipeHandler]: %s, %v", message, err)

	// Send the JSON response with the status code and error message.
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
		"data":    nil,
	})
}

// VoteRecipe handles the HTTP request for voting a recipe.
// It parses the recipe UUID from the URL parameter, gets the user UUID from the context,
// binds the incoming JSON payload to a RequestVoteRecipe struct,
// and calls the recipe service to vote the recipe.
// If any error occurs during this process, it returns the corresponding status code and error message.
// If the recipe is voted successfully, it returns a 200 OK status.
func (r *recipeHandler) VoteRecipe(c *gin.Context) {
	// Parse the recipe UUID from the URL parameter.
	recipeUUID, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		handleError(c, http.StatusBadRequest, "Invalid UUID format", err)
		return
	}

	// Get user UUID from context.
	userUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.MustGet("userUUID")))

	// Bind the incoming JSON payload to a RequestVoteRecipe struct.
	reqVote := &entity.RequestVoteRecipe{}
	if err := c.ShouldBindJSON(reqVote); err != nil {
		handleError(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Vote the recipe.
	statusCode, err := r.recipeService.VoteRecipe(c, userUUID, recipeUUID, reqVote.Vote)
	if err != nil {
		handleError(c, statusCode, "An error occurred while voting the recipe", err)
		return
	}

	// Return a successful response.
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Recipe voted successfully",
	})
}
