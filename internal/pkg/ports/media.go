package ports

import (
	"errors"

	"github.com/emur-uy/backend/internal/pkg/entity"
)

// ErrInvalidInput is an error instance to be returned when an invalid input is provided to a function or method.
var ErrInvalidInput = errors.New("invalid input")

// MediaRepository defines an interface for accessing the media data store.
// It is a contract for all database operations related to Media data.
type MediaRepository interface {
	// CreateWithOmit creates a new Media record while omitting specific fields.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// Delete removes a Media record from the data store.
	// Returns an error if the operation fails.
	Delete(value interface{}) error

	// Find retrieves all Media records that match the given conditions.
	// Returns an error if the operation fails.
	Find(model interface{}, dest interface{}, conditions ...interface{}) error
}

// MediaService is an interface defining a contract for business logic operators related to Media.
// It works with the entity layer to manipulate Media data.
type MediaService interface {
	// CreateMedia creates a new Media entity.
	// Returns an error if the operation fails.
	CreateMedia(media *entity.Media) error

	// DeleteMedia removes an existing Media entity.
	// Returns an error if the operation fails.
	DeleteMedia(media *entity.Media) error

	// FindByMediaID retrieves a Media entity based on its ID.
	// Returns an error if the operation fails.
	FindByMediaID(id int, i *entity.Media) error
}

// ReminderMediaRepository defines an interface for accessing the reminder_media data store.
// It is a contract for all database operations related to ReminderMedia data.
type ReminderMediaRepository interface {
	// Create creates a new ReminderMedia record in the data store.
	// Returns an error if the operation fails.
	Create(value interface{}) error

	// Delete removes a ReminderMedia record from the data store.
	// Returns an error if the operation fails.
	Delete(value interface{}) error

	// Find retrieves all ReminderMedia records that match the given conditions.
	// Returns an error if the operation fails.
	Find(model interface{}, dest interface{}, conditions ...interface{}) error
}

// ReminderMediaService is an interface defining a contract for business logic operators related to ReminderMedia.
// It works with the entity layer to manipulate ReminderMedia data.
type ReminderMediaService interface {
	// CreateReminderMedia creates a new ReminderMedia entity.
	// Returns an error if the operation fails.
	CreateReminderMedia(reminderMedia *entity.ReminderMedia) error

	// DeleteReminderMedia removes an existing ReminderMedia entity.
	// Returns an error if the operation fails.
	DeleteReminderMedia(reminderMedia *entity.ReminderMedia) error

	// FindByReminderID retrieves ReminderMedia entities based on the Reminder ID.
	// Returns an error if the operation fails.
	FindByReminderID(id int, i *[]*entity.ReminderMedia) error
}
