package media

import (
	"fmt"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
)

type service struct {
	repo ports.MediaRepository
}

func (s *service) FindByMediaID(mediaID int, i *entity.Media) error {
	return s.repo.Find(&entity.Media{}, &i, "id = ?", mediaID)
}

// NewService returns a new instance of the media service with the given media repository.
func NewService(mediaRepo ports.MediaRepository) ports.MediaService {
	return &service{
		repo: mediaRepo,
	}
}

// CreateMedia is the service for creating a media and saving it in the database
func (s *service) CreateMedia(media *entity.Media) error {
	// Save the media to the database
	err := s.repo.CreateWithOmit("uuid", media)
	if err != nil {
		return fmt.Errorf("error creating media: %s", err)
	}
	return nil
}

// DeleteMedia is the service for deleting a media from the database
func (s *service) DeleteMedia(media *entity.Media) error {
	// Delete the media from the database
	err := s.repo.Delete(media)
	if err != nil {
		return fmt.Errorf("error deleting media: %s", err)
	}
	return nil
}
