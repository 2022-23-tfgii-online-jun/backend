package ports

import (
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AnswerRepository is the interface that defines the methods for accessing the answer data store.
type AnswerRepository interface {
	FindByUUID(uuid uuid.UUID, out interface{}) (interface{}, error)

	Create(value interface{}) error

	// CreateWithOmit creates a new answer record while omitting specific fields.
	// Returns an error if the operation fails.
	CreateWithOmit(omit string, value interface{}) error

	// UpdateAnswer updates an existing answer record with the provided answer data.
	// Returns an error if the operation fails.
	Update(value interface{}) error

	// First retrieves the first record that matches the given conditions from the database
	// Returns an error if the operation fails.
	First(out interface{}, conditions ...interface{}) error

	Find(out interface{}, conditions ...interface{}) error

	Delete(out interface{}) error
}

// AnswerService is the interface that defines the methods for managing answers in the application.
type AnswerService interface {
	CreateAnswer(c *gin.Context, userUUID uuid.UUID, questionUUID uuid.UUID, createReq *entity.RequestCreateAnswer) (*entity.Answer, error)
	// UpdateAnswer(answerUUID uuid.UUID, updateReq *entity.RequestUpdateAnswer) (*entity.Answer, error)
	// DeleteAnswer(c *gin.Context, answerUUID uuid.UUID) error
	GetAllAnswers(questionUUID uuid.UUID) ([]*entity.Answer, error)
}
	