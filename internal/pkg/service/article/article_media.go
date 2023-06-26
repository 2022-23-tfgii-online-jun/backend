package article

import (
	"fmt"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
)

// service struct holds the necessary dependencies for the article media service
type articleMediaService struct {
	repo ports.ArticleMediaRepository
}

func (s *articleMediaService) FindByArticleID(articleID int, i *[]*entity.ArticleMedia) error {
	return s.repo.Find(&entity.ArticleMedia{}, i, "article_id = ?", articleID)
}

func (s *articleMediaService) DeleteArticleMedia(articleMedia *entity.ArticleMedia) error {
	// Delete the media from the database
	err := s.repo.Delete(articleMedia)
	if err != nil {
		return fmt.Errorf("error deleting article media: %s", err)
	}
	return nil
}

// NewArticleMediaService returns a new instance of the articleMedia service with the given articleMedia repository.
func NewArticleMediaService(repo ports.ArticleMediaRepository) ports.ArticleMediaService {
	return &articleMediaService{
		repo: repo,
	}
}

// CreateArticleMedia creates a new article_media association and saves it in the repository.
func (s *articleMediaService) CreateArticleMedia(articleMedia *entity.ArticleMedia) error {
	if articleMedia == nil {
		return ports.ErrInvalidInput
	}
	return s.repo.Create(articleMedia)
}
