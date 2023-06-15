package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AnswerRepository is an interface that acts as a contract for the data access layer,
// requiring implementations to provide methods for querying and modifying answer data.
type AnswerRepository interface {
	// FindByUUID retrieves an Answer based on its UUID.
	// Returns the answer and an error if any occurred.
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// Create takes a new Answer value and adds it to the data store.
	// Returns an error if the operation fails.
	Create(value interface{}) error

	// CreateWithOmit creates a new answer record while omitting specific fields.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// Update updates an existing Answer value in the data store.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the database.
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	// Find retrieves all Answer values that match the given conditions from the database.
	// Returns an error if the operation fails.
	Find(out interface{}, conditions ...interface{}) error

	// Delete removes an existing Answer from the data store.
	// Returns an error if the operation fails.
	Delete(out interface{}) error
}

// AnswerService is an interface defining a contract for business logic operators related to Answers.
// It works with the entity layer to manipulate Answer data.
type AnswerService interface {
	// CreateAnswer takes the user's UUID, the question's UUID, and a request to create an Answer.
	// Returns the status and an error if any occurred.
	CreateAnswer(c *gin.Context, userUUID uuid.UUID, questionUUID uuid.UUID, createReq *entity.RequestCreateAnswer) (int, error)
}
