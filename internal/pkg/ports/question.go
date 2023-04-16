package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// QuestionRepository is the interface that defines the methods for accessing the question data store.
type QuestionRepository interface {
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	Create(value interface{}) error

	// CreateWithOmit creates a new question record while omitting specific fields.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// UpdateQuestion updates an existing question record with the provided question data.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the database
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	Find(out interface{}, conditions ...interface{}) error

	Delete(out interface{}) error
}

// QuestionService is the interface that defines the methods for managing questions in the application.
type QuestionService interface {
	CreateQuestion(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateQuestion) (*entity.Question, error)
	// UpdateQuestion(questionUUID uuid.UUID, updateReq *entity.RequestUpdateQuestion) (*entity.Question, error)
	// DeleteQuestion(c *gin.Context, questionUUID uuid.UUID) error
	GetAllQuestions() ([]*entity.Question, error)
}
