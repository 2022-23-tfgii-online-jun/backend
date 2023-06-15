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

// categoryHandler type contains an instance of CategoryService.
type categoryHandler struct {
	categoryService ports.CategoryService
}

// newHandler is a constructor function for initializing categoryHandler with the given CategoryService.
// The return is a pointer to an categoryHandler instance.
func newHandler(categoryService ports.CategoryService) *categoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}
}

// CreateCategory handles the HTTP request for creating a category.
// It binds the incoming JSON payload to the reqCreate struct.
// If any error occurs during this process, it will return a 400 Bad Request status.
// If the category is created successfully, it will return a 200 OK status with the created category.
func (c *categoryHandler) CreateCategory(ctx *gin.Context) {
	// Declare a variable for the incoming request payload.
	reqCreate := &entity.Category{}

	// Bind incoming JSON payload to the reqCreate struct.
	if err := ctx.ShouldBindJSON(reqCreate); err != nil {
		handleError(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Create the category and store it in the database using service.
	createdCategory, err := c.categoryService.CreateCategory(ctx, reqCreate)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "An error occurred while creating the category", err)
		return
	}

	// Return a successful response with the created category.
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Category created successfully",
		"data":    createdCategory,
	})
}

// GetAllCategories handles the HTTP request for getting all categories.
// It retrieves all categories from the database.
// If any error occurs during this process, it will return a 500 Internal Server Error status.
// If the categories are retrieved successfully, it will return a 200 OK status with the retrieved categories.
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

// UpdateCategory handles the HTTP request for updating a category.
// It parses the category UUID from the URL parameter and binds the incoming JSON payload to the reqUpdate struct.
// If any error occurs during this process, it will return a 400 Bad Request status.
// If the category is updated successfully, it will return a 200 OK status with the updated category.
func (c *categoryHandler) UpdateCategory(ctx *gin.Context) {
	// Parse the category UUID from the URL parameter.
	categoryUUID, err := uuid.Parse(ctx.Param("uuid"))
	if err != nil {
		handleError(ctx, http.StatusBadRequest, "Invalid UUID format", err)
		return
	}

	// Bind the incoming JSON payload to an UpdateCategory struct.
	reqUpdate := &entity.Category{}
	if err := ctx.ShouldBindJSON(reqUpdate); err != nil {
		handleError(ctx, http.StatusBadRequest, "Invalid input", err)
		return
	}

	// Update the category in the database.
	updatedCategory, err := c.categoryService.UpdateCategory(categoryUUID, reqUpdate)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "An error occurred while updating the category", err)
		return
	}

	// Return a successful response with the updated category.
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Category updated successfully",
		"data":    updatedCategory,
	})
}

// DeleteCategory handles the HTTP request for deleting a category.
// It gets the category UUID from the path parameter and calls the categoryService to delete the category.
// If any error occurs during this process, it will return the corresponding status code and error message.
// If the category is deleted successfully, it will return a 200 OK status.
func (c *categoryHandler) DeleteCategory(ctx *gin.Context) {
	// Get category UUID from path parameter.
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

// handleError is a generic error handler that logs the error and responds with the corresponding status code and error message.
func handleError(ctx *gin.Context, statusCode int, message string, err error) {
	// Log the error message and the error itself.
	log.Printf("[CategoryHandler]: %s, %v", message, err)

	// Send the JSON response with the status code and error message.
	ctx.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
		"data":    nil,
	})
}
