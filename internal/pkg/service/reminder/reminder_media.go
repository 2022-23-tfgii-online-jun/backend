package reminder

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
)

type reminderMediaService struct {
	repo ports.ReminderMediaRepository
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
