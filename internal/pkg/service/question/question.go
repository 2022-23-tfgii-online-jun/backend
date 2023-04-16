package question

import (
	"fmt"

	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/emur-uy/backend/internal/pkg/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type service struct {
	repo ports.QuestionRepository
}

// NewService returns a new instance of the question service with the given question repository.
func NewService(questionRepo ports.QuestionRepository) ports.QuestionService {
	return &service{
		repo: questionRepo,
	}
}

// CreateQuestion is the service for creating a question and saving it in the database
func (s *service) CreateQuestion(c *gin.Context, userUUID uuid.UUID, createReq *entity.RequestCreateQuestion) (*entity.Question, error) {
	user := &entity.User{}

	// Find user by UUID
	foundUser, err := s.repo.FindByUUID(userUUID, user)
	if err != nil {
		// Return error if the user is not found
		return nil, err
	}
	// Perform type assertion to convert foundUser to *entity.User
	user, ok := foundUser.(*entity.User)
	if !ok {
		return nil, fmt.Errorf("type assertion failed")
	}

	// Create a new question
	question := &entity.Question{
		Text: createReq.Text,
	}

	// Save the question to the database
	err = s.repo.CreateWithOmit("uuid", question)
	if err != nil {
		return nil, fmt.Errorf("error creating question: %s", err)
	}

	return question, nil
}

// GetAllQuestions returns all questions stored in the database
func (s *service) GetAllQuestions() ([]*entity.Question, error) {
	// Get all questions from the database
	var questions []*entity.Question
	if err := s.repo.Find(&questions); err != nil {
		return nil, err
	}

	return questions, nil
}
