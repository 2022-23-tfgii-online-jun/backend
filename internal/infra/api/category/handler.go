package category

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type categoryHandler struct {
	categoryService ports.CategoryService
}

// Constructor function for categoryHandler struct
func newHandler(categoryService ports.CategoryService) *categoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}
}

// CreateCategory handler for creating a category
func (c *categoryHandler) CreateCategory(ctx *gin.Context) {
	// Declare a variable for the incoming request payload
	reqCreate := &entity.Category{}

	// Get user uuid from context
	// userUUID, _ := uuid.Parse(fmt.Sprintf("%v", ctx.MustGet("userUUID")))

	// Bind incoming form-data payload to the reqCreate struct.
	if err := ctx.ShouldBindJSON(reqCreate); err != nil {
		handleError(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Create the category and store it in the database.
	createdCategory, err := c.categoryService.CreateCategory(ctx, reqCreate)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "An error occurred while creating the category", err)
		return
	}

	// Return a successful response.
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Category created successfully",
		"data":    createdCategory,
	})
}

// GetAllCategories handles the HTTP request for getting all categories.
func (c *categoryHandler) GetAllCategories(ctx *gin.Context) {
	// Get all categories from the database.
	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "An error occurred while getting the categories", err)
		return
	}

	// Return a successful response with the retrieved categories.
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Categories retrieved successfully",
		"data":    categories,
	})
}

// UpdateCategory handler for updating a category
func (c *categoryHandler) UpdateCategory(ctx *gin.Context) {
	// Parse the category UUID from the URL parameter.
	categoryUUID, err := uuid.Parse(ctx.Param("uuid"))
	if err != nil {
		handleError(ctx, http.StatusBadRequest, "Invalid UUID format", err)
		return
	}

	// Bind the incoming JSON payload to an UpdateCategory struct.
	reqUpdate := &entity.Category{}
	if err := ctx.ShouldBind(reqUpdate); err != nil {
		handleError(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Update the category in the database.
	updatedCategory, err := c.categoryService.UpdateCategory(categoryUUID, reqUpdate)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "An error occurred while updating the category", err)
		return
	}

	// Return a successful response.
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Category updated successfully",
		"data":    updatedCategory,
	})
}

// DeleteCategory handler for deleting a category
func (c *categoryHandler) DeleteCategory(ctx *gin.Context) {
	// Get category uuid from path parameter
	categoryUUID, _ := uuid.Parse(fmt.Sprintf("%v", ctx.Param("uuid")))

	// Delete the category from the database.
	statusCode, err := c.categoryService.DeleteCategory(ctx, categoryUUID)
	if err != nil {
		handleError(ctx, statusCode, "An error occurred while deleting the category", err)
		return
	}

	// Return a successful response.
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Category deleted successfully",
	})
}

// handleError is a generic error handler that logs the error and responds
func handleError(ctx *gin.Context, statusCode int, message string, err error) {
	// Log the error message and the error itself
	log.Printf("[CategoryHandler]: %s, %v", message, err)

	// Send the JSON response with the status code and error message
	ctx.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
		"data":    nil,
	})
}
