package reminder

import (
	"fmt"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
)

type reminderMediaService struct {
	repo ports.ReminderMediaRepository
}

func (s *reminderMediaService) FindByReminderID(reminderID int, i *[]*entity.ReminderMedia) error {
	return s.repo.Find(&entity.ReminderMedia{}, &i, "reminder_id = ?", reminderID)
}

func (s *reminderMediaService) DeleteReminderMedia(reminderMedia *entity.ReminderMedia) error {
	// Delete the media from the database
	err := s.repo.Delete(reminderMedia)
	if err != nil {
		return fmt.Errorf("error deleting reminder media: %s", err)
	}
	return nil
}

// NewReminderMediaService returns a new instance of the reminderMedia service with the given reminderMedia repository.
func NewReminderMediaService(repo ports.ReminderMediaRepository) ports.ReminderMediaService {
	return &reminderMediaService{
		repo: repo,
	}
}

// CreateReminderMedia creates a new reminder_media association and saves it in the repository.
func (s *reminderMediaService) CreateReminderMedia(reminderMedia *entity.ReminderMedia) error {
	if reminderMedia == nil {
		return ports.ErrInvalidInput
	}
	return s.repo.Create(reminderMedia)
}
