package recipe

import (
	"fmt"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
)

// service struct holds the necessary dependencies for the recipe media service
type recipeMediaService struct {
	repo ports.RecipeMediaRepository
}

func (s *recipeMediaService) FindByRecipeID(recipeID int, i *[]*entity.RecipeMedia) error {
	return s.repo.Find(&entity.RecipeMedia{}, i, "recipe_id = ?", recipeID)
}

func (s *recipeMediaService) DeleteRecipeMedia(recipeMedia *entity.RecipeMedia) error {
	// Delete the media from the database
	err := s.repo.Delete(recipeMedia)
	if err != nil {
		return fmt.Errorf("error deleting recipe media: %s", err)
	}
	return nil
}

// NewRecipeMediaService returns a new instance of the recipeMedia service with the given recipeMedia repository.
func NewRecipeMediaService(repo ports.RecipeMediaRepository) ports.RecipeMediaService {
	return &recipeMediaService{
		repo: repo,
	}
}

// CreateRecipeMedia creates a new recipe_media association and saves it in the repository.
func (s *recipeMediaService) CreateRecipeMedia(recipeMedia *entity.RecipeMedia) error {
	if recipeMedia == nil {
		return ports.ErrInvalidInput
	}
	return s.repo.Create(recipeMedia)
}
