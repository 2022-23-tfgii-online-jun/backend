package ports

import (
	"errors"

	"github.com/emur-uy/backend/internal/pkg/entity"
)

// ErrInvalidInput is an error indicating that the input is invalid.
var ErrInvalidInput = errors.New("invalid input")

// MediaRepository defines the interface for interacting with the media data store.
type MediaRepository interface {
	CreateWithOmit(omit string, value interface{}) error
	// Any other methods you might need for media management.
}

// MediaService defines the interface for managing media in the application.
type MediaService interface {
	CreateMedia(media *entity.Media) error
	// Any other methods you might need for media management.
}

// ReminderMediaRepository defines the interface for interacting with the reminder_media data store.
type ReminderMediaRepository interface {
	Create(value interface{}) error
	// Any other methods you might need for reminder_media management.
}

// ReminderMediaService defines the interface for managing reminder_media in the application.
type ReminderMediaService interface {
	CreateReminderMedia(reminderMedia *entity.ReminderMedia) error
	// Any other methods you might need for reminder_media management.
}
