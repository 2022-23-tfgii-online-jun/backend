package category

import (
	"fmt"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type service struct {
	repo ports.CategoryRepository
}

// NewService returns a new instance of the category service with the given category repository.
func NewService(categoryRepo ports.CategoryRepository) ports.CategoryService {
	return &service{
		repo: categoryRepo,
	}
}

// CreateCategory is the service for creating a category and saving it in the database
func (s *service) CreateCategory(c *gin.Context, createReq *entity.Category) (int, error) {
	// Create a new category
	category := &entity.Category{
		Name: createReq.Name,
	}

	// Save the category to the database
	err := s.repo.CreateWithOmit("uuid", category)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error creating category: %s", err)
	}

	// Return the HTTP OK status code if the creation is successful
	return http.StatusOK, nil
}

// GetAllCategories returns all categories stored in the database
func (s *service) GetAllCategories() ([]*entity.Category, error) {
	// Get all categories from the database
	var categories []*entity.Category
	if err := s.repo.Find(&categories); err != nil {
		return nil, err
	}

	return categories, nil
}

// UpdateCategory is the service for updating a category in the database
func (s *service) UpdateCategory(categoryUUID uuid.UUID, updateReq *entity.Category) (int, error) {
	// Find the existing category by UUID
	category := &entity.Category{}
	foundCategory, err := s.repo.FindByUUID(categoryUUID, category)
	if err != nil {
		// Return error if the category is not found
		return http.StatusNotFound, err
	}
	// Perform type assertion to convert foundCategory to *entity.Category
	category, ok := foundCategory.(*entity.Category)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	// Update the category fields with the new data from the update request
	category.Name = updateReq.Name

	// Update the category in the database
	err = s.repo.Update(category)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error updating category: %s", err)
	}

	// Return the HTTP OK status code if the update is successful
	return http.StatusOK, nil
}

// DeleteCategory deletes a category from the database by its UUID.
func (s *service) DeleteCategory(c *gin.Context, categoryUUID uuid.UUID) (int, error) {
	// Retrieve the category from the repository by its UUID.
	category := &entity.Category{}
	foundCategory, err := s.repo.FindByUUID(categoryUUID, category)
	if err != nil {
		// Return an error response if the category is not found.
		return http.StatusNotFound, fmt.Errorf("category not found")
	}

	// Perform type assertion to convert foundCategory to *entity.Category.
	category, ok := foundCategory.(*entity.Category)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("type assertion failed")
	}

	// Delete the category from the repository.
	err = s.repo.Delete(category)
	if err != nil {
		// Return an error response if there was an issue deleting the category.
		return http.StatusInternalServerError, fmt.Errorf("failed to delete category")
	}

	// Return a success response.
	return http.StatusOK, nil
}
