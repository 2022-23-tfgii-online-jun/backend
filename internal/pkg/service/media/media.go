package media

import (
	"errors"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
)

var (
	ErrCreatingMedia = errors.New("error creating media")
	ErrDeletingMedia = errors.New("error deleting media")
)

// service struct holds the necessary dependencies for the media service
type service struct {
	repo ports.MediaRepository
}

// NewService returns a new instance of the media service with the given media repository.
func NewService(mediaRepo ports.MediaRepository) ports.MediaService {
	return &service{
		repo: mediaRepo,
	}
}

// FindByMediaID finds a media by its ID and populates the provided media entity with the result.
func (s *service) FindByMediaID(mediaID int, i *entity.Media) error {
	return s.repo.Find(&entity.Media{}, &i, "id = ?", mediaID)
}

// CreateMedia is the service for creating a media and saving it in the database
func (s *service) CreateMedia(media *entity.Media) error {
	// Save the media to the database, omitting the uuid field
	err := s.repo.CreateWithOmit("uuid", media)
	if err != nil {
		return ErrCreatingMedia
	}
	return nil
}

// DeleteMedia is the service for deleting a media from the database
func (s *service) DeleteMedia(media *entity.Media) error {
	// Delete the media from the database
	err := s.repo.Delete(media)
	if err != nil {
		return ErrDeletingMedia
	}
	return nil
}
