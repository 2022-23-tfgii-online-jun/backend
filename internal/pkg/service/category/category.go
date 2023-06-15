package category

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
	ErrTypeAssertion    = errors.New("type assertion failed")
	ErrCreatingCategory = errors.New("error creating category")
	ErrUpdatingCategory = errors.New("error updating category")
	ErrDeletingCategory = errors.New("failed to delete category")
)

// service struct holds required dependencies for the category service
type service struct {
	repo ports.CategoryRepository
}

// NewService creates a new instance of the category service with the provided category repository.
func NewService(categoryRepo ports.CategoryRepository) ports.CategoryService {
	return &service{
		repo: categoryRepo,
	}
}

// CreateCategory is a service function for creating a category and saving it in the database
func (s *service) CreateCategory(c *gin.Context, createReq *entity.Category) (int, error) {
	category := &entity.Category{
		Name: createReq.Name,
	}

	err := s.repo.CreateWithOmit("uuid", category)
	if err != nil {
		return http.StatusInternalServerError, ErrCreatingCategory
	}

	return http.StatusOK, nil
}

// GetAllCategories retrieves all categories stored in the database
func (s *service) GetAllCategories() ([]*entity.Category, error) {
	var categories []*entity.Category
	if err := s.repo.Find(&categories); err != nil {
		return nil, err
	}

	return categories, nil
}

// UpdateCategory is a service function for updating a category in the database
func (s *service) UpdateCategory(categoryUUID uuid.UUID, updateReq *entity.Category) (int, error) {
	category := &entity.Category{}
	foundCategory, err := s.repo.FindByUUID(categoryUUID, category)
	if err != nil {
		return http.StatusNotFound, err
	}

	category, ok := foundCategory.(*entity.Category)
	if !ok {
		return http.StatusInternalServerError, ErrTypeAssertion
	}

	category.Name = updateReq.Name

	err = s.repo.Update(category)
	if err != nil {
		return http.StatusInternalServerError, ErrUpdatingCategory
	}

	return http.StatusOK, nil
}

// DeleteCategory is a service function to delete a category from the database by its UUID.
func (s *service) DeleteCategory(c *gin.Context, categoryUUID uuid.UUID) (int, error) {
	category := &entity.Category{}
	foundCategory, err := s.repo.FindByUUID(categoryUUID, category)
	if err != nil {
		return http.StatusNotFound, fmt.Errorf("category not found")
	}

	category, ok := foundCategory.(*entity.Category)
	if !ok {
		return http.StatusInternalServerError, ErrTypeAssertion
	}

	err = s.repo.Delete(category)
	if err != nil {
		return http.StatusInternalServerError, ErrDeletingCategory
	}

	return http.StatusOK, nil
}
