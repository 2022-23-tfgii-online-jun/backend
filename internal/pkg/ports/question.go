package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// QuestionRepository defines the interface for interacting with the Question data store.
// It outlines the methods that are needed for adding, updating, retrieving and deleting Question records.
type QuestionRepository interface {
	// FindByUUID locates a Question in the data store by its UUID.
	// Returns the found record and an error if the operation fails.
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	// Create inserts a new Question record into the data store.
	// Returns an error if the operation fails.
	Create(value interface{}) error

	// CreateWithOmit creates a new Question record in the data store while ignoring specific fields.
	// This is useful when you want to exclude certain fields from being affected by the operation.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// Update modifies an existing Question record in the data store with the provided data.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first Question record from the data store that matches the provided conditions.
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	// Find retrieves Question records from the data store that match the given conditions.
	// Returns an error if the operation fails.
	Find(out interface{}, conditions ...interface{}) error

	// Delete removes a Question record from the data store.
	// Returns an error if the operation fails.
	Delete(out interface{}) error
}

// QuestionService defines the methods for managing Question data within the application.
// It handles the business logic associated with Question data.
type QuestionService interface {
	// CreateQuestion creates a new Question using the provided request data and user UUID.
	// Returns the created Question and an error if the operation fails.
	CreateQuestion(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateQuestion) (*entity.Question, error)

	// UpdateQuestion is commented out, but when active, it would update an existing Question.
	// Returns the updated Question and an error if the operation fails.

	// DeleteQuestion is commented out, but when active, it would delete a Question based on the provided UUID.
	// Returns an error if the operation fails.

	// GetAllQuestions retrieves all Question records.
	// Returns a slice of Questions and an error if the operation fails.
	GetAllQuestions() ([]*entity.Question, error)

	GetAllQuestionsAndAnswers() ([]*entity.QuestionAndAnswers, error)
}
