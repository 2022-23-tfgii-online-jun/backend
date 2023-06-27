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

// RecipeMediaRepository defines an interface for accessing the recipe_media data store.
// It is a contract for all database operations related to RecipeMedia data.
type RecipeMediaRepository interface {
	// Create creates a new RecipeMedia record in the data store.
	// Returns an error if the operation fails.
	Create(value interface{}) error

	// Delete removes a RecipeMedia record from the data store.
	// Returns an error if the operation fails.
	Delete(value interface{}) error

	// Find retrieves all RecipeMedia records that match the given conditions.
	// Returns an error if the operation fails.
	Find(model interface{}, dest interface{}, conditions ...interface{}) error
}

// RecipeMediaService is an interface defining a contract for business logic operators related to RecipeMedia.
// It works with the entity layer to manipulate RecipeMedia data.
type RecipeMediaService interface {
	// CreateRecipeMedia creates a new RecipeMedia entity.
	// Returns an error if the operation fails.
	CreateRecipeMedia(recipeMedia *entity.RecipeMedia) error

	// DeleteRecipeMedia removes an existing RecipeMedia entity.
	// Returns an error if the operation fails.
	DeleteRecipeMedia(recipeMedia *entity.RecipeMedia) error

	// FindByRecipeID retrieves RecipeMedia entities based on the Recipe ID.
	// Returns an error if the operation fails.
	FindByRecipeID(id int, i *[]*entity.RecipeMedia) error
}

// ArticleMediaRepository defines an interface for accessing the article_media data store.
// It is a contract for all database operations related to ArticleMedia data.
type ArticleMediaRepository interface {
	// Create creates a new ArticleMedia record in the data store.
	// Returns an error if the operation fails.
	Create(value interface{}) error

	// Delete removes an ArticleMedia record from the data store.
	// Returns an error if the operation fails.
	Delete(value interface{}) error

	// Find retrieves all ArticleMedia records that match the given conditions.
	// Returns an error if the operation fails.
	Find(model interface{}, dest interface{}, conditions ...interface{}) error
}

// ArticleMediaService is an interface defining a contract for business logic operators related to ArticleMedia.
// It works with the entity layer to manipulate ArticleMedia data.
type ArticleMediaService interface {
	// CreateArticleMedia creates a new ArticleMedia entity.
	// Returns an error if the operation fails.
	CreateArticleMedia(articleMedia *entity.ArticleMedia) error

	// DeleteArticleMedia removes an existing ArticleMedia entity.
	// Returns an error if the operation fails.
	DeleteArticleMedia(articleMedia *entity.ArticleMedia) error

	// FindByArticleID retrieves ArticleMedia entities based on the Article ID.
	// Returns an error if the operation fails.
	FindByArticleID(id int, i *[]*entity.ArticleMedia) error
}
