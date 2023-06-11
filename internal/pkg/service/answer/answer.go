package answer

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	ErrTypeAssertion  = errors.New("type assertion failed")
	ErrCreatingAnswer = errors.New("error creating answer")
)

// service is the main structure for the answer service which uses ports.AnswerRepository for data access
type service struct {
	repo ports.AnswerRepository
}

// NewService returns a new instance of the answer service with the given answer repository.
// It is used to control the flow of data to and from the data sources.
func NewService(answerRepo ports.AnswerRepository) ports.AnswerService {
	return &service{
		repo: answerRepo,
	}
}

// CreateAnswer is the service for creating an answer and saving it in the database.
// It handles the business logic of validating and storing the answer.
func (s *service) CreateAnswer(c *gin.Context, userUUID uuid.UUID, questionUUID uuid.UUID, createReq *entity.RequestCreateAnswer) (int, error) {
	user := &entity.User{}

	// Find user by UUID
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		// Return error if the user is not found
		return http.StatusNotFound, fmt.Errorf("user not found: %w", err)
	}
	user, ok := foundUser.(*entity.User)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("%w: %v", ErrTypeAssertion, foundUser)
	}

	question := &entity.Question{}

	// Find question by UUID
	foundQuestion, err := s.repo.FindByUUID(questionUUID, question)
	if err != nil {
		// Return error if the question is not found
		return http.StatusNotFound, fmt.Errorf("question not found: %w", err)
	}
	question, ok = foundQuestion.(*entity.Question)
	if !ok {
		return http.StatusInternalServerError, fmt.Errorf("%w: %v", ErrTypeAssertion, foundQuestion)
	}

	// Create a new answer
	answer := &entity.Answer{
		UserID:     user.ID,
		QuestionID: question.ID,
		Text:       createReq.Text,
		IsPublic:   true,
	}

	// Save the answer to the database
	err = s.repo.CreateWithOmit("uuid", answer)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("%w: %v", ErrCreatingAnswer, err)
	}

	// Return the HTTP OK status code if the update is successful
	return http.StatusOK, nil
}
